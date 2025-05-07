package user

import (
	"errors"
	"time"

	"gorm.io/gorm"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/wataee/GoQuestions/config"
	"github.com/wataee/GoQuestions/internal/database/repository"
	"github.com/wataee/GoQuestions/internal/models"
)

var errInvalidClaims = errors.New("невалидные claims") 

type UserService interface {
	Login(input models.UserInput) (models.TokenPair, error)
	RefreshToken(refreshToken string) (models.TokenPair, error)
	GetProfile(userID int) (models.ProfileDTO, error)

}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Login(input models.UserInput) (models.TokenPair, error) {
	user, err := s.repo.GetByUsername(input.Username)
	if err == gorm.ErrRecordNotFound {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
		if err != nil {
			return models.TokenPair{}, err
		}

		input.Password = string(hashedPassword)

		userId, err := s.repo.CreateUser(input)
		if err != nil {
			return models.TokenPair{}, err
		}
		return s.GenerateTokenPair(userId, input.Username, input.Role), nil
	}
	if err != nil {
		return models.TokenPair{}, nil
	}
	// если зареган

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return models.TokenPair{}, err
	}
	if input.Role != user.Role {
		return models.TokenPair{}, errors.New("не совпадает роль из БД с введённой ролью")
	}
	return s.GenerateTokenPair(int(user.ID), input.Username, input.Role), nil
}

func (s *userService) GetProfile(userID int) (models.ProfileDTO, error){
	user, err := s.repo.GetByID(userID)
	if err != nil {
		return models.ProfileDTO{}, err
	}
	return models.ProfileDTO{
		ID: user.ID,
		Username: user.Username,
		Role: user.Role,
		Answers: user.Answers,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (s *userService)GenerateTokenPair(userID int, username string, role string) models.TokenPair {
	accessToken, _ := s.generateToken(userID, username, role, 2 * time.Hour)
	refreshToken, _ := s.generateToken(userID, username, role, 7 * 24 * time.Hour)

	return models.TokenPair{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}
}

func (s *userService)generateToken(userID int, username string, role string, ttl time.Duration) (string, error) {
	claims := models.UserClaims{
		UserID: userID,
		Username: username,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer: "goquestions",
		},
	}
	
	createToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := createToken.SignedString(config.JwtKey)
	if err != nil {
		return "", err
	}
	
	return tokenString, nil
}

func (s *userService)RefreshToken(refreshToken string) (models.TokenPair, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &models.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return config.JwtKey, nil
	})
	if err != nil || !token.Valid {
		return models.TokenPair{}, err
	}

	claims, ok := token.Claims.(*models.UserClaims)
	if !ok {
		return models.TokenPair{}, errInvalidClaims
	}

	return s.GenerateTokenPair(claims.UserID, claims.Username, claims.Role), nil

}
 

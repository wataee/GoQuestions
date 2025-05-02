package user

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/wataee/GoQuestions/config"
	"github.com/wataee/GoQuestions/internal/database/repository"
	"github.com/wataee/GoQuestions/internal/models"
)

var errInvalidClaims = errors.New("невалидные claims") 

type UserService interface {
	Login(input models.UserInput) (TokenPair, error)
	RefreshToken(refreshToken string) (TokenPair, error)

}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Login(input models.UserInput) (TokenPair, error) {
	isUserInDb, err := s.repo.FindByUsername(input.Username)

	if err != nil {
		return TokenPair{}, err
	}

	if !isUserInDb {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
		if err != nil {
			return TokenPair{}, err
		}

		input.Password = string(hashedPassword)
		
		userId, err := s.repo.CreateUser(input)
		if err != nil {
			return TokenPair{}, err
		}
		return s.GenerateTokenPair(userId, input.Username, input.Role), nil
	}
	// если зареган
	user, err := s.repo.GetByUsername(input.Username)
	if err != nil {
		return TokenPair{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return TokenPair{}, err
	}
	fmt.Println("Если зареган")
	return s.GenerateTokenPair(int(user.ID), input.Username, input.Role), nil
}

type UserClaims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

type TokenPair struct {
	AccessToken string `json:"access_token"`
	RefreshToken string  `json:"refresh_token"`
}

func (s *userService)GenerateTokenPair(userID int, username string, role string) TokenPair {
	accessToken, _ := s.generateToken(userID, username, role, 2 * time.Hour)
	refreshToken, _ := s.generateToken(userID, username, role, 7 * 24 * time.Hour)

	return TokenPair{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}
}

func (s *userService)generateToken(userID int, username string, role string, ttl time.Duration) (string, error) {
	claims := UserClaims{
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

func (s *userService)RefreshToken(refreshToken string) (TokenPair, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return config.JwtKey, nil
	})
	if err != nil || !token.Valid {
		return TokenPair{}, err
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return TokenPair{}, errInvalidClaims
	}

	return s.GenerateTokenPair(claims.UserID, claims.Username, claims.Role), nil

}
 

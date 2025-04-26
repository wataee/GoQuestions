package user

import (

)

type UserService interface {
	login(input UserInput) (string, error)
}


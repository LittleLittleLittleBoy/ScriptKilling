package service

import (
	"log"

	"github.com/littlelittlelittleboy/scriptkilling/constants"
	"github.com/littlelittlelittleboy/scriptkilling/db"
	"github.com/littlelittlelittleboy/scriptkilling/models"
)

type UserService struct{}

func (s UserService) Login(username string, password string) constants.LoginStatus {
	user := models.User{
		Username: username,
		Password: password,
	}
	has, err := db.Engine.Get(&user)
	if err != nil {
		log.Printf("db error {%s}", err)
	}

	if !has {
		return constants.LOGIN_PASSWORD_ERROR
	}

	return constants.LOGIN_SUCCESS
}

func (s UserService) Register(username string, password string) constants.LoginStatus {
	user := models.User{
		Username: username,
		Password: password,
	}
	count, err := db.Engine.Insert(&user)
	if err != nil {
		log.Printf("db error {%s}", err)
	}

	if count == 0 {
		return constants.LOGIN_ACCOUNT_EXISTED
	}

	return constants.LOGIN_SUCCESS
}

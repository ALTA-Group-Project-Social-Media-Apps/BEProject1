package services

import (
	"errors"

	"github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/config"
	"github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/features/user/domain"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &userService{
		qry: repo,
	}
}

func (us *userService) Delete(ID uint) error {
	err := us.qry.Delete(ID)
	if err != nil {
		log.Error(err.Error())
		return errors.New("no data")
	}

	return nil
}

func (us *userService) AddUser(newUser domain.Core) (domain.Core, error) {
	var temp string
	if newUser.Password != "" {
		generate, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
		newUser.Password = string(generate)
	}
	if newUser.Password == "" {
		temp = newUser.Password
	} else {
		temp = "error"
	}
	res, err := us.qry.Insert(newUser)

	if err != nil {
		if temp == "" {
			return domain.Core{}, errors.New("cannot encript password")
		}
		return domain.Core{}, errors.New(config.DUPLICATED_DATA)
	}

	return res, nil
}

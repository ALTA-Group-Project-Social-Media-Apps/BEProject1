package domain

import "github.com/labstack/echo/v4"

type Core struct {
	ID       uint
	Username string
	Email    string
	Password string
	Photo    string
}

type Repository interface {
	Delete(ID uint) error
}

type Service interface {
	Delete(id uint) error
}

type Handler interface {
	DeleteByID() echo.HandlerFunc
}

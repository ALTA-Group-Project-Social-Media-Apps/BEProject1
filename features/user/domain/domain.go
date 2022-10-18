package domain

type Core struct {
	ID       uint
	Username string
	Email    string
	Password string
	Photo    string
}

type Repository interface {
}

type Service interface {
}

type Handler interface {
}

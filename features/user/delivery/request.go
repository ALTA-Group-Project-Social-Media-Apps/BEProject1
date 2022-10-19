package delivery

import "github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/features/user/domain"

type RegisterFormat struct {
	Username string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case RegisterFormat:
		cnv := i.(RegisterFormat)
		return domain.Core{Username: cnv.Username, Email: cnv.Email, Password: cnv.Password}
	}

	return domain.Core{}
}

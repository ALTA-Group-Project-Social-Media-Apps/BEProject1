package repository

import (
	"github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/features/user/domain"

	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(dbConn *gorm.DB) domain.Repository {
	return &repoQuery{
		db: dbConn,
	}
}

func (rq *repoQuery) Delete(ID uint) error {
	var resQry User
	if err := rq.db.Delete(&resQry, "ID = ?", ID).Error; err != nil {
		return err
	}

	return nil
}

func (rq *repoQuery) Insert(newUser domain.Core) (domain.Core, error) {
	var cnv User
	cnv = FromDomain(newUser)
	if err := rq.db.Create(&cnv).Error; err != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	newUser = ToDomain(cnv)
	return newUser, nil
}

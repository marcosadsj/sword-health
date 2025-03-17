package repository

import "gorm.io/gorm"

type IRepository interface {
	Teste()
}

type GOORMRepository struct {
	DB *gorm.DB
}

func (r GOORMRepository) Teste() {

}

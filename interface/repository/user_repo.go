package repository

import (
	"github.com/IbrahimAbunaib/go-mux-api/Documents/Clean_Architect_Api/Domain/model"
	"github.com/IbrahimAbunaib/go-mux-api/Documents/Clean_Architect_Api/Usecase/repository"
	"github.com/jinzhu/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) FindAll(u []*model.User) ([]model.User, error) {
	err := ur.db.Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *userRepository) Create(u *model.User) (*model.User, error) {
	if err := ur.db.Create(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

package repository

import "github.com/IbrahimAbunaib/go-mux-api/Documents/Clean_Architect_Api/Domain/model"

type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
	Create(u []*model.User) ([]*model.User, error)
}

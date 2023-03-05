package presenter

import (
	"github.com/IbrahimAbunaib/go-mux-api/Documents/Clean_Architect_Api/Domain/model"

	"github.com/IbrahimAbunaib/go-mux-api/Documents/Clean_Architect_Api/usecase/presenter"
)

type userPresenter struct{}

func NewUserPresenter() presenter.UserPresenter {
	return &userPresenter{}
}

func (up *userPresenter) ResponseUsers(us []*model.User) []*model.User {
	for _, u := range us {
		u.Name = "Mr." + u.Name
	}
	return us
}

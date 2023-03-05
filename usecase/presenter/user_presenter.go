package presenter

import "github.com/IbrahimAbunaib/go-mux-api/Documents/Clean_Architect_Api/Domain/model"

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
}

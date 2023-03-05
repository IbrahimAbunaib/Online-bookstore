package registry

import (
	"github.com/IbrahimAbunaib/go-mux-api/Documents/Clean_Architect_Api/interface/controller"
	ip "github.com/IbrahimAbunaib/go-mux-api/Documents/Clean_Architect_Api/interface/presenter"
	ir "github.com/IbrahimAbunaib/go-mux-api/Documents/Clean_Architect_Api/interface/repository"
	"github.com/IbrahimAbunaib/go-mux-api/Documents/Clean_Architect_Api/usecase/interactor"
)

func (r *registry) NewUserController() controller.UserController {
	userInteractor := interactor.NewUserInteractor(
		ir.NewUserRepository(r.db),
		ip.NewUserPresenter(),
		ir.NewDBRepository(r.db),
	)
	return controller.NewUserController(userInteractor)
}

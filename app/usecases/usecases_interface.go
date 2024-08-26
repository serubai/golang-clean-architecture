package usecases

import "github.com/ubaidillahhf/go-clarch/app/infra/repository"

type AppUseCase struct {
	ProductUsecase IProductUsecase
	UserUsecase    IUserUsecase
}

func NewAppUseCase(
	ProductRepo repository.IProductRepository,
	UserRepo repository.IUserRepository,
) AppUseCase {
	return AppUseCase{
		ProductUsecase: NewProductUsecase(&ProductRepo),
		UserUsecase:    NewUserUsecase(&UserRepo),
	}
}

package usecases

import (
	"context"

	"github.com/ubaidillahhf/go-clarch/app/domain"
	"github.com/ubaidillahhf/go-clarch/app/infra/exception"
	"github.com/ubaidillahhf/go-clarch/app/infra/repository"
)

type IProductUsecase interface {
	Create(ctx context.Context, request domain.CreateProductRequest) (domain.Product, *exception.Error)
	List(ctx context.Context) ([]domain.Product, *exception.Error)
}

func NewProductUsecase(productRepository *repository.IProductRepository) IProductUsecase {
	return &productUsecase{
		ProductRepository: *productRepository,
	}
}

type productUsecase struct {
	ProductRepository repository.IProductRepository
}

func (service *productUsecase) Create(ctx context.Context, request domain.CreateProductRequest) (res domain.Product, err *exception.Error) {

	newProduct := domain.Product{
		Name:     request.Name,
		Price:    request.Price,
		Quantity: request.Quantity,
	}

	p, pErr := service.ProductRepository.Insert(ctx, newProduct)
	if pErr != nil {
		return res, pErr
	}

	return p, nil
}

func (service *productUsecase) List(ctx context.Context) (responses []domain.Product, err *exception.Error) {
	products, pErr := service.ProductRepository.FindAll(ctx)
	if pErr != nil {
		return responses, pErr
	}

	for _, product := range products {
		responses = append(responses, domain.Product{
			Id:       product.Id,
			Name:     product.Name,
			Price:    product.Price,
			Quantity: product.Quantity,
		})
	}
	return responses, nil
}

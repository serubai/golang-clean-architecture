package handler

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/ubaidillahhf/go-clarch/app/domain"
	"github.com/ubaidillahhf/go-clarch/app/infra/exception"
	"github.com/ubaidillahhf/go-clarch/app/infra/presenter"
	"github.com/ubaidillahhf/go-clarch/app/usecases"
)

type IProductHandler interface {
	Create(c *fiber.Ctx) error
	List(c *fiber.Ctx) error
}

type productHandler struct {
	pUsecase usecases.IProductUsecase
}

func NewProductHandler(pUsecase *usecases.IProductUsecase) IProductHandler {
	return &productHandler{
		pUsecase: *pUsecase,
	}
}

//		@Summary        Create new product
//		@Description    Create new product
//		@Tags           Products
//		@Accept         json
//		@Produce        json
//	 @Param data body domain.Product true "body payload"
//		@Security       JWT
//		@Success        200 {object} presenter.SuccessResponse{data=domain.Product}
//		@Router         /products [post]
func (co *productHandler) Create(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var request domain.CreateProductRequest
	if err := c.BodyParser(&request); err != nil {
		return c.JSON(presenter.Error(err.Error(), nil, exception.BadRequestError))
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(request); err != nil {
		return c.JSON(presenter.Error(err.Error(), nil, exception.BadRequestError))
	}

	res, resErr := co.pUsecase.Create(ctx, request)
	if resErr != nil {
		return c.JSON(presenter.Error(resErr.Err.Error(), nil, resErr.Code))
	}

	return c.JSON(presenter.Success("Success", res, nil))
}

func (co *productHandler) List(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	responses, err := co.pUsecase.List(ctx)
	if err != nil {
		return c.JSON(presenter.Error(err.Err.Error(), nil, err.Code))
	}

	return c.JSON(presenter.Success("Success", responses, nil))
}

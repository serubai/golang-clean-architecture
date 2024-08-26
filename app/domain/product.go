package domain

type Product struct {
	Id       string `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Price    int64  `json:"price" bson:"price"`
	Quantity int32  `json:"quantity" bson:"quantity"`
}

type CreateProductRequest struct {
	Name     string `validate:"required"`
	Price    int64  `validate:"required"`
	Quantity int32  `validate:"required"`
}

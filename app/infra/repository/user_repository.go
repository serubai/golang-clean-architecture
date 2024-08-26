package repository

import (
	"context"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/ubaidillahhf/go-clarch/app/domain"
	"github.com/ubaidillahhf/go-clarch/app/infra/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserRepository interface {
	Insert(ctx context.Context, newData domain.User) (domain.User, *exception.Error)
	FindByIdentifier(ctx context.Context, username, email string) (domain.User, *exception.Error)
}

func NewUserRepository(database *mongo.Database) IUserRepository {
	return &userRepository{
		Collection: database.Collection("users"),
	}
}

type userRepository struct {
	Collection *mongo.Collection
}

func (repo *userRepository) Insert(ctx context.Context, newData domain.User) (res domain.User, err *exception.Error) {

	oneDoc := domain.User{
		Id:             gonanoid.Must(),
		Email:          newData.Email,
		Password:       newData.Password,
		FavoritePhrase: newData.FavoritePhrase,
	}

	data, dataErr := repo.Collection.InsertOne(ctx, oneDoc)
	if dataErr != nil {
		return res, &exception.Error{
			Code: exception.IntenalError,
			Err:  dataErr,
		}
	}

	newId := data.InsertedID
	newData.Id = newId.(string)

	return newData, nil
}

func (repo *userRepository) FindByIdentifier(ctx context.Context, username, email string) (res domain.User, err *exception.Error) {
	c := repo.Collection.FindOne(ctx, bson.M{"email": email, "username": username})

	c.Decode(&res)

	return
}

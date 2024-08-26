package domain

type User struct {
	Id             string `json:"id" bson:"_id"`
	Username       string `json:"username" bson:"username"`
	Fullname       string `json:"fullname" bson:"fullname"`
	Email          string `json:"email" bson:"email"`
	Password       string `json:"password" bson:"password"`
	FavoritePhrase string `json:"favorite_phrase" bson:"favorite_phrase"`
}

type (
	RegisterRequest struct {
		Username       string `json:"username"`
		Fullname       string `json:"fullname" validate:"required"`
		Email          string `json:"email" validate:"required,email"`
		Password       string `json:"password" validate:"required"`
		FavoritePhrase string `json:"favorite_phrase" validate:"required"`
	}

	RegisterResponse struct {
		Id    string `json:"id"`
		Email string `json:"email"`
	}

	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	LoginResponse struct {
		Email string `json:"email"`
		Token string `json:"token"`
	}
)

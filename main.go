package main

import (
	"github.com/ubaidillahhf/go-clarch/app/infra/config"
	"github.com/ubaidillahhf/go-clarch/app/infra/repository"
	"github.com/ubaidillahhf/go-clarch/app/infra/router"
	"github.com/ubaidillahhf/go-clarch/app/usecases"
	_ "github.com/ubaidillahhf/go-clarch/docs"
)

// @title			Go Clarch Boilerplate
// @version			1.0
// @description		This is a golang clean architecture boilerplate using fiber
// @contact.name	Ubaidillah Hakim Fadly
// @contact.url		https://ubed.dev
// @contact.email	ubai@codespace.id
// @license.name	MIT License
// @license.url		https://github.com/satria-research/go-clarch?tab=MIT-1-ov-file
// @host			localhost
// @BasePath		/v1
//
//	@schemes                     http https
//	@securityDefinitions.apiKey  JWT
//	@in                          header
//	@name                        Authorization
//	@description                 JWT security accessToken. Please add it in the format "Bearer {AccessToken}" to authorize your requests.
func main() {
	// load config
	configuration := config.New(".env")

	// error monitoring
	config.SentryInit(configuration)

	// conn mongo
	database := config.NewMongoDatabase(configuration)

	// Setup Repository
	productRepository := repository.NewProductRepository(database)
	userRepository := repository.NewUserRepository(database)

	// Setup Service
	useCase := usecases.NewAppUseCase(
		productRepository,
		userRepository,
	)

	router.Init(useCase, configuration)
}

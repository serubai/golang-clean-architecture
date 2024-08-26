package config

import (
	"context"
	"strconv"
	"time"

	"github.com/ubaidillahhf/go-clarch/app/infra/utility/constants"
	logx "github.com/ubaidillahhf/go-clarch/app/infra/utility/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoDatabase(configuration IConfig) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	mongoPoolMin, mpmErr := strconv.Atoi(configuration.Get("MONGO_POOL_MIN"))
	if mpmErr != nil {
		panic("mongoPoolMin unknown")
	}

	mongoPoolMax, poolMaxErr := strconv.Atoi(configuration.Get("MONGO_POOL_MAX"))
	if poolMaxErr != nil {
		panic("poolMaxErr unknown")
	}

	mongoMaxIdleTime, maxIdleTimeErr := strconv.Atoi(configuration.Get("MONGO_MAX_IDLE_TIME_SECOND"))
	if maxIdleTimeErr != nil {
		panic("maxIdleTimeErr unknown")
	}

	option := options.Client().
		ApplyURI(configuration.Get("MONGO_URI")).
		SetMinPoolSize(uint64(mongoPoolMin)).
		SetMaxPoolSize(uint64(mongoPoolMax)).
		SetMaxConnIdleTime(time.Duration(mongoMaxIdleTime) * time.Second)

	client, clientErr := mongo.NewClient(option)
	if clientErr != nil {
		panic("Failed to connect to database!")
	}

	if err := client.Connect(ctx); err != nil {
		panic("Failed to connect to database!")
	}

	if err := client.Ping(ctx, readpref.Nearest()); err != nil {
		logx.Create().Error().Msg("Connect Mongo DB Failed")
		panic(constants.FAILED_CONNECT_DB)
	}

	logx.Create().Info().Msg("Connected to MongoDB success")

	return client.Database(configuration.Get("MONGO_DATABASE"))
}

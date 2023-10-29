package db

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/iimeta/iim-api/internal/config"
	"github.com/iimeta/iim-api/utility/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	client          *mongo.Client
	DefaultDatabase string
)

func init() {

	ctx := gctx.New()
	var err error

	uri, err := config.Get(ctx, "mongodb.uri")
	if err != nil {
		logger.Error(ctx, err)
	}

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri.String()))
	if err != nil {
		panic(err)
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		panic(fmt.Sprint("MongoDB", err))
	}

	logger.Info(ctx, "MongoDB Successfully connected and pinged.")

	database, err := config.Get(ctx, "mongodb.database")
	if err != nil {
		logger.Error(ctx, err)
	}

	DefaultDatabase = database.String()
}

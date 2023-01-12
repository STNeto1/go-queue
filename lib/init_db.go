package lib

import (
	"_models/ent"
	"context"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func InitDB(logger *zap.Logger) *ent.Client {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=queue password=postgres sslmode=disable")

	if err != nil {
		logger.Fatal("failed opening connection to postgres", zap.Error(err))
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		logger.Fatal("failed creating schema resources", zap.Error(err))
	}

	return client
}

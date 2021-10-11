package share

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	MongoRead, MongoWrite *mongo.Client
	Ctx                   *context.Context
}

func NewRepository(mongoRead, mongoWrite *mongo.Client, ctx *context.Context) *Repository {
	return &Repository{
		MongoRead:  mongoRead,
		MongoWrite: mongoWrite,
		Ctx:        ctx,
	}
}

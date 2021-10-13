package share

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	MongoRead, MongoWrite *mongo.Database
	Ctx                   *context.Context
}

func NewRepository(mongoRead, mongoWrite *mongo.Database, ctx *context.Context) *Repository {
	return &Repository{
		MongoRead:  mongoRead,
		MongoWrite: mongoWrite,
		Ctx:        ctx,
	}
}

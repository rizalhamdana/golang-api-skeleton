package main

import (
	"context"
	"fmt"

	"github.com/rizalhamdana/golang-api-skeleton/config"

	"github.com/rizalhamdana/golang-api-skeleton/modules/share"

	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	DBRead       *mongo.Client
	DBWrite      *mongo.Client
	MongoContext context.Context
}

func MakeHandler() *Service {

	ctx := context.Background()

	dbRead, err := config.MongoRead(&ctx)
	if err != nil {
		fmt.Printf("Failed to connect to DB, msg: %s \n", err.Error())
		return nil
	}

	dbWrite, err := config.MongoWrite(&ctx)
	if err != nil {
		fmt.Printf("Failed to connect to DB, msg: %s \n", err.Error())
		return nil
	}

	_ = share.NewRepository(dbRead, dbWrite, &ctx)

	return &Service{
		DBRead:  dbRead,
		DBWrite: dbWrite,
	}
}

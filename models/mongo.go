package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Find() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		Username: "admin",
		Password: "root",
	}).ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println("连接失败")
	}
	Collection := client.Database("admin").Collection("user")

	res, _ := Collection.InsertOne(ctx, bson.M{"name": "zs", "age": "27"})
	id := res.InsertedID
	log.Println(id)
}

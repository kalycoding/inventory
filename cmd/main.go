package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kalycoding/inventory/datastore"
	"github.com/kalycoding/inventory/handler"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var storage *datastore.InventoryMongoStorage

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	serverAddr := "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000"
	//client, err := datastore.CreateMongoClient(ctx, serverAddr)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(serverAddr))
	if err != nil {
		log.Fatalln("Error connecting to MongoDB")
	}
	storage = datastore.NewInventoryMongoStorage(client)

	paymentHandler := handler.NewInventoryHandler(storage)
	r := gin.Default()

	r.POST("api/inventory", paymentHandler.CreateInventory)
	r.POST("api/category", paymentHandler.CreateCategory)
	r.GET("api/category", paymentHandler.GetAllCategories)
	r.Run()
}

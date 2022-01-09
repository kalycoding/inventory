package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kalycoding/inventory/datastore"
	"github.com/kalycoding/inventory/handler"
)

var storage *datastore.InventoryMongoStorage

func main() {
	serverAddr := os.Getenv("MONGO_URI")
	client, err := datastore.CreateMongoClient(serverAddr)
	if err != nil {
		log.Fatalln("Error connecting to MongoDB")
	}
	storage = datastore.NewInventoryMongoStorage(client)

	paymentHandler := handler.NewInventoryHandler(storage)
	r := gin.Default()

	r.POST("api/inventory", paymentHandler.CreateInventory)
}

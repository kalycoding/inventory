package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kalycoding/inventory/datastore"
	"github.com/kalycoding/inventory/handler"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var storage *datastore.InventoryMongoStorage
var port = os.Getenv("PORT")

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

	/* CORS Allow All Origin to be true */
	r.Use(cors.New(CorsMiddleware()))

	// Category Endpoint
	r.POST("api/category", paymentHandler.CreateCategory)
	r.GET("api/category", paymentHandler.GetAllCategories)
	r.GET("api/category/:catId", paymentHandler.GetCategory)
	r.DELETE("api/category/:catId", paymentHandler.DeleteCategory)

	// Product Endpoint
	r.POST("api/product", paymentHandler.CreateProduct)
	r.GET("api/product", paymentHandler.GetAllProduct)
	r.GET("api/product/:prodId", paymentHandler.GetProduct)
	r.DELETE("api/product/:prodId", paymentHandler.DeleteProduct)

	// Inventory Endpoint
	r.POST("api/inventory", paymentHandler.CreateInventory)
	r.GET("api/inventory", paymentHandler.GetAllInventory)
	r.GET("api/inventory/:invId", paymentHandler.GetInventory)
	r.DELETE("api/inventory/:invId", paymentHandler.DeleteInventory)
	r.PUT("api/inventory/:invId", paymentHandler.UpdateInventory)

	//Export product to CSV Endpoint
	r.GET("api/inventory/export/csv", paymentHandler.ExportInventoryToCSV)

	port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r.Run(":" + port)
}

/* CorsMiddleware set Allow All Origin to true */
func CorsMiddleware() cors.Config {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	return config
}

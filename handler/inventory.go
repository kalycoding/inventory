package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kalycoding/inventory/datastore"
	"github.com/kalycoding/inventory/inventory"
)

type InventoryHandler struct {
	inventory *datastore.InventoryMongoStorage
}

func NewInventoryHandler(i *datastore.InventoryMongoStorage) *InventoryHandler {
	return &InventoryHandler{
		inventory: i,
	}
}

func (i InventoryHandler) CreateInventory(c *gin.Context) {
	inv := &inventory.Inventory{}
	if err := c.BindJSON(&inv); err != nil {
		panic(err)
		return
	}
	invResult, inv, err := i.inventory.CreateInventory(inv)
	if err != nil {
		log.Fatalf("Error")
	}
	c.JSON(201, invResult)
}

package handler

import (
	//"log"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/kalycoding/inventory/datastore"
	"github.com/kalycoding/inventory/inventory"
	//"github.com/kalycoding/inventory/inventory"
)

type InventoryHandler struct {
	inventory *datastore.InventoryMongoStorage
}

func NewInventoryHandler(i *datastore.InventoryMongoStorage) *InventoryHandler {
	return &InventoryHandler{
		inventory: i,
	}
}

func (i InventoryHandler) CreateCategory(c *gin.Context) {
	cat := &inventory.Category{}
	if err := c.BindJSON(&cat); err != nil {
		log.Fatalf("Error %v", err)
	}
	_, _, err := i.inventory.CreateCategory(c, cat)
	if err != nil {
		c.JSON(400, gin.H{"error": "supplier name muist "})
		return
	}

	c.JSON(201, cat)
}

func (i InventoryHandler) GetAllCategories(c *gin.Context) {
	cursor, err := i.inventory.GetAllCategories(c)
	if err != nil {
		c.JSON(404, gin.H{"error": "supplier name muist "})
		return
	}

	var cat []inventory.Category
	if err = cursor.All(c, &cat); err != nil {
		panic(err)
	}

	c.JSON(200, cat)
}

func (i InventoryHandler) CreateInventory(c *gin.Context) {
	inv := &inventory.Inventory{}
	if err := c.BindJSON(&inv); err != nil {
		log.Fatalf("Error %v", err)
	}
	// invResult, cat, err := i.inventory.CreateCategory(cat)
	// if err != nil {
	// 	c.JSON(400, err)
	// }
	c.JSON(201, inv)
}

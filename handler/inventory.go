package handler

import (
	//"log"

	"encoding/csv"
	"fmt"
	"log"
	"os"

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

// Category Handler
func (i InventoryHandler) CreateCategory(c *gin.Context) {
	cat := &inventory.Category{}
	if err := c.BindJSON(&cat); err != nil {
		log.Fatalf("Error %v", err)
	}
	_, cat, err := i.inventory.CreateCategory(c, cat)
	if err != nil {
		c.JSON(200, gin.H{"error": "Category Already Exist"})
		return
	}
	c.JSON(201, cat)
}

func (i InventoryHandler) GetAllCategories(c *gin.Context) {
	cursor, err := i.inventory.GetAllCategories(c)
	if err != nil {
		c.JSON(404, gin.H{"error": "supplier name must "})
		return
	}

	var cat []inventory.Category
	if err = cursor.All(c, &cat); err != nil {
		panic(err)
	}

	c.JSON(200, cat)
}

func (i InventoryHandler) GetCategory(c *gin.Context) {
	catId := c.Param("catId")
	cursor, err := i.inventory.GetCategory(c, catId)
	if err != nil {
		c.JSON(404, gin.H{"error": "not found "})
		return
	}

	var cat inventory.Category
	if err = cursor.Decode(&cat); err != nil {
		panic(err)
	}

	c.JSON(200, cat)
}

func (i InventoryHandler) DeleteCategory(c *gin.Context) {
	catId := c.Param("catId")
	cursor, err := i.inventory.DeleteCategory(c, catId)
	if err != nil {
		c.JSON(404, gin.H{"error": "not found "})
		return
	}

	var cat inventory.Category
	if err = cursor.Decode(&cat); err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{"delete": "done", "data": cat})
}

// End Category Handler

//Product Handler
func (i InventoryHandler) CreateProduct(c *gin.Context) {
	category := &inventory.Product{}
	catId := c.Query("catId")
	if catId == "" {
		c.JSON(400, gin.H{"error": "Supplier category object"})
	}
	if err := c.BindJSON(&category); err != nil {
		log.Fatalf("Error %v", err)
	}
	_, pro, err := i.inventory.CreateProduct(c, category, catId)
	if err != nil {
		c.JSON(200, gin.H{"error": "Product Already Exist"})
		return
	}
	c.JSON(201, pro)
}

func (i InventoryHandler) GetProduct(c *gin.Context) {
	prodId := c.Param("prodId")
	cursor, err := i.inventory.GetProduct(c, prodId)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	var cat inventory.Category
	if err = cursor.Decode(&cat); err != nil {
		panic(err)
	}

	c.JSON(200, cat)
}

func (i InventoryHandler) GetAllProduct(c *gin.Context) {
	cursor, err := i.inventory.GetAllProduct(c)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	var prod []inventory.Product
	if err = cursor.All(c, &prod); err != nil {
		panic(err)
	}

	c.JSON(200, prod)
}

func (i InventoryHandler) DeleteProduct(c *gin.Context) {
	prodId := c.Param("prodId")
	cursor, err := i.inventory.DeleteProduct(c, prodId)
	if err != nil {
		c.JSON(404, gin.H{"error": "not found "})
		return
	}

	var cat inventory.Category
	if err = cursor.Decode(&cat); err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{"delete": "done", "data": cat})
}

//End Product Handler

//Inventory Handler
func (i InventoryHandler) CreateInventory(c *gin.Context) {
	inventory := &inventory.Inventory{}
	prodId := c.Query("prodId")
	if err := c.BindJSON(&inventory); err != nil {
		log.Fatalf("Error %v", err)
	}
	_, pro, err := i.inventory.CreateInventory(c, inventory, prodId)
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(201, pro)
}

func (i InventoryHandler) GetAllInventory(c *gin.Context) {
	cursor, err := i.inventory.GetAllInventory(c)
	if err != nil {
		c.JSON(404, gin.H{"error": "supplier name must "})
		return
	}

	var prod []inventory.Inventory

	if err = cursor.All(c, &prod); err != nil {
		panic(err)
	}
	for _, v := range prod {
		fmt.Println(v.Product.Name)
	}
	c.JSON(200, prod)
}

func (i InventoryHandler) GetInventory(c *gin.Context) {
	invId := c.Param("invId")
	cursor, err := i.inventory.GetInventory(c, invId)
	if err != nil {
		c.JSON(404, gin.H{"error": "not found "})
		return
	}

	var inv inventory.Inventory
	if err = cursor.Decode(&inv); err != nil {
		panic(err)
	}

	c.JSON(200, inv)
}

func (i InventoryHandler) DeleteInventory(c *gin.Context) {
	invId := c.Param("invId")
	cursor, err := i.inventory.DeleteInventory(c, invId)
	if err != nil {
		c.JSON(404, gin.H{"error": "not found "})
		return
	}

	var inv inventory.Inventory
	if err = cursor.Decode(&inv); err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{"delete": "done", "data": inv})
}

func (i InventoryHandler) UpdateInventory(c *gin.Context) {
	invId := c.Param("invId")
	inv := &inventory.Inventory{}
	if err := c.BindJSON(&inv); err != nil {
		log.Fatalf("Error %v", err)
	}

	_, res, err := i.inventory.UpdateInventory(c, inv, invId)
	if err != nil {
		c.JSON(400, gin.H{"Error": "failed updating inventory"})
		return
	}
	c.JSON(200, res)
}

func (i InventoryHandler) ExportInventoryToCSV(c *gin.Context) {
	cursor, err := i.inventory.GetAllInventory(c)
	if err != nil {
		c.JSON(404, gin.H{"error": "supplier name must "})
		return
	}

	var prod []inventory.Inventory

	if err = cursor.All(c, &prod); err != nil {
		panic(err)
	}

	records := [][]string{
		{"product", "Quantity", "Cost Price",
			"Stock Level", "Selling Price", "Supplier",
			"Product Category"},
	}
	for _, v := range prod {
		records = append(records, []string{
			v.Product.Name,
			fmt.Sprint(v.Quantity),
			fmt.Sprint(v.CostPrice),
			fmt.Sprint(v.StockLevel),
			fmt.Sprint(v.SellingPrice),
			v.Supplier,
			v.Product.Category.Name,
		})
	}
	fmt.Println(records)
	f, err := os.Create("inventory.csv")
	defer f.Close()

	if err != nil {

		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
	c.JSON(200, gin.H{"status": "inventory csv exported"})
}

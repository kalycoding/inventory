package inventory

import (
	"errors"
	"time"
)

type Category struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Kind        string    `json:"kind"`
	CreatedAt   time.Time `json:"createdAt" data:"readonly"`
	UpdatedAt   time.Time `json:"updatedAt" data:"readonly"`
	IsDeleted   bool      `json:"isDeleted"`
}

type Product struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Name        string    `json:"name"`
	Category    Category  `json:"category"`
	Description string    `json:"description"`
	Kind        string    `json:"kind"`
	CreatedAt   time.Time `json:"createdAt" data:"readonly"`
	UpdatedAt   time.Time `json:"updatedAt" data:"readonly"`
	IsDeleted   bool      `json:"isDeleted"`
}

type Inventory struct {
	ID           string    `json:"id" bson:"_id,omitempty"`
	Product      Product   `json:"product"`
	Kind         string    `json:"kind"`
	CostPrice    float64   `json:"costPrice"`
	Quantity     int       `json:"quantity"`
	StockLevel   int       `json:"stockLevel"`
	SellingPrice float64   `json:"sellingPrice"`
	Supplier     string    `json:"supplier"`
	CreatedAt    time.Time `json:"createdAt" data:"readonly"`
	UpdatedAt    time.Time `json:"updatedAt" data:"readonly"`
	IsDeleted    bool      `json:"isDeleted"`
}
type TimeStamp struct {
	CreatedAt time.Time `json:"createdAt" data:"readonly"`
	UpdatedAt time.Time `json:"updatedAt" data:"readonly"`
	IsDeleted bool      `json:"isDeleted"`
}

func (inventory *Inventory) Validate() error {
	if inventory.CostPrice < 1 || inventory.Quantity < 1 || inventory.SellingPrice < 1 {
		return errors.New("Product is Empty")
	}
	return nil
}

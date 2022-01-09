package inventory

import (
	"errors"
	"time"
)

type Category struct {
	ID          string `json:"id" data:"readonly"`
	Name        string `json:"name"`
	Description string `json:"description"`
	TimeStamp
}

type Product struct {
	ID          string `json:"id" data:"readonly"`
	Name        string `json:"name"`
	Description string `json:"description"`
	TimeStamp
}

type Inventory struct {
	Product      Product `json:"product"`
	CostPrice    float64 `json:"costPrice"`
	Quantity     int     `json:"quantity"`
	RestockLevel int     `json:"restockLevel"`
	SellingPrice float64 `json:"sellingPrice"`
	Supplier     string  `json:"supplier"`
	TimeStamp
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

package datastore

import (
	"context"
	"errors"

	"github.com/kalycoding/inventory/inventory"
	"go.mongodb.org/mongo-driver/mongo"
)

type InventoryMongoStorage struct {
	*mongo.Client
}

const (
	INVENTORYDATABASE   = "inventory"
	INVENTORYCOLLECTION = "inventory"
)

func NewInventoryMongoStorage(store *mongo.Client) *InventoryMongoStorage {
	return &InventoryMongoStorage{
		store,
	}
}

func (e *InventoryMongoStorage) InventoryCollection() *mongo.Collection {
	return e.Client.Database(INVENTORYDATABASE).Collection(INVENTORYCOLLECTION)
}

func (i *InventoryMongoStorage,
) CreateInventory(payload *inventory.Inventory,
) (*mongo.InsertOneResult, *inventory.Inventory, error) {
	if payload.Supplier == "" {
		return nil, nil, errors.New("supplier must be supplier")
	}
	ref := i.InventoryCollection()
	result, err := ref.InsertOne(context.Background(), payload)
	if err != nil {
		return nil, nil, errors.New(("unable to create Inventory"))
	}
	return result, payload, nil
}

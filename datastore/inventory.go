package datastore

import (
	"context"
	"errors"
	"time"

	"github.com/kalycoding/inventory/inventory"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type InventoryMongoStorage struct {
	*mongo.Client
}

const (
	INVENTORYDATABASE   = "inventory"
	INVENTORYCOLLECTION = "inventoryItems"
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
	// if payload.Supplier == "" {
	// 	return nil, nil, errors.New("supplier must be supplier")
	// }
	ref := i.InventoryCollection()
	result, err := ref.InsertOne(context.Background(), payload)
	if err != nil {
		return nil, nil, errors.New(("unable to create Inventory"))
	}
	return result, payload, nil
}

func (i *InventoryMongoStorage,
) CreateCategory(ctx context.Context, payload *inventory.Category,
) (*mongo.InsertOneResult, *inventory.Category, error) {

	if payload.Name == "" {
		return nil, nil, errors.New("supplier must be supplier")
	}
	//fmt.Println(payload)
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	ref := i.InventoryCollection()
	payload.CreatedAt = time.Now()
	res, err := ref.InsertOne(ctx, payload)

	if err != nil {
		return nil, nil, errors.New(("unable to create Inventory"))
	}
	//fmt.Println(result)
	return res, payload, nil
}

func (i *InventoryMongoStorage,
) GetAllCategories(ctx context.Context,
) (*mongo.Cursor, error) {

	filter := bson.D{{}}
	//fmt.Println(payload)
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	ref := i.InventoryCollection()
	res, err := ref.Find(ctx, filter)

	if err != nil {
		return nil, errors.New(("unable to fetch categories"))
	}
	//fmt.Println(result)
	return res, nil
}

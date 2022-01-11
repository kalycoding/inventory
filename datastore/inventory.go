package datastore

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/kalycoding/inventory/inventory"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

//Inventory CRUD
func (i *InventoryMongoStorage,
) CreateProduct(ctx context.Context, payload *inventory.Product, catId string,
) (*mongo.InsertOneResult, *inventory.Product, error) {
	// if payload.Supplier == "" {
	// 	return nil, nil, errors.New("supplier must be supplier")
	// }
	cat := inventory.Category{}
	ref := i.InventoryCollection()
	objectId, err := primitive.ObjectIDFromHex(catId)
	if err != nil {
		log.Println("Invalid id")
	}

	filter := bson.D{{"_id", objectId}}
	err = ref.FindOne(ctx, filter).Decode(&cat)

	if err != nil {
		return nil, nil, errors.New("category don't exist")
	}
	proCheck := inventory.Product{}
	prodFilter := bson.D{{"name", payload.Name}}
	_ = ref.FindOne(ctx, prodFilter).Decode(&proCheck)
	if proCheck.Name != "" {
		return nil, nil, errors.New("product already exist")
	}

	payload.Category = cat
	payload.CreatedAt = time.Now()
	payload.Kind = "product"
	result, err := ref.InsertOne(ctx, payload)
	if err != nil {
		return nil, nil, errors.New(("unable to create Inventory"))
	}
	return result, payload, nil
}

func (i *InventoryMongoStorage,
) GetProduct(ctx context.Context, prodId string,
) (*mongo.SingleResult, error) {

	// convert id string to ObjectId
	objectId, err := primitive.ObjectIDFromHex(prodId)
	if err != nil {
		log.Println("Invalid id")
	}

	filter := bson.D{{"_id", objectId}}

	ref := i.InventoryCollection()
	res := ref.FindOne(ctx, filter)
	var cat inventory.Category
	err = ref.FindOne(ctx, filter).Decode(&cat)
	if err != nil {
		return nil, errors.New("not Found")
	}
	return res, nil
}

func (i *InventoryMongoStorage,
) GetAllProduct(ctx context.Context,
) (*mongo.Cursor, error) {

	filter := bson.D{{"kind", "product"}}
	//fmt.Println(payload)
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	ref := i.InventoryCollection()
	res, err := ref.Find(ctx, filter)

	if err != nil {
		return nil, errors.New(("unable to fetch product"))
	}
	//fmt.Println(result)
	return res, nil
}

func (i *InventoryMongoStorage,
) DeleteProduct(ctx context.Context, catId string,
) (*mongo.SingleResult, error) {

	objectId, err := primitive.ObjectIDFromHex(catId)
	if err != nil {
		log.Println("Invalid id")
	}

	filter := bson.D{{"_id", objectId}}

	ref := i.InventoryCollection()
	res := ref.FindOne(ctx, filter)
	var cat inventory.Category
	err = ref.FindOneAndDelete(ctx, filter).Decode(&cat)
	if err != nil {
		return nil, errors.New("not Found")
	}
	return res, nil
}

// Category CRUD
func (i *InventoryMongoStorage,
) CreateCategory(ctx context.Context, payload *inventory.Category,
) (*mongo.InsertOneResult, *inventory.Category, error) {

	//fmt.Println(payload)
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	ref := i.InventoryCollection()
	catCheck := inventory.Category{}
	prodFilter := bson.D{{"name", payload.Name}}
	_ = ref.FindOne(ctx, prodFilter).Decode(&catCheck)
	if catCheck.Name != "" {
		return nil, nil, errors.New("category already exist")
	}
	payload.Kind = "category"
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

	filter := bson.D{{"kind", "category"}}
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

func (i *InventoryMongoStorage,
) GetCategory(ctx context.Context, catId string,
) (*mongo.SingleResult, error) {

	// convert id string to ObjectId
	objectId, err := primitive.ObjectIDFromHex(catId)
	if err != nil {
		log.Println("Invalid id")
	}

	filter := bson.D{{"_id", objectId}}

	ref := i.InventoryCollection()
	res := ref.FindOne(ctx, filter)
	var cat inventory.Category
	err = ref.FindOne(ctx, filter).Decode(&cat)
	if err != nil {
		return nil, errors.New("not Found")
	}
	return res, nil
}

func (i *InventoryMongoStorage,
) DeleteCategory(ctx context.Context, catId string,
) (*mongo.SingleResult, error) {

	objectId, err := primitive.ObjectIDFromHex(catId)
	if err != nil {
		log.Println("Invalid id")
	}

	filter := bson.D{{"_id", objectId}}

	ref := i.InventoryCollection()
	res := ref.FindOne(ctx, filter)
	var cat inventory.Category
	err = ref.FindOneAndDelete(ctx, filter).Decode(&cat)
	if err != nil {
		return nil, errors.New("not Found")
	}
	return res, nil
}

// End Category CRUD

// Inventory Datastore
func (i *InventoryMongoStorage,
) CreateInventory(ctx context.Context, payload *inventory.Inventory, productId string,
) (*mongo.InsertOneResult, *inventory.Inventory, error) {
	objectId, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		log.Println("Invalid id")
	}
	prod := inventory.Product{}
	filter := bson.D{{"_id", objectId}}
	ref := i.InventoryCollection()
	err = ref.FindOne(ctx, filter).Decode(&prod)
	if err != nil {
		return nil, nil, errors.New("product not found")
	}
	payload.Product = prod
	payload.CreatedAt = time.Now()
	payload.Kind = "inventory"
	res, err := ref.InsertOne(ctx, payload)
	if err != nil {
		return nil, nil, errors.New("failed Creating inventory")
	}
	return res, payload, nil
}

func (i *InventoryMongoStorage,
) GetAllInventory(ctx context.Context,
) (*mongo.Cursor, error) {

	filter := bson.D{{"kind", "inventory"}}
	//fmt.Println(payload)
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	ref := i.InventoryCollection()
	res, err := ref.Find(ctx, filter)

	if err != nil {
		return nil, errors.New(("unable to fetch inventory"))
	}
	//fmt.Println(result)
	return res, nil
}

func (i *InventoryMongoStorage,
) GetInventory(ctx context.Context, invId string,
) (*mongo.SingleResult, error) {

	// convert id string to ObjectId
	objectId, err := primitive.ObjectIDFromHex(invId)
	if err != nil {
		log.Println("Invalid id")
	}

	filter := bson.D{{"_id", objectId}}

	ref := i.InventoryCollection()
	res := ref.FindOne(ctx, filter)
	var inv inventory.Inventory
	err = ref.FindOne(ctx, filter).Decode(&inv)
	if err != nil {
		return nil, errors.New("not Found")
	}
	return res, nil
}

func (i *InventoryMongoStorage,
) DeleteInventory(ctx context.Context, invId string,
) (*mongo.SingleResult, error) {

	objectId, err := primitive.ObjectIDFromHex(invId)
	if err != nil {
		log.Println("Invalid id")
	}

	filter := bson.D{{"_id", objectId}}

	ref := i.InventoryCollection()
	res := ref.FindOne(ctx, filter)
	var inv inventory.Inventory
	err = ref.FindOneAndDelete(ctx, filter).Decode(&inv)
	if err != nil {
		return nil, errors.New("not Found")
	}
	return res, nil
}

func (i *InventoryMongoStorage,
) UpdateInventory(ctx context.Context, payload *inventory.Inventory, invId string,
) (*inventory.Inventory, *mongo.UpdateResult, error) {

	// convert id string to ObjectId
	objectId, err := primitive.ObjectIDFromHex(invId)
	if err != nil {
		log.Println("Invalid id")
	}

	//filter := bson.D{{"_id", objectId}}

	ref := i.InventoryCollection()
	supplier := payload.Supplier
	costPrice := payload.CostPrice
	updatedAt := time.Now()
	quantity := payload.Quantity
	restock := payload.StockLevel
	fmt.Println(restock)
	sellingPrice := payload.SellingPrice
	res, err := ref.UpdateOne(
		ctx,
		bson.M{"_id": objectId},
		bson.D{
			{"$set", bson.D{{"costPrice", costPrice}, {"supplier", supplier}, {"updatedAt", updatedAt},
				{"quantity", quantity}, {"stockLevel", restock}, {"sellingPrice", sellingPrice}}},
		},
	)
	if err != nil {
		return nil, nil, errors.New("not Found")
	}
	return payload, res, nil
}


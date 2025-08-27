package repository

import (
	"context"
	"golang-crud-basic/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository interface {
    Create(product *model.Product) error
    GetAll() ([]model.Product, error)
    GetByID(id string) (*model.Product, error)
    Update(id string, product *model.Product) error
    Delete(id string) error
}

type productMongoRepository struct {
    collection *mongo.Collection
}

func NewProductMongoRepository(db *mongo.Database) ProductRepository {
    return &productMongoRepository{collection: db.Collection("product")}
}


func (r *productMongoRepository) Create(product *model.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if product.ID.IsZero() {
		product.ID = primitive.NewObjectID()
	}
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, product)
	return err
}


func (r *productMongoRepository) GetAll() ([]model.Product, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    cursor, err := r.collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var products []model.Product
    if err := cursor.All(ctx, &products); err != nil {
        return nil, err
    }
    return products, nil
}

func (r *productMongoRepository) GetByID(id string) (*model.Product, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var product model.Product
    err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
    if err != nil {
        return nil, err
    }
    return &product, nil
}

func (r *productMongoRepository) Update(id string, product *model.Product) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    product.UpdatedAt = time.Now()
    _, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": product})
    return err
}

func (r *productMongoRepository) Delete(id string) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
    return err
}

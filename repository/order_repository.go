package repository

import (
	"context"
	"golang-crud-basic/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository interface {
    Create(order *model.Order) error
    GetAll() ([]model.Order, error)
    GetByID(id string) (*model.Order, error)
}

type orderMongoRepository struct {
    collection *mongo.Collection
}

func NewOrderMongoRepository(db *mongo.Database) OrderRepository {
    return &orderMongoRepository{collection: db.Collection("order")}
}

func (r *orderMongoRepository) Create(order *model.Order) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    order.CreatedAt = time.Now()
    order.UpdatedAt = time.Now()

    res, err := r.collection.InsertOne(ctx, order)
    if err != nil {
        return err
    }

    if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
        order.ID = oid 
    }

    return nil
}


func (r *orderMongoRepository) GetAll() ([]model.Order, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    cursor, err := r.collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var orders []model.Order
    if err = cursor.All(ctx, &orders); err != nil {
        return nil, err
    }

    return orders, nil
}

func (r *orderMongoRepository) GetByID(id string) (*model.Order, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }

    var order model.Order
    err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&order)
    if err != nil {
        return nil, err
    }

    return &order, nil
}

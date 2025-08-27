package repository

import (
	"context"
	"errors"
	"time"

	"golang-crud-basic/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdminRepository interface {
	GetAll() ([]model.Admin, error)
	GetByEmail(email string) (*model.Admin, error)
	Create(admin *model.Admin) error
	UpdateByEmail(email string, admin *model.Admin) error
	DeleteByEmail(email string) error
}

type adminMongoRepository struct {
	collection *mongo.Collection
}

func NewAdminMongoRepository(db *mongo.Database) AdminRepository {
	return &adminMongoRepository{collection: db.Collection("admin")}
}

func (r *adminMongoRepository) GetAll() ([]model.Admin, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var admins []model.Admin
	if err = cursor.All(ctx, &admins); err != nil {
		return nil, err
	}
	return admins, nil
}

func (r *adminMongoRepository) GetByEmail(email string) (*model.Admin, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var admin model.Admin
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&admin)
	if err != nil {
		return nil, errors.New("admin not found")
	}
	return &admin, nil
}

func (r *adminMongoRepository) Create(admin *model.Admin) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	admin.CreatedAt = time.Now()
	admin.UpdatedAt = time.Now()

	res, err := r.collection.InsertOne(ctx, admin)
	admin.ID = res.InsertedID.(primitive.ObjectID)
	return err
}

func (r *adminMongoRepository) UpdateByEmail(email string, admin *model.Admin) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	admin.UpdatedAt = time.Now()
	_, err := r.collection.UpdateOne(ctx, bson.M{"email": email}, bson.M{"$set": admin})
	return err
}

func (r *adminMongoRepository) DeleteByEmail(email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := r.collection.DeleteOne(ctx, bson.M{"email": email})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("admin not found")
	}
	return nil
}

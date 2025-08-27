package repository

import (
	"context"
	"golang-crud-basic/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecruiterRepository interface {
	Create(recruiter *model.Recruiter) error
	GetAll() ([]model.Recruiter, error)
	GetByID(id string) (*model.Recruiter, error)
	Update(id string, recruiter *model.Recruiter) error
	Delete(id string) error
}

type recruiterMongoRepository struct {
	collection *mongo.Collection
}

func NewRecruiterMongoRepository(db *mongo.Database) RecruiterRepository {
	return &recruiterMongoRepository{collection: db.Collection("recruiter")}
}

func (r *recruiterMongoRepository) Create(recruiter *model.Recruiter) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if recruiter.ID.IsZero() {
		recruiter.ID = primitive.NewObjectID()
	}
	recruiter.CreatedAt = time.Now()
	recruiter.UpdatedAt = time.Now()

	res, err := r.collection.InsertOne(ctx, recruiter)

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
        recruiter.ID = oid 
    }

	return err
}

func (r *recruiterMongoRepository) GetAll() ([]model.Recruiter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var recruiters []model.Recruiter
	if err := cursor.All(ctx, &recruiters); err != nil {
		return nil, err
	}

	return recruiters, nil
}

func (r *recruiterMongoRepository) GetByID(id string) (*model.Recruiter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var recruiter model.Recruiter
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&recruiter)
	if err != nil {
		return nil, err
	}

	return &recruiter, nil
}

func (r *recruiterMongoRepository) Update(id string, recruiter *model.Recruiter) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	recruiter.UpdatedAt = time.Now()
	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": recruiter})
	return err
}

func (r *recruiterMongoRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

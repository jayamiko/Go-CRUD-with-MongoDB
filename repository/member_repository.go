package repository

import (
	"context"
	"errors"
	"golang-crud-basic/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MemberRepository interface {
	GetAll() ([]model.Member, error)
	GetByRecruiterID(id string) (*model.Member, error)
	Create(member *model.Member) error
	UpdateByRecruiter(id string, member *model.Member) error
	Delete(id string) error

	ExistsByEmail(email string) (bool, error)
}

type memberMongoRepository struct {
	collection *mongo.Collection
}

func NewMemberMongoRepository(db *mongo.Database) MemberRepository {
	return &memberMongoRepository{
		collection: db.Collection("member"),
	}
}

func (r *memberMongoRepository) GetAll() ([]model.Member, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var members []model.Member
	if err = cursor.All(ctx, &members); err != nil {
		return nil, err
	}

	return members, nil
}

func (r *memberMongoRepository) GetByRecruiterID(recruiterID string) (*model.Member, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	recruiterObjID, err := primitive.ObjectIDFromHex(recruiterID)
	if err != nil {
		return nil, err
	}

	var member model.Member
	err = r.collection.FindOne(ctx, bson.M{"recruiterId": recruiterObjID}).Decode(&member)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, model.ErrMemberNotFound
		}
		return nil, err
	}

	return &member, nil
}

func (r *memberMongoRepository) Create(member *model.Member) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := r.collection.InsertOne(ctx, member)
	if err != nil {
		return err
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		member.ID = oid 
	}

	return nil
}


func (r *memberMongoRepository) UpdateByRecruiter(recruiterID string, member *model.Member) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	recruiterObjID, err := primitive.ObjectIDFromHex(recruiterID)
	if err != nil {
		return err
	}

	member.UpdatedAt = time.Now()

	filter := bson.M{
		"recruiterId": recruiterObjID,
	}

	update := bson.M{
		"$set": bson.M{
			"statusAktivasi": member.StatusAktivasi,
			"email":          member.Email,
			"password":       member.Password,
			"updatedAt":      member.UpdatedAt,
		},
	}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("member not found with this recruiterId")
	}

	return nil
}


func (r *memberMongoRepository) Delete(recruiterID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	recruiterObjID, err := primitive.ObjectIDFromHex(recruiterID)
	if err != nil {
		return err
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"recruiterId": recruiterObjID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return model.ErrMemberNotFound
	}

	return nil
}

func (r *memberMongoRepository) ExistsByEmail(email string) (bool, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    count, err := r.collection.CountDocuments(ctx, bson.M{"email": email})
    if err != nil {
        return false, err
    }

    return count > 0, nil
}


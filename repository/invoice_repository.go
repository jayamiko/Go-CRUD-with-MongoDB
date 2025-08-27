package repository

import (
	"context"
	"time"

	"golang-crud-basic/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type InvoiceRepository interface {
	Create(invoice *model.Invoice) error
	GetAll() ([]model.Invoice, error)
	GetByID(id string) (*model.Invoice, error)
	Delete(id string) error
}

type invoiceMongoRepository struct {
	collection *mongo.Collection
}

func NewInvoiceMongoRepository(db *mongo.Database) InvoiceRepository {
	return &invoiceMongoRepository{collection: db.Collection("invoice")}
}

func (r *invoiceMongoRepository) Create(invoice *model.Invoice) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if invoice.ID.IsZero() {
		invoice.ID = primitive.NewObjectID()
	}
	now := time.Now()
	invoice.CreatedAt = now
	invoice.UpdatedAt = now

	_, err := r.collection.InsertOne(ctx, invoice)
	return err
}

func (r *invoiceMongoRepository) GetAll() ([]model.Invoice, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var invoices []model.Invoice
	if err = cur.All(ctx, &invoices); err != nil {
		return nil, err
	}
	return invoices, nil
}

func (r *invoiceMongoRepository) GetByID(id string) (*model.Invoice, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var invoice model.Invoice
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&invoice)
	if err != nil {
		return nil, err
	}
	return &invoice, nil
}

func (r *invoiceMongoRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

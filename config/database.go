package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB() *mongo.Database {
	dbName := os.Getenv("DATABASE_NAME") 
	if dbName == "" {
		dbName = "db_test" 
	}
	mongoURI := os.Getenv("MONGO_URI") 
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("❌ Failed to create MongoDB client:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		log.Fatal("❌ Failed to connect to MongoDB:", err)
	}

	return client.Database(dbName)
}

package main

import (
	"context"
	"fmt"
	"golang-crud-basic/presenter"
	"golang-crud-basic/repository"
	"golang-crud-basic/usecase"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	r := mux.NewRouter()
	
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017" 
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("db_test")

	memberCollection := db.Collection("member")
	adminCollection := db.Collection("admin")
	
	authHandler := &presenter.AuthHandler{
		MemberCollection: memberCollection,
		AdminCollection:  adminCollection,
	}
	
	memberRepo := repository.NewMemberMongoRepository(db)
	memberUC := usecase.NewMemberUsecase(memberRepo)
	
	adminRepo := repository.NewAdminMongoRepository(db)
	adminUC := usecase.NewAdminUsecase(adminRepo)

	productRepo := repository.NewProductMongoRepository(db)
	productUC := usecase.NewProductUsecase(productRepo)
	
	orderRepo := repository.NewOrderMongoRepository(db)
	orderUC := usecase.NewOrderUsecase(orderRepo)

	invoiceRepo := repository.NewInvoiceMongoRepository(db)
	invoiceUC := usecase.NewInvoiceUsecase(invoiceRepo)

	recruiterRepo := repository.NewRecruiterMongoRepository(db)
	recruiterUC := usecase.NewRecruiterUsecase(recruiterRepo)

	presenter.NewAuthHandler(r, authHandler)
	presenter.NewMemberHandler(r, memberUC)
	presenter.NewAdminHandler(r, adminUC)
	presenter.NewProductHandler(r, productUC)
	presenter.NewOrderHandler(r, orderUC)
	presenter.NewRecruiterHandler(r, recruiterUC)
	presenter.NewInvoiceHandler(r, invoiceUC)
	
	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

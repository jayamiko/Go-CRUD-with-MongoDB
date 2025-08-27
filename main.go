package main

import (
	"fmt"
	"golang-crud-basic/config"
	"golang-crud-basic/presenter"
	"golang-crud-basic/repository"
	"golang-crud-basic/routes"
	"golang-crud-basic/usecase"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	
	db := config.ConnectMongoDB()

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

	// register all routes
	routes.RegisterRoutes(r, memberUC, adminUC, productUC, orderUC, recruiterUC, invoiceUC, authHandler)
	
	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

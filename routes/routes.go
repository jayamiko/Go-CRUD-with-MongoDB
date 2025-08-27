package routes

import (
	"golang-crud-basic/presenter"
	"golang-crud-basic/usecase"

	"github.com/gorilla/mux"
)

func RegisterRoutes(
	router *mux.Router,
	memberUC usecase.MemberUsecase,
	adminUC usecase.AdminUsecase,
	productUC usecase.ProductUsecase,
	orderUC usecase.OrderUsecase,
	recruiterUC usecase.RecruiterUsecase,
	invoiceUC usecase.InvoiceUsecase,
	authHandler *presenter.AuthHandler,
) {
	// AUTH
	router.HandleFunc("/auth/login", authHandler.Login).Methods("POST")

	// MEMBER
	memberHandler := &presenter.MemberHandler{Usecase: memberUC}
	router.HandleFunc("/members", memberHandler.GetAll).Methods("GET")
	router.HandleFunc("/members/{recruiterId}", memberHandler.GetByRecruiter).Methods("GET")
	router.HandleFunc("/members", memberHandler.Create).Methods("POST")
	router.HandleFunc("/members/{recruiterId}", memberHandler.UpdateByRecruiter).Methods("PUT")
	router.HandleFunc("/members/{recruiterId}", memberHandler.DeleteByRecruiter).Methods("DELETE")

	// ADMIN
	adminHandler := &presenter.AdminHandler{Usecase: adminUC}
	router.HandleFunc("/admins", adminHandler.GetAll).Methods("GET")
	router.HandleFunc("/admins/{email}", adminHandler.GetByEmail).Methods("GET")
	router.HandleFunc("/admins", adminHandler.Create).Methods("POST")
	router.HandleFunc("/admins/{email}", adminHandler.UpdateByEmail).Methods("PUT")
	router.HandleFunc("/admins/{email}", adminHandler.DeleteByEmail).Methods("DELETE")

	// PRODUCT
	productHandler := &presenter.ProductHandler{Usecase: productUC}
	router.HandleFunc("/products", productHandler.GetAll).Methods("GET")
	router.HandleFunc("/products/{id}", productHandler.GetByID).Methods("GET")
	router.HandleFunc("/products", productHandler.Create).Methods("POST")
	router.HandleFunc("/products/{id}", productHandler.Update).Methods("PUT")
	router.HandleFunc("/products/{id}", productHandler.Delete).Methods("DELETE")

	// ORDER
	orderHandler := &presenter.OrderHandler{Usecase: orderUC}
	router.HandleFunc("/orders", orderHandler.GetAll).Methods("GET")
	router.HandleFunc("/orders/{id}", orderHandler.GetByID).Methods("GET")
	router.HandleFunc("/orders", orderHandler.Create).Methods("POST")

	// RECRUITER
	recruiterHandler := &presenter.RecruiterHandler{Usecase: recruiterUC}
	router.HandleFunc("/recruiters", recruiterHandler.GetAll).Methods("GET")
	router.HandleFunc("/recruiters/{id}", recruiterHandler.GetByID).Methods("GET")
	router.HandleFunc("/recruiters", recruiterHandler.Create).Methods("POST")
	router.HandleFunc("/recruiters/{id}", recruiterHandler.Update).Methods("PUT")
	router.HandleFunc("/recruiters/{id}", recruiterHandler.Delete).Methods("DELETE")

	// INVOICE
	invoiceHandler := &presenter.InvoiceHandler{Usecase: invoiceUC}
	router.HandleFunc("/invoices", invoiceHandler.GetAll).Methods("GET")
	router.HandleFunc("/invoices/{id}", invoiceHandler.GetByID).Methods("GET")
	router.HandleFunc("/invoices", invoiceHandler.Create).Methods("POST")
	router.HandleFunc("/invoices/{id}", invoiceHandler.Delete).Methods("DELETE")
}

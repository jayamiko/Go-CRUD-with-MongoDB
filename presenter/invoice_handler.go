package presenter

import (
	"encoding/json"
	"net/http"

	"golang-crud-basic/model"
	"golang-crud-basic/usecase"

	"github.com/gorilla/mux"
)

type InvoiceHandler struct {
	Usecase usecase.InvoiceUsecase
}

func NewInvoiceHandler(router *mux.Router, uc usecase.InvoiceUsecase) {
	handler := &InvoiceHandler{Usecase: uc}

	router.HandleFunc("/invoices", handler.Create).Methods("POST")
	router.HandleFunc("/invoices", handler.GetAll).Methods("GET")
	router.HandleFunc("/invoices/{id}", handler.GetByID).Methods("GET")
	router.HandleFunc("/invoices/{id}", handler.Delete).Methods("DELETE")
}

func (h *InvoiceHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var invoice model.Invoice
	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Usecase.Create(&invoice); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(invoice)
}

func (h *InvoiceHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	invoices, err := h.Usecase.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(invoices)
}

func (h *InvoiceHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	invoice, err := h.Usecase.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(invoice)
}

func (h *InvoiceHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	if err := h.Usecase.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

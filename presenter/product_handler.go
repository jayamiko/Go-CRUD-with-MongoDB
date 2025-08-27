package presenter

import (
	"encoding/json"
	"golang-crud-basic/model"
	"golang-crud-basic/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

type ProductHandler struct {
    Usecase usecase.ProductUsecase
}

func NewProductHandler(router *mux.Router, uc usecase.ProductUsecase) {
    handler := &ProductHandler{Usecase: uc}

    router.HandleFunc("/products", handler.GetAll).Methods("GET")
    router.HandleFunc("/products/{id}", handler.GetByID).Methods("GET")
    router.HandleFunc("/products", handler.Create).Methods("POST")
    router.HandleFunc("/products/{id}", handler.Update).Methods("PUT")
    router.HandleFunc("/products/{id}", handler.Delete).Methods("DELETE")
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    products, err := h.Usecase.GetAll()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    id := mux.Vars(r)["id"]
    product, err := h.Usecase.GetByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var product model.Product

    if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := h.Usecase.Create(&product); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(product) 
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    id := mux.Vars(r)["id"]
    var product model.Product
    if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := h.Usecase.Update(id, &product); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    if err := h.Usecase.Delete(id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

package presenter

import (
	"encoding/json"
	"golang-crud-basic/model"
	"golang-crud-basic/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

type OrderHandler struct {
    Usecase usecase.OrderUsecase
}

func (h *OrderHandler) GetAll(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    orders, err := h.Usecase.GetAll()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(orders)
}

func (h *OrderHandler) GetByID(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    id := mux.Vars(r)["id"]

    order, err := h.Usecase.GetByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(order)
}

func (h *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var order model.Order
    if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := h.Usecase.Create(&order); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(order)
}

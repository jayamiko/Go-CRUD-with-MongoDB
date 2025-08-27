package presenter

import (
	"encoding/json"
	"golang-crud-basic/model"
	"golang-crud-basic/usecase"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type AdminHandler struct {
	Usecase usecase.AdminUsecase
}

func (h *AdminHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	admins, _ := h.Usecase.GetAll()
	json.NewEncoder(w).Encode(admins)
}

func (h *AdminHandler) GetByEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	email := mux.Vars(r)["email"]
	admin, err := h.Usecase.GetByEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(admin)
}

func (h *AdminHandler) Create(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var admin model.Admin
    if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if admin.Password == "" {
        http.Error(w, "password is required", http.StatusBadRequest)
        return
    }

    hashed, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "failed to hash password", http.StatusInternalServerError)
        return
    }
    admin.Password = string(hashed) 

    if err := h.Usecase.Create(&admin); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    admin.Password = "" 
    json.NewEncoder(w).Encode(admin)
}

func (h *AdminHandler) UpdateByEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	email := mux.Vars(r)["email"]
	var admin model.Admin
	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.Usecase.UpdateByEmail(email, &admin)
	json.NewEncoder(w).Encode(admin)
}

func (h *AdminHandler) DeleteByEmail(w http.ResponseWriter, r *http.Request) {
	email := mux.Vars(r)["email"]
	err := h.Usecase.DeleteByEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

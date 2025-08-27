package presenter

import (
	"encoding/json"
	"golang-crud-basic/model"
	"golang-crud-basic/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

type RecruiterHandler struct {
	Usecase usecase.RecruiterUsecase
}

func (h *RecruiterHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var recruiter model.Recruiter
	if err := json.NewDecoder(r.Body).Decode(&recruiter); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Usecase.Create(&recruiter); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(recruiter)
}

func (h *RecruiterHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	recruiters, err := h.Usecase.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(recruiters)
}

func (h *RecruiterHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	recruiter, err := h.Usecase.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(recruiter)
}

func (h *RecruiterHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	var recruiter model.Recruiter
	if err := json.NewDecoder(r.Body).Decode(&recruiter); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Usecase.Update(id, &recruiter); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(recruiter)
}

func (h *RecruiterHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	if err := h.Usecase.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "deleted"})
}

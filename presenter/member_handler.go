package presenter

import (
	"encoding/json"
	"errors"
	"golang-crud-basic/model"
	"golang-crud-basic/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

type MemberHandler struct {
	Usecase usecase.MemberUsecase
}


func NewMemberHandler(router *mux.Router, uc usecase.MemberUsecase) {
	handler := &MemberHandler{Usecase: uc}

	// CRUD routes
	router.HandleFunc("/members", handler.GetAll).Methods("GET")
	router.HandleFunc("/members/{recruiterId}", handler.GetByRecruiter).Methods("GET")
	router.HandleFunc("/members", handler.Create).Methods("POST")
	router.HandleFunc("/members/{recruiterId}", handler.UpdateByRecruiter).Methods("PUT")
	router.HandleFunc("/members/{recruiterId}", handler.DeleteByRecruiter).Methods("DELETE")
}

func removePassword(member *model.Member) *model.Member {
	if member != nil {
		member.Password = ""
	}
	return member
}

func removePasswordSlice(members []model.Member) []model.Member {
	for i := range members {
		members[i].Password = ""
	}
	return members
}

func (h *MemberHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	members, err := h.Usecase.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(removePasswordSlice(members))
}

func (h *MemberHandler) GetByRecruiter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	recruiterID := mux.Vars(r)["recruiterId"]

	member, err := h.Usecase.GetByRecruiterID(recruiterID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(removePassword(member))
}

func (h *MemberHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var member model.Member

	if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Usecase.Create(&member); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(removePassword(&member))
}

func (h *MemberHandler) UpdateByRecruiter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	recruiterID := mux.Vars(r)["recruiterId"]

	var member model.Member
	if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Usecase.UpdateByRecruiter(recruiterID, &member); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedMember, err := h.Usecase.GetByRecruiterID(recruiterID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(removePassword(updatedMember))
}

// DELETE member by recruiterId
func (h *MemberHandler) DeleteByRecruiter(w http.ResponseWriter, r *http.Request) {
	recruiterID := mux.Vars(r)["recruiterId"]

	if err := h.Usecase.Delete(recruiterID); err != nil {
		if errors.Is(err, model.ErrMemberNotFound) { // undefined: model.ErrMemberNotFound
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

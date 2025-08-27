package presenter

import (
	"context"
	"encoding/json"
	"golang-crud-basic/model"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	MemberCollection *mongo.Collection
	AdminCollection  *mongo.Collection
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET")) 

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var member model.Member
	err := h.MemberCollection.FindOne(ctx, bson.M{"email": req.Email}).Decode(&member)
	if err == nil {
		if bcrypt.CompareHashAndPassword([]byte(member.Password), []byte(req.Password)) != nil {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}

		claims := jwt.MapClaims{
			"userType":    "member",
			"recruiterId": member.RecruiterID.Hex(),
			"email":       member.Email,
			"exp":         time.Now().Add(24 * time.Hour).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenStr, err := token.SignedString(jwtSecret)
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}

		member.Password = ""

		resp := map[string]interface{}{
			"user": map[string]interface{}{
				"_id":           member.ID.Hex(),
				"recruiterId":   member.RecruiterID.Hex(),
				"statusAktivasi": member.StatusAktivasi,
				"email":         member.Email,
				"createdAt":     member.CreatedAt,
				"updatedAt":     member.UpdatedAt,
			},
			"token": tokenStr,
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	var admin model.Admin
	err = h.AdminCollection.FindOne(ctx, bson.M{"email": req.Email}).Decode(&admin)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password)) != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	claims := jwt.MapClaims{
		"userType": "admin",
		"email":    admin.Email,
		"role":     admin.Role,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtSecret)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	admin.Password = ""

	resp := map[string]interface{}{
		"user": map[string]interface{}{
			"_id":       admin.ID.Hex(),
			"email":     admin.Email,
			"role":      admin.Role,
			"createdAt": admin.CreatedAt,
			"updatedAt": admin.UpdatedAt,
		},
		"token": tokenStr,
	}

	json.NewEncoder(w).Encode(resp)
}

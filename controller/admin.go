package controller

import (
	"context"
	"encoding/json"
	"fmt"

	"log"

	"net/http"

	"github.com/Kdsingh333/HousewareHQ/database"
	"github.com/Kdsingh333/HousewareHQ/middleware"
	"github.com/Kdsingh333/HousewareHQ/models"
	"github.com/dgrijalva/jwt-go"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Collection
var ctx = context.TODO()

func init() {
	db = database.Setup()
}

// ---------------------User Add by admin---------------------

func AdminAdd(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("Token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value

	claims := &models.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if claims.Role != "admin" {
		json.NewEncoder(w).Encode(map[string]string{"error": "You don't have permission of admin level"})
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Error in reading body"})
		return
	}

	var result bson.M
	ok := db.FindOne(ctx, bson.D{{
		Key:   "username",
		Value: user.Username,
	}}).Decode(&result)

	if ok != mongo.ErrNoDocuments {
		json.NewEncoder(w).Encode(map[string]string{
			"error":  "This username is already present",
			"error2": " Error occurred while querying the database",
		})
		return
	}

	user.Password, err = middleware.GeneratehashPassword(user.Password)
	if err != nil {
		log.Fatalln("error in password hash")
	}

	insertResult, err := db.InsertOne(ctx, &user)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Error occure during inserting data",
		})
		return
	}
	fmt.Println("Inserted user with ID:", insertResult.InsertedID)

}

// -------------------------------User Delete by Admin----------

func AdminDelete(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("Token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value

	claims := &models.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if claims.Role != "admin" {
		json.NewEncoder(w).Encode(map[string]string{"error": "You don't have permission of admin level"})
		return
	}

	var username models.Name
	if err := json.NewDecoder(r.Body).Decode(&username); err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Error in reading body"})
		return
	}
	result, err := db.DeleteOne(ctx, bson.M{"username": username.Username})
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Error occurred while deleting the user"})
		return
	}

	if result.DeletedCount == 0 {
		json.NewEncoder(w).Encode(map[string]string{"error": "User not found"})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}

// / -------------------------getting all user list--------------------
func GetUser(w http.ResponseWriter, r *http.Request) {
	var allUser []models.Name

	result, err := db.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Error occurred while Finding All the user", http.StatusInternalServerError)
		return
	}

	for result.Next(ctx) {
		var usr models.Name
		if err := result.Decode(&usr); err != nil {
			http.Error(w, "Error occurred while decoding user document", http.StatusInternalServerError)
			return
		}
		allUser = append(allUser, usr)

	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(allUser)

}

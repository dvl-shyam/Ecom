package main

import (
	// "fmt"
	"context"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var templates *template.Template

func getCollection() *mongo.Collection {
	client, err := ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB in getCollection: %v", err)
		return nil
	}
	return client.Database("testdb").Collection("people")
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title":   "Welcome",
		"Message": "Welcome To this website",
	}
	templates.ExecuteTemplate(w, "base.html", data)
}

func Login(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "login.html", nil)
}

func Products(w http.ResponseWriter, r *http.Request) {
	collection := client.Database("htmlTemplate").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if r.Method == http.MethodPost {
		var input struct {
			Name  string `json:"name"`
			Price int    `json:"price"`
		}

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		if input.Name == "" || input.Price == 0 {
			http.Error(w, "Product name and price are required", http.StatusBadRequest)
			return
		}

		_, err := collection.InsertOne(ctx, map[string]interface{}{
			"name":  input.Name,
			"price": input.Price,
		})
		if err != nil {
			http.Error(w, "Failed to save product", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "Product created successfully"})
		return
	}

	cursor, err := collection.Find(ctx, map[string]interface{}{})
	if err != nil {
		http.Error(w, "Database query failed", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var products []map[string]interface{}
	for cursor.Next(ctx) {
		var product map[string]interface{}
		cursor.Decode(&product)
		products = append(products, product)
	}

	if r.Header.Get("Accept") == "application/json" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
		return
	}

	data := map[string]interface{}{
		"Title":    "Products",
		"Products": products,
	}
	templates.ExecuteTemplate(w, "base.html", data)
}


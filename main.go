package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	// "os"
)

func main() {

	templates = template.Must(template.ParseGlob("templates/*.html"))
	type Product struct {
		Name  string `bson:"name"`
		Price int    `bson:"price"`
	}

	// check := func(err error) {
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	client, err := ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v", err)
	}
	defer client.Disconnect(nil)

	// h1 := func(w http.ResponseWriter, r *http.Request) {

		// templat := template.Must(template.ParseFiles("index.html"))
		// templat.Execute(w ,nil)

		// t, err := template.New("shyam kuntal").Parse(`{{define "T" }}Hello,{{.}} {{end}}`)
		// check(err)
		// err = t.ExecuteTemplate(w, "T", "<script>alert('you have been pwned')</script>")

		// t, err := template.New("example").Parse("<h1>{{.Title}}</h1>")
		// check(err)
		// data := map[string]string{"Title": "Hello, World!"}
		// err = t.Execute(w, data)

	// }

	// http.HandleFunc("/", h1)
	http.HandleFunc("GET /", Home)
	http.HandleFunc("GET /login", Login)
	http.HandleFunc("GET /products", Products)
	http.HandleFunc("POST /products", Products)


	// PORT := os.Getenv("PORT")
	fmt.Printf("Server is running on Port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

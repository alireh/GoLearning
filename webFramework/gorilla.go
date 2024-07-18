package webFramework

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"learning/webFramework/handlers"
)

//installation commands
// go get -u github.com/gorilla/mux
// go get -u github.com/lib/pq

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func TestGorilla() {
	fmt.Println("----------------------------------------Start Gorilla----------------------------------------")
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "postgres")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/users", handlers.GetUsers(db)).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.GetUser(db)).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUser(db)).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.UpdateUser(db)).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.DeleteUser(db)).Methods("DELETE")

	log.Println("Server started at :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
	fmt.Println("----------------------------------------End Gorilla  ----------------------------------------")
}

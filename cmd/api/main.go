package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"time"
)

func main() {
	// 1. Load Environment Variables
	_ = godotenv.Load()

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	dbPass := os.Getenv("DB_PASSWORD")

	// 2. Connect to Database
	connStr := fmt.Sprintf("postgres://pet_admin:%s@%s:5432/shelter_db?sslmode=disable", dbPass, dbHost)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Could not connect to DB: ", err)
	}
	fmt.Printf("Connected to Database on %s!\n", dbHost)

	// 3. Initialize Application Struct
	// This injects the database into our handlers
	app := &Application{
		DB: db,
	}

	// 4. Start Server using the Routes defined in routes.go
	port := "8080"
	fmt.Println("Server starting on port " + port + "...")

	// app.routes() comes from the routes.go file!
	err = http.ListenAndServe(":"+port, app.routes())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Worker started...")

	// TODO: Phase 4 - Add loop to process background jobs
	for {
		fmt.Println("Checking for jobs...")
		time.Sleep(1 * time.Hour)
	}
}

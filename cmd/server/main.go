package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/database"
	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/router"
	"github.com/joho/godotenv"
)


func main() {
	_ = godotenv.Load(".env")
	API_URL := os.Getenv("REACT_APP_API_URL")
	API_PORT := os.Getenv("PORT")

	// Get DB
	db, err := database.GetDB()
	if err != nil {
		log.Fatalf("Database failed: %v", err)
	}
	defer db.Close()

	// Routing Setup
	r := router.Setup(db)
	
	fmt.Printf("Listening at %s", API_URL)
	log.Fatalln(http.ListenAndServe(":" + API_PORT, r))
}

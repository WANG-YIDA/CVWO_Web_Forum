package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/database"
	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/router"
)


func main() {
	API_DOMAIN := os.Getenv("REACT_APP_API_DOMAIN")
    API_PORT := os.Getenv("PORT")

	// Get DB
	db, err := database.GetDB()
	if err != nil {
		log.Fatalf("Database failed: %v", err)
	}
	defer db.Close()

	// Routing Setup
	r := router.Setup(db)
	
	fmt.Printf("Listening on port %s at %s", API_PORT, API_DOMAIN)
	log.Fatalln(http.ListenAndServe(":" + API_PORT, r))
}

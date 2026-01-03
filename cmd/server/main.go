package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/database"
	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/router"
)


func main() {
	// Get DB
	db, err := database.GetDB()
	if err != nil {
		log.Fatalf("Database failed: %v", err)
	}
	defer db.Close()

	// Routing Setup
	r := router.Setup(db)
	
	fmt.Print("Listening on port 8000 at http://localhost:8000")
	log.Fatalln(http.ListenAndServe(":8000", r))
}

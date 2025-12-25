package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/router"
)


func main() {
	r := router.Setup()
	db, err := database.GetDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Conn.Close()

	fmt.Print("Listening on port 8000 at http://localhost:8000")
	log.Fatalln(http.ListenAndServe(":8000", r))
}

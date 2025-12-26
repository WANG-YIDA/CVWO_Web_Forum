package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/database"
	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/router"
)


func main() {
	r := router.Setup()
	db, err := database.GetDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	fmt.Print("Listening on port 8000 at http://localhost:8000")
	log.Fatalln(http.ListenAndServe(":8000", r))
}

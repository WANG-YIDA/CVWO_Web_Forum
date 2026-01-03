package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/router"
)


func main() {
	r, err := router.Setup()
	if err != nil {
		log.Fatalf("server failed: %v", err)
	}
	
	fmt.Print("Listening on port 8000 at http://localhost:8000")
	log.Fatalln(http.ListenAndServe(":8000", r))
}

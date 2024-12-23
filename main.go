package main

import (
	"log"
	"net/http"

	"github.com/munraitoo13/goapi/handlers"
)

const PORT = "6000"

func main() {
	http.HandleFunc("/client/profile", handlers.HandleClientProfile) // handlefunc for endpoints. path, handler

	log.Println("Server running on port:", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil)) // listen and serve starts the server
}

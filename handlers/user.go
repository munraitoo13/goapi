package handlers

import (
	"encoding/json"
	"net/http"

	db "github.com/munraitoo13/goapi/database"
)

// pointer to request and response writer to the response (w write, r read)
func HandleClientProfile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetClient(w, r)
	case http.MethodPatch:
		PatchClient(w, r)
	default:
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
	}
}

func GetClient(w http.ResponseWriter, r *http.Request) {
	var clientId = r.URL.Query().Get("clientId") // gets the url query
	client, ok := db.Clients[clientId]           // gets the client from db

	// chekcs for errors or empty query
	if !ok || clientId == "" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// sets the header type
	w.Header().Set("Content-Type", "application/json")

	// response to be written
	response := db.Client{
		Id:    client.Id,
		Email: client.Email,
		Name:  client.Name,
		Age:   client.Age,
	}

	// encodes the response to
	json.NewEncoder(w).Encode(response)
}

func PatchClient(w http.ResponseWriter, r *http.Request) {
	var clientId = r.URL.Query().Get("clientId") // gets clientId query
	client, ok := db.Clients[clientId]           // gets the client

	// validate client
	if !ok || clientId == "" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// decode payload and treats error
	var payloadData db.Client
	err := json.NewDecoder(r.Body).Decode(&payloadData)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close() // closes the body after reading because it's a stream of data

	// changes the values
	client.Email = payloadData.Email
	client.Age = payloadData.Age
	client.Name = payloadData.Name

	// updates client in db with local client
	db.Clients[clientId] = client

	// writes success status
	w.WriteHeader(http.StatusOK)
}

package handlers

import (
	"encoding/json"
	"net/http"

	db "github.com/munraitoo13/goapi/database"
)

// handle all client methods
func HandleClientProfile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetClient(w, r) // if req is get
	case http.MethodPatch:
		PatchClient(w, r) // if req is patch
	default:
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed) // if req is not the 2 above
	}
}

func GetClient(w http.ResponseWriter, r *http.Request) {
	client := r.Context().Value("client").(db.Client) // gets the client from middleware context

	w.Header().Set("Content-Type", "application/json") // sets the header type

	// response to be written
	response := db.Client{
		Id:    client.Id,
		Email: client.Email,
		Name:  client.Name,
		Age:   client.Age,
	}

	json.NewEncoder(w).Encode(response) // writes the encoded response to client
}

func PatchClient(w http.ResponseWriter, r *http.Request) {
	client := r.Context().Value("client").(db.Client) // gets the client from middleware context

	var payloadData db.Client                           // initialize payload data
	err := json.NewDecoder(r.Body).Decode(&payloadData) // decode the json payload and add err var

	// check for errors
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	defer r.Body.Close() // closes the body stream after read

	// changes the values on local variable
	client.Email = payloadData.Email
	client.Age = payloadData.Age
	client.Name = payloadData.Name

	db.Clients[client.Id] = client // updates client in db with local variable

	w.WriteHeader(http.StatusOK) // writes success status
}

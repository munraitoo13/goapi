package main

import (
	"net/http"
	"regexp"
)

type userHandler struct {
}

var (
	listUsersRe  = regexp.MustCompile(`^\/users[\/]*$`)
	getUserRe    = regexp.MustCompile(`^\/users\/(\d+)$`)
	createUserRe = regexp.MustCompile(`^\/users[\/]*$`)
)

func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && listUsersRe.MatchString(r.URL.Path):
		h.List(w, r)
		return
	case r.Method == http.MethodGet && getUserRe.MatchString(r.URL.Path):
		h.Get(w, r)
		return
	case r.Method == http.MethodPost && createUserRe.MatchString(r.URL.Path):
		h.Create(w, r)
		return
	default:
		notFound(w, r)
		return
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound) // sends the 404 or something else
	w.Write([]byte("not found"))       // write always is a byte
}

func (h *userHandler) List(w http.ResponseWriter, r *http.Request)   {}
func (h *userHandler) Get(w http.ResponseWriter, r *http.Request)    {}
func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {}

func main() {
	mux := http.ServeMux()

	mux.Handle("/users/", &userHandler{}) // first the route then the handler.

	http.ListenAndServe("localhost:8000", mux)
}

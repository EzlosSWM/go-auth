package server

import (
	"encoding/json"
	"form/internal/domain"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var (
	errInvalidMethod = map[string]string{"err": "invalid method"}
)

// RegisterRoutes what do you think this does?
func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handler)

	mux.HandleFunc("/user", s.handleUser)

	return mux
}

// handlerUser handles all user cases
// register & login users
func (s *Server) handleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST": 
		s.registerUser(w,r)
	default: 
		WriteJSON(w, http.StatusMethodNotAllowed, errInvalidMethod)
	}
}

// registerUser creates a new user and stores to db
func (s *Server) registerUser(w http.ResponseWriter, r *http.Request) {
	first := r.FormValue("firstname")
	last := r.FormValue("lastname")
	email := r.FormValue("email")
	pw := r.FormValue("password")
	pwConf := r.FormValue("passwordConf")

	msg, ok := validateRegistration(first,last,email,pw,pwConf)
	if !ok {
		WriteJSON(w, http.StatusBadRequest, msg)
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)

	u := domain.NewUser(first, last, email, password)


	// logic to add to all users array
	// send to db


	jsonResp, err := json.Marshal(u)
	if err != nil {
		log.Fatalf("err handling JSON marshal. Err: %v\n", err)
	}

	_, _ = w.Write(jsonResp)
}

// TODO: all cases
// validateRegistration obviously checks if the user form is correct
func validateRegistration(first, last, email, password, passwordConf string) (map[string]string, bool) {
	// fix this shit haha
	if len([]byte(password)) <= 7 || len([]byte(passwordConf)) <= 7 || password != passwordConf {
		return map[string]string{"err": "password invalid"}, false
	}

	return nil, true
}


// handler serves index
func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, "hey bbg")
}

// WriteJSON sends JSON data 
func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.WriteHeader(status)
	jsonResp, _ := json.Marshal(v)
	_, _ = w.Write(jsonResp)
}

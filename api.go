package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func writeJSON(res http.ResponseWriter, status int, value any) error {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	return json.NewEncoder(res).Encode(value)
}

type ApiError struct {
	Error string
}

type apiFunc func(res http.ResponseWriter, req *http.Request) error

func makeHTTPHandler(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			writeJSON(w, http.StatusInternalServerError, ApiError{err.Error()})
		}
	}
}

type APIServer struct {
	listenAddress string
}

func NewAPIServer(listenAddress string) *APIServer {
	return &APIServer{listenAddress}
}

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Do stuff here
        log.Println(r.RequestURI)
        // Call the next handler, which can be another middleware in the chain, or the final handler.
        next.ServeHTTP(w, r)
    })
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()

	router.HandleFunc("/accounts", makeHTTPHandler(s.handleAccounts))
	router.HandleFunc("/accounts/{id}", makeHTTPHandler(s.handleGetAccount));
	router.Use(mux.CORSMethodMiddleware(router))
	router.Use(loggingMiddleware)

	log.Println("Starting server on", s.listenAddress)

	return http.ListenAndServe(s.listenAddress, router)
}

func (s *APIServer) handleAccounts(res http.ResponseWriter, req *http.Request) error {
	// Handle accounts logic here
	if req.Method == "GET" {
		fmt.Println("GET")
		return s.handleGetAccount(res, req)
	}

	if req.Method == "POST" {
		return s.handleCreateAccount(res, req)
	}

	if req.Method == "DELETE" {
		return s.handleDeleteAccount(res, req)
	}

	return fmt.Errorf("Unsupported method %s", req.Method)
}

func (s *APIServer) handleGetAccount(res http.ResponseWriter, req *http.Request) error {
	// account := NewAccount("John", "Doe")
	
	id := mux.Vars(req)["id"]
	fmt.Println("ID:", id)
	
	return writeJSON(res, http.StatusOK, &Account{})
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

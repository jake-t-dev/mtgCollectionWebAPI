package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type APIServer struct {
	listenAddr string
}
type APIFunc func(http.ResponseWriter, *http.Request) error

type APIError struct {
	Error string `json:"error"`
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{listenAddr: listenAddr}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/home", makeHTTPHandleFunc(s.handleReq)) //sample handler

	log.Printf("JSON API server running on port: %s", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func makeHTTPHandleFunc(handler APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			// error handling
			WriteJSON(w, http.StatusBadRequest, APIError{err.Error()})
		}
	}
}

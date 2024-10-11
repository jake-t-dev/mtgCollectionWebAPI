package main

import (
	"fmt"
	"net/http"
)

func (s *APIServer) handleReq(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetReq(w, r)
	case "POST":
		return s.handlePostReq(w, r)
	case "DELETE":
		return s.handleDeleteReq(w, r)
	case "PUT":
		return s.handlePutReq(w, r)
	}
	return fmt.Errorf("unsupported method: %s", r.Method)
}

func (s *APIServer) handlePostReq(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIServer) handleGetReq(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIServer) handlePutReq(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIServer) handleDeleteReq(w http.ResponseWriter, r *http.Request) error {
	return nil
}

package main

import (
	"log"
	"net/http"
)

type ApplicationErrors struct {
}

func (appEr *ApplicationErrors) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error- %s path: %s error: %s", r.Method, r.URL.Path, err)
	writeJSONError(w, http.StatusInternalServerError, "the server encountered a problem")
}

func (appEr *ApplicationErrors) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request error- %s path: %s error: %s", r.Method, r.URL.Path, err)
	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (appEr *ApplicationErrors) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("resource not found error- %s path: %s error: %s", r.Method, r.URL.Path, err)
	writeJSONError(w, http.StatusNotFound, "resource not found")
}

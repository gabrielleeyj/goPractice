package controllers

import (
	"encoding/json"
	"library.com/library/services"
	"net/http"
)

type Controller interface {
	List(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	service services.Service
}

func (c *controller) List(w http.ResponseWriter, r *http.Request) {
	books, err := c.service.ListAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	body, err := json.Marshal(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func NewController(s services.Service) Controller {
	return &controller{
		service: s,
	}
}

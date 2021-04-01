package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"library.com/library/services"
	"net/http"
)

type Controller interface {
	List(w http.ResponseWriter, r *http.Request)
	Reserve(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	service services.Service
}

func (c *controller) Reserve(w http.ResponseWriter, r *http.Request) {
	isbn := chi.URLParam(r, "isbn")

	book, err := c.service.Reserve(isbn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(book)
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

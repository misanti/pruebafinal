package controllers

import (
	"encoding/json"
	"net/http"
	"pruebafinal/models"
	"pruebafinal/usecases/buyer"
	"regexp"
	"strconv"
)

type buyerController struct {
	buyerIDPattern *regexp.Regexp
}

func (bc buyerController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/buyers" {
		switch r.Method {
		case http.MethodGet:
			bc.GetAll(w, r)
		case http.MethodPost:
			bc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := bc.buyerIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		switch r.Method {
		case http.MethodGet:
			bc.Get(id, w)
		case http.MethodDelete:
			bc.delete(id, w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func (bc *buyerController) GetAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(buyer.GetAll(), w)
}
func (bc *buyerController) Get(id int, w http.ResponseWriter) {
	u, err := buyer.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(u, w)
}
func (bc *buyerController) post(w http.ResponseWriter, r *http.Request) {
	u, err := bc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse Buyer object"))
	}
	u, err = buyer.Save(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(u, w)
}
func (bc *buyerController) delete(id int, w http.ResponseWriter) {
	err := buyer.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (bc *buyerController) parseRequest(r *http.Request) (models.Buyer, error) {
	dec := json.NewDecoder(r.Body)
	var u models.Buyer
	err := dec.Decode(&u)
	if err != nil {
		return models.Buyer{}, err
	}
	return u, nil
}

func newBuyerController() *buyerController {
	return &buyerController{
		buyerIDPattern: regexp.MustCompile(`/buyers/(\d+)/?`),
	}
}

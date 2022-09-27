package handlers

import (
	"encoding/json"
	"github.com/WenLopes/recent/usecases/history"
	"log"
	"net/http"
)

type AllHandler struct {
	repo history.Repository
}

func NewAllHandler(repository history.Repository) *AllHandler {
	handler := &AllHandler{
		repo: repository,
	}
	return handler
}

func (handler *AllHandler) GetAll() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		all := handler.repo.GetAll()
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(200)

		if all != nil {
			if err := json.NewEncoder(writer).Encode(all); err != nil {
				log.Fatal(err)
			}
		}
	}
}

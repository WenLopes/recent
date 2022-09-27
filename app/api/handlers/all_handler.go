package handlers

import (
	"encoding/json"
	"github.com/WenLopes/recent/usecases/users_history"
	"log"
	"net/http"
)

type AllHandler struct {
	usersHistoryService users_history.UsersHistory
}

func NewAllHandler(usersHistoryService users_history.UsersHistory) *AllHandler {
	handler := &AllHandler{
		usersHistoryService: usersHistoryService,
	}
	return handler
}

func (handler *AllHandler) GetAll() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		all := handler.usersHistoryService.GetByUserId("1")
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(200)

		if err := json.NewEncoder(writer).Encode(all); err != nil {
			log.Fatal(err)
		}
	}
}

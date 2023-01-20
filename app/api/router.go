package api

import "github.com/gorilla/mux"

//TODO: Utilizar mecanismo de rotas
func (s *Server) Router(r *mux.Router) {
	r.HandleFunc("/v1/all/pix_key", s.handlers.AllHandler.GetAll()).Methods("GET").Name("all")
}

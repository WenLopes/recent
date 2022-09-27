package api

type Server struct {
	handlers *Handlers
}

func NewServer(handlers *Handlers) *Server {
	server := &Server{}
	server.handlers = handlers
	return server
}

//TODO: Criar o server usando essa função
func WithHandlers(handlers *Handlers) func(server *Server) error {
	return func(server *Server) error {
		server.handlers = handlers
		return nil
	}
}

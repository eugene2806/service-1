package builder

import (
	"github.com/gorilla/mux"
	"net/http"
	"service-1/internal/transport/server"
)

type HandlerBuilder struct {
	server *server.Server
}

func NewHandlerBuilder(server *server.Server) *HandlerBuilder {
	return &HandlerBuilder{
		server: server,
	}
}
func (h *HandlerBuilder) BuildHandler() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/ping", h.server.GetPing).Methods(http.MethodGet)

	return router
}

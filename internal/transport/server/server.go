package server

import (
	"log"
	"net/http"
	"service-1/internal/repository"
)

type Server struct {
	ProjectRepository *repository.ProjectRepository
}

func BuildServer(rep *repository.ProjectRepository) *Server {
	return &Server{
		ProjectRepository: rep,
	}
}

func (s *Server) GetPing(W http.ResponseWriter, r *http.Request) {
	log.Println("SUPER PING")
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("content-type", "application/json")
}

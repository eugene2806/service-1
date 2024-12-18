package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"service-1/internal/builder"
	"service-1/internal/repository"
	"service-1/internal/transport/server"
	"service-1/storage"
)

func restCmd(stor *storage.Storage) *cobra.Command {

	return &cobra.Command{
		Use:   "rest",
		Short: "A brief description of your command",
		Run: func(cmd *cobra.Command, args []string) {
			reposit := repository.NewProjectRepository(stor)

			serv := server.BuildServer(reposit)

			handler := builder.NewHandlerBuilder(serv).BuildHandler()

			restServer := builder.BuildRestServer("50051", handler)

			log.Println("rest server start...")

			log.Fatal(restServer.ListenAndServe())
		},
	}
}

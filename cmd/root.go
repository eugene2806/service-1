package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"service-1/internal/config"
	"service-1/storage"
)

const migratePath = "file://migrate"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "service-1",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	cfg := config.NewStorageConfig()
	stor := storage.NewStorage(cfg)
	stor.Open()
	defer stor.Close()

	migrate := migrateCmd()

	migrate.AddCommand(migrateUpCmd(stor.GetMigrationURL(), migratePath))
	migrate.AddCommand(migrateDownCmd(stor.GetMigrationURL(), migratePath))

	rootCmd.AddCommand(restCmd(stor))
	rootCmd.AddCommand(migrate)

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {

}

package cmd

import (
	"example/service/cmd/database"
	"example/service/cmd/services"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "Example Services",
		Short: "Example",
		Long:  "Example - Backend Services",
	}
)

func Execute() {
	//Register command
	rootCmd.AddCommand(database.MigrationCommand)
	rootCmd.AddCommand(services.StartCmd())

	services.StartCmd().Flags().StringP("config", "c", "config/file", "Config dir i.e. config/file")

	if err := rootCmd.Execute(); err != nil {
		log.Fatalln("Error: \n", err.Error())
		os.Exit(-1)
	}
}
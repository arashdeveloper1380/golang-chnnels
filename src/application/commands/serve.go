package commands

import (
	"fiber-gorm-channel-ecommerce/src/pkg/boot"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve app on dev server",
	Long:  "Application will be served on host and port defined in config.yml file",
	Run: func(cmd *cobra.Command, args []string) {
		serveServer()
	},
}

func serveServer() {
	boot.Serve()
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

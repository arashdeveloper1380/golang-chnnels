package commands

import (
	"fiber-gorm-channel-ecommerce/src/pkg/boot"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Table migration",
	Long:  "Table migration",
	Run: func(cmd *cobra.Command, args []string) {
		migrate()
	},
}

func migrate() {
	boot.Migrate()
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

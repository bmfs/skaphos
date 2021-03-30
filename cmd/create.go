package cmd

import (
	"log"

	"github.com/bmfs/skaphos/pkg/creator"
	"github.com/spf13/cobra"
)

// createCmd scaffolds a new project based on a template
var createCmd = &cobra.Command{
	Use:   "create <source> [<destination>]",
	Short: "Creates a new project from a template",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		source := args[0]
		destination := "."
		if len(args) >= 2 {
			destination = args[1]
		}
		err := creator.Create(source, destination)
		if err != nil {
			log.Fatalf("failed to create: %s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

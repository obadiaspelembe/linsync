package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "linsync",
	Short: "linsync is a tool for syncronizing linode object storage files with your disk folder",
	Long:  `linsync is a tool for syncronizing linode object storage files with your disk folder.`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO sync functionality
		fmt.Println("Hello from linsync")
	},
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

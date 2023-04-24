package cmd

import (
	"linsync/helper"

	"github.com/spf13/cobra"
)

var pushCommand = &cobra.Command{
	Use: "push",
	Args:  cobra.ExactArgs(1),
	Short: "Push file from local dir to bucket [filename]",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {

			var filename = args[0]

			helper.PushFile(filename)
		}

	},
}
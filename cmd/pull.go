package cmd

import (
	"linsync/helper"

	"github.com/spf13/cobra"
)

var pullCommand = &cobra.Command{
	Use: "pull",
	Args:  cobra.ExactArgs(1),
	Short: "Pull file from linode bucket to dir [filename]",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {

			var filename = args[0]

			helper.PullFile(filename)
		}

	},
}
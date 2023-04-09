package cmd

import (
	"fmt"
	"os"

	"linsync/helper"

	"github.com/spf13/cobra"
)

const SHORT_DESCRIPTION = "linsync is a tool for syncronizing linode object storage files with your disk folder."

var rootCommand = &cobra.Command{
	Use:   "linsync",
	Short: SHORT_DESCRIPTION,
	Long:  SHORT_DESCRIPTION,

	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 2 {
			var subCommands = args[0]
			var filename = args[1]

			switch subCommands {
			case "push":
				helper.PushFile(filename)
				break
			case "pull":
				helper.PullFile(filename)
				break
			default:
				fmt.Println("unknown command")
			}

			return
		}

		fmt.Println(SHORT_DESCRIPTION)

	},
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

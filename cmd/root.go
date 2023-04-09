package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"linsync/helper"
)

var rootCommand = &cobra.Command{
	Use:   "linsync",
	Short: "linsync is a tool for syncronizing linode object storage files with your disk folder",
	Long:  `linsync is a tool for syncronizing linode object storage files with your disk folder.`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO sync functionality
		// push filename
		// pull filename
		fmt.Println(args[0]);

		var subCommands = args[0];
		var filename = args[1];

		switch subCommands {
		case "push":
			helper.PushFile(filename)
		    break;
		case "pull":
			fmt.Println("commands: pull" )
			break;
		default:
			fmt.Println("unknown command")
		}
	},
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

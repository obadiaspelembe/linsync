package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

const SHORT_DESCRIPTION = "linsync is a tool for syncronizing linode object storage files with your disk folder."

var version = "0.0.1-alpha"
var rootCommand = &cobra.Command{
	Use:     "linsync",
	Version: version,
	Short:   SHORT_DESCRIPTION,
	Long:    SHORT_DESCRIPTION,
}

func Execute() { 
	rootCommand.AddCommand(pullCommand)
	rootCommand.AddCommand(pushCommand) 
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

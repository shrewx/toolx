package main

import (
	"fmt"
	"github.com/shrewx/toolx/cmd"
	"github.com/shrewx/toolx/cmd/gen"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "toolx",
}

func init() {
	rootCmd.AddCommand(gen.CmdGen)
	rootCmd.AddCommand(cmd.Swagger())
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

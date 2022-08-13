package main

import (
	"fmt"
	"github.com/shrewx/toolx/gen"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "tools",
}

func init() {
	rootCmd.AddCommand(gen.CmdGen)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

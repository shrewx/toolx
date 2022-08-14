package gen

import "github.com/spf13/cobra"

var CmdGen = &cobra.Command{
	Use:   "gen",
	Short: "generate file command",
}

func init() {
	CmdGen.AddCommand(statusErrorCommand())
}

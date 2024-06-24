package gen

import (
	"github.com/go-courier/packagesx"
	"github.com/shrewx/ginx/pkg/statuserror"
	"github.com/spf13/cobra"
	"os"
)

func statusErrorCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "error",
		Short: "generate status error file",
		Run: func(cmd *cobra.Command, args []string) {
			pwd, _ := os.Getwd()
			pkg, err := packagesx.Load(pwd)
			if err != nil {
				panic(err)
			}

			g := statuserror.NewStatusErrorGenerator(pkg)
			g.Scan(args...)
			g.Output(pwd)
		},
	}
}

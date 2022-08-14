package main

import (
	"github.com/shrewx/ginx"
	"github.com/shrewx/toolx/__example__/router"
	"github.com/spf13/cobra"
)

var Config = struct {
	Server *ginx.Server `yaml:"server"`
}{}

func main() {

	ginx.Execute(func(cmd *cobra.Command, args []string) {

		ginx.Conf(&Config)

		ginx.RunServer(Config.Server, router.V0Router)
	})

}

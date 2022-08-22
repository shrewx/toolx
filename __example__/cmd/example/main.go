package main

import (
	"github.com/shrewx/ginx"
	"github.com/shrewx/ginx/pkg/conf/server"
	"github.com/shrewx/ginx/pkg/service_discovery"
	"github.com/shrewx/toolx/__example__/router"
	"github.com/spf13/cobra"
)

var Config = struct {
	Server *server.Config `yaml:"server"`
}{}

func main() {

	ginx.Launch(func(cmd *cobra.Command, args []string) {

		ginx.Parse(&Config)

		ginx.Register(service_discovery.NewConsul())

		ginx.RunServer(Config.Server, router.V0Router)
	})

}

package commands

import (
	"github.com/spf13/cobra"

	"github.com/pdxjohnny/discovery/discovery"
	"github.com/pdxjohnny/discovery/frontend"
	"github.com/pdxjohnny/discovery/key"
	"github.com/pdxjohnny/discovery/proxy"
)

var Commands = []*cobra.Command{
	&cobra.Command{
		Use:   "discovery",
		Short: "Discovery service commands",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			discovery.Run()
		},
	},
	&cobra.Command{
		Use:   "proxy",
		Short: "Reverse porxy to frontends",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			proxy.Run()
		},
	},
	&cobra.Command{
		Use:   "frontend",
		Short: "Frontend http(s) server to accept files",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			frontend.Run()
		},
	},
	&cobra.Command{
		Use:   "key",
		Short: "Generate RSA keys",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			key.Run()
		},
	},
}

func init() {
	ConfigDefaults(Commands...)
}

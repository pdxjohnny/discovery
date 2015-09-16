package commands

import (
	"github.com/spf13/cobra"

	"github.com/pdxjohnny/frontend/discovery"
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
}

func init() {
	ConfigDefaults(Commands...)
}

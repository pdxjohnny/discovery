package main

import (
	"runtime"

	"github.com/spf13/cobra"

	"github.com/pdxjohnny/discovery/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var rootCmd = &cobra.Command{Use: "discovery"}
	rootCmd.AddCommand(commands.Commands...)
	rootCmd.Execute()
}

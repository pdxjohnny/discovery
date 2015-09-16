package main

import (
	"runtime"

	"github.com/spf13/cobra"

	"github.com/pdxjohnny/frontend/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var rootCmd = &cobra.Command{Use: "frontend"}
	rootCmd.AddCommand(commands.Commands...)
	rootCmd.Execute()
}

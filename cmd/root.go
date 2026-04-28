package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tumeraltunbas/nomad/internal"
)

var rootCmd = &cobra.Command{
	Use:   internal.SupportedCommands.Base,
	Short: "Nomad is a CLI tool that keeps your AI development environment consistent across every machine you work on.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Nomad! Type 'nomad --help' to see what it can do.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

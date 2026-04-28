package cmd

import (
	"fmt"
	"slices"

	"github.com/spf13/cobra"
	"github.com/tumeraltunbas/nomad/internal"
)

var pullCmd = &cobra.Command{
	Use: internal.SupportedCommands.Pull,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		tool, _ := cmd.Flags().GetString(internal.ToolFlag.Long)
		storage, _ := cmd.Flags().GetString(internal.StorageFlag.Long)

		hasValidTool := slices.Contains(internal.SupportedTools, tool)
		hasValidStorage := slices.Contains(internal.SupportedStorages, storage)

		if hasValidTool == true && hasValidStorage == true {
			return nil
		}

		return fmt.Errorf(
			"invalid tool or storage. valid tools: %v, valid storages: %v",
			internal.SupportedTools,
			internal.SupportedStorages,
		)
	},

	Run: func(cmd *cobra.Command, args []string) {
		toolValue, _ := cmd.Flags().GetString(internal.ToolFlag.Long)
		storageValue, _ := cmd.Flags().GetString(internal.StorageFlag.Long)

		fmt.Println("Pull command runned", toolValue, storageValue)
	},
}

var tool string
var storage string

func init() {
	rootCmd.AddCommand(pullCmd)

	pullCmd.Flags().StringVarP(&tool, internal.ToolFlag.Long, internal.ToolFlag.Short, "", internal.ToolFlag.Usage)
	pullCmd.Flags().StringVarP(&storage, internal.StorageFlag.Long, internal.StorageFlag.Short, "", internal.StorageFlag.Usage)

	pullCmd.MarkFlagRequired(internal.ToolFlag.Long)
	pullCmd.MarkFlagRequired(internal.StorageFlag.Long)
}

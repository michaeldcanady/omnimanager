package command

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "agent",
		Short: "",
		Long:  ``,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	syncCmd.Flags().Bool("device", true, "Specify to sync only device policies")
	applyCmd.Flags().Bool("device", true, "Specify to apply only device policies")
	policyCmd.AddCommand(syncCmd, applyCmd)
	rootCmd.AddCommand(policyCmd)
}

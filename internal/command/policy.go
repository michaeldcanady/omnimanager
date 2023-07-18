package command

import (
	"fmt"

	"github.com/michaeldcanady/omnimanageragent/internal/localmanager"
	"github.com/michaeldcanady/omnimanageragent/internal/remotemanager"
	"github.com/spf13/cobra"
)

var policyCmd = &cobra.Command{
	Use:   "policy",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var apiEndpoint = "/policies"

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync agent policy",
	Run: func(cmd *cobra.Command, args []string) {
		_, _ = cmd.Flags().GetBool("device")
		localmgr, err := localmanager.NewLocalManager()
		if err != nil {
			fmt.Printf("failed to create LocalManager: %w", err)
			return
		}
		
		remotemgr := remotemanager.NewRemoteManager(apiEndpoint)

		//Get Device Configurations
		policies, err := remotemgr.GetDeviceConfigurations()
		if err != nil {
			fmt.Printf("failed to get device configurations: %w", err)
			return
		}

		//cache policies
		for _, policy := range policies {
			err = localmgr.CacheDeviceConfiguration(&policy)
			if err != nil {
				fmt.Printf("failed to cache device configuration: %w", err)
				return
			}
		}
	},
}

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply agent policy",
	Run: func(cmd *cobra.Command, args []string) {
		_, _ = cmd.Flags().GetBool("device")
	},
}

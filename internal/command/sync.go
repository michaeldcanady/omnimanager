package command

import "github.com/spf13/cobra"

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Synchronize device configurations between local and remote managers.",
	Long:  `Usage: omnimanager sync`,
	Run: func(cmd *cobra.Command, args []string) {
		//localmgr, err := localmanager.NewLocalManager()
		//if err != nil {
		//fmt.Errorf("failed to create LocalManager: %w", err)
		//	return 1
		//}
		//remotemgr := remotemanager.NewRemoteManager(apiEndpoint)

		//policies, err := remotemgr.GetDeviceConfigurations()
		//if err != nil {
		//fmt.Errorf("failed to get device configurations: %w", err)
		//	return 1
		//}

		// cache policies
		//for _, policy := range policies {
		//	err = localmgr.CacheDeviceConfiguration(&policy)
		//	if err != nil {
		//fmt.Errorf("failed to cache device configuration: %w", err)
		//		return 1
		//	}
		//}
	},
}

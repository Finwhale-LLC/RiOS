package cmd

import (
	"github.com/spf13/cobra"
)

var (
	apiEndpoint string
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "rios-worker",
	Short: "RiOS Worker CLI - Contribute GPU power to the RiOS network",
	Long: `RiOS Worker CLI allows you to contribute your GPU computing power
to the RiOS decentralized network and earn $ROS tokens as rewards.`,
	Run: func(cmd *cobra.Command, args []string) {
		PrintBanner()
		cmd.Help()
	},
}

// Execute executes the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&apiEndpoint, "api", "http://localhost:3000", "RiOS API endpoint")
}


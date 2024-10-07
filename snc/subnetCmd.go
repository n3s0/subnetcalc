package snc

import (
	"github.com/spf13/cobra"
)

<<<<<<< HEAD
var subnetCmd = &cobra.Command {
  Use: "subnet [subnet/CIDR]",
  Short: "NOT COMPLETE: Calculate and display subnet and subnet ranges",
  Long: "Calculate and display subnet and subnet ranges.",
  Run: func(cmd *cobra.Command, args []string) {
  },
=======
var subnetCmd = &cobra.Command{
	Use:   "subnet [subnet/CIDR]",
	Short: "Calculate and display subnet and subnet ranges",
	Long:  "Calculate and display subnet and subnet ranges.",
	Run: func(cmd *cobra.Command, args []string) {
	},
>>>>>>> 8a102b8 (Updating subnet calculator feature)
}

func init() {
	rootCmd.AddCommand(subnetCmd)
}

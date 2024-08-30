package snc

import (
  "github.com/spf13/cobra"
)

var subnetCmd = &cobra.Command {
  Use: "subnet [subnet/CIDR]",
  Short: "Calculate and display subnet and subnet ranges",
  Long: "Calculate and display subnet and subnet ranges.",
  Run: func(cmd *cobra.Command, args []string) {
  },
}

func init() {
  rootCmd.AddCommand(subnetCmd)
}

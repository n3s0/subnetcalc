package snc

import (
  "fmt"
  "github.com/spf13/cobra"
  "os"
)

var rootCmd = &cobra.Command {
  Use: "subnetcalc",
  Short: "Subnet calculator and reference",
  Long: "A command line subnet calculator and info reference.",
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Fprintf(os.Stderr, "Application threw error: %s\n", err)
    os.Exit(1)
  }
}

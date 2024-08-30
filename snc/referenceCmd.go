package snc

import (
	"github.com/spf13/cobra"
)

// Will be updating these variables for uniformity.
var subnetCalcTable bool
var decToBinTable bool
var classfulRangesTable bool
var reservedRangesTable bool
var ethernetCableLengthChart bool

var referenceCmd = &cobra.Command{
	Use:     "[ref|reference]",
	Aliases: []string{"ref"},
	Short:   "Reference charts",
	Long:    "Reference for subnet charts, classful ranges, reserved ranges, and decimal to binary.",
	Args:    cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if subnetCalcTable {
			SubnetChart()
		}

		if decToBinTable {
			DecimalToBinary()
		}

		if classfulRangesTable {
			ClassfulRanges()
		}

		if reservedRangesTable {
			ReservedRanges()
		}

		if ethernetCableLengthChart {
			EthernetCableLength()
		}
	},
}

func init() {
	rootCmd.AddCommand(referenceCmd)
	referenceCmd.Flags().BoolVarP(&subnetCalcTable, "subnet-chart", "", false, "Subnet chart")
	referenceCmd.Flags().BoolVarP(&decToBinTable, "decimal-to-binary", "", false, "Decimal to binary chart")
	referenceCmd.Flags().BoolVarP(&classfulRangesTable, "classful-ranges", "", false, "Classful ranges chart")
	referenceCmd.Flags().BoolVarP(&reservedRangesTable, "reserved-ranges", "", false, "Reserved ranges chart")
	referenceCmd.Flags().BoolVarP(&ethernetCableLengthChart, "ethernet-cable-length", "", false, "Ethernet cable length chart")
}

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var version = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Long:  "Show PinEnc version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("PinEnc Version v0.1.220513")
	},
}

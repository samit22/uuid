package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// empty represents command generating empty UUID.
var version = &cobra.Command{
	Use:   "version",
	Short: "Current version of uuid.",
	Long:  `Returns the version of uuid.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(getVersion())
	},
}

func getVersion() string {
	return Version
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// empty represents command generating empty UUID.
var empty = &cobra.Command{
	Use:     "empty",
	Aliases: []string{"nil"},
	Short:   "Generate empty uuid",
	Long: `Generate empty:

Generates empty (all-zeroes) uuid
`,
	Run: func(cmd *cobra.Command, args []string) {
		const emptyUUID = "00000000-0000-0000-0000-000000000000"

		fmt.Println(emptyUUID)
	},
}

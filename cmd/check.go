package cmd

import (
	"fmt"
	"regexp"

	"github.com/spf13/cobra"
)

// check represents a command validating whether the provided input is a valid UUID
var check = &cobra.Command{
	Use:   "check",
	Short: "Checks whether the provided input is a valid UUID",
	Long: `Checks whether the provided input is a valid UUID,
while using regular expression pattern matching.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		uuidPattern := "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
		re := regexp.MustCompile(uuidPattern)
		if re.MatchString(args[0]) {
			fmt.Println("valid")
			return
		}
		fmt.Println("invalid")
	},
}

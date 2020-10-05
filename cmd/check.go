package cmd

import (
	"fmt"
	"regexp"

	"github.com/spf13/cobra"
)

// check represents a command validating whether the provided input is a valid UUID
var check = &cobra.Command{
	Use:   "check",
	Short: "Checks whether each of the provided arguments is a valid UUID",
	Long: `Checks whether each of the provided arguments is a valid UUID,
while using regular expression pattern matching.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			if Check(arg) {
				fmt.Printf("%s is a valid UUID\n", arg)
			} else {
				fmt.Printf("%s is not a valid UUID\n", arg)
			}
		}
	},
}

func Check(uuid string) bool {
	uuidPattern := "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
	re := regexp.MustCompile(uuidPattern)
	if re.MatchString(uuid) {
		return true
	}
	return false
}

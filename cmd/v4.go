/*
Copyright © 2020 Samit samitghimire@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samit22/uuid/generator"
	"github.com/spf13/cobra"
)

// v4Cmd represents the v4 command
var v4Cmd = &cobra.Command{
	Use:   "v4",
	Short: "Generate UUID v4",
	Long: `Generate v4 UUID:

By default, a single UUID is generated. If a number is passed as an argument,
that amount of UUIDs will be generated.

Example: uuid v4 10
`,
	Run: func(cmd *cobra.Command, args []string) {
		ug := generator.Generate{}

		num, err := parseArgs(args)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		msg := "UUID"
		if num > 1 {
			msg = "UUIDs"
		}
		fmt.Printf("Generating %d %s\n\n", num, msg)
		uuids, err := generateUUIDV4(num, ug)
		if err != nil {
			fmt.Printf("Generating UUIDs failed: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(strings.Join(uuids, "\n"))
	},
}

func parseArgs(args []string) (n int, err error) {
	if len(args) == 0 {
		return 1, nil
	}

	num, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, errors.New("argument must be an integer")
	}
	if num < 1 {
		return 0, errors.New("please enter a number greater than 0")

	}

	return num, nil
}

func generateUUIDV4(n int, g generator.GeneratorV4) (res []string, err error) {
	res = make([]string, n)
	for i := 0; i < n; i++ {
		res[i], err = g.V4()
		if err != nil {
			return nil, err
		}
	}
	return
}

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
	"strconv"
	"strings"

	"github.com/samit22/uuid/generator"

	"github.com/spf13/cobra"
)

//GeneratorV4 interface to interact with uuid generator
type GeneratorV4 interface {
	V4() string
}

// v4Cmd represents the v4 command
var v4Cmd = &cobra.Command{
	Use:   "v4",
	Short: "Generate uuid v4",
	Long: `Generated v4 uuid:

By default single uuid is generated.
Multiple uuids is generated giving the argument the number of uuids to generate.
Example: uuid v4 10
`,
	Run: func(cmd *cobra.Command, args []string) {
		ug := generator.Generate{}
		generateUUIDV4(args, ug)
	},
}

func generateUUIDV4(args []string, g GeneratorV4) (uuids []string, err error) {
	var (
		num = 1
		msg = "UUID"
	)
	if len(args) != 0 {
		n := args[0]
		num, err = strconv.Atoi(n)
		if err != nil {
			errMsg := "Argument must a integer, please enter number between 1 to 100"
			fmt.Println(errMsg)
			err = errors.New(errMsg)
			return
		}
		if num < 1 || num > 100 {
			errMsg := "Please enter number between 1 to 100"
			fmt.Println(errMsg)
			err = errors.New(errMsg)
			return
		}
	}
	if num > 1 {
		msg = "UUIDS"
	}
	fmt.Printf("Generating %d %s \n \n", num, msg)

	uuids = generate(num, g)
	fmt.Println(strings.Join(uuids, "\n"))
	return
}

func generate(n int, g GeneratorV4) (res []string) {
	res = make([]string, n)
	for i := 0; i < n; i++ {
		res[i] = g.V4()
	}
	return
}

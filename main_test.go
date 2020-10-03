package main

import "os"

func ExampleGenerateEmpty() {
	// Pass argument to test 'uuid empty' command.
	os.Args = append(os.Args, "empty")

	main()
	// Output: 00000000-0000-0000-0000-000000000000
}

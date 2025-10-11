package main

import (
	"fmt"
	"os"
)

func main() {
	if err := Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Linting completed successfully with no issues.")
	os.Exit(0)
}

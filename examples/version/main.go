package main

import (
	"fmt"

	"github.com/agussyahrilmubarok/gohelp/version"
)

func main() {
	fmt.Printf("Package Version: %v\n", version.Version)
	fmt.Printf("Package Commit: %v\n", version.Commit)
	fmt.Printf("Package BuildDate: %v\n", version.BuildDate)
}

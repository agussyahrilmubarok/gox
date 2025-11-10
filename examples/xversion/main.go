package main

import (
	"fmt"

	"github.com/agussyahrilmubarok/gox/xversion"
)

func main() {
	fmt.Printf("Package Version: %v\n", xversion.Version)
	fmt.Printf("Package Commit: %v\n", xversion.Commit)
	fmt.Printf("Package BuildDate: %v\n", xversion.BuildDate)
}

package main

import (
	"fmt"

	"github.com/agussyahrilmubarok/gohelp/xstringutil"
)

func main() {
	strArr := []string{"hello_world", "user_name", "go-helper"}

	for _, v := range strArr {
		fmt.Println("====", v, "====")
		fmt.Println("CamelCase:", xstringutil.ToCamelCase(v))
		fmt.Println("SnakeCase:", xstringutil.ToSnakeCase(v))
		fmt.Println("Reversed :", xstringutil.Reverse(v))
		fmt.Println()
	}
}

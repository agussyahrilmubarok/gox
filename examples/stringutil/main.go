package main

import (
	"fmt"

	"github.com/agussyahrilmubarok/gohelp/stringutil"
)

func main() {
	strArr := []string{"hello_world", "user_name", "go-helper"}

	for _, v := range strArr {
		fmt.Println("====", v, "====")
		fmt.Println("CamelCase:", stringutil.ToCamelCase(v))
		fmt.Println("SnakeCase:", stringutil.ToSnakeCase(v))
		fmt.Println("Reversed :", stringutil.Reverse(v))
		fmt.Println()
	}
}

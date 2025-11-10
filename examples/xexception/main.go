package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/agussyahrilmubarok/gox/xexception"
)

func simulateLogin(username, password string) error {
	if username == "" || password == "" {
		return xexception.NewHTTPBadRequest("username or password cannot be empty", nil)
	}

	if username != "admin" || password != "1234" {
		return xexception.NewHTTPUnauthorized("invalid username or password", errors.New("auth failed"))
	}

	if username == "panic" {
		return xexception.NewHTTPInternal("unexpected server error", errors.New("database connection lost"))
	}

	return nil
}

func handleLogin(username, password string) {
	err := simulateLogin(username, password)
	if err != nil {
		if httpErr, ok := err.(*xexception.Http); ok {
			fmt.Printf("❌ HTTP Error [%d %s]: %s\n",
				httpErr.Code,
				http.StatusText(httpErr.Code),
				httpErr.Message,
			)
			if httpErr.Err != nil {
				fmt.Println("↳ Root cause:", httpErr.Err)
			}
			return
		}

		fmt.Println("❌ Unexpected error:", err)
		return
	}

	fmt.Println("✅ Login successful!")
}

func main() {
	fmt.Println("=== Example: Login with empty username ===")
	handleLogin("", "1234")

	fmt.Println("\n=== Example: Login with wrong credentials ===")
	handleLogin("guest", "wrong")

	fmt.Println("\n=== Example: Login causing internal error ===")
	handleLogin("panic", "1234")

	fmt.Println("\n=== Example: Successful login ===")
	handleLogin("admin", "1234")
}

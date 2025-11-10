package xstringutil_test

import (
	"testing"

	"github.com/agussyahrilmubarok/gohelp/xstringutil"
)

func TestToCamelCase(t *testing.T) {
	tests := map[string]string{
		"hello_world":  "HelloWorld",
		"user_name":    "UserName",
		"go-helper":    "GoHelper",
		"":             "",
		"alreadyCamel": "Alreadycamel",
	}

	for input, expected := range tests {
		got := xstringutil.ToCamelCase(input)
		if got != expected {
			t.Errorf("ToCamelCase(%q) = %q; want %q", input, got, expected)
		}
	}
}

func TestToSnakeCase(t *testing.T) {
	tests := map[string]string{
		"HelloWorld": "hello_world",
		"UserName":   "user_name",
		"simple":     "simple",
		"":           "",
	}

	for input, expected := range tests {
		got := xstringutil.ToSnakeCase(input)
		if got != expected {
			t.Errorf("ToSnakeCase(%q) = %q; want %q", input, got, expected)
		}
	}
}

func TestContainsIgnoreCase(t *testing.T) {
	if !xstringutil.ContainsIgnoreCase("HelloWorld", "world") {
		t.Error("Expected true, got false")
	}
	if xstringutil.ContainsIgnoreCase("HelloWorld", "planet") {
		t.Error("Expected false, got true")
	}
}

func TestReverse(t *testing.T) {
	tests := map[string]string{
		"hello": "olleh",
		"abc":   "cba",
		"":      "",
		"a":     "a",
	}

	for input, expected := range tests {
		got := xstringutil.Reverse(input)
		if got != expected {
			t.Errorf("Reverse(%q) = %q; want %q", input, got, expected)
		}
	}
}

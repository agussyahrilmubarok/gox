package xversion_test

import (
	"testing"

	"github.com/agussyahrilmubarok/gox/pkg/xversion"
)

func TestInfo(t *testing.T) {
	got := xversion.Info()
	if got == "" {
		t.Error("Info() should not return empty string")
	}
}

func TestDefaultValues(t *testing.T) {
	if xversion.Version == "" || xversion.Commit == "" || xversion.BuildDate == "" {
		t.Error("Default version variables should be non-empty")
	}
}

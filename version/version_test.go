package version_test

import (
	"testing"

	"github.com/agussyahrilmubarok/gohelp/version"
)

func TestInfo(t *testing.T) {
	got := version.Info()
	if got == "" {
		t.Error("Info() should not return empty string")
	}
}

func TestDefaultValues(t *testing.T) {
	if version.Version == "" || version.Commit == "" || version.BuildDate == "" {
		t.Error("Default version variables should be non-empty")
	}
}

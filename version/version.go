package version

import "fmt"

var (
	Version   = "dev"
	Commit    = "none"
	BuildDate = "unknown"
)

func Info() string {
	return fmt.Sprintf("Version: %s, Commit: %s, Build Date: %s", Version, Commit, BuildDate)
}

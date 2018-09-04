package version

import (
	"fmt"

	"github.com/Masterminds/semver"
)

var (
	Version string
	GitHash string
)

// String prints the build information for the bot
func String() string {
	_, err := semver.NewVersion(Version)
	if err != nil {
		Version = "Dev build (no valid version)"
	}
	if len(GitHash) == 0 {
		GitHash = "N/A"
	}
	return fmt.Sprintf("Version : %s\nGit Hash: %s\n", Version, GitHash)
}
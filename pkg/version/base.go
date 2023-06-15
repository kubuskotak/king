// Package version provides meta description of the service.
package version

// These should be set via go build -ldflags -X 'vvv'.
var (
	GitCommit   string
	GitDescribe string
	BuildTime   string

	Version    = "alpha-test"
	Prerelease = "dev1"
	Metadata   = ""

	GoVersion = "unknown"
	BuildUser = "unknown"
)

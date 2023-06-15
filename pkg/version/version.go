// Package version provides meta description of the service.
package version

import (
	"bytes"
	"fmt"
)

// Info is the result of an info.
type Info struct {
	Revision          string `json:"revision,omitempty"`
	Version           string `json:"version,omitempty"`
	VersionPrerelease string `json:"version_prerelease,omitempty"`
	VersionMetadata   string `json:"version_metadata,omitempty"`
	BuildDate         string `json:"build_date,omitempty"`
}

// GetVersion represent all info data struct.
func GetVersion() *Info {
	ver := Version
	rel := Prerelease
	md := Metadata
	if GitDescribe != "" {
		ver = GitDescribe
	}
	if GitDescribe == "" && rel == "" && Prerelease != "" {
		rel = "dev"
	}

	return &Info{
		Revision:          GitCommit,
		Version:           ver,
		VersionPrerelease: rel,
		VersionMetadata:   md,
		BuildDate:         BuildTime,
	}
}

// VersionNumber get number version of the service.
func (c *Info) VersionNumber() string {
	if Version == "unknown" && Prerelease == "unknown" {
		return "(version unknown)"
	}

	version := c.Version

	if c.VersionPrerelease != "" {
		version = fmt.Sprintf("%s-%s", version, c.VersionPrerelease)
	}

	if c.VersionMetadata != "" {
		version = fmt.Sprintf("%s+%s", version, c.VersionMetadata)
	}

	return version
}

// FullVersionNumber get full version of the service.
func (c *Info) FullVersionNumber(rev bool) string {
	var versionString bytes.Buffer

	if Version == "unknown" && Prerelease == "unknown" {
		return "Ymir (version unknown)"
	}

	_, _ = fmt.Fprintf(&versionString, "Ymir v%s", c.Version)
	if c.VersionPrerelease != "" {
		_, _ = fmt.Fprintf(&versionString, "-%s", c.VersionPrerelease)
	}

	if c.VersionMetadata != "" {
		_, _ = fmt.Fprintf(&versionString, "+%s", c.VersionMetadata)
	}

	if rev && c.Revision != "" {
		_, _ = fmt.Fprintf(&versionString, " (%s)", c.Revision)
	}

	if c.BuildDate != "" {
		_, _ = fmt.Fprintf(&versionString, ", built %s", c.BuildDate)
	}

	return versionString.String()
}

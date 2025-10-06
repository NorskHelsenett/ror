package rorversion

import (
	"fmt"
	"strings"
)

const (
	DefaultVersion = "0.0.0-dev"
	DefaultCommit  = "FFFFFF"
	DefaultLibVer  = "N/A"
)

var (
	Version string
	Commit  string
	LibVer  string
)

type RorVersion struct {
	Version string
	Commit  string
	LibVer  string
}

func NewRorVersion(version string, commit string) RorVersion {
	if Version == "" {
		Version = version
	}
	if Commit == "" {
		Commit = commit
	}
	if LibVer == "" {
		LibVer = "N/A"
	}
	return GetRorVersion()

}

func GetRorVersion() RorVersion {
	if Version == "" {
		Version = DefaultVersion
	}
	if Commit == "" {
		Commit = DefaultCommit
	}
	if LibVer == "" {
		LibVer = DefaultLibVer
	}
	return RorVersion{
		Version: Version,
		Commit:  Commit,
		LibVer:  LibVer,
	}
}

func (v RorVersion) GetVersion() string {
	return v.Version
}

func (v RorVersion) GetCommit() string {
	return v.Commit
}
func (v RorVersion) GetLibVer() string {
	return v.LibVer
}

func (v RorVersion) GetMajorVersion() string {
	split := strings.Split(v.Version, ".")
	return split[0]
}
func (v RorVersion) GetMinorVersion() string {
	split := strings.Split(v.Version, ".")
	return split[1]
}
func (v RorVersion) GetPatchVersion() string {
	split := strings.Split(v.Version, ".")
	return split[2]
}

// GetVersionWithCommit returns version-commit
// in the format 1.0.0-abcdefg
func (v RorVersion) GetVersionWithCommit() string {
	return fmt.Sprintf("%s-%s", v.Version, v.Commit)
}

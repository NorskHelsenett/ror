package rorversion

import (
	"fmt"
	"strings"
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
	if Version != "" {
		version = Version
	}
	if Commit != "" {
		commit = Commit
	}
	if LibVer == "" {
		LibVer = "N/A"
	}
	return RorVersion{
		Version: version,
		Commit:  commit,
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
func (v RorVersion) GetVersionWithCommit() string {
	return fmt.Sprintf("%s-%s", v.Version, v.Commit)
}

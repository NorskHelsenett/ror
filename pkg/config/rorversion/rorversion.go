package rorversion

import (
	"fmt"
	"strings"
)

type RorVersion struct {
	Version string
	Commit  string
}

func NewRorVersion(version string, commit string) RorVersion {
	return RorVersion{
		Version: version,
		Commit:  commit,
	}
}

func (v RorVersion) GetVersion() string {
	return v.Version
}

func (v RorVersion) GetCommit() string {
	return v.Commit
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

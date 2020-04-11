package auth0

import (
	"strconv"

	"github.com/hashicorp/go-version"
)

const v = "4.3.0"

var semver *version.Version

func init() {
	semver = version.Must(version.NewVersion(v))
}

// Version of this library.
func Version() string {
	return semver.String()
}

// VersionMajor is the major version of this library
func VersionMajor() string {
	return strconv.Itoa(semver.Segments()[0])
}

// SemVer returns the version parsed as semver.
func SemVer() *version.Version {
	return semver
}

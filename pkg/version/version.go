package version

/*
Folloing value is defined during linking in makefile as follows:

VERSION := $(shell git describe --tags --abbrev=0)
BUILD_OPTS := -ldflags "-X 'inspection/pkg/version.MajorMinorRevision=$(VERSION)'"
*/
var (
	MajorMinorRevision = "X.X.X"
	Build              = "0"
)

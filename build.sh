#!/bin/bash


VERSION=$(git describe --tags --abbrev=0)

#echo ${VERSION}
#BUILD_OPTS=-ldflags "-X 'inspection/pkg/version.MajorMinorRevision=${VERSION}'"
#echo ${BUILD_OPTS}

go build -ldflags "-X 'inspection/pkg/version.MajorMinorRevision=${VERSION}'" ./cmd/ginspection
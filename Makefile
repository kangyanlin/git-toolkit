BUILD_VERSION	:= $(shell cat version)
BUILD_TIME		:= $(shell date "+%F %T")
COMMIT_SHA1		:= $(shell git rev-parse HEAD)

all:
	gox	-osarch="darwin/amd64 linux/386 linux/amd64" \
		-output="dist/{{.Dir}}_{{.OS}}_{{.Arch}}" \
		-ldflags "-X 'github.com/tonydeng/git-toolkit/cmd.Version=${BUILD_VERSION}' \
		-X 'github.com/tonydeng/git-toolkit/cmd.BuildTime=${BUILD_TIME}' \
		-X 'github.com/tonydeng/git-toolkit/cmd.CommitID=${COMMIT_SHA1}'"

clean:
	rm -rf dist

install:
	go install -ldflags "-X 'github.com/tonydeng/git-toolkit/cmd.Version=${BUILD_VERSION}' \
       					 -X 'github.com/tonydeng/git-toolkit/cmd.BuildTime=${BUILD_TIME}' \
       					 -X 'github.com/tonydeng/git-toolkit/cmd.CommitID=${COMMIT_SHA1}'"

.PHONY: all clean install

.EXPORT_ALL_VARIABLES:

GO111MODULE = on
GOPROXY = https://goproxy.io
GOSUMDB = sum.golang.google.cn
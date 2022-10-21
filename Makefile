CGO_CPPFLAGS ?= ${CPPFLAGS}
export CGO_CPPFLAGS
CGO_CFLAGS ?= ${CFLAGS}
export CGO_CFLAGS
CGO_LDFLAGS ?= $(filter -g -L% -l% -O%,${LDFLAGS})
export CGO_LDFLAGS

.PHONY: build
build:
	GOOS= GOARCH= GOARM= GOFLAGS= CGO_ENABLED= go build -o ohmyflux ./cmd/ohmyflux

archive_linux_amd64:
	tar czf ohmyflux_linux_amd64.tar.gz ohmyflux

archive_linux_arm64:
	tar czf ohmyflux_linux_arm64.tar.gz ohmyflux

archive_macOs_amd64:
	tar czf ohmyflux_macOs_amd64.tar.gz ohmyflux

archive_macOs_amd64:
	tar czf ohmyflux_macOs_arm64.tar.gz ohmyflux

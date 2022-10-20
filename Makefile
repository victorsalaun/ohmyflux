CGO_CPPFLAGS ?= ${CPPFLAGS}
export CGO_CPPFLAGS
CGO_CFLAGS ?= ${CFLAGS}
export CGO_CFLAGS
CGO_LDFLAGS ?= $(filter -g -L% -l% -O%,${LDFLAGS})
export CGO_LDFLAGS

.PHONY: build
build:
	GOOS= GOARCH= GOARM= GOFLAGS= CGO_ENABLED= go build -o ohmyflux ./cmd/ohmyflux

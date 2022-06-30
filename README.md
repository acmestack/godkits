# godkits

[![license card](https://img.shields.io/badge/License-Apache%202.0-brightgreen.svg?label=license)](https://github.com/acmestack/godkits/blob/main/LICENSE)
[![go version](https://img.shields.io/github/go-mod/go-version/acmestack/godkits)](#)
[![go report](https://goreportcard.com/badge/github.com/acmestack/godkits)](https://goreportcard.com/report/github.com/acmestack/godkits)
[![codecov report](https://codecov.io/gh/acmestack/godkits/branch/main/graph/badge.svg)](https://codecov.io/gh/acmestack/godkits)
[![workflow](https://github.com/acmestack/godkits/actions/workflows/go.yml/badge.svg?event=push)](#)
[![Go Reference](https://pkg.go.dev/badge/github.com/acmestack/godkits.svg)](https://pkg.go.dev/github.com/acmestack/godkits)
[![build docs](https://github.com/acmestack/godkits/actions/workflows/build-docs.yaml/badge.svg)](https://github.com/acmestack/godkits/actions/workflows/build-docs.yaml)
[![lasted release](https://img.shields.io/github/v/release/acmestack/godkits?label=lasted)](https://github.com/acmestack/godkits/releases)


the golang development toolkits.

## Naming

### Go SDK Extensions
- package start with `gox`, others same with the go sdk
- using the `x` suffix, means **extension**, like `listx`, `stringx`
- `test` in the same place

### Other Extensions
- package start with **feature name**, like `http`


## Usage
```go

import "github.com/acmestack/godkits/log/log"

func TestLog(t *testing.T) {
	log.Info("test")
}
```

## Code Comment

### Code comment with method
```go
// TestLog test log
//  @params t tests params
func TestLog(t *testing.T) {
log.Info("test")
}

// NewAsyncWriter Write data with Buffer, this Writer and Closer is thread safety, but WriteCloser parameters not safety.
//  @param w       Writer
//  @param bufSize accept buffer max length
//  @param block   if true, overflow buffer size, will blocking, if false will occur error
//  @return *AsyncLogWriter
func NewAsyncWriter(w io.Writer, bufSize int, block bool) *AsyncLogWriter {
}
```

## Stargazers over time

[![Stargazers over time](https://starchart.cc/acmestack/godkits.svg)](https://starchart.cc/acmestack/godkits)

## Contribute and Support

- [How to Contribute](https://acmestack.org/docs/contributing/guide/)
- [Document](https://pkg.go.dev/github.com/acmestack/godkits)

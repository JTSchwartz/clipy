Clipy
===

[![GoDoc](https://godoc.org/github.com/jtschwartz/clipy?status.svg)](https://pkg.go.dev/github.com/jtschwartz/clipy)
[![Go Report Card](https://goreportcard.com/badge/jtschwartz/clipy)](https://goreportcard.com/report/jtschwartz/clipy)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/JTSchwartz/clipy)
![GitHub Release Date](https://img.shields.io/github/release-date/JTSchwartz/clipy)
![GitHub issues](https://img.shields.io/github/issues/JTSchwartz/clipy)

Clipy is a small command line utility to sending data to your clipboard.

## Usage

Clipy can send data to your clipboard either via standard input (pipes) or from a file:
```shell
echo "Send me to your clipboard" | clipy
# Or
clipy /path/to/file.ext
```
In most cases, a user would have clipy as the final step of the process and would prefer clipy not to write out whatever data it took in. However, if you choose to put clipy somewhere else in a list of pipe-connected utilities, simply use the `-o / --output` flag:
```shell
clipy -o path/to/file.ext | grep ...
```

## Install

```shell
go install github.com/jtschwartz/clipy@latest
```

Installation requires Go 1.16+ and ensure that `$GOPATH/bin` is apart of your `PATH`.
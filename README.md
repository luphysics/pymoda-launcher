# PyMODA launcher

This repository contains the PyMODA launcher for macOS/Linux, which is written in Go.

## Purpose 

The PyMODA launcher is a small executable which acts as a wrapper around the PyMODA implementation. 

When PyMODA performs an update, it downloads and extracts the new release into a new folder alongside the current release. When the PyMODA launcher is next opened, it will launch the new release, providing a seamless update experience with zero downtime.

## Developer notes

Go compiles to native executables, and it has excellent support for cross-compilation.

To compile the launcher for your platform, install Go and run:

```
go build launcher.go
```

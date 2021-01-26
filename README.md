# Gomock-proj 
[![Build Status](https://travis-ci.com/ShanQincheng/gomock-proj.svg?branch=master)](https://travis-ci.com/ShanQincheng/gomock-proj)
[![Go Report Card](https://goreportcard.com/badge/github.com/shanqincheng/gomock-proj)](https://goreportcard.com/report/github.com/shanqincheng/gomock-proj)

# Overview
Gomock-proj is a tool for mocking an entire **Go** project in one go.
It generates mock files under the directory `$pwd/test/mocks`. If 
the directory doesn't exist, gomock-proj will create it.

Gomock-proj provides:
* Traverse and mock Go files in entire project.
* Traverse and mock Go files under specific directory.

# Concepts
Gomock-proj concurrently invokes 
[gomock](https://github.com/golang/mock) command
`mockgen -source=foo.go -destination=test/mocks/foo.go` to 
generate mock files.

After all mock files have been generated, Gomock-proj invokes command
`goimport -w test/mocks` to remove unnecessary imports in mock files.

# Installing
```
go get golang.org/x/tools/cmd/goimports

GO111MODULE=on go get github.com/golang/mock/mockgen@v1.4.4

GO111MODULE=on go get github.com/shanqincheng/gomock-proj
```

# Getting Started
Gomock-proj has only one operation: **mock**, with one
necessary flag: **--dir / -d**.

*Following command will **create two directories `test` and 
`test/mocks`** if it doesn't exist in current path, and then 
**create mock files in `test/mocks`** which as same as original Go source 
file in Name and Relative Path*.

Example ( mock Go files in entire project ):
```
cd toYourGoProjectRootDir
gomock-proj mock --dir="."
```

Example ( mock Go files in specific dirs: *foo* and *bar* ):
```
cd toYourGoProjectRootDir
gomock-proj mock --dir="foo,bar"
```

# License

Gomock-proj is released under the Apache 2.0 license. See [LICENSE.txt](https://github.com/spf13/cobra/blob/master/LICENSE.txt)
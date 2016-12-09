# gowrap
A packaging tool for apps developed in Go. 

## Features
* Creates multi-platform builds without a Makefile.
* Versions your app during the build (uses Git `tag` for release, Git commit `hash` for development).
* Adds builds to compressed archives (`.zip` for Windows builds, `.tar.gz` for others).
* Generates cleanly structured archive file names in the form of `appname-version-goos-goarch.ext`.
* Runs on any Go-supported platform.

## Installation
Run in terminal:
```sh
$ go get -u github.com/samherrmann/gowrap
```

In order to use `gowrap` conveniently from any project directory, make sure that either `$GOBIN`, if it is set, or `$GOPATH/bin` is in your system `$PATH`.

## Usage
Run in terminal:
```sh
$ cd project-root # directory containing main.go
$ gowrap
```
When you run `gowrap`, it first looks for a `gowrap.json` file in the current working directory. If the file does not exist, `gowrap` will generate a sample file for you and exit. 

### gowrap.json
The following is a sample `gowrap.json` file:
```json
{
    "targets": [
        "linux-amd64"
    ]
}
```
Each string element in the `targets` array must be of the form `$GOOS-$GOARCH`. Edit the array and add all the target platforms on which you want to deploy your application. The valid combinations of `$GOOS` and `$GOARCH` are documented [here](https://golang.org/doc/install/source#environment).

Once you have added all the target platforms for your application to the `gowrap.json` file, re-run `gowrap` in the terminal. `gowrap` will now generate a build for all the target platforms and write them to the `/dist` folder.

### Version
In order to have `gowrap` embed a version number into your app, simply add a `version` field to your `main` package:
```go
package main

import "fmt"

var version string

func main() {
    fmt.Println("Version: " + version) // not required
}
```
During the build, `gowrap` will set the value of `version` to the Git `tag`, if one exists at the HEAD, or else the Git commit `hash` of the HEAD. It is common to use Git tags to mark release points (ex: v1.0.0), therefore the described `gowrap` behaviour creates a clear differentiation between a development build and a release build.

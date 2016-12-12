package gotools

import (
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"strings"
)

const (
	envVarKeyGOOS       = "GOOS"
	envVarKeyGOARCH     = "GOARCH"
	envVarKeyGOHOSTOS   = "GOHOSTOS"
	envVarKeyGOHOSTARCH = "GOHOSTARCH"
)

var (
	// GOHOSTOS is the operating system of the host.
	GOHOSTOS string

	// GOHOSTARCH the compilation architecture of the host.
	GOHOSTARCH string

	// SupportedPlatforms is a list of all the target operating
	// systems and compilation architectures
	SupportedPlatforms *[]Platform
)

func init() {
	os, err := initGoHostOS()
	if err != nil {
		panic(err)
	}
	arch, err := initGoHostArch()
	if err != nil {
		panic(err)
	}
	platforms, err := initSupportedPlatforms()
	if err != nil {
		panic(err)
	}
	GOHOSTOS = os
	GOHOSTARCH = arch
	SupportedPlatforms = platforms
}

// initGoHostOS returns the value of GOHOSTOS
func initGoHostOS() (string, error) {
	return goEnvVar(envVarKeyGOHOSTOS)
}

// GoHostArch returns the value of GOHOSTARCH
func initGoHostArch() (string, error) {
	return goEnvVar(envVarKeyGOHOSTARCH)
}

// initSupportedPlatforms returns a list of all the target operating
// systems and compilation architectures.

// This function requires at least Go version go1.7 to be installed.
// The 'go tool dist list' command that this function calls was introduced
// in the following commit:
// https://github.com/golang/go/commit/c3ecded729214abf8a146902741cd6f9d257f68c
func initSupportedPlatforms() (*[]Platform, error) {
	p := &[]Platform{}

	out, err := exec.Command("go", "tool", "dist", "list", "-json").Output()
	if err != nil {
		return p, errors.New("Error while getting list of supported platforms: " +
			err.Error() +
			". Are you running at least Go version go1.7?")
	}

	err = json.Unmarshal(out, p)
	return p, err
}

// Generate executes the command "go generate"
func Generate() error {
	_, err := cmdOutput("go", "generate")
	return err
}

// Build executes the command "go build" for the desired
// target OS and architecture, and writes the generated
// executable to the 'outDir' directory.
func Build(args ...string) error {
	args = append([]string{"build"}, args...)
	_, err := cmdOutput("go", args...)
	return err
}

// GoOS returns the value of GOOS
func GoOS() (string, error) {
	return goEnvVar(envVarKeyGOOS)
}

// SetGoOS sets the value of GOOS
func SetGoOS(goos string) error {
	return os.Setenv(envVarKeyGOOS, goos)
}

// GoArch returns the value of GOARCH
func GoArch() (string, error) {
	return goEnvVar(envVarKeyGOARCH)
}

// SetGoArch sets the value of GOARCH
func SetGoArch(goarch string) error {
	return os.Setenv(envVarKeyGOARCH, goarch)
}

// ExeSuffix returns ".exe" if the GOOS
// environment variable is set to
// "windows".
func ExeSuffix() (string, error) {
	goos, err := GoOS()
	if err != nil {
		return "", err
	}
	if goos == "windows" {
		return ".exe", nil
	}
	return "", nil
}

// goEnvVar returns the value of the provided Go
// environment variable
func goEnvVar(key string) (string, error) {
	return cmdOutput("go", "env", key)
}

// cmdOutput executes the command specified by name and
// returns the first line of the standard output.
func cmdOutput(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.New(string(out))
	}

	return strings.TrimRight(string(out), "\n"), err
}

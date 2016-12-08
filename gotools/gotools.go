package gotools

import (
	"encoding/json"
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
func initSupportedPlatforms() (*[]Platform, error) {
	p := &[]Platform{}

	out, err := exec.Command("go", "tool", "dist", "list", "-json").Output()
	if err != nil {
		return p, err
	}

	err = json.Unmarshal(out, p)
	return p, err
}

// Generate executes the command "go generate"
func Generate() error {
	return exec.Command("go", "generate").Run()
}

// Build executes the command "go build" for the desired
// target OS and architecture, and writes the generated
// executable to the 'outDir' directory.
func Build(args ...string) error {
	args = append([]string{"build"}, args...)
	return exec.Command("go", args...).Run()
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
	out, err := exec.Command(name, args...).Output()
	return strings.TrimRight(string(out), "\n"), err
}

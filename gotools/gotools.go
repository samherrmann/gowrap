package gotools

import (
	"os"
	"os/exec"
	"strings"
)

const (
	cmdName         = "go"
	envVarKeyGOOS   = "GOOS"
	envVarKeyGOARCH = "GOARCH"
)

// Generate executes the command "go generate"
func Generate() error {
	return exec.Command(cmdName, "generate").Run()
}

// Build executes the command "go build" for the desired
// target OS and architecture, and writes the generated
// executable to the 'outDir' directory.
func Build(args ...string) error {
	args = append([]string{"build"}, args...)
	return exec.Command(cmdName, args...).Run()
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
	return cmdOutput(cmdName, "env", key)
}

// cmdOutput executes the command specified by name and
// returns the first line of the standard output.
func cmdOutput(name string, args ...string) (string, error) {
	out, err := exec.Command(name, args...).Output()
	return strings.TrimRight(string(out), "\n"), err
}

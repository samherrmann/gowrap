package gotools

import (
	"os"

	"github.com/samherrmann/gowrap/exec"
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

// GetGoOS returns the value of GOOS
func GetGoOS() (string, error) {
	return getGoEnvVar(envVarKeyGOOS)
}

// GetGoArch returns the value of GOARCH
func GetGoArch() (string, error) {
	return getGoEnvVar(envVarKeyGOARCH)
}

// SetGoOS returns the value of GOOS
func SetGoOS(goos string) error {
	return os.Setenv(envVarKeyGOOS, goos)
}

// SetGoArch returns the value of GOARCH
func SetGoArch(goarch string) error {
	return os.Setenv(envVarKeyGOARCH, goarch)
}

// ExeSuffix returns ".exe" if the GOOS
// environment variable is set to
// "windows".
func ExeSuffix() (string, error) {
	goos, err := GetGoOS()
	if err != nil {
		return "", err
	}
	if goos == "windows" {
		return ".exe", nil
	}
	return "", nil
}

// getGoEnvVar returns the value of the provided Go
// environment variable
func getGoEnvVar(envVar string) (string, error) {
	return exec.Command(cmdName, "env", envVar).OutputLine()
}

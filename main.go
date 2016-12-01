package main

import "fmt"

var (
	// appName is the name of the
	// application to be built.
	appName = currentFolderName()

	// appVersion is the version of
	// the application to be built.
	appVersion = gitVersion()
)

func main() {

	config, err := readOrSaveConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	runGoBuildChain(config)
}

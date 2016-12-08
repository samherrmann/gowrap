package gotools

import (
	"testing"
)

func TestSupportedPlatforms(t *testing.T) {
	platforms, err := initSupportedPlatforms()
	if err != nil {
		t.Fatal("SupportedPlatforms returned an error: " + err.Error())
	}
	if len(*platforms) == 0 {
		t.Error("SupportedPlatforms returned a zero-length list.")
	}
}

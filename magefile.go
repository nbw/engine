// +build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/sh"
)

// Install builds and installs the engine in $GOPATH/bin
func Install() error {
	fmt.Println("installing engine")
	err := sh.Run("go", "install", "-v", "github.com/battlesnakeio/engine")
	if err != nil {
		return err
	}
	return nil
}

// Install builds and installs the engine-cli in $GOPATH/bin
func InstallCLI() error {
	fmt.Println("installing engine")
	err := sh.Run("go", "install", "-v", "github.com/battlesnakeio/engine/cmd/engine-cli")
	if err != nil {
		return err
	}
	return nil
}

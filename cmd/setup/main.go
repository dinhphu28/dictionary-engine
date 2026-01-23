package main

import (
	"fmt"
	"log"

	"github.com/dinhphu28/dictionary/internal/setup"
)

func main() {
	osinfo := setup.DetectOS()
	fmt.Println("Detected OS:", osinfo.Name)
	fmt.Println("Installing dictionary...")

	if osinfo.IsLinux {
		paths := setup.InstallationPaths()

		if err := setup.Install(paths); err != nil {
			log.Fatalf("install failed: %v", err)
		}

		setup.InstallNativeMessagingManifests(paths.BinPath)

		fmt.Println("✅ Installation complete")
		fmt.Println("Make sure ~/.local/bin is in your PATH")
	}
}

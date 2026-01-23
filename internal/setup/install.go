package setup

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/dinhphu28/dictionary/internal/config"
)

func Install(paths Paths) error {
	// 1. Install binary
	binPath := paths.BinPath
	log.Printf("BIN PATH: %v", binPath)
	if err := installBinary(binPath); err != nil {
		return err
	}

	// 2. Install config
	if err := WriteConfigToml(
		filepath.Join(paths.ConfigDir, config.ConfigFile),
		DefaultConfig(),
		false,
	); err != nil {
		return err
	}

	// 3. Install resources
	resourcesPath := DefaultConfig().Paths.Resources
	log.Printf("RESOURCES PATH: %v", resourcesPath)
	if err := installResources(resourcesPath); err != nil {
		return err
	}

	return nil
}

func installBinary(path string) error {
	if err := copyFile(
		DefaultPaths().BinPath,
		path,
		0o755,
	); err != nil {
		return fmt.Errorf("install binary: %w", err)
	}
	return nil
}

func InstallNativeMessagingManifests(bin string) {
	chrome, firefox := detectBrowsers()

	if chrome {
		manifest := chromeManifest(bin, "kpgiaenkniiaacjbiipbmcdjfbjmgmll")
		if err := installChromeManifest(manifest); err != nil {
			log.Fatalf("install chrome native messaging manifest failed: %v", err)
		}
	}

	if firefox {
		manifest := firefoxManifest(bin, "503e78dec27c89515afd99f62ecf12e3305a204d@temporary-addon")
		if err := installFirefoxManifest(manifest); err != nil {
			log.Fatalf("install firefox native messaging manifest failed: %v", err)
		}
	}
}

func installResources(resourcesPath string) error {
	if _, err := os.Stat("./resources"); err == nil {
		if err := copyDir("./resources", resourcesPath); err != nil {
			return err
		}
	}
	return nil
}

package portable

import (
	"os"

	"github.com/dinhphu28/dictionary/internal/startup"
)

func IsPortable() bool {
	return hasConfig()
}

func hasConfig() bool {
	configPath := startup.ResolvePath("config.json")
	return fileExists(configPath)
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

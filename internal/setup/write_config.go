package setup

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/dinhphu28/dictionary/internal/config"
)

func WriteConfigToml(path string, cfg *config.Config, overwrite bool) error {
	// Ensure directory exists
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("create config dir: %w", err)
	}

	// Refuse to overwrite unless explicitly allowed
	if !overwrite {
		if _, err := os.Stat(path); err == nil {
			return fmt.Errorf("config already exists: %s", path)
		}
	}

	f, err := os.OpenFile(
		path,
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0o644,
	)
	if err != nil {
		return fmt.Errorf("open config file: %w", err)
	}
	defer f.Close()

	enc := toml.NewEncoder(f)
	enc.Indent = "  "

	if err := enc.Encode(cfg); err != nil {
		return fmt.Errorf("encode toml: %w", err)
	}

	return nil
}

func DefaultConfig() *config.Config {
	cfg := &config.Config{
		Version: 1,
		Priority: []string{
			"oxford_american",
			"vietname_english",
			"oxford_british_dictionary",
			"oxford_american_writer_thesaurus",
			"oxford_british_thesaurus",
		},
	}

	cfg.Paths.Resources = filepath.Join(InstallationPaths().DataDir, "resources")

	return cfg
}

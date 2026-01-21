package config

const (
	ConfigFile = "config.toml"
)

type Config struct {
	Version  int      `toml:"version"`
	Priority []string `toml:"priority"`

	Paths struct {
		Resources string `toml:"resources"`
	} `toml:"paths"`
}

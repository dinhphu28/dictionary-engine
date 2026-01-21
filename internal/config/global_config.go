package config

type GlobalConfig struct {
	Priority []string `json:"priority"`
	DataDir  string   `json:"data_dir"`
}

package dictionary

import (
	"log"
	"path/filepath"

	"github.com/dinhphu28/dictionary/database"
	"github.com/dinhphu28/dictionary/internal/config"
	"github.com/dinhphu28/dictionary/internal/setup"
	"github.com/dinhphu28/dictionary/internal/startup"
)

var (
	cfg               config.Config
	dictionaries      []database.Dictionary
	approximateLookup *ApproximateLookup
)

func StartEngine() {
	loadConfigAndData()

	dictionaryLookup := NewDictionaryLookup(dictionaries, cfg.Priority)
	approximateLookup = NewApproximateLookup(*dictionaryLookup)
}

func Ready() bool {
	return len(dictionaries) > 0
}

func LoadedDictionaries() int {
	return len(dictionaries)
}

func GetApproximateLookup() ApproximateLookup {
	return *approximateLookup
}

func Lookup(q string) (LookupResultWithSuggestion, error) {
	result, err := approximateLookup.LookupWithSuggestion(q)
	if err != nil {
		return LookupResultWithSuggestion{}, err
	}
	return result, nil
}

func loadConfigAndData() {
	paths := setup.DefaultPaths()
	configPathToml := filepath.Join(paths.ConfigDir, "config.toml")

	if err := config.LoadConfig(configPathToml); err != nil {
		log.Fatal("failed to load config:", err)
	}
	cfg := config.GetConfig()
	log.Printf("CONFIG TOML: %v", cfg)

	resourcesPath := cfg.Paths.Resources
	if !filepath.IsAbs(resourcesPath) {
		resourcesPath = startup.ResolvePath(resourcesPath)
	}

	if err := database.LoadDictionaries(resourcesPath); err != nil {
		log.Fatal("failed to load dictionaries:", err)
	}
	dictionaries = database.GetDictionaries()

	log.Printf("Loaded %d dictionaries\n", len(dictionaries))
}

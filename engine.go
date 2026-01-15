package dictionary

import (
	"log"
	"path/filepath"

	"github.com/dinhphu28/dictionary/database"
	"github.com/dinhphu28/dictionary/internal/config"
	"github.com/dinhphu28/dictionary/internal/setup"
)

var (
	dictionaries      []database.Dictionary
	approximateLookup *ApproximateLookup
)

func StartEngine() {
	paths := setup.DefaultPaths()
	configPath := filepath.Join(paths.ConfigDir, "config.json")
	resourcesPath := filepath.Join(paths.DataDir, "resources")

	if err := config.LoadConfig(configPath); err != nil {
		log.Fatal("failed to load config:", err)
	}
	globalConfig := config.GetGlobalConfig()

	if err := database.LoadDictionaries(resourcesPath); err != nil {
		log.Fatal("failed to load dictionaries:", err)
	}
	dictionaries = database.GetDictionaries()

	log.Printf("Loaded %d dictionaries\n", len(dictionaries))

	dictionaryLookup := NewDictionaryLookup(dictionaries, globalConfig)
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

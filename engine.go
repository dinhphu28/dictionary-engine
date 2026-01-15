package dictionary

import (
	"log"
	"path/filepath"

	"dinhphu28.com/dictionary/database"
	"dinhphu28.com/dictionary/dictionary"
	"dinhphu28.com/dictionary/internal/config"
	"dinhphu28.com/dictionary/internal/setup"
)

var (
	dictionaries      []database.Dictionary
	approximateLookup *dictionary.ApproximateLookup
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

	dictionaryLookup := dictionary.NewDictionaryLookup(dictionaries, globalConfig)
	approximateLookup = dictionary.NewApproximateLookup(*dictionaryLookup)
}

func Ready() bool {
	return len(dictionaries) > 0
}

func LoadedDictionaries() int {
	return len(dictionaries)
}

func GetApproximateLookup() dictionary.ApproximateLookup {
	return *approximateLookup
}

func Lookup(q string) (dictionary.LookupResultWithSuggestion, error) {
	result, err := approximateLookup.LookupWithSuggestion(q)
	if err != nil {
		return dictionary.LookupResultWithSuggestion{}, err
	}
	return result, nil
}

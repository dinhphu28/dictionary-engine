package main

import (
	"fmt"
	"log"
	"os"

	_ "modernc.org/sqlite"

	"github.com/dinhphu28/dictionary"
	"github.com/dinhphu28/dictionary/doctor"
	"github.com/dinhphu28/dictionary/internal/api"
	"github.com/dinhphu28/dictionary/native"
)

func main() {
	log.SetOutput(os.Stderr)

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "doctor":
			doctor.RunDoctor()
			return

		case "native":
			native.RunNative()
			return

		case "http":
			runHTTP()
			return
		}
	}

	// Default behavior
	native.RunNative()
}

func runHTTP() {
	fmt.Println("HTTP mode")

	dictionary.StartEngine()
	lookupHandlerV2 := api.NewLookupHandlerV2()
	router := api.NewRouter(*lookupHandlerV2)
	router.StartAPIServer()
}

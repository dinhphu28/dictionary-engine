package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	_ "modernc.org/sqlite"

	"dinhphu28.com/dictionary"
	"dinhphu28.com/dictionary/doctor"
	"dinhphu28.com/dictionary/internal/api"
	"dinhphu28.com/dictionary/internal/native"
)

func main() {
	log.SetOutput(os.Stderr)

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "doctor":
			doctor.RunDoctor()
			return

		case "native":
			runNative()
			return

		case "http":
			runHTTP()
			return
		}
	}

	// Default behavior
	runNative()
}

func runNative() {
	// 🔒 CRITICAL: never write logs to stdout
	log.SetOutput(os.Stderr)
	log.Println("Native host started")

	dictionary.StartEngine()
	// approximateLookup := dictionary.GetApproximateLookup()

	ready := dictionary.Ready()
	loadedDictionaries := dictionary.LoadedDictionaries()

	for {
		raw, err := native.ReadMessage()
		if err != nil {
			if err == io.EOF {
				log.Println("Chrome disconnected, exiting")
				return
			}
			log.Printf("read error: %v", err)
			return
		}

		var req native.Request
		if err := json.Unmarshal(raw, &req); err != nil {
			log.Printf("bad request: %v", err)
			_ = native.WriteMessage(native.Response{
				Type:    native.Error,
				Message: "invalid request",
			})
			continue
		}

		log.Printf("received: %+v", req)

		switch req.Type {

		case native.Ping:
			_ = native.WriteMessage(native.Response{
				Type:    native.Pong,
				Ready:   ready,
				Message: "Dictionaries loaded: " + strconv.Itoa(loadedDictionaries),
			})

		case native.Lookup:
			// 🔁 TEMP: fake result to prove Chrome works
			// result, err := approximateLookup.LookupWithSuggestion(req.Query)
			result, err := dictionary.Lookup(req.Query)
			if err != nil {
				_ = native.WriteMessage(native.Response{
					Type:    native.Error,
					Message: "lookup error: " + err.Error(),
				})
				continue
			}
			_ = native.WriteMessage(native.Response{
				Type:   native.Result,
				Ready:  true,
				Query:  req.Query,
				Result: result,
			})

		default:
			_ = native.WriteMessage(native.Response{
				Type:    native.Error,
				Message: "unknown message type",
			})
		}
	}
}

func runHTTP() {
	// your existing HTTP server logic
	fmt.Println("HTTP mode")

	dictionary.StartEngine()
	// approximateLookup := dictionary.GetApproximateLookup()
	//
	// lookupHandlerV2 := api.NewLookupHandlerV2(approximateLookup)
	lookupHandlerV2 := api.NewLookupHandlerV2()
	router := api.NewRouter(*lookupHandlerV2)
	router.StartAPIServer()
}

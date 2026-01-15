package native

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/dinhphu28/dictionary"
)

func RunNative() {
	// 🔒 CRITICAL: never write logs to stdout
	log.SetOutput(os.Stderr)
	log.Println("Native host started")

	dictionary.StartEngine()
	// approximateLookup := dictionary.GetApproximateLookup()

	ready := dictionary.Ready()
	loadedDictionaries := dictionary.LoadedDictionaries()

	for {
		raw, err := ReadMessage()
		if err != nil {
			if err == io.EOF {
				log.Println("Chrome disconnected, exiting")
				return
			}
			log.Printf("read error: %v", err)
			return
		}

		var req Request
		if err := json.Unmarshal(raw, &req); err != nil {
			log.Printf("bad request: %v", err)
			_ = WriteMessage(Response{
				Type:    Error,
				Message: "invalid request",
			})
			continue
		}

		log.Printf("received: %+v", req)

		switch req.Type {

		case Ping:
			_ = WriteMessage(Response{
				Type:    Pong,
				Ready:   ready,
				Message: "Dictionaries loaded: " + strconv.Itoa(loadedDictionaries),
			})

		case Lookup:
			// 🔁 TEMP: fake result to prove Chrome works
			// result, err := approximateLookup.LookupWithSuggestion(req.Query)
			result, err := dictionary.Lookup(req.Query)
			if err != nil {
				_ = WriteMessage(Response{
					Type:    Error,
					Message: "lookup error: " + err.Error(),
				})
				continue
			}
			_ = WriteMessage(Response{
				Type:   Result,
				Ready:  true,
				Query:  req.Query,
				Result: result,
			})

		default:
			_ = WriteMessage(Response{
				Type:    Error,
				Message: "unknown message type",
			})
		}
	}
}

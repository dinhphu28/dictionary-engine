package native

import "github.com/dinhphu28/dictionary"

type MessageRequestType int

const (
	UnknownRequestType MessageRequestType = iota
	Lookup
	Ping
)

type MessageResponseType int

const (
	UnknownResponseType MessageResponseType = iota
	Result
	Error
	Pong
	Loading
)

type Request struct {
	Type  MessageRequestType `json:"type"`  // "lookup" | "ping"
	Query string             `json:"query"` // word to look up
}

type Response struct {
	Type    MessageResponseType                   `json:"type"`  // "result" | "error" | "pong" | "loading"
	Query   string                                `json:"query"` // echoed word
	Ready   bool                                  `json:"ready"`
	Result  dictionary.LookupResultWithSuggestion `json:"result"`
	Message string                                `json:"message"`
}

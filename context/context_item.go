package context

import ai "github.com/sashabaranov/go-openai"

// Item provides context to the AI in a specific message
type Item interface {
	// Build returns a ChatCompletionMessage that can be used to send to the AI
	Build() (ai.ChatCompletionMessage, error)

    // Name returns the name of the context item
    Name() string
}

// GenericItem is a generic context item used for embedding in other structs
type ItemBase struct {
	// Name is the name of the context item
	Name string
	// Explanation is the explanation of the context item
	Explanation string
	// Priority is used to determine the order of the context items
    // Higher number -> earlier in the request
    // The position can _slightly_ affet how much attention the AI pays to the context item
	Priority int
	// Content is the content of the context item
	Content string
}

type ItemConfigBase struct {
    Enabled bool `toml:"enabled"`
    Priority int `toml:"priority"`
}

const CONTEXT_PREIX = "Context:"
const FORMAT_STRING = "%s %s\nExplanation: %s\nContent:\n%s"

package context

import (
	"fmt"
	"os"

	ai "github.com/sashabaranov/go-openai"
)

type Shell struct {
	ItemBase
}

type ShellConfig struct {
	ItemConfigBase
}

var DEFAULT_SHELL_CONFIG = ShellConfig{
	ItemConfigBase: ItemConfigBase{
		Enabled:  true,
		Priority: 11,
	},
}

func NewShell(cfg ShellConfig) (*Shell, error) {
	if cfg.Enabled {
		return &Shell{
			ItemBase: ItemBase{
				Name:        "What shell we are in",
				Explanation: "Content of the SHELL environment variable",
				Priority:    cfg.Priority,
				Content:     "",
			},
		}, nil
	} else {
		return &Shell{}, fmt.Errorf("Shell is disabled")
	}
}

func (s Shell) Name() string {
	return s.ItemBase.Name
}

func (s *Shell) Build() (msg ai.ChatCompletionMessage, err error) {
	s.Content = os.Getenv("SHELL")

	msg = ai.ChatCompletionMessage{
		Role:    ai.ChatMessageRoleUser,
		Content: fmt.Sprintf(FORMAT_STRING, CONTEXT_PREIX, s.Name(), s.Explanation, s.Content),
	}

	return
}

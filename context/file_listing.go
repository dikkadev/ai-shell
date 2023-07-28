package context

import (
	"fmt"
	"os"
	"time"

	ai "github.com/sashabaranov/go-openai"
)

type FileListing struct {
	ItemBase
}

type FileListingConfig struct {
	ItemConfigBase
}

var DEFAULT_FILE_LISTING_CONFIG = FileListingConfig{
	ItemConfigBase: ItemConfigBase{
		Enabled:  true,
		Priority: 10,
	},
}

// Returns a new FileListing context item with empty content
func NewFileListing(cfg FileListingConfig) (*FileListing, error) {
	if cfg.Enabled {
		return &FileListing{
			ItemBase: ItemBase{
				Name:        "Current Directory File Listing",
				Explanation: "Formatted list of files in the cwd",
				Priority:    cfg.Priority,
				Content:     "",
			},
		}, nil
	} else {
		return &FileListing{}, fmt.Errorf("FileListing is disabled")
	}
}

func (f FileListing) Name() string {
    return f.ItemBase.Name
}

// Fills content of FileListing and returns message
func (f *FileListing) Build() (msg ai.ChatCompletionMessage, err error) {
	content := ""

	files, err := os.ReadDir(".")
	if err != nil {
		return msg, fmt.Errorf("Error reading directory: %s", err)
	}

	for _, f := range files {
		fileInfo, err := f.Info()
		if err != nil {
			continue
		}

		content += fmt.Sprintf("%s %dB %s %s\n", fileInfo.Mode(), fileInfo.Size(), fileInfo.ModTime().Format(time.RFC3339), f.Name())
	}

    f.Content = content

	msg = ai.ChatCompletionMessage{
		Role:    ai.ChatMessageRoleUser,
		Content: fmt.Sprintf(FORMAT_STRING, CONTEXT_PREIX, f.Name(), f.Explanation, f.Content),
	}

	return
}

package chat

import (
	ctx2 "context"
	"fmt"
	"time"

	"github.com/i582/cfmt/cmd/cfmt"
	ai "github.com/sashabaranov/go-openai"
	"github.com/sett17/ai-shell/cli"
	"github.com/sett17/ai-shell/context"
)

type Chat struct {
	Context     []context.Item
	Instruction string

	client      *ai.Client
    messages    []ai.ChatCompletionMessage
}

func New(instruction string, token string) Chat {
	return Chat{
		Context:     []context.Item{},
		Instruction: instruction,
		client:      ai.NewClient(token),
	}
}

func (c *Chat) AddContext(item context.Item, err error) {
	if err != nil {
		cli.Warning(err.Error())
		return
	}
	c.Context = append(c.Context, item)
}

func (c *Chat) Execute() (string, error) {
	c.messages = make([]ai.ChatCompletionMessage, 2)
	c.messages[0] = SYSTEM_MESSAGE
	c.messages[1] = ai.ChatCompletionMessage{
		Role:    ai.ChatMessageRoleUser,
		Content: fmt.Sprintf("Instruction: %s", c.Instruction),
	}

	for _, item := range c.Context {
		cli.Dbg(fmt.Sprintf("Building context item '%s'", item.Name()))
		start := time.Now()
		msg, err := item.Build()
		cli.Dbg(fmt.Sprintf("\ttook %s", time.Since(start)))
		if err != nil {
			return "", err
		}
		c.messages = append(c.messages, msg)
	}

	cfmt.Printf("{{|}}::green Suggested command ")

	sw := cli.NewStopWatch(time.Millisecond)
	sw.Start()

	resp, err := c.client.CreateChatCompletion(ctx2.Background(), ai.ChatCompletionRequest{
		Model:       ai.GPT3Dot5Turbo16K,
		Messages:    c.messages,
		Temperature: 0.69,
		N:           1,
	})
	if err != nil {
		return "", err
	}
	sw.Stop()
    c.messages = append(c.messages, resp.Choices[0].Message)

	fmt.Printf("\033[100m%s\033[0m\n", resp.Choices[0].Message.Content)

	return resp.Choices[0].Message.Content, nil
}

func (c *Chat) Revise(revision string) (string, error) {
    c.messages = append(c.messages, ai.ChatCompletionMessage{
        Role:    ai.ChatMessageRoleUser,
        Content: revision,
    })

	cfmt.Printf("{{|}}::green Suggested command ")

	sw := cli.NewStopWatch(time.Millisecond)
	sw.Start()

    resp, err := c.client.CreateChatCompletion(ctx2.Background(), ai.ChatCompletionRequest{
        Model:       ai.GPT3Dot5Turbo16K,
        Messages:    c.messages,
        Temperature: 0.69,
        N:           1,
    })
    if err != nil {
        return "", err
    }

    sw.Stop()
    c.messages = append(c.messages, resp.Choices[0].Message)

	fmt.Printf("\033[100m%s\033[0m\n", resp.Choices[0].Message.Content)

    return resp.Choices[0].Message.Content, nil
}

var SYSTEM_MESSAGE = ai.ChatCompletionMessage{
	Role: ai.ChatMessageRoleSystem,
	Content: `As an AI, you are programmed to generate precise console commands based on user instructions for any shell type. These instructions are your main focus. Along with these instructions, you may receive 'context' c.messages containing various details such as the file listing of the current directory, specific contents of files, command history, and the shell identifier. This context should be used as needed to create the most suitable console command, but remember, it may or may not be helpful in generating the command. The user's instruction should always be your primary guide.
In every situation, your output should be exclusively the console command, with no added text or explanations. Reiterating for clarity, your output, crafted using any relevant user instructions and possibly the provided context, should contain no other information apart from the command itself. Even when the instructions or context might be ambiguous, you are to give it your best try without asking for any additional clarification or providing any explanation - just the console command.`,
}

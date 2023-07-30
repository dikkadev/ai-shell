package cli

import (
	"os"

	"github.com/i582/cfmt/cmd/cfmt"
)

func help(_ string) error {
	cfmt.Printf(`
{{Usage:}}::underline ai-shell [options] instruction

`)

	cfmt.Printf(`
This is a command-line tool that uses OpenAI's GPT-3.5-turbo to generate precise console commands based on your instructions. It takes in your commands, gathers some context, and lets an AI model figure out the best course of action.
`)

	cfmt.Printf(`
{{It is recommended to create a symlink called ai to the ai-shell executable}}::bgYellow|black
`)

	cfmt.Printf(`

{{Options:}}::underline
`)
	for _, arg := range ProgArgs {
		cfmt.Printf("  {{-%s}}::green, {{--%s}}::green\t{{%s}}::gray\n", arg.Short, arg.Long, arg.Help)
	}
	cfmt.Printf("  {{-%s}}::green, {{--%s}}::green\t{{%s}}::gray\n", "h", "help", "Show this help message")
	cfmt.Printf("  {{-%s}}::green, {{--%s}}::green\t{{%s}}::gray\n", VersionProgArg.Short, VersionProgArg.Long, VersionProgArg.Help)

	os.Exit(0)
	return nil
}

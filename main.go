package main

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/sett17/ai-shell/chat"
	"github.com/sett17/ai-shell/cli"
	"github.com/sett17/ai-shell/config"
	"github.com/sett17/ai-shell/context"
	"github.com/sett17/ai-shell/global"
)

func main() {
	cli.ParseForHelp(os.Args[1:])
	cli.ParseForVersion(os.Args[1:])

	cfg, err := config.Load()
	if err != nil {
		cli.Error((fmt.Errorf("Error while decoding config (try deleting it)\n%w", err)), true)
	}
	global.Cfg = cfg

	instruction := cli.Parse(os.Args[1:])

	key := os.Getenv("OPENAI_API_KEY")
	if key == "" {
		cli.Error(fmt.Errorf("OPENAI_API_KEY environment variable not set"), true)
	}

	chat := chat.New(instruction, key)

	chat.AddContext(context.NewFileListing(global.Cfg.FileListingConfig))
    chat.AddContext(context.NewShell(global.Cfg.ShellConfig))

	isRevision := false
	revision := ""
    loop:
	for {
        cmd := ""
		if isRevision {
			cmd, err = chat.Revise(revision)
			if err != nil {
				cli.Error(err, true)
			}
		} else {
			cmd, err = chat.Execute()
			if err != nil {
				cli.Error(err, true)
			}
		}

		answers := struct {
			Task string
		}{}

		err = survey.Ask(
			[]*survey.Question{
				{
					Name: "task",
					Prompt: &survey.Select{
						Message: "What to do now?",
						Options: []string{"Execute", "Edit & Execute", "Revise"},
						Default: "Execute",
					},
				},
			},
			&answers)
		if err != nil {
			cli.Error(err, true)
		}

		switch answers.Task {
		case "Execute":
            err := ExecuteCommand(cmd)
            if err != nil {
                cli.Error(err, true)
            }
            break loop
		case "Edit & Execute":
            cli.Warning("Not implemented yet, sorry")
            break loop
		case "Revise":
			isRevision = true
			answers := struct {
				Revision string
			}{}
			err = survey.Ask(
				[]*survey.Question{
					{
						Name:   "revision",
						Prompt: &survey.Input{Message: "Instruction: "},
					},
				},
				&answers)
            if err != nil {
                cli.Error(err, true)
            }
            revision = answers.Revision
		}

	}
}

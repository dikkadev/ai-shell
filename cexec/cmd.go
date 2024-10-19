package cexec

import (
	"os"
	"os/exec"

	"github.com/dikkadev/ai-shell/cli"
)

type CmdExecutor struct {
    file string
}

func (c *CmdExecutor) Create(content string) error {
    f, err := os.CreateTemp("", "ai-shell*.bat")
    if err != nil {
        return err
    }
    cli.Dbg("Created temp file: " + f.Name())
    defer f.Close()
    c.file = f.Name()

    _, err = f.WriteString("@echo off\n" + content + "\n@echo on\n")
    return err
}

func (c *CmdExecutor) Execute() error {
    _, err := os.Stat(c.file)
    if os.IsNotExist(err) {
        return err
    }

    defer func() {
        os.Remove(c.file)
        cli.Dbg("Removed temp file: " + c.file)
    }()
    cmd := exec.Command("cmd", "/c", c.file)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    cli.Dbg("Executing script file")
    return cmd.Run()
}


func (c *CmdExecutor) Edit() error {
    editor := os.Getenv("EDITOR")
    if editor == "" {
        editor = "notepad.exe"
    }
    cli.Dbg("Editing script file with " + editor)
    return exec.Command(editor, c.file).Run()
}


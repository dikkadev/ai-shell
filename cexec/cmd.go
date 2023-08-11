package cexec

import (
	"os"
	"os/exec"
)


type CmdExecutor struct {
    file string
}

func (c *CmdExecutor) Create(content string) error {
    f, err := os.CreateTemp("", "ai-shell*.bat")
    if err != nil {
        return err
    }
    defer f.Close()
    c.file = f.Name()

    _, err = f.WriteString("@echo off\n" + content + "\n@echo on\n")
    return err
}

func (c *CmdExecutor) Execute() error {
    defer os.Remove(c.file)
    cmd := exec.Command("cmd", "/c", c.file)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    return cmd.Run()
}

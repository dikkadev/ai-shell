package cexec

import (
	"os"
	"os/exec"
)


type BashExecutor struct {
    file string
}

func (b *BashExecutor) Create(content string) error {
    f, err := os.CreateTemp("", "ai-shell*.sh")
    if err != nil {
        return err
    }
    defer f.Close()
    b.file = f.Name()

    _, err = f.WriteString("shopt -s expand_aliases\n" + content)
    return err
}

func (b *BashExecutor) Execute() error {
    defer os.Remove(b.file)
    cmd := exec.Command("/bin/bash", "-e", b.file)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    return cmd.Run()
}

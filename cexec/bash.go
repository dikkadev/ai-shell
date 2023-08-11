package cexec

import (
	"os"
	"os/exec"

	"github.com/sett17/ai-shell/cli"
)


type BashExecutor struct {
    file string
}

func (b *BashExecutor) Create(content string) error {
    f, err := os.CreateTemp("", "ai-shell*.sh")
    if err != nil {
        return err
    }
    cli.Dbg("Created temp file: " + f.Name())
    defer f.Close()
    b.file = f.Name()

    _, err = f.WriteString("shopt -s expand_aliases\n" + content)
    return err
}

func (b *BashExecutor) Execute() error {
    defer func() {
        os.Remove(b.file)
        cli.Dbg("Removed temp file: " + b.file)
    }()
    cmd := exec.Command("/bin/bash", "-e", b.file)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    cli.Dbg("Executing script file")
    return cmd.Run()
}

func (b *BashExecutor) Edit() error {
    editor := os.Getenv("EDITOR")
    if editor == "" {
        editor = "vim"
    }
    cli.Dbg("Editing script file with " + editor)
    return exec.Command(editor, b.file).Run()
}

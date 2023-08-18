package cexec

import (
	"os"
	"os/exec"

	"github.com/sett17/ai-shell/cli"
)

type PwshExecutor struct {
    file string
}

func (p *PwshExecutor) Create(content string) error {
    f, err := os.CreateTemp("", "ai-shell*.ps1")
    if err != nil {
        return err
    }
    cli.Dbg("Created temp file: " + f.Name())
    defer f.Close()
    p.file = f.Name()

    _, err = f.WriteString(content)
    return err
}

func (p *PwshExecutor) Execute() error {
    _, err := os.Stat(p.file)
    if os.IsNotExist(err) {
        return err
    }

    defer func() {
        os.Remove(p.file)
        cli.Dbg("Removed temp file: " + p.file)
    }()
    cmd := exec.Command("pwsh", "-File", p.file)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    cli.Dbg("Executing script file")
    return cmd.Run()
}

func (p *PwshExecutor) Edit() error {
    editor := os.Getenv("EDITOR")
    if editor == "" {
        editor = "notepad.exe"
    }
    cli.Dbg("Editing script file with " + editor)
    return exec.Command(editor, p.file).Run()
}

package cexec

import (
	"os"
	"runtime"
)

type Executor interface {
    Create(content string) error
    Execute() error
}

func ChooseExecutor() Executor {
	if runtime.GOOS != "windows" {
        return &BashExecutor{}
	} else {
		if shell := os.Getenv("ComSpec"); shell != "" {
            return &CmdExecutor{}
		} else {
            //TODO powershell
            return &CmdExecutor{}
		}
	}
}

func ExecuteCommand(cmd string, exe Executor) error {
    err := exe.Create(cmd)
    if err != nil {
        return err
    }
    return exe.Execute()
}

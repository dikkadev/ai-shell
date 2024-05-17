package cexec

import (
	"os"
	"runtime"
	"strings"

	"github.com/sett17/ai-shell/cli"
	"github.com/shirou/gopsutil/process"
)

type Executor interface {
	Create(content string) error
	Execute() error
	Edit() error
	AddToHistory(string) error
}

func ChooseExecutor() Executor {
	if runtime.GOOS != "windows" {
		cli.Dbg("Chose bash executor")
		return &BashExecutor{}
	} else {
		p, err := process.NewProcess(int32(os.Getpid()))
		if err != nil {
			panic(err)
		}

		ppid, err := p.Ppid()
		if err != nil {
			panic(err)
		}
		pp, err := process.NewProcess(ppid)
		if err != nil {
			panic(err)
		}

		name, err := pp.Name()
		if err != nil {
			panic(err)
		}

		if strings.Contains(strings.ToLower(name), "pwsh") {
			cli.Dbg("Chose powershell executor")
			return &PwshExecutor{}
		} else {
			cli.Dbg("Chose cmd executor")
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

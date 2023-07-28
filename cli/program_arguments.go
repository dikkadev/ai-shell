package cli

import (
	"os"

	"github.com/i582/cfmt/cmd/cfmt"
	"github.com/sett17/ai-shell/global"
)

type progArgFunction func(after string) error

type ProgArg struct {
	Name    string
	Short   string
	Long    string
	Help    string
	HasData bool

	Func progArgFunction
}

var DebugProgArg = ProgArg{
	Name:  "Debug",
	Short: "d",
	Long:  "debug",
	Help:  "Enabled Debug output",
	Func: func(_ string) error {
		global.Cfg.Debug = true
		Dbg("Debug output enabled")
        return nil
	},
}

var HelpProgArg = ProgArg{
	Name:    "Help",
	Short:   "h",
	Long:    "help",
	Help:    "",
	HasData: false,
	Func:    help,
}

var VersionProgArg = ProgArg{
	Name:  "Version",
	Short: "v",
	Long:  "version",
	Help:  "Show the version of this program",
	HasData: false,
	Func: func(_ string) error {
		cfmt.Printf("AI-shell version %s\n", global.Version)
		os.Exit(0)
		return nil
	},
}

var ProgArgs = []ProgArg{
}

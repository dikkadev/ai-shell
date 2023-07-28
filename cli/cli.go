package cli

import (
	"strings"

	"github.com/alexeyco/simpletable"
	"github.com/i582/cfmt/cmd/cfmt"
)

func Parse(args []string) (instruction string) {
    for i := 0; i < len(args); i++ {
        arg := args[i]
        if strings.HasPrefix(arg, "-") || strings.HasPrefix(arg, "--") {
            arg = strings.TrimLeft(arg, "-")
            for _, pArg := range ProgArgs {
                if pArg.Short == arg || pArg.Long == arg {
                    argData := ""
                    if pArg.HasData {
                        if i+1 < len(args) {
                            argData = args[i+1]
                            i++
                        }
                    }
                    defer pArg.Func(argData)
                }
            }
        } else {
            instruction += arg + " "
        }
    }
    return
}

func ParseForDebug(args []string) {
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") || strings.HasPrefix(arg, "--") {
			arg = strings.TrimLeft(arg, "-")
			if arg == DebugProgArg.Short || arg == DebugProgArg.Long {
				DebugProgArg.Func("")
				return
			}
		}
	}
}

func ParseForHelp(args []string) {
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") || strings.HasPrefix(arg, "--") {
			arg = strings.TrimLeft(arg, "-")
			if arg == HelpProgArg.Short || arg == HelpProgArg.Long {
				HelpProgArg.Func("")
				return
			}
		}
	}
}

func ParseForVersion(args []string) {
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") || strings.HasPrefix(arg, "--") {
			arg = strings.TrimLeft(arg, "-")
			if arg == VersionProgArg.Short || arg == VersionProgArg.Long {
                VersionProgArg.Func("")
				return
			}
		}
	}
}

var roundTableStyle = &simpletable.Style{
	Border: &simpletable.BorderStyle{
		TopLeft:            "╭",
		Top:                "─",
		TopRight:           "╮",
		Right:              "│",
		BottomRight:        "╯",
		Bottom:             "─",
		BottomLeft:         "╰",
		Left:               "│",
		TopIntersection:    "┬",
		BottomIntersection: "┴",
	},
	Divider: &simpletable.DividerStyle{
		Left:         "├",
		Center:       "─",
		Right:        "┤",
		Intersection: "┼",
	},
	Cell: "│",
}

var Logo = cfmt.Sprint(`
             ______     __                             
            /\  __ \   /\ \                            
            \ \  __ \  \ \ \                           
             \ \_\ \_\  \ \_\                          
              \/_/\/_/   \/_/                          
 ______     __  __     ______     __         __        
/\  ___\   /\ \_\ \   /\  ___\   /\ \       /\ \       
\ \___  \  \ \  __ \  \ \  __\   \ \ \____  \ \ \____  
 \/\_____\  \ \_\ \_\  \ \_____\  \ \_____\  \ \_____\ 
  \/_____/   \/_/\/_/   \/_____/   \/_____/   \/_____/ 

  by Sett   {{https://github.com/Sett17/ai-shell}}::green
`)

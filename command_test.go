package main

import (
	"reflect"
	"testing"
)

func TestSplitCommand(t *testing.T) {
	tests := []struct {
		name     string
		command  string
		wantCmd  string
		wantArgs []string
	}{
		// Basic Tests
		{
			name:     "NoArgs",
			command:  "ls",
			wantCmd:  "ls",
			wantArgs: []string{},
		},
		{
			name:     "SingleArg",
			command:  "ls -l",
			wantCmd:  "ls",
			wantArgs: []string{"-l"},
		},
		{
			name:     "MultipleArgs",
			command:  "tar -xzf file.tar.gz",
			wantCmd:  "tar",
			wantArgs: []string{"-xzf", "file.tar.gz"},
		},

		// Quotes
		{
			name:     "SingleQuotedArgs",
			command:  "cd 'my dir'",
			wantCmd:  "cd",
			wantArgs: []string{"my dir"},
		},
		{
			name:     "SingleQuotedArgsWithSpaces",
			command:  "rm 'my dir' 'my other dir'",
			wantCmd:  "rm",
			wantArgs: []string{"my dir", "my other dir"},
		},
		{
			name:     "SingleQuotedArgsWithSpacesAndEscapedQuotes",
			command:  "rm 'my dir' 'my other dir' 'my \"other\" other dir'",
			wantCmd:  "rm",
			wantArgs: []string{"my dir", "my other dir", "my \"other\" other dir"},
		},
		{
			name:     "DoubleQuotedArgs",
			command:  "cd \"my dir\"",
			wantCmd:  "cd",
			wantArgs: []string{"my dir"},
		},
		{
			name:     "DoubleQuotedArgsWithSpaces",
			command:  "rm \"my dir\" \"my other dir\"",
			wantCmd:  "rm",
			wantArgs: []string{"my dir", "my other dir"},
		},
		{
			name:     "DoubleQuotedArgsWithSpacesAndEscapedQuotes",
			command:  "rm \"my dir\" \"my other dir\" \"my \\\"other\\\" other dir\"",
			wantCmd:  "rm",
			wantArgs: []string{"my dir", "my other dir", "my \"other\" other dir"},
		},

		// Spaces
		{
			name:     "LeadingSpaces",
			command:  "   ls -l",
			wantCmd:  "ls",
			wantArgs: []string{"-l"},
		},
		{
			name:     "TrailingSpaces",
			command:  "ls -l   ",
			wantCmd:  "ls",
			wantArgs: []string{"-l"},
		},
		{
			name:     "ConsecutiveSpaces",
			command:  "ls   -l",
			wantCmd:  "ls",
			wantArgs: []string{"-l"},
		},

		// Edge Cases
		{
			name:     "EmptyCommand",
			command:  "",
			wantCmd:  "",
			wantArgs: []string{},
		},
		{
			name:     "ArgsWithEquals",
			command:  "command --arg=value",
			wantCmd:  "command",
			wantArgs: []string{"--arg=value"},
		},
		{
			name:     "SpecialCharacterArgs",
			command:  "command &|",
			wantCmd:  "command",
			wantArgs: []string{"&|"},
		},
		{
			name:     "MixedQuotesInArgs",
			command:  "command 'arg\"arg'",
			wantCmd:  "command",
			wantArgs: []string{"arg\"arg"},
		},
		{
			name:     "EmptyArgs",
			command:  "command ''",
			wantCmd:  "command",
			wantArgs: []string{""},
		},

		// Escaped Characters
		{
			name:     "EscapedSpaces",
			command:  `command arg\\ arg`,
			wantCmd:  "command",
			wantArgs: []string{`arg\ arg`},
		},
		{
			name:     "EscapedQuotes",
			command:  "command \\\"arg\\\"",
			wantCmd:  "command",
			wantArgs: []string{"\"arg\""},
		},
		{
			name:     "EscapedBackslashes",
			command:  "command arg\\\\",
			wantCmd:  "command",
			wantArgs: []string{"arg\\"},
		},
		{
			name:     "MixedQuotesAndEscapes",
			command:  "command \"arg'arg\\\"arg\"",
			wantCmd:  "command",
			wantArgs: []string{"arg'arg\"arg"},
		},
		{
			name:     "EscapedCharacters",
			command:  `echo -e "\e[31m\e[43m\e[1mrainbow\e[0m"`,
			wantCmd:  "echo",
			wantArgs: []string{"-e", "\\e[31m\\e[43m\\e[1mrainbow\\e[0m"},
		},
		{
			name:     "BackslashEscapedQuotes",
			command:  `command "arg\"arg"`,
			wantCmd:  "command",
			wantArgs: []string{"arg\"arg"},
		},
		{
			name:     "BackslashEscapedSingleQuotes",
			command:  `command 'arg\'arg'`,
			wantCmd:  "command",
			wantArgs: []string{"arg'arg"},
		},
		{
			name:     "BackslashEscapedBackslashes",
			command:  `command "arg\\arg"`,
			wantCmd:  "command",
			wantArgs: []string{"arg\\arg"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCmd, gotArgs := splitCommand(tt.command)
			if gotCmd != tt.wantCmd {
				t.Errorf("splitCommand() gotCmd = %+v, wantCmd = %+v", gotCmd, tt.wantCmd)
			}
			if !reflect.DeepEqual(gotArgs, tt.wantArgs) {
				t.Errorf("splitCommand() gotArgs = %+v, wantArgs = %+v", gotArgs, tt.wantArgs)
			}
		})
	}
}

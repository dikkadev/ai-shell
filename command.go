package main

import (
	"os"
	"os/exec"
	"strings"
	"unicode"
)

func ExecuteCommand(command string) error {
	file, args := splitCommand(command)
	cmd := exec.Command(file, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func EditCommand(cmd *string) (err error) {
	return nil
}

// Written by AI to pass the test suite
func splitCommand(input string) (string, []string) {
	// Trim the input
	input = strings.TrimSpace(input)

	// Prepare variables for parsing
	args := []string{}
	var curArg strings.Builder
	inDoubleQuotes, inSingleQuotes, escaped := false, false, false

	// Iterate over each character in the string
	for i, r := range input {
		switch {
		// If it's a backslash, check the next character
		case r == '\\':
			// If it's not the last character and the next character is a quote or a backslash, enable escaping
			if i+1 < len(input) && (input[i+1] == '"' || input[i+1] == '\'' || input[i+1] == '\\') {
				escaped = true
			} else if i+1 < len(input) && unicode.IsSpace(rune(input[i+1])) {
				// If it's followed by a space, treat it as a literal character
				curArg.WriteRune(r)
			} else {
				// Else, treat it as a normal character
				curArg.WriteRune(r)
			}
		// If it's a quote and not escaped, toggle quote tracking
		case r == '"' && !escaped && !inSingleQuotes:
			inDoubleQuotes = !inDoubleQuotes
			if inDoubleQuotes {
				curArg.Reset()
			} else {
				args = append(args, curArg.String())
				curArg.Reset()
			}
		case r == '\'' && !escaped && !inDoubleQuotes:
			inSingleQuotes = !inSingleQuotes
			if inSingleQuotes {
				curArg.Reset()
			} else {
				args = append(args, curArg.String())
				curArg.Reset()
			}
		// If it's a space and we're not inside quotes or escaped, append the current arg and start a new one
		case unicode.IsSpace(r) && !inSingleQuotes && !inDoubleQuotes && !escaped:
			if curArg.Len() > 0 {
				args = append(args, curArg.String())
				curArg.Reset()
			}
		// If it's any other character, just add it to the current arg
		default:
			curArg.WriteRune(r)
			escaped = false
		}
	}

	// Don't forget to append the last argument
	if curArg.Len() > 0 {
		args = append(args, curArg.String())
	}

	if len(args) > 0 {
		return args[0], args[1:]
	}

	return "", []string{}
}

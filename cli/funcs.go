package cli

import (
	"fmt"
	"os"

	"github.com/i582/cfmt/cmd/cfmt"
)

func help(_ string) error {
	cfmt.Printf(`
{{Usage:}}::underline dev [OPTIONS] OPERATIONS+ARGUMENT..
`)
	fmt.Println(Logo)
	cfmt.Printf(`
Tired of the weird quirks of make? Annoyed of making typos in long chained
commands, or getting to them via reverse-i-search? Well, here is a solution
that comes as just an executable for each platform and with an extensive
help command.

{{Operation Options:}}::underline
 These options can be used as letters in the Devfile

`)

	cfmt.Printf(`

{{Options:}}::underline
`)
	for _, arg := range ProgArgs {
		cfmt.Printf("  {{-%s}}::purple, {{--%s}}::purple\t{{%s}}::gray\n", arg.Short, arg.Long, arg.Help)
	}
	cfmt.Printf("  {{-%s}}::purple, {{--%s}}::purple\t{{%s}}::gray\n", "h", "help", "Show this help message")
	cfmt.Printf("  {{-%s}}::purple, {{--%s}}::purple\t{{%s}}::gray\n", VersionProgArg.Short, VersionProgArg.Long, VersionProgArg.Help)

	cfmt.Printf(`

{{Operations:}}::underline
  {{OPERATIONS+ARGUMENT...}}::purple  {{Which operation to execute.}}::gray
                          {{Arguments should be in same order as in Devfile}}::gray
`)
	os.Exit(0)
    return nil
}

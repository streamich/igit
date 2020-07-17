package igit

import (
	"flag"
	"fmt"
	"os"
)

var helpMessage string = `Usage: igit [OPTIONS]

Interactie git.

Semantic commits through an interactive interface.

Options:
  --scope Wheter to ask for commit scope.
  --help  Show this message and exit.
`

// CliFlags holds parsed all flags user can provide using CLI arguments.
type CliFlags struct {
	help  bool
	scope bool
}

// ParseFlags parses use provided CLI arguments.
func ParseFlags() *CliFlags {
	var scope = flag.Bool("scope", false, "Wheter to ask for commit scope.")
	var help = flag.Bool("help", false, "Show this message and exit.")

	flag.Parse()

	if *help {
		fmt.Print(helpMessage)
		os.Exit(0)
	}

	flags := &CliFlags{
		help:  *help,
		scope: *scope,
	}
	return flags
}

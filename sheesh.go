package main

import (
	"github.com/ariary/quicli/pkg/quicli"
	"github.com/ariary/sheesh/pkg/sheesh"
)

func main() {
	cli := quicli.Cli{
		Usage:       "sheesh [command] [flags]",
		Description: "Better-than-an-alias generator",
		Flags: quicli.Flags{
			{Name: "file", Default: ".sheesh.yml", Description: "sheesh configuration file", ForSubcommand: quicli.SubcommandSet{"create", "addflag", "setscript"}},
		},
		Function: sheesh.Generate,
		Subcommands: quicli.Subcommands{
			{Name: "create", Description: "Create sheesh command", Function: sheesh.Create},
			{Name: "addflag", Description: "Add a flag to existing command", Function: sheesh.AddFlag},
			{Name: "setscript", Description: "Set script of an existing command", Function: sheesh.SetScript},
		},
	}
	cli.RunWithSubcommand()
}

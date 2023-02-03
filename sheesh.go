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
			{Name: "file", Default: ".sheesh.yml", Description: "sheesh configuration file", ForSubcommand: quicli.SubcommandSet{"setcommand", "setflag", "setscript"}},
			{Name: "command", Default: "", Description: "sheesh command to target", ForSubcommand: quicli.SubcommandSet{"setcommand", "setflag", "setscript"}},
			{Name: "remove", Default: false, Description: "remove object", ForSubcommand: quicli.SubcommandSet{"setcommand", "setflag", "setscript"}},
			{Name: "script", Default: "", Description: "scritp content", ForSubcommand: quicli.SubcommandSet{"setscript"}},
		},
		Function: sheesh.Generate,
		Subcommands: quicli.Subcommands{
			{Name: "setcommand", Description: "Set sheesh command", Function: sheesh.SetCommand},
			{Name: "setflag", Description: "Set a flag to existing command", Function: sheesh.SetFlag},
			{Name: "setscript", Description: "Set script of an existing command", Function: sheesh.SetScript},
		},
	}
	cli.RunWithSubcommand()
}

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"gopkg.in/yaml.v3"

	"github.com/ariary/quicli/pkg/quicli"
)

type Flag struct {
	Name     string
	Argument bool
}

type Command struct {
	Name   string
	Flags  []Flag
	Script string
}

type Config struct {
	Commands []Command `yaml:"commands"`
}

// Create: create a sheesh command file
func Create(cfg quicli.Config) {
	fmt.Println("toto")
}

// Generate: generate output of a sheesh command
func Generate(cfg quicli.Config) {
	file := cfg.GetStringFlag("file")

	//read file
	yfile, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	//unmarshall yaml
	data := Config{}
	err = yaml.Unmarshal(yfile, &data)
	if err != nil {
		panic(err)
	}

	if len(data.Commands) == 0 {
		fmt.Println("No command found in yaml", file)
		os.Exit(1)
	}

	var output string
	for i := 0; i < len(data.Commands); i++ {
		output += processCommand(data.Commands[i])
		fmt.Println(output)
	}
}

// AddFlag: add a flag of a sheesh command
func AddFlag(cfg quicli.Config) {
	fmt.Println("toto")
}

// SetScript: set the script of of a sheesh command
func SetScript(cfg quicli.Config) {
	fmt.Println("toto")
}

func main() {
	cli := quicli.Cli{
		Usage:       "sheesh [command] [flags]",
		Description: "Better-than-an-alias generator",
		Flags: quicli.Flags{
			{Name: "file", Default: ".sheesh.yml", Description: "sheesh configuration file", ForSubcommand: quicli.SubcommandSet{"create", "addflag", "setscript"}},
		},
		Function: Generate,
		Subcommands: quicli.Subcommands{
			{Name: "create", Description: "Create sheesh command", Function: Create},
			{Name: "addflag", Description: "Add a flag to existing command", Function: AddFlag},
			{Name: "setscript", Description: "Set script of an existing command", Function: SetScript},
		},
	}
	cli.RunWithSubcommand()
}

func processCommand(c Command) (out string) {
	// fmt.Println(c.Script)

	// last stage: command + completion
	var outputBuffer bytes.Buffer
	output := Output{"toto", "function", "completion"}

	ot, err := template.New("output").Parse(outputTpl)
	if err != nil {
		panic(err)
	}
	err = ot.Execute(&outputBuffer, output)
	if err != nil {
		panic(err)
	}
	out = outputBuffer.String()
	return out
}

// template
type Output struct {
	CommandName       string
	Command           string
	CommandCompletion string
}

var outputTpl = `{{ .Command}}
{{ .CommandCompletion}}
compdef _{{ .CommandName}} {{ .CommandName}}
`

var commandTpl = `
ddf
`

var commandCompletionTpl = ``

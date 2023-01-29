package sheesh

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ariary/quicli/pkg/quicli"
	"gopkg.in/yaml.v3"
)

type Flag struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	NoArgs      bool     `yaml:"noarg"`
	File        bool     `yaml:"file"`
	Predefined  []string `yaml:"predefined"`
}

type Command struct {
	Name   string `yaml:"name"`
	Flags  []Flag `yaml:"flags"`
	Script string `yaml:"script"`
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
		output += ProcessCommand(data.Commands[i])
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

// ProcessCommand: take the command and return all the sheesh output (function + completion)
func ProcessCommand(c Command) (out string) {
	// fmt.Println(c.Script)

	commandCompletion := MarshallCompletion(c)

	out = MarshallOutput(c.Name, "titi", commandCompletion)

	return out
}

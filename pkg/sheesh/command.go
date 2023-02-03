package sheesh

import (
	"bytes"
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

func replaceCommand(commands []Command, c Command) (nCommands []Command) {
	nCommands = commands
	commandName := c.Name
	for i := 0; i < len(commands); i++ {
		if commands[i].Name == commandName {
			nCommands[i] = c
		}
	}
	return nCommands
}

// getCommandsFromFile: parse yaml file to retrieve commands list
func getCommandsFromFile(filename string) (commands Config, err error) {
	yfile, err := ioutil.ReadFile(filename)
	if err != nil {
		return commands, err
	}
	commands = Config{}
	err = yaml.Unmarshal(yfile, &commands)
	if err != nil {
		return commands, err
	}
	return commands, err
}

// getCommandByName: retrieve specified command
func getCommandByName(commands []Command, commandName string) (c Command) {
	for i := 0; i < len(commands); i++ {
		if commands[i].Name == commandName {
			return commands[i]
		}
	}
	return c
}

// getCommandByNameFromFile: retrieve specified command in yaml file
func getCommandByNameFromFile(filename string, commandName string) (c Command) {
	commands, err := getCommandsFromFile(filename)
	if err != nil {
		panic(err)
	}
	c = getCommandByName(commands.Commands, commandName)
	return c
}

// getCommandsFromFile: parse yaml file to retrieve commands list
func getCommandByNameFromCfg(cfg quicli.Config) (c Command) {
	cName := cfg.GetStringFlag("command")
	if cName == "" {
		fmt.Println("No command name provided. Exit. use --command")
		os.Exit(1)
	}
	file := cfg.GetStringFlag("file")
	command := getCommandByNameFromFile(file, cName)
	if command.Name == "" {
		fmt.Println("Command", cName, "not found in", file)
		os.Exit(1)
	}
	return command
}

func writeCommandsToFile(commands Config, filename string) (err error) {
	var commandsBuffer bytes.Buffer
	yamlEncoder := yaml.NewEncoder(&commandsBuffer)
	yamlEncoder.SetIndent(2)
	yamlEncoder.Encode(&commands)
	return ioutil.WriteFile(filename, commandsBuffer.Bytes(), 0644)
}

// SetCommand: create a sheesh command (add it to others if they exist)
func SetCommand(cfg quicli.Config) {
	cName := cfg.GetStringFlag("command")
	file := cfg.GetStringFlag("file")
	if cName == "" {
		fmt.Println("No command name provided. Exit. use --command")
		os.Exit(1)
	}
	// get commands
	commands, err := getCommandsFromFile(file)
	if err != nil {
		panic(err)
	}
	if cfg.GetBoolFlag("remove") {
		//remove command
		for i := 0; i < len(commands.Commands); i++ {
			if commands.Commands[i].Name == cName {
				commands.Commands = append(commands.Commands[:i], commands.Commands[i+1:]...)
				fmt.Println("Command", cName, "removed.")
				return
			}
		}
	} else {
		// add command
		for i := 0; i < len(commands.Commands); i++ {
			if commands.Commands[i].Name == cName {
				fmt.Println("Command", cName, "already exist.")
				os.Exit(0)
			}
		}
		command := Command{Name: cName}
		commands.Commands = append(commands.Commands, command)
	}

	if err := writeCommandsToFile(commands, file); err != nil {
		panic(err)
	}
	fmt.Println("Command", cName, "added.")
}

// Generate: generate output of a sheesh command
func Generate(cfg quicli.Config) {
	file := cfg.GetStringFlag("file")

	// get commands
	cfgYaml, err := getCommandsFromFile(file)
	if err != nil {
		panic(err)
	}

	if len(cfgYaml.Commands) == 0 {
		fmt.Println("No command found in yaml", file)
		os.Exit(1)
	}

	var output string
	for i := 0; i < len(cfgYaml.Commands); i++ {
		if cfg.GetStringFlag("command") != "" && cfgYaml.Commands[i].Name != cfg.GetStringFlag("command") {
			continue
		}
		output += ProcessCommand(cfgYaml.Commands[i])
		fmt.Println(output)
	}
}

// SetFlag: add a flag of a sheesh command
func SetFlag(cfg quicli.Config) {
	command := getCommandByNameFromCfg(cfg)
	if command.Name == "" {
		fmt.Println("Command", cfg.GetStringFlag("command"), "not found in", cfg.GetStringFlag("file"))
		os.Exit(1)
	}
}

// SetScript: set the script of of a sheesh command
func SetScript(cfg quicli.Config) {
	filename := cfg.GetStringFlag("file")
	command := getCommandByNameFromCfg(cfg)
	if command.Name == "" {
		fmt.Println("Command", cfg.GetStringFlag("command"), "not found in", filename)
		os.Exit(1)
	}
	command.Script = cfg.GetStringFlag("script")

	commands, err := getCommandsFromFile(filename)
	if err != nil {
		panic(err)
	}
	commands.Commands = replaceCommand(commands.Commands, command)
	if err := writeCommandsToFile(commands, filename); err != nil {
		panic(err)
	}
	fmt.Println(cfg.GetStringFlag("command"), "script changed.")

}

// ProcessCommand: take the command and return all the sheesh output (function + completion)
func ProcessCommand(c Command) (out string) {
	// fmt.Println(c.Script)

	commandCompletion := MarshallCompletion(c)

	commandContent := MarshallCommandContent(c)

	out = MarshallOutput(c.Name, commandContent, commandCompletion)

	return out
}

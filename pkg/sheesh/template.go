package sheesh

import (
	"bytes"
	"strings"
	"text/template"
)

type Output struct {
	CommandName       string
	Command           string
	CommandCompletion string
}

var outputTpl = `{{ .Command}}
{{ .CommandCompletion}}
compdef _{{ .CommandName}} {{ .CommandName}}
`

// MarshallOutput: take command and construct shell output accordingly
func MarshallOutput(cmdName, cmdContent, cmdCompletion string) (out string) {
	// last stage: command + completion
	var outputBuffer bytes.Buffer
	output := Output{cmdName, cmdContent, cmdCompletion}

	ot, err := template.New("output").Parse(outputTpl)
	if err != nil {
		panic(err)
	}
	err = ot.Execute(&outputBuffer, output)
	if err != nil {
		panic(err)
	}
	out = outputBuffer.String()
	return
}

type Completion struct {
	CommandName      string
	FlagsDefinitions string
	FlagCases        string
}

var completionTpl = `_{{ .CommandName}}() { 
  _arguments {{ .FlagsDefinitions}}
  case "$state" in
	  {{ .FlagCases}}
  esac
}  
`

func MarshallCompletion(c Command) (out string) {
	var completionBuffer bytes.Buffer

	fDef := MarshallFlagDefinitions(c.Flags)
	fCases := MarshallFlagCases(c.Flags)

	completion := Completion{c.Name, fDef, fCases}

	ct, err := template.New("completion").Parse(completionTpl)
	if err != nil {
		panic(err)
	}
	err = ct.Execute(&completionBuffer, completion)
	if err != nil {
		panic(err)
	}
	out = completionBuffer.String()
	return
}

func MarshallFlagDefinitions(flags []Flag) (out string) {
	var def []string
	for i := 0; i < len(flags); i++ {
		f := flags[i]
		var hyphen string
		if len(f.Name) > 1 {
			hyphen = "--"
		} else {
			hyphen = "-"
		}
		d := "'" + hyphen + f.Name + "[" + f.Description + "]:" + f.Name + ":->" + f.Name + "'"
		def = append(def, d)
	}
	out = strings.Join(def, " ")
	return out
}

var flagCaseTpl = `	{{ .FlagName}})
		{{ .CaseDef}}
	;;`

type FlagCase struct {
	FlagName string
	CaseDef  string
}

func MarshallFlagCases(flags []Flag) (out string) {
	var cases []string
	for i := 0; i < len(flags); i++ {
		f := flags[i]
		var flagCase bytes.Buffer

		var caseDef string
		if !f.NoArgs {
			if f.File {
				caseDef = "_files"
			}
			if len(f.Predefined) != 0 {
				caseDef = "_values '" + f.Name + "' " + strings.Join(f.Predefined, " ")
			}
		} else {
			caseDef = "_values '" + f.Name + "'"
		}

		caseStrct := FlagCase{f.Name, caseDef}

		fct, err := template.New("flag case").Parse(flagCaseTpl)
		if err != nil {
			panic(err)
		}
		err = fct.Execute(&flagCase, caseStrct)
		if err != nil {
			panic(err)
		}
		cases = append(cases, flagCase.String())
	}
	out = strings.Join(cases, "\n")
	return out
}

// COMMAND CONTENT

type CommandContent struct {
	CommandName      string
	FlagInitVars     string
	FlagContentCases string
	Script           string
}

var contentTpl = `{{ .CommandName}}(){
{{ .FlagInitVars}}
while true; do
	case "$1" in
	{{ .FlagContentCases}}
	-- ) shift; break ;;
	* ) break ;;
	esac
done

{{ .Script}}
}
`

func MarshallCommandContent(c Command) (out string) {
	var contentBuffer bytes.Buffer

	var initFlagVars, flagCases []string
	for i := 0; i < len(c.Flags); i++ {
		init := MarshallFlagInitVar(c.Flags[i])
		initFlagVars = append(initFlagVars, init)
		flagCase := MarshallFlagCase(c.Flags[i])
		flagCases = append(flagCases, flagCase)
	}

	content := CommandContent{c.Name, strings.Join(initFlagVars, "\n"), strings.Join(flagCases, "\n"), c.Script}

	ct, err := template.New("command content").Parse(contentTpl)
	if err != nil {
		panic(err)
	}
	err = ct.Execute(&contentBuffer, content)
	if err != nil {
		panic(err)
	}
	out = contentBuffer.String()
	return
}

type FlagInit struct {
	VarName      string
	DefaultValue string
}

var flagInitTpl = `{{ .VarName}}={{ .DefaultValue}}`

func MarshallFlagInitVar(f Flag) (out string) {
	var initBuffer bytes.Buffer
	var defaultV string
	if f.NoArgs {
		defaultV = "false"
	}
	flagInit := FlagInit{strings.ToUpper(f.Name), defaultV}

	ct, err := template.New("flag var init").Parse(flagInitTpl)
	if err != nil {
		panic(err)
	}
	err = ct.Execute(&initBuffer, flagInit)
	if err != nil {
		panic(err)
	}
	out = initBuffer.String()
	return
}

type FlagCaseCommand struct {
	FlagName string
	VarName  string
	VarValue string
	Hyphen   string
	Shift    string
}

var flagCaseCommandTpl = `{{ .Hyphen}}{{ .FlagName}}) {{ .VarName}}={{ .VarValue}}; shift {{ .Shift}};;`

func MarshallFlagCase(f Flag) (out string) {

	// 	-d | --debug ) DEBUG=true; shift ;;
	// -n | --name ) NAME="$2"; shift 2 ;;

	var caseBuffer bytes.Buffer
	var value, shift, hyphen string
	if f.NoArgs {
		value = "true"
	} else {
		value = "\"$2\""
		shift = "2"
	}
	if len(f.Name) > 1 {
		hyphen = "--"
	} else {
		hyphen = "-"
	}
	flagCase := FlagCaseCommand{f.Name, strings.ToUpper(f.Name), value, hyphen, shift}

	ct, err := template.New("flag var init").Parse(flagCaseCommandTpl)
	if err != nil {
		panic(err)
	}
	err = ct.Execute(&caseBuffer, flagCase)
	if err != nil {
		panic(err)
	}
	out = caseBuffer.String()
	return
}

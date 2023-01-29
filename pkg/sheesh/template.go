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
		d := "'-" + f.Name + "[" + f.Description + "]:" + f.Name + ":->" + f.Name + "'"
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
				caseDef = "_values '" + f.Name + "' " + strings.Join(f.Predefined, ",")
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

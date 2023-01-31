
<div align=center>
<strong>·</strong> 🪂 <strong>·</strong>
<br><strong>Generate only with:</strong>
<code>source <(sheesh)</code>
<br><strong>·</strong> 🪂 <strong>·</strong>
</div>


## G O !
You have two options:
* Use `sheesh` to produce command
* Define command within a yaml

### generate with `sheesh`

```shell
sheesh create "hello" --script "echo 'hello'"
sheesh addflag --command "hello" --name "who"
sheesh addflag --command "hello" --name "random" --no-argument
sheesh setscript --script "if [ "$RANDOM" = true ]; then WHO='toto';fi;echo 'hello ${WHO}'"
source <(sheesh)
```

### generate with yaml file

1. Create `.sheesh.yml` file
2. Launch:
```shell
source <(sheesh)
```

An `.sheesh.yml` example producing the same command as the above section:
```yaml
---
commands:
  - name: hello
    flags:
      - name: who
        description: "determine to whom speak"
        predefined:
          - "toto"
          - "titi"
      - name: random
        noarg: true
      - name: save
        description: "file to save output"
        file: true
    script: |
      if [ "$RANDOM" = true ];then
      WHO='toto'
      fi
      echo 'hello ${WHO}'
```

## T I P S 🎩

* very useful when you are testing api with `curl`
* want to have command in all shell, add this to `.${SHELL}rc`: `source <(sheesh --file [PATH_TO_SHEESHYAML]`)


## Limits/Improvement
* Only for zsh
* Use uppercase flag name in your script to use it value (`$FLAGNAME`)
* No default value
* Can't use var with "-" (hence flag too)
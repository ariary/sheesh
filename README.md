
<div align=center>

<br><strong>Generate only with:</strong>
<code>source <(sheesh)</code>

<img width=500 src="https://user-images.githubusercontent.com/14805903/215820436-2e7d553e-48c0-4232-b286-f21ae8c3ef1e.png">
<br><strong>·</strong> 🪂 <strong>·</strong><br>
</div>


## G O !
You have two options:
* Use `sheesh` to produce command
* Define command within a yaml

### generate with `sheesh`

```shell
sheesh setcommand --command "hello" --script "echo 'hello'"
sheesh setflag --command "hello" --name "who" --predefined "toto,titi,tata"
sheesh setflag --command "hello" --name "random" --noargs
sheesh setscript --command "hello" --script "if [ \"$RANDOM\" = true ]; then WHO='toto';fi;echo 'hello ${WHO}'"
source <(sheesh)
```

### generate with yaml file

1. Create `.sheesh.yml` file
2. Launch:
```shell
source <(sheesh)
```

An `.sheesh.yml` example producing some api call:
```yaml
---
commands:
  - name: api-postman
    flags:
      - name: stealth
        description: "change User-Agent"
        noarg: true
      - name: token
        predefined:
          - "THISISAADMINTOKEN"
          - "Dzdk7e0987djjdzz87dz"
      - name: save
        description: "file to save output"
        file: true
    script: |
      USERAGENT="curl 2.0/7"
      if [ "$STEALTH" = true ] ; then
          USERAGENT="not a hacker"
      fi
      if [ -n "$SAVE" ];then
        curl -H "User-Agent: ${USERAGENT}" -H "Authorization: Bearer ${TOKEN}" http://postman-echo.com/get > "${SAVE}"
      else
        curl -H "User-Agent: ${USERAGENT}" -H "Authorization: Bearer ${TOKEN}" http://postman-echo.com/get
      fi
```

## T I P S 🎩

* very useful when you are testing api with `curl`
* want to have command in all shell, add this to `.${SHELL}rc`: `source <(sheesh --file [PATH_TO_SHEESHYAML]`)


## Limits/Improvement
* Only for zsh
* Use uppercase flag name in your script to use it value (`$FLAGNAME`)
* No default value
* Can't use var with "-" (hence flag too)

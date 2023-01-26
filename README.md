
<div align=center>
<strong>·</strong> 🪸 <strong>·</strong>
<br><strong>Generate only with:</strong>
<code>source <(sheesh)</code>
<br><strong>·</strong> 🪸 <strong>·</strong>
</div>


## G O !
You have two options:
* Use `sheesh` to produce command
* Define command within a yaml

### generate with `sheesh`

```shell
sheesh create "hello" --script "echo 'hello'"
sheesh addflag --command "hello" --name "who"
sheesh setscript --script "echo 'hello {{who}}'"
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
    script: |
      echo 'hello {{who}}'
```

## T I P S 🎩

* very useful when you are testing api with `curl`
* want to have command in all shell, add this to `.${SHELL}rc`: `source <(sheesh --path [PATH_TO_SHEESHYAML]`)

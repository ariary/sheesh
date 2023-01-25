

## G O !

You have two options:
* Use `sheesh` to produce command
* Define command within a yaml

### g e n e r a t e with `sheesh`

```shell
sheesh create "hello" --script "echo 'hello'"
sheesh addflag --command "hello" --name "who"
sheesh setscript --script "echo 'hello {{who}}'"
source <(sheesh)
```

### g e n e r a t e with yaml file

1. Create `.sheesh.yml` file
2. Launch:
```shell
source <(sheesh completion)
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

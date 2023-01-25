

## G O !

You have two options:
* Use `sheesh` to produce command
* Define command with a yaml

### g e n e r a t e with `sheesh`

```shell
sheesh create "hello" --script "echo 'hello'"
sheesh addflag --command "hello" --name "who"
sheesh setscript --script "echo 'hello {{who}}'"
source <(sheesh completion)
```

### g e n e r a t e with yaml file

1. Create `.sheesh.yml` file
2. Launch:
```shell
source <(sheesh completion)
```

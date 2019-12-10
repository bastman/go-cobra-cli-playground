# go-cobra-cli-playground
how to create cli apps (golang: cobra + viper) ?

how to create something like ... ?
```
e.g.: $ kubectl describe deployment <DEPLOYMENT_NAME> -o yaml
```



## install golang (osx)

- [docs/InstallGo.md](docs/InstallGo.md)

## build cli-apps in go

- cobra + viper:
    - https://github.com/spf13/cobra
    - https://github.com/spf13/viper
    
- cobra app generator:
    - https://github.com/spf13/cobra/blob/master/cobra/README.md 
    
    Note: does not support go modules yet!   

## learn cobra + viper    

- [docs/LearnCobra.md](docs/LearnCobra.md)
      
    
## example
```
# create a new app "example-001"
$ make cobra.app.create COBRA_APP=example-001

# add command "hello" to app "example-001"
$ make cobra.command.create COBRA_APP=example-001 COBRA_COMMAND="hello"

```
    

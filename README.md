# go-cobra-cli-playground
how to create cli apps (golang, cobra)

## install golang (osx)

- [docs/InstallGo.md](docs/InstallGo.md)

## build cli-apps in go

- cobra + viper:
    - https://github.com/spf13/cobra
    - https://github.com/spf13/viper
    
- cobra generator:
    - https://github.com/spf13/cobra/blob/master/cobra/README.md 
    Note: does not support go modules yet!   
    
- tutorials
    - https://ordina-jworks.github.io/development/2018/10/20/make-your-own-cli-with-golang-and-cobra.html
    - https://github.com/nickgeudens/hello-cli
      
    
## example
```
# create a new app "example-001"
$ make cobra.app.create COBRA_APP=example-001

# add command "hello" to app "example-001"
$ make cobra.command.create COBRA_APP=example-001 COBRA_COMMAND="hello"

```
    

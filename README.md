# go-cobra-cli-playground
how to create cli apps (golang: cobra + viper) ?

how to create something like ... ?
```
e.g.: $ kubectl describe deployment <DEPLOYMENT_NAME> -o yaml
```

## install golang (osx)

- [docs/InstallGo.md](docs/InstallGo.md)

## build cli-apps in go (cobra + viper)

- cobra: https://github.com/spf13/cobra

    "Cobra is both a library for creating powerful modern CLI applications as well as a program to generate applications and command files."
    
- cobra app generator: https://github.com/spf13/cobra/blob/master/cobra/README.md 
    - Note: does not support go modules out of the box yet!  
    
- viper: https://github.com/spf13/viper

    "The viper library is famous for configuration solution for go application. It reads from JSON, TOML, YAML, HCL, envfile and Java properties config files."
 
see: https://towardsdatascience.com/how-to-create-a-cli-in-golang-with-cobra-d729641c7177

## learn cobra + viper    

- [docs/LearnCobra.md](docs/LearnCobra.md)
      
    
## example
```
# create a new app "example-001"
$ make cobra.app.create COBRA_APP=example-001

# add command "hello" to app "example-001"
$ make cobra.command.create COBRA_APP=example-001 COBRA_COMMAND="hello"

```
    

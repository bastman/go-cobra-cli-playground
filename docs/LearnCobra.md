# Learn Cobra

- cobra + viper:
    - https://github.com/spf13/cobra
    - https://github.com/spf13/viper
    
- cobra generator:
    - https://github.com/spf13/cobra/blob/master/cobra/README.md 
    Note: does not support go modules yet!   
    
- tutorials
    - https://ordina-jworks.github.io/development/2018/10/20/make-your-own-cli-with-golang-and-cobra.html
    - https://github.com/nickgeudens/hello-cli
    - https://towardsdatascience.com/how-to-create-a-cli-in-golang-with-cobra-d729641c7177

## basic concepts

```
e.g.: kubectl describe deployment <DEPLOYMENT_NAME> -o yaml
```

$ RootCommand Command SubCommand Args --Flags

    ("RootCommand" aka "AppName")

$ AppName Command Args --Flags 
    
    ("Args" aka "positional arguments")

$ AppName Command SubCommand Args --Flags 

    ("Flags" aka "named parameters/options")
    
In short ...
- Commands represent actions
- Args are things
- Flags are modifiers for those actions
    
````
e.g.: 
$ ./app say hello --name=Homer
$ ./app say hello --help
````

### flags: local vs persistent flags
see: https://towardsdatascience.com/how-to-create-a-cli-in-golang-with-cobra-d729641c7177

- persistent flags: This flag will be available to the command it is assigned as well as all the child or subcommands of the command
````
e.g.: $ ./app say hello --help
````

- local flags: This flag is only available to the command which it is assigned to
````
e.g.: $ ./app say hello --name=Homer
````


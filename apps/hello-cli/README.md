# Example

- from
    - https://ordina-jworks.github.io/development/2018/10/20/make-your-own-cli-with-golang-and-cobra.html
    - https://github.com/nickgeudens/hello-cli

```
$ cd <this-folder>

# compile 
$ go build

# run
$ ./app --help

$ ./app say --help

$ ./app say hello --help
$ ./app say hello --name=Homer              # result: Hello Homer
$ ./app say hello --config ./config.yml     # result: Howdy Billy
```

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
    
## example
```
$ cobra init --pkg-name github.com/spf13/newApp newApp && cd newApp
$ cobra add serve
$ go install
$ go build

```
    

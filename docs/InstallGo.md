
# install golang (osx)

file: $HOME/.bash_profile
```
### golang development ###
# see: https://ahmadawais.com/install-go-lang-on-macos-with-homebrew/
export GOPATH="${HOME}/.go"
export GOROOT="$(brew --prefix golang)/libexec"
export PATH="$PATH:${GOPATH}/bin:${GOROOT}/bin"
#test -d "${GOPATH}" || mkdir "${GOPATH}"
#test -d "${GOPATH}/src/github.com" || mkdir -p "${GOPATH}/src/github.com"

```

prepare GOPATH
```
echo "GOPATH: $GOPATH"
echo "GOROOT: $GOROOT"
$ test -d "${GOPATH}" || mkdir "${GOPATH}"
$ test -d "${GOPATH}/src/github.com" || mkdir -p "${GOPATH}/src/github.com"

```

install golang
```
$ brew install go
$ go version
$ go get golang.org/x/tools/cmd/godoc
```

# go modules / vendoring
- https://github.com/golang/go/wiki/Modules#how-do-i-use-vendoring-with-modules-is-vendoring-going-away

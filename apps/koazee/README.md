# App: koazee

Example: filter-map-reduce using koazee library

see: 
- https://github.com/wesovilabs/koazee
- wiki: https://github.com/wesovilabs/koazee/wiki

note: koazee provides kind of lazy evaluated stream (see Java Streams Api, Kotlin Sequence / Flow)

"Koazee is a StreamLike, Immutable, Lazy Loading and smart Golang Library to deal with slices.
Visit the Koazee wiki to find out what Koazee can do."

WARN: Does not yet work properly with go mod. see: https://github.com/wesovilabs/koazee/issues/54

## in action:

go-funk + koazee: https://github.com/nabinno/dojo/blob/master/project_euler/problem_4.go

## alternatives:

- https://medium.com/wesovilabs/koazee-vs-go-funk-vs-go-linq-caf8ef18584e

- go-funk: https://github.com/thoas/go-funk
- Koazee: https://github.com/wesovilabs/koazee
- Go-Linq: https://github.com/ahmetb/go-linq
- glow: https://github.com/chrislusf/glow   

## Basic examples:

```
$ cd <this-folder>

# compile 
$ go build

# run
$ ./app --help

$ ./app run
```

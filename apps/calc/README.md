# Calc

- see: https://towardsdatascience.com/how-to-create-a-cli-in-golang-with-cobra-d729641c7177

```
$ cd apps/calc

# compile 
$ go build

# run
$ ./app --help
$ ./app add 1 2 3               # result: 6
$ ./app add 1.2 2.5             # result: 3 (args treated as int types)
$ ./app add 1.2 2.5 --float     # result: 3.7

$ ./app add --help
$ ./app add even 1 2 3 4 5 6    # result: 12 (reason: 2+4+6=12)

```

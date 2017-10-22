# Playing With Go

In the program folder (`./src/hello`)

```
$ go build
```

Run without compilation:

```
$ go run main.go
```

Format:

```
$ gofmt -w main.go
```

Import:

```
$ goimport -w main.go
```

## Dependency management

```
$ brew install dep
```

* `$GOPATH` can point to `~/projects/go`
* Add `$GOPATH/bin` to `$PATH`
* All your go project should be in `~/projects/go/src`
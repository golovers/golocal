# GOLOCAL

Use local version of your code when working with [glide](https://github.com/Masterminds/glide) or [godep](https://github.com/golang/dep)

## Install:

```shell
go get github.com/golovers/golocal && go install github.com/golovers/golocal
```

## Usage
```shell
Usage:
  golocal [option] (args)

Options:
  -add:    Add local package(s) to configuration file
  -remove: Remove local package(s) from configuration file
  -list:   List all configured local packages
  -up:     Vendoring (copy) configured local packages to current vendor dicrectory
  -clear:  Remove all configured local packages from configuration file
  
Examples:
  ./golocal -list
  ./golocal -add "github.com/Sirupsen/logrus,github.com/mattn/go-sqlite3,google.golang.org/grpc"
  ./golocal -remove "github.com/Sirupsen/logrus,github.com/mattn/go-sqlite3"
  ./golocal -up
  ./golocal -clear
```

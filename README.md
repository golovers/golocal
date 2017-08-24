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

1. Help: 

   ```shell
   golocal -h
   ```

2. Configure to use local for some packages:

   ```shell
   golocal -add "github.com/package1,github.com/package2,gitlab.com/package3"
   ```

   Or you can update the configuration file directly at `vendor/vendor.local`


3. Pull local code to override remote code

   ```shell
   golocal -up
   ```

   ​

# GOLOCAL

Use local version of your code when working with [glide](https://github.com/Masterminds/glide) or [godep](https://github.com/golang/dep)

## Install:

```shell
go get github.com/golovers/golocal && go install github.com/golovers/golocal
```

## Usage

1. Help: 

   ```shell
   golocal -h
   ```

2. Configure to use local for some packages:

   ```shell
   golocal -import "github.com/package1,github.com/package2,gitlab.com/package3"
   ```

   Or you can update the configuration file directly at `vendor/vendor.local`


3. Pull local code to override remote code

   ```shell
   golocal -up
   ```

   ​
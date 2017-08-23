package main

import (
	"flag"
	"strings"
	"github.com/golovers/golocal/golo"
)

func main() {
	importPackages := flag.String("import", "", "comma-separated list of packages that use local")
	up := flag.Bool("up", false, "override the vendors by local")

	lis := flag.Bool("list", false, "list of use local packages")
	flag.Parse()

	// List
	if *lis {
		golo.List()
	}
	// Update the package list
	locals := strings.Split(*importPackages, ",")
	if *importPackages != "" && len(locals) > 0 {
		golo.UseLocal(locals)
	}
	// Update vendor
	if *up {
		golo.Up()
	}
}

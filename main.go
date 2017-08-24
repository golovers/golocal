package main

import (
	"flag"
	"github.com/golovers/golocal/golo"
	"strings"
)

func main() {
	fAdd := flag.String("add", "", "Add list of packages that use local. Comma separated.")
	fRemove := flag.String("remove", "", "Remove list of packages that used as local. Comma separated.")
	fUp := flag.Bool("up", false, "Override the vendor by local packages")
	fList := flag.Bool("list", false, "List of use local packages")
	fClear := flag.Bool("clear", false, "Remove all local packages from config file")
	flag.Parse()

	// List all configured local packages
	if *fList {
		golo.List()
		return
	}
	// Clear all packages from config list
	if *fClear {
		golo.Clear()
		return
	}
	// Remove local packages from config list
	if *fRemove != "" {
		golo.Remove(strings.Split(*fRemove, ","))
		return
	}
	// Add new local packages to config list
	if *fAdd != "" {
		golo.Add(strings.Split(*fAdd, ","))
		return
	}
	// Override vendor by configured local packages
	if *fUp {
		golo.Up()
	}
}

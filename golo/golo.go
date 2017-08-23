package golo

import (
	"strings"
	"log"
	"path"
	"os"
	"github.com/cf-guardian/guardian/kernel/fileutils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"go/build"
)

var (
	CONFIG_NAME = "vendor/vendor.local"
)

type Config struct {
	UseLocal []string
}

// Print the use local list
func List() {
	conf := getConf()
	if len(conf.UseLocal) > 0 {
		log.Printf("List packages that used as local: \n%v", strings.Join(conf.UseLocal, "\n"))
	} else {
		log.Println("No packages are used as local")
	}
}

// Write the config
func UseLocal(names []string) {
	conf := Config{
		UseLocal: names,
	}
	if data, err := yaml.Marshal(conf); err != nil {
		log.Fatal(err)
	} else {
		if err := ioutil.WriteFile(CONFIG_NAME, data, 0755); err != nil {
			log.Fatal(err)
		}
	}
}

// Load config
func getConf() Config {
	conf := Config{}
	if f, err := ioutil.ReadFile(CONFIG_NAME); err != nil {
		if os.IsNotExist(err) {
			return Config{}
		}
	} else {
		yaml.Unmarshal(f, &conf)
	}
	return conf
}

// Update the local repo as configured
func Up() {
	srcPath := path.Join(build.Default.GOPATH, "src")
	currDir, _ := os.Getwd()
	dstPath := path.Join(currDir, "vendor")
	conf := getConf()

	for _, pkg := range conf.UseLocal {
		src := path.Join(srcPath, pkg)
		dst := path.Join(dstPath, pkg)
		log.Println("Use local for: ", dst)
		os.RemoveAll(dst)
		CopyDir(src, dst)
	}
}

// Get parent dir
func ParentDir(dir string) string {
	if strings.HasSuffix(dir, "/") {
		dir = dir[:strings.LastIndex(dir, "/")]
	}
	return path.Dir(dir)
}

// Copy whole src folder to dst folder
func CopyDir(src string, dst string) {
	if _, err := os.Stat(dst); err != nil && os.IsNotExist(err) {
		os.MkdirAll(dst, 0755)
	}

	f := fileutils.New()
	f.Copy(ParentDir(dst), src)
}

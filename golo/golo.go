package golo

import (
	"github.com/cf-guardian/guardian/kernel/fileutils"
	"go/build"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

var (
	CONFIG_NAME = "vendor/vendor.local"
)

type Config struct {
	UseLocal []string
}

// List prints the configured local packages
func List() {
	conf := getConf()
	if len(conf.UseLocal) > 0 {
		log.Printf("List packages that used as local: \n%v", strings.Join(conf.UseLocal, "\n"))
	} else {
		log.Println("No packages are used as local")
	}
}

// Clear all configured local packages from config file
func Clear() {
	cfgFile := path.Join(getWd(), CONFIG_NAME)
	if !isExists(cfgFile) {
		log.Printf("Config file is not existed (%s)", cfgFile)
		return
	}
	if err := os.Remove(path.Join(getWd(), CONFIG_NAME)); err != nil {
		log.Printf("Failed to clear configuration: %s", err.Error())
	} else {
		log.Print("All local packages configuration cleared")
		// TODO: May remove source codes, too
	}
}

// Remove local package(s) from config file
func Remove(pkgs []string) {
	cfg := getConf()

	for _, pkg := range pkgs {
		if idx, ok := contains(cfg.UseLocal, pkg); ok {
			cfg.UseLocal = append(cfg.UseLocal[:idx], cfg.UseLocal[idx+1:]...)
			log.Printf("Removed %s package", pkg)
			continue
		}
		log.Printf("Ignored %s package", pkg)
	}
	cfg.writeToFile()
}

// Write the config
func Add(pkgs []string) {
	cfg := getConf()

	for _, pkg := range pkgs {
		if isExists(path.Join(getGoPath(), pkg)) {
			if _, ok := contains(cfg.UseLocal, pkg); !ok {
				cfg.UseLocal = append(cfg.UseLocal, pkg)
			}
			log.Printf("Added %s package", pkg)
			continue
		}
		log.Printf("Ignored %s package", pkg)
	}
	cfg.writeToFile()
}

// Update the local repo as configured
func Up() {
	conf := getConf()
	gp := getGoPath()
	vendor := path.Join(getWd(), "vendor")

	for _, pkg := range conf.UseLocal {
		src := path.Join(gp, pkg)
		dst := path.Join(vendor, pkg)
		if err := os.RemoveAll(dst); err != nil {
			log.Panicf("Failed to remove %s directory", dst)
		}
		copyDir(src, dst)
		log.Printf("Vendored %s package", pkg)
	}
}

// Get parent dir
func parentDir(dir string) string {
	if strings.HasSuffix(dir, "/") {
		dir = dir[:strings.LastIndex(dir, "/")]
	}
	return path.Dir(dir)
}

// Copy whole src folder to dst folder
func copyDir(src string, dst string) {
	if !isExists(dst) {
		if err := os.MkdirAll(dst, 0755); err != nil {
			log.Panicf("Failed to create %s directory", dst)
		}
	}

	f := fileutils.New()
	if err :=f.Copy(parentDir(dst), src); err!= nil {
		log.Panicf("Failed to copy %s directory to %s", src, dst)
	}
}

// Load config
func getConf() *Config {
	conf := &Config{}
	f, err := ioutil.ReadFile(CONFIG_NAME)
	if err != nil {
		return &Config{}
	}
	yaml.Unmarshal(f, &conf)
	return conf
}

func getGoPath() string {
	p := os.Getenv("GOPATH")
	if p == "" {
		p = build.Default.GOPATH
	}
	return path.Join(p, "src")
}

func getWd() string {
	pwd, err := os.Getwd()
	if err != nil {
		log.Panic("Failed to get working directory")
	}
	return pwd
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func contains(pkgs []string, pkg string) (int, bool) {
	for i, v := range pkgs {
		if v == pkg {
			return i, true
		}
	}
	return -1, false
}

func (c *Config) writeToFile() {
	data, err := yaml.Marshal(c)
	if err != nil {
		log.Panic(err)
	}
	if err := ioutil.WriteFile(CONFIG_NAME, data, 0755); err != nil {
		log.Panic(err)
	}
}

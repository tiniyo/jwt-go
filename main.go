package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/tiniyo/jwt-go/modules/log"
	"github.com/tiniyo/jwt-go/router"
)

const (
	DefaultConfFilePath = "configs/config.toml"
)

var (
	confFilePath string
	help         bool
)

func init() {
	flag.StringVar(&confFilePath, "c", DefaultConfFilePath, "Config Path")
	flag.StringVar(&confFilePath, "config", DefaultConfFilePath, "Config Path")
	flag.BoolVar(&help, "h", false, "Show help message")
	flag.BoolVar(&help, "help", false, "Show help message")
	flag.Parse()
	flag.Usage = usage
}

func usage() {
	s := `jwt-go : a Application for JWT Genration and Verification
		Usage: JWT [Options...]
		Options:
    		-c,  -config=<path>           Config path of the site. Default is configs/config.toml.
		Other options:
    		-h,  -help                  Show help message.
		`
	fmt.Printf(s)
	os.Exit(0)
}


func main() {

	log.Info("Application for JWT Genration and Verification")

	if help {
		usage()
		return
	}
	log.Debugf("run with conf:%s", confFilePath)
	router.RunSubdomains(confFilePath)
}


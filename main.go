package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/cjlapao/common-go/helper"
	"github.com/cjlapao/common-go/log"
	"github.com/cjlapao/common-go/version"
	"github.com/cjlapao/postmaster-server-go/nginx"
	"github.com/cjlapao/postmaster-server-go/system"
	ncparser "github.com/yangchenxing/go-nginx-conf-parser"
)

var logger = log.Get()
var ver = version.Get()

func main() {
	ver.Name = "Postmaster Admin API"
	ver.Author = "Carlos Lapao"
	ver.License = "MIT"
	ver.Rev = 1
	ver.PrintHeader()

	helpArg := helper.GetFlagSwitch("help", false)

	module := GetModuleArgument()
	if module == "" {
		os.Exit(0)
	}

	switch module {
	case "api":
		logger.Info("api")
	case "install":
		system.InstallDefault()
		nginx.Install()
		os.Exit(0)
	case "remove":
		system.RemoveDefault()
		nginx.Remove()
		system.RunAutoRemove()
		os.Exit(0)
	case "nginx":
		conf := nginx.Configuration{}
		err := conf.Parse("/etc/nginx/nginx.conf")
		if err != nil {
			logger.Error(err.Error())
		}
		// conf.User = "www wwww"
		// conf.Events = &nginx.Event{}
		// conf.Events.WorkerConnections = 4096
		result := conf.ToString()
		logger.Info(result)
		os.Exit(0)
	case "test":
		dat, err := ioutil.ReadFile("/etc/nginx/nginx.conf")
		test, err := ncparser.Parse(dat)
		if err != nil {
			fmt.Printf("there was an error, %v", err.Error())
		}
		fmt.Println(test[0].Words[0])
		os.Exit(0)

	default:
	}

	if helpArg {
		os.Exit(0)
	}
}

func GetModuleArgument() string {
	args := os.Args[1:]

	if len(args) == 0 || strings.HasPrefix(args[0], "-") {
	}

	return args[0]
}

func GetCommandArgument() string {
	args := os.Args[2:]

	if len(args) == 0 || strings.HasPrefix(args[0], "-") {
		return ""
	}

	return args[0]
}

func GetSubCommandArgument() string {
	args := os.Args[3:]

	if len(args) == 0 || strings.HasPrefix(args[0], "-") {
		return ""
	}

	return args[0]
}

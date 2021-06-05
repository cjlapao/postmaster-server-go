package main

import (
	"os"
	"strings"

	"github.com/cjlapao/common-go/helper"
	"github.com/cjlapao/common-go/log"
	"github.com/cjlapao/common-go/version"
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
		logger.Info("testing")
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

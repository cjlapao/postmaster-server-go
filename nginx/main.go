package nginx

import (
	"github.com/cjlapao/common-go/commands"
	"github.com/cjlapao/common-go/log"
)

var logger = log.Get()

func Install() {
	logger.Info("Installing Nginx in the System...")
	output, err := commands.Execute("apt-get", "install", "--assume-yes", "nginx-full")
	if err != nil {
		logger.FatalError(err, "There was an error installing nginx in your system")
	}
	logger.Info(output)
}

func Remove() {
	logger.Info("Removing Nginx from the System...")
	output, err := commands.Execute("apt-get", "purge", "--assume-yes", "nginx-full", "nginx-common")
	logger.Info(output)
	if err != nil {
		logger.FatalError(err, "There was an error removing nginx from your system")
	}
}

package system

import (
	"github.com/cjlapao/common-go/commands"
	"github.com/cjlapao/common-go/log"
)

var logger = log.Get()

func Update() {
	logger.Info("Updating System...")
	output, err := commands.Execute("apt-get", "update")

	if err != nil {
		logger.FatalError(err, "There was an error updating your system")
	}
	logger.Info(output)
}

func Upgrade() {
	logger.Info("Upgrading your System...")
	output, err := commands.Execute("apt-get", "upgrade", "--assume-yes")

	if err != nil {
		logger.FatalError(err, "There was an error upgrading your system")
	}
	logger.Info(output)
}

func InstallDefault() {
	Update()
	Upgrade()

	logger.Info("Installing default dependencies in System...")
	output, err := commands.Execute("apt-get", "install", "--assume-yes",
		"moreutils",
		"mutt",
		"sipcalc",
		"python3-pip")

	if err != nil {
		logger.FatalError(err, "There was an error installing nginx in your system")
	}
	logger.Info(output)
}

func RemoveDefault() {

	logger.Info("Removing default dependencies in System...")
	output, err := commands.Execute("apt-get", "purge", "--assume-yes", "--purge",
		"moreutils",
		"mutt",
		"sipcalc",
		"python3-pip")

	if err != nil {
		logger.FatalError(err, "There was an error installing nginx in your system")
	}

	logger.Info(output)
}

func RunAutoRemove() {
	logger.Info("Removing unecessary packages from the System...")
	output, err := commands.Execute("apt-get", "autoremove", "--assume-yes")
	if err != nil {
		logger.FatalError(err, "There was an error removing some unecessary dependencies from your system")
	}
	logger.Info(output)
}

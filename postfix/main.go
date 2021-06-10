package postfix

import(
	"github.com/cjlapao/common-go/log"
)

var logger = log.Get()

func Install() {
	logger.Info("Installing Postfix in the System")
}
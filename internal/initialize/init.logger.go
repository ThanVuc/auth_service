package initialize

import (
	"auth_service/global"
	"auth_service/pkg/loggers"
)

func InitLogger() {
	global.Logger = loggers.NewLogger(
		global.Config.Log,
	)
}

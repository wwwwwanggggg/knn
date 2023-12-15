package config

import (
	"final/service/validator"
)

func init() {
	initConfig()
	initLogger()
	validator.InitValidator(Config.AppLanguage)
}

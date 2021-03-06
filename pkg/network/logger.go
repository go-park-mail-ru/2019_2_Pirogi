package network

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"go.uber.org/zap"
)

func CreateLogger() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.DisableStacktrace = true
	cfg.OutputPaths = []string{
		"stdout",
		configs.Default.AccessLog,
	}
	cfg.ErrorOutputPaths = []string{
		"stderr",
		configs.Default.ErrorLog,
	}
	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	zap.ReplaceGlobals(logger)
	return logger, nil
}

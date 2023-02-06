package infra

import (
	"github.com/jmendezm/drone-dispatch-controller/config"
	log "github.com/sirupsen/logrus"
	"io"
)

func InitLog(config *config.Config) {
	var err error
	var hook *SyslogHook
	switch config.LogLevel {
	case 0:
		log.SetLevel(log.DebugLevel)
	case 1:
		log.SetLevel(log.InfoLevel)
	case 2:
		log.SetLevel(log.WarnLevel)
	case 3:
		log.SetLevel(log.ErrorLevel)
	case 4:
		log.SetLevel(log.FatalLevel)
	case 5:
		log.SetLevel(log.TraceLevel)
	}
	if err == nil {
		log.AddHook(hook)
	}
	log.SetFormatter(&log.JSONFormatter{
		DisableTimestamp: true,
		PrettyPrint:      false,
		FieldMap: log.FieldMap{
			log.FieldKeyLevel: "lv",
		},
	})
	if !config.ShowLogs {
		log.SetOutput(io.Discard)
	}
}

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
		hook, err = NewSyslogHook("udp", config.SysLogAddr, LOG_DEBUG, "")
	case 1:
		log.SetLevel(log.InfoLevel)
		hook, err = NewSyslogHook("udp", config.SysLogAddr, LOG_INFO, "")
	case 2:
		log.SetLevel(log.WarnLevel)
		hook, err = NewSyslogHook("udp", config.SysLogAddr, LOG_WARNING, "")
	case 3:
		log.SetLevel(log.ErrorLevel)
		hook, err = NewSyslogHook("udp", config.SysLogAddr, LOG_ERR, "")
	case 4:
		log.SetLevel(log.FatalLevel)
		hook, err = NewSyslogHook("udp", config.SysLogAddr, LOG_EMERG, "")
	case 5:
		log.SetLevel(log.TraceLevel)
		hook, err = NewSyslogHook("udp", config.SysLogAddr, LOG_ALERT, "")
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

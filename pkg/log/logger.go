package log

import (
	log "github.com/cihub/seelog"
)

func InitLogger() error {
	logger_, err := log.LoggerFromConfigAsFile("resource/config/seelog.xml")
	if err != nil {
		return err
	}
	log.ReplaceLogger(logger_)

	return err
}

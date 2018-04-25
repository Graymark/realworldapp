package logger

import (
	"log"

	"github.com/Graymark/realworldapp/domain"
	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	env    string
	Logger *logrus.Logger
}

type CredentialsGetter interface {
	GetCredentials() string
}

func NewLogger(env, podName string) *LogrusLogger {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.Printf("=== STARTING "+podName+" Logger with logLevel=%s", logger.Level)

	return &LogrusLogger{env: env, Logger: logger}
}

func (l LogrusLogger) Log(args ...interface{}) {
	if l.Logger == nil {
		return
	}
	// for stackDriver the log should contain 2 args :
	// first : the error (of uc.Error type),
	// second : an additional message (most likely the request that lead to the error)
	if len(args) == 2 {
		castedError, ok := args[0].(error)
		if ok {
			l.newLog(castedError, args[1])
		} else {
			l.Logger.Info(args)
		}
		return
	}
	l.Logger.Info(args)
}

func (l LogrusLogger) newLog(err error, usecase interface{}) {
	switch v := err.(type) {
	case *domain.Error:
		f := logrus.Fields{
			"type": v.ErrorType.String(),
			"mess": v.Message,
		}
		f["env"] = l.env

		if v.Additional != "" {
			f["additional"] = v.Additional
		}
		ll := l.Logger.WithFields(f)

		switch v.ErrorLevel {
		case domain.ErrDebug:
			ll.Debug(usecase)
		case domain.ErrInfo:
			ll.Info(usecase)
		case domain.ErrWarn:
			ll.Warn(usecase)
		case domain.ErrError:
			ll.Error(usecase)
		case domain.ErrFatal:
			ll.Fatal(usecase)
		}
	default:
		l.Logger.WithError(err).Error(usecase)
	}
}

type SimpleLogger struct{}

func (l SimpleLogger) Log(args ...interface{}) {
	log.Println(args)
}

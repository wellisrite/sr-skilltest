package log

import (
	"fmt"
	"os"
	"sr-skilltest/internal/model/constant"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

//Event is public function to create logging
func Event(traceId string, text ...string) {
	msgText := "[" + strings.Join(text, "][") + "]"
	logrus.SetLevel(logrus.InfoLevel)
	logrus.WithField("trace_id", traceId).Info(msgText)
}

//Message is public function to create logging
func Message(traceId string, text ...string) {
	msgText := "[" + strings.Join(text, "][") + "]"
	logrus.SetLevel(logrus.TraceLevel)
	logrus.WithField("trace_id", traceId).Trace(msgText)
}

//Warning is public function to create logging
func Warning(traceId string, text ...string) {
	msgText := "[" + strings.Join(text, "][") + "]"
	logrus.SetLevel(logrus.WarnLevel)
	logrus.WithField("trace_id", traceId).Warn(msgText)
}

//Error is public function to create logging
func Error(traceId string, err error, text ...string) {
	msgText := "[" + strings.Join(text, "][") + "]"
	errText := errors.Wrap(err, err.Error())
	stackTrace := "[" + strings.Replace(strings.Replace(fmt.Sprintf("%+v", errText), "\n\t", " ", -1),
		"\n", " | ", -1) + "]"
	logrus.SetLevel(logrus.ErrorLevel)
	logrus.WithField("trace_id", traceId).WithError(err).Error(msgText + stackTrace)
}

//Fatal is public function to create logging
func Fatal(traceId string, err error, text ...string) {
	msgText := "[" + strings.Join(text, "][") + "]"
	errText := errors.Wrap(err, err.Error())
	stackTrace := "[" + strings.Replace(strings.Replace(fmt.Sprintf("%+v", errText), "\n\t", " ", -1),
		"\n", " | ", -1) + "]"
	logrus.SetLevel(logrus.FatalLevel)
	logrus.WithField("trace_id", traceId).WithError(err).Fatal(msgText + stackTrace)
}

// SetupLogging is used to set up logging system
func SetupLogging(mode string) {
	logrus.SetOutput(os.Stdout)
	if strings.ToLower(mode) == constant.MODE_PRODUCTION {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		})
	}
}

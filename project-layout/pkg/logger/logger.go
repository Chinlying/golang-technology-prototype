// Package logger Log implementation for all microservices in the project.
// Log functions can be called through the convenience interfaces
// logger.Debugf(), logger.Errorf(), logger.Panicf()
//
// Deliberately reduces the interface to only Debugf, Errorf and Panicf.
// The other log levels are discouraged (see fdc Software Engineering Standards
// for details)
package logger

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/grpclog"
)

// DefaultTimeFormat is the default time format to be used with DataDog logs
const DefaultTimeFormat = time.RFC3339Nano

var (
	logger      *DefaultLogger
	serviceName string
	loggerMu    sync.Mutex
)

// SetDefaultLogger overrides the logger that is used in the convenience interface
// that can be accessed from everywhere. Changing the default logger should be
// done very early in the programs main function.
func SetDefaultLogger(l *logrus.Logger) {
	loggerMu.Lock()
	defer loggerMu.Unlock()

	logger = &DefaultLogger{l}
	// independent from our log level we are not interested in log messages but errors for 3rd party libraries
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(io.Discard, io.Discard, l.Writer()))
}

// GetDefaultLogger returns the logger currently used in the convenience interface.
// It is mainly provided so you can reset the it after changing it in a test.
func GetDefaultLogger() *DefaultLogger {
	loggerMu.Lock()
	defer loggerMu.Unlock()

	return logger
}

// SetServiceName allows overwriting default service name 'hostname'
func SetServiceName(newServiceName string) {
	serviceName = newServiceName
}

func init() { //nolint
	SetLogLevel(os.Getenv("LOG_LEVEL"))
	hostname, err := os.Hostname()
	if err != nil {
		serviceName = "<unknown hostname>"
	} else {
		serviceName = hostname
	}
	SetServiceName(serviceName)
	logger := logrus.StandardLogger()
	logger.SetFormatter(&logrus.JSONFormatter{
		// Note: The Stackdriver documentation makes it sound like the RFC3339
		// timestamp can't have sub-second precision (and suggests different
		// formats), but in practice it seems to be working fine
		// https://cloud.google.com/logging/docs/agent/configuration#timestamp-processing
		TimestampFormat: DefaultTimeFormat,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyMsg: "message",
		},
	})
	logger.AddHook(&serviceHook{})

	SetDefaultLogger(logger)
}

type DefaultLogger struct {
	logger *logrus.Logger
}

func (l *DefaultLogger) Debugf(format string, v ...interface{}) {
	l.logger.Debugf(format, v...)
}

func (l *DefaultLogger) Errorf(format string, v ...interface{}) {
	l.logger.Errorf(format, v...)
}

func (l *DefaultLogger) Infof(format string, v ...interface{}) {
	l.logger.Infof(format, v...)
}

func (l *DefaultLogger) Warnf(format string, v ...interface{}) {
	l.logger.Warnf(format, v...)
}

func (l *DefaultLogger) Panicf(format string, v ...interface{}) {
	l.logger.Panicf(format, v...)
}

func (l *DefaultLogger) Fatalf(format string, v ...interface{}) {
	l.logger.Fatalf(format, v...)
}

// WithError adds a error to a log entry
func (l *DefaultLogger) WithError(err error) *logrus.Entry {
	return l.logger.WithError(err)
}

// WithField adds a field to a log entry
func (l *DefaultLogger) WithField(key string, value interface{}) *logrus.Entry {
	return l.logger.WithField(key, value)
}

// WithContext adds the context to a log entry
func (l *DefaultLogger) WithContext(ctx context.Context) *logrus.Entry {
	return l.logger.WithContext(ctx)
}

// SetLogLevel allows overwriting default log level 'Error'
func SetLogLevel(logLevel string) {
	logrus.SetLevel(getLogrusLogLevel(logLevel))
}

func getLogrusLogLevel(logLevel string) logrus.Level {
	switch logLevel {
	case "": // not set
		return logrus.InfoLevel
	case "panic":
		return logrus.PanicLevel
	case "fatal":
		return logrus.FatalLevel
	case "error":
		return logrus.ErrorLevel
	case "warn":
		return logrus.WarnLevel
	case "info":
		return logrus.InfoLevel
	case "debug":
		return logrus.DebugLevel
	case "trace":
		return logrus.TraceLevel
	}

	panic(fmt.Sprintf("LOG_LEVEL %s is not known", logLevel))
}

// structured logging support
type serviceHook struct {
}

func (*serviceHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (*serviceHook) Fire(e *logrus.Entry) error {
	if e != nil {
		e.Data["service"] = fmt.Sprintf("%+v", serviceName)
	}
	return nil
}

var (
	logrusToGCE = map[logrus.Level]string{
		logrus.TraceLevel: "NOTICE",
		logrus.DebugLevel: "DEBUG",
		logrus.InfoLevel:  "INFO",
		logrus.WarnLevel:  "WARNING",
		logrus.ErrorLevel: "ERROR",
		logrus.FatalLevel: "CRITICAL",
		logrus.PanicLevel: "EMERGENCY",
	}
)

// GetGCELogLevel returns current log level as GCE string
func GetGCELogLevel() string {
	return logrusToGCE[logrus.GetLevel()]
}

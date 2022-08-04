package logging

import (
	"github.com/nats-io/nats-server/v2/logger"
)

// NATSLogger - uses the nats-server logging code
type NATSLogger struct {
	logger       *logger.Logger
	traceEnabled bool
}

// NewNATSLogger creates a new logger that uses the nats-server library
func NewNATSLogger(conf Config) Logger {
	l := logger.NewStdLogger(conf.Time, conf.Debug, conf.Trace, conf.Colors, conf.PID)
	return &NATSLogger{
		logger:       l,
		traceEnabled: conf.Trace,
	}
}

// Noticef  forwards to the nats logger
func (logger *NATSLogger) Noticef(format string, v ...interface{}) {
	logger.logger.Noticef(format, v...)
}

// Warnf forwards to the nats logger
func (logger *NATSLogger) Warnf(format string, v ...interface{}) {
	logger.logger.Warnf(format, v...)
}

// Fatalf forwards to the nats logger
func (logger *NATSLogger) Fatalf(format string, v ...interface{}) {
	logger.logger.Fatalf(format, v...)
}

// Errorf forwards to the nats logger
func (logger *NATSLogger) Errorf(format string, v ...interface{}) {
	logger.logger.Errorf(format, v...)
}

// Debugf forwards to the nats logger
func (logger *NATSLogger) Debugf(format string, v ...interface{}) {
	logger.logger.Debugf(format, v...)
}

// Tracef forwards to the nats logger
func (logger *NATSLogger) Tracef(format string, v ...interface{}) {
	logger.logger.Tracef(format, v...)
}

// TraceEnabled returns true if tracing is configured, useful for fast path logging
func (logger *NATSLogger) TraceEnabled() bool {
	return logger.traceEnabled
}

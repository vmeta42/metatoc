package logging

// Config defines logging flags for the NATS logger.
type Config struct {
	Time   bool
	Debug  bool
	Trace  bool
	Colors bool
	PID    bool
}

// Logger interface.
type Logger interface {
	// Log a notice statement
	Noticef(format string, v ...interface{})
	// Log a warning statement
	Warnf(format string, v ...interface{})
	// Log a fatal error
	Fatalf(format string, v ...interface{})
	// Log an error
	Errorf(format string, v ...interface{})
	// Log a debug statement
	Debugf(format string, v ...interface{})
	// Log a trace statement
	Tracef(format string, v ...interface{})

	TraceEnabled() bool
}

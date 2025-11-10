package xlogrus

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

// NewLogger creates a logrus.Logger that writes logs to both a file and stdout.
// logPath: path to the log file (will create directories automatically)
// logLevel: string representation of log level (info, debug, warn, etc.)
func NewLogger(logPath, logLevel string) (*logrus.Logger, error) {
	// Ensure the directory for the log file exists
	logDir := filepath.Dir(logPath)
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		return nil, err
	}

	// Open the log file for writing (append if exists)
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	// Create a new logrus logger
	logger := logrus.New()

	// Set output to stdout as primary output
	// File output is handled by the writerHook
	logger.SetOutput(os.Stdout)

	// Add a hook to write logs to file
	logger.AddHook(&writerHook{
		Writer:    logFile,
		LogLevels: logrus.AllLevels, // log all levels to file
	})

	// Parse and set log level, default to InfoLevel if invalid
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)

	// Set log format (TextFormatter with full timestamp)
	// Change to JSONFormatter if structured logs are preferred
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	return logger, nil
}

// writerHook is a simple logrus hook to write log entries to a file
type writerHook struct {
	Writer    *os.File       // destination file
	LogLevels []logrus.Level // levels to write
}

// Fire is called by logrus when a log entry is fired
func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String() // convert entry to string
	if err != nil {
		return err
	}
	_, err = hook.Writer.Write([]byte(line)) // write to file
	return err
}

// Levels returns the log levels that this hook handles
func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}

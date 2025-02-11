package logger

import (
	"io"
	"os"
	"path/filepath"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

type LogConfig struct {
	LogFile   string
	LogLevel  string
	LogType   string
	LogDir    string
	MaxSize   int
	MaxBackup int
	MaxAge    int
}

func SetAppLogger(cfg LogConfig) {
	// Set log formatter
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02T15:04:05.999999999Z07:00",
	})

	// Set log level
	switch cfg.LogLevel {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}

	// Set log output
	switch cfg.LogType {
	case "FILE":
		logFile := createNewLogFile(cfg)
		logrus.SetOutput(io.MultiWriter(os.Stdout, logFile))
	default:
		logrus.SetOutput(os.Stdout)
	}
}

func createNewLogFile(cfg LogConfig) *lumberjack.Logger {
	// Ensure log directory exists
	if err := os.MkdirAll(cfg.LogDir, 0755); err != nil {
		logrus.Fatalf("Failed to create log directory: %v", err)
	}

	// Main log file path
	logFileName := filepath.Join(cfg.LogDir, cfg.LogFile)

	return &lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    cfg.MaxSize,   // MB
		MaxBackups: cfg.MaxBackup, // Number of backups
		MaxAge:     cfg.MaxAge,    // Days
		Compress:   true,
		LocalTime:  true,
	}
}

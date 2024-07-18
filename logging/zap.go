package logging

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// installation commands
// go get -u go.uber.org/zap
func TestZap() {
	fmt.Println("----------------------------------------Start Zap----------------------------------------")
	// Creating a basic logger
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	// Simple logging example
	logger.Info("This is an info message")

	// Creating a more configurable logger
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	customLogger, err := config.Build()
	if err != nil {
		panic(err)
	}
	defer customLogger.Sync()

	// Logging different levels
	customLogger.Debug("This is a debug message")
	customLogger.Info("This is an info message")
	customLogger.Warn("This is a warning message")
	customLogger.Error("This is an error message")

	// Structured logging
	customLogger.Info("Structured log example",
		zap.String("url", "http://example.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", 5))

	// Logging with fields
	sugar := customLogger.Sugar()
	sugar.Infow("Failed to fetch URL",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", "5s")

	// Logging with format
	sugar.Infof("Failed to fetch URL: %s", "http://google.com")
	fmt.Println("----------------------------------------End Zap  ----------------------------------------")
}

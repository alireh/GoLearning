package logging

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// installation command
// go get github.com/rs/zerolog/log
func TestZerolog() {
	fmt.Println("----------------------------------------Start Zerolog----------------------------------------")
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Configure zerolog to write to the console with pretty printing
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

	log.Info().Msg("This is an info message")
	log.Debug().Msg("This is a debug message that won't be shown because the global level is set to Info")
	log.Warn().Msg("This is a warning message")
	log.Error().Msg("This is an error message")

	// Structured logging with fields
	log.Info().
		Str("component", "main").
		Int("attempt", 1).
		Dur("backoff", time.Second).
		Msg("Retrying in 1 second")

	// Log with a context
	logger := log.With().
		Str("component", "worker").
		Logger()
	logger.Info().Msg("Starting worker")

	// Example of logging an error with a stack trace
	err := someFunction()
	if err != nil {
		log.Error().
			Err(err).
			Msg("An error occurred")
	}
	fmt.Println("----------------------------------------End Zerolog  ----------------------------------------")
}

func someFunction() error {
	// Simulate an error
	return fmt.Errorf("some error")
}

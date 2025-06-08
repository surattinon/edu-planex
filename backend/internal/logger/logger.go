package logger

import (
    "os"
    "time"

    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)

func Init(debug bool) {
    // ConsoleWriter for pretty colors
    console := zerolog.ConsoleWriter{
        Out:        os.Stdout,
        TimeFormat: time.RFC3339,
        NoColor:    false,
    }

    // Set the global minimum level
    if debug {
        zerolog.SetGlobalLevel(zerolog.DebugLevel)
    } else {
        zerolog.SetGlobalLevel(zerolog.InfoLevel)
    }

    // Build the global logger with console output, timestamps, and caller info
    log.Logger = log.Output(console).
        With().
        Timestamp().
        Caller().
        Logger()
}

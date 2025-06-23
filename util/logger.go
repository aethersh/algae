package util

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var Logger = zerolog.New(
	zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.DateTime},
).Level(zerolog.TraceLevel).With().Timestamp().Caller().Logger()

var RequestLogger = zerolog.New(
	zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.DateTime},
).Level(zerolog.TraceLevel).With().Timestamp().Logger()

package errors

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"os"
)

const fatalMsg = "Something went wrong and EasyDonate had to stop."

// Logger is the system-wide logger
var Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

// Fatal terminates the program and exits with a non-zero code. It should be rarely used.
func Fatal(err error) {
	Logger.Fatal().Err(err).Msg(fatalMsg)
}

// FatalMsg terminates the program and exits with a non-zero code. It should be rarely used.
func FatalMsg(err error, friendlyMessage string) {
	Logger.Fatal().Err(err).Msg(friendlyMessage)
}

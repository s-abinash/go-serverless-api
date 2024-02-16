package utils

import (
	"go-serverless-api/config"
	"io"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var (
	Logger zerolog.Logger
)

var once sync.Once

func LoggerInit() {
	env := config.GetConfig()
	once.Do(func() {

		var output io.Writer
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.TimeFieldFormat = time.RFC3339

		logLevel, err := strconv.Atoi(env.LogLevel)
		if err != nil {
			logLevel = int(zerolog.InfoLevel)
		}

		if env.Env == "local" {
			output = zerolog.ConsoleWriter{
				Out:        os.Stdout,
				TimeFormat: time.RFC3339,
			}
		} else {
			output = os.Stdout
		}

		Logger = zerolog.New(output).Level(zerolog.Level(logLevel)).With().Timestamp().Caller().Logger()
	})
}

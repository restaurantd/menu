package loger

import (
	"github.com/rs/zerolog"
	"os"
)

type Loger struct {
	Logger *zerolog.Logger
}

func New() Loger {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return Loger{Logger: &logger}
}

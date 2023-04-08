package logx

import "github.com/rs/zerolog"

type Level int8

const (
	LevelDebug Level = Level(zerolog.DebugLevel)
	InfoLevel  Level = Level(zerolog.InfoLevel)
	WarnLevel  Level = Level(zerolog.WarnLevel)
	ErrorLevel Level = Level(zerolog.ErrorLevel)
	FatalLevel Level = Level(zerolog.FatalLevel)
	PanicLevel Level = Level(zerolog.PanicLevel)
	NoLevel    Level = Level(zerolog.NoLevel)
	Disabled   Level = Level(zerolog.Disabled)
	TraceLevel Level = Level(zerolog.TraceLevel)
)

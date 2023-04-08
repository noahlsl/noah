package logx

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"os"
	"strconv"
)

var (
	l = zerolog.Logger{}
)

func init() {
	l = zerolog.New(os.Stdout)
}
func initLog(c *Conf) zerolog.Logger {
	log := zerolog.New(os.Stderr).
		With().
		Timestamp().
		Str("service", c.Server).
		CallerWithSkipFrameCount(2).
		Logger()
	for _, field := range c.GlobalField {
		log.With().Str(field.Key, fmt.Sprintf("%v", field.Value))
	}

	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return washPath(file) + ":" + strconv.Itoa(line)
	}

	// 设置日志等级
	log = l.Level(zerolog.Level(c.Level))
	// 使用官方提供的，输出更友好
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	return log
}

func MustUpSet(c *Conf) {
	l = initLog(c)
}

func Error(err error) {
	l.Err(err).Caller(1).Send()
}
func Errorf(msg string, args ...interface{}) {
	l.Error().Caller(1).Msg(fmt.Sprintf(msg, args...))
}
func Errorw(msg string, fields ...Field) {
	l.Error().Caller(1).Fields(ToMap(fields...)).Msg(msg)
}

func Info(msg string) {
	l.Info().Caller(1).Msg(msg)
}
func Infof(msg string, args ...interface{}) {
	l.Info().Caller(1).Msg(fmt.Sprintf(msg, args...))
}
func Infow(msg string, fields ...Field) {
	l.Info().Caller(1).Fields(ToMap(fields...)).Msg(msg)
}

func Debug(msg string) {
	l.Debug().Caller(1).Msg(msg)
}
func Debugf(msg string, args ...interface{}) {
	l.Debug().Caller(1).Msg(fmt.Sprintf(msg, args...))
}
func Debugw(msg string, fields ...Field) {
	l.Debug().Caller(1).Fields(ToMap(fields...)).Msg(msg)
}

func Warn(msg string) {
	l.Warn().Caller(1).Msg(msg)
}
func Warnf(msg string, args ...interface{}) {
	l.Debug().Caller(1).Msg(fmt.Sprintf(msg, args...))
}
func Warnw(msg string, fields ...Field) {
	l.Debug().Caller(1).Fields(ToMap(fields...)).Msg(msg)
}

func Fatal(msg string) {
	l.Fatal().Caller(1).Msg(msg)
}
func Fatalf(msg string, args ...interface{}) {
	l.Debug().Caller(1).Msg(fmt.Sprintf(msg, args...))
}
func Fatalw(msg string, fields ...Field) {
	l.Debug().Caller(1).Fields(ToMap(fields...)).Msg(msg)
}

func Panic(msg string) {
	l.Panic().Caller(1).Msg(msg)
}
func Panicf(msg string, args ...interface{}) {
	l.Debug().Caller(1).Msg(fmt.Sprintf(msg, args...))
}
func Panicw(msg string, fields ...Field) {
	l.Debug().Caller(1).Fields(ToMap(fields...)).Msg(msg)
}

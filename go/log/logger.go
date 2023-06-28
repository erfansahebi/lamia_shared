package log

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

type Logger interface {
	Debugf(ctx context.Context, format string, args ...interface{})
	Infof(ctx context.Context, format string, args ...interface{})
	Warnf(ctx context.Context, format string, args ...interface{})
	Errorf(ctx context.Context, format string, args ...interface{})
	Fatalf(ctx context.Context, format string, args ...interface{})

	WithField(key string, value interface{}) Logger
	WithError(err error) Logger
	WithFields(fields Fields) Logger
	NoCaller() Logger
	SkipLevel(n int) Logger

	// Skip(skip bool) Logger
}

var defaultLogger zerolog.Logger

func init() {
	defaultLogger = zerolog.New(os.Stderr)
	defaultLogger.Level(zerolog.InfoLevel)
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
}

func NewLogger() Logger {
	return newLogger()
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	newLogger().SkipLevel(3).Debugf(ctx, format, args...)
}
func Infof(ctx context.Context, format string, args ...interface{}) {
	newLogger().SkipLevel(3).Infof(ctx, format, args...)
}
func Warnf(ctx context.Context, format string, args ...interface{}) {
	newLogger().SkipLevel(3).Warnf(ctx, format, args...)
}
func Errorf(ctx context.Context, format string, args ...interface{}) {
	newLogger().SkipLevel(3).Errorf(ctx, format, args...)
}

func Fatalf(ctx context.Context, format string, args ...interface{}) {
	newLogger().SkipLevel(3).Fatalf(ctx, format, args...)
}

func WithField(key string, value interface{}) Logger {
	return newLogger().WithField(key, value)
}
func WithError(err error) Logger {
	return newLogger().WithError(err)
}
func WithFields(fields Fields) Logger {
	return newLogger().WithFields(fields)
}

type logger struct {
	zLogger zerolog.Logger

	err       error
	fields    Fields
	skipLevel int
	noCaller  bool
}

func newLogger() Logger {
	if os.Getenv("DEBUG") == "true" {
		return newConsoleLogger()
	}
	return &logger{
		zLogger: defaultLogger,
		fields:  make(Fields),
	}
}

func (l *logger) SkipLevel(n int) Logger {
	return l.skip(n)
}

func (l *logger) skip(level int) *logger {
	if level > l.skipLevel {
		l.skipLevel = level
	}
	return l
}
func (l *logger) Debugf(ctx context.Context, format string, args ...interface{}) {
	l.skip(2).flush(ctx, l.zLogger.Debug, format, args...)
}
func (l *logger) Infof(ctx context.Context, format string, args ...interface{}) {
	l.skip(2).flush(ctx, l.zLogger.Info, format, args...)
}
func (l *logger) Warnf(ctx context.Context, format string, args ...interface{}) {
	l.skip(2).flush(ctx, l.zLogger.Warn, format, args...)
}
func (l *logger) Errorf(ctx context.Context, format string, args ...interface{}) {
	l.skip(2).flush(ctx, l.zLogger.Error, format, args...)
}

func (l *logger) Fatalf(ctx context.Context, format string, args ...interface{}) {
	l.skip(2).flush(ctx, l.zLogger.Fatal, format, args...)
}

func (l *logger) NoCaller() Logger {
	l.noCaller = true
	return l
}

func (l *logger) WithField(key string, value interface{}) Logger {
	l.fields[key] = value
	return l
}
func (l *logger) WithError(err error) Logger {
	l.err = err
	return l
}
func (l *logger) WithFields(fields Fields) Logger {
	l.fields.Merge(fields)
	return l
}

func (l *logger) flush(ctx context.Context, ev func() *zerolog.Event, format string, args ...interface{}) {
	e := ev().Timestamp().Fields(map[string]interface{}(l.fields))
	if !l.noCaller {
		e = e.Caller(l.skipLevel)
	}
	e.Stack().Err(l.err).Msgf(format, args...)
}

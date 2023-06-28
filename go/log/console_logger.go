package log

import (
	"context"
	"fmt"
	"io"
	"os"
	"runtime"
	"time"
)

type consoleLogger struct {
	w         io.Writer
	err       error
	fields    Fields
	skipLevel int
	noCaller  bool
}

func newConsoleLogger() *consoleLogger {
	return &consoleLogger{
		w:      os.Stderr,
		fields: make(Fields),
	}
}

func (l *consoleLogger) skip(level int) *consoleLogger {
	if level > l.skipLevel {
		l.skipLevel = level
	}
	return l
}
func (l *consoleLogger) Debugf(ctx context.Context, format string, args ...interface{}) {
	l.skip(2).flush(ctx, format, args...)
}
func (l *consoleLogger) Infof(ctx context.Context, format string, args ...interface{}) {
	l.skip(2).flush(ctx, format, args...)
}
func (l *consoleLogger) Warnf(ctx context.Context, format string, args ...interface{}) {
	l.skip(2).flush(ctx, format, args...)
}
func (l *consoleLogger) Errorf(ctx context.Context, format string, args ...interface{}) {
	l.skip(2).flush(ctx, format, args...)
}

func (l *consoleLogger) Fatalf(ctx context.Context, format string, args ...interface{}) {
	l.skip(2).flush(ctx, format, args...)
}

func (l *consoleLogger) NoCaller() Logger {
	l.noCaller = true
	return l
}

func (l *consoleLogger) SkipLevel(n int) Logger {
	l.skipLevel = n
	return l
}

func (l *consoleLogger) WithField(key string, value interface{}) Logger {
	l.fields[key] = value
	return l
}
func (l *consoleLogger) WithError(err error) Logger {
	l.err = err
	return l
}
func (l *consoleLogger) WithFields(fields Fields) Logger {
	l.fields.Merge(fields)
	return l
}

func (l *consoleLogger) flush(ctx context.Context, format string, args ...interface{}) {
	fields := ""
	if len(l.fields) > 0 {
		fields = fmt.Sprintf("%v", l.fields)
	}

	err := ""
	if l.err != nil {
		err = fmt.Sprintf("err: %v", l.err)
	}

	caller := ""
	_, file, line, ok := runtime.Caller(l.skipLevel)
	if ok {
		caller = fmt.Sprintf("%s:%d", file, line)
	}

	fmt.Fprintf(l.w, "%s %s %s %v %s\n", time.Now().Format(time.RFC822), fmt.Sprintf(format, args...), err, fields, caller)
}

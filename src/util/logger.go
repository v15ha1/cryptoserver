package util

import (
	"context"
	"io"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

//CustomLogFormatter ...
type CustomLogFormatter struct {
	Formatter logrus.Formatter
}

//GetLogger ...
func (f *CustomLogFormatter) GetLogger(writer io.Writer, app string, logLevel string) *logrus.Entry {
	logger := logrus.New()
	logger.SetOutput(writer)

	level := logrus.InfoLevel
	switch logLevel {
	case "debug":
		level = logrus.DebugLevel
	case "trace":
		level = logrus.TraceLevel
	case "error":
		level = logrus.ErrorLevel
	case "info":
		level = logrus.InfoLevel
	}

	logger.SetLevel(level)

	fields := map[string]interface{}{
		"application": app,
	}
	entry := logrus.NewEntry(logger).WithFields(fields)
	logger.SetFormatter(f)
	return entry
}

//NewCustomLogFormatter ...
func NewCustomLogFormatter() *CustomLogFormatter {
	jsonFormatter := new(logrus.JSONFormatter)
	jsonFormatter.TimestampFormat = "2006-01-02 15:04:05"
	return &CustomLogFormatter{
		Formatter: jsonFormatter,
	}
}

// Format renders a log entry.
func (f *CustomLogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	line, fileName, funcName := f.getCaller(7)
	fields := logrus.Fields{
		"file":        fileName,
		"line_number": line,
		"func":        funcName,
		//		"env":         env,
		//		"host":        host,
		//		"servergroup": serverGroup,
	}
	for key, value := range entry.Data {
		fields[key] = value
	}
	entry.Data = fields
	return f.Formatter.Format(entry)
}

func (f *CustomLogFormatter) getCaller(skip int) (line int, fileName, funcName string) {
	pc, fileName, line, ok := runtime.Caller(skip)
	if !ok {
		return -1, "", ""
	}
	fileName = filepath.Base(fileName)
	funcName = runtime.FuncForPC(pc).Name()
	funcIndex := strings.LastIndex(funcName, ".")
	funcName = funcName[funcIndex+1:]
	return
}

//Logger logger
func Logger(ctx context.Context) *logrus.Entry {
	if entry, ok := ctx.Value("loggerKey").(*logrus.Entry); ok {
		return entry
	}
	return nil
}

//WithLogger ...
func WithLogger(ctx context.Context, entry *logrus.Entry) context.Context {
	return context.WithValue(ctx, "loggerKey", entry)
}

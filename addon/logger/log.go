package logger

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/pingcap/ticp/config"

	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type logCtxKeyType struct{}
type Fields log.Fields

var logCtxKey logCtxKeyType

type LogRecord struct {
	defaultLogEntry *log.Entry
}

var logRecord *LogRecord

const (
	// LogDebug debug level
	LogDebug = "debug"
	// LogInfo info level
	LogInfo = "info"
	// LogWarn warn level
	LogWarn = "warn"
	// LogError error level
	LogError = "error"
	// LogFatal fatal level
	LogFatal = "fatal"
)

const (
	// OutputConsole log console output
	OutputConsole = "console"
	// OutputFile log file output
	OutputFile = "file"
)

const (
	// RecordSysField record sys name
	RecordSysField = "sys"
	// RecordModField record mod name
	RecordModField = "mod"
	// RecordFunField record fun name
	RecordFunField = "fun"
	// RecordFileField record file name
	RecordFileField = "file"
	// RecordLineField record line number
	RecordLineField = "line"
)

func Init() {
	if logRecord == nil {
		logger := log.New()

		// Get global log configuration
		conf := config.GetLogConfig()

		// Set log format
		logger.SetFormatter(&log.JSONFormatter{})
		// Set log level
		logger.SetLevel(getLogLevel(conf.LogLevel))

		// Define output type writer
		writers := []io.Writer{os.Stdout}

		// Determine whether the log output contains the file type
		if strings.Contains(strings.ToLower(conf.LogOutput), OutputFile) {
			writers = append(writers, getFileOutput(conf))
		}
		// If console is not included, remove the os.Stdout output
		if !strings.Contains(strings.ToLower(conf.LogOutput), OutputConsole) {
			writers = writers[1:]
		}
		// Set log output
		logger.SetOutput(io.MultiWriter(writers...))

		logRecord = &LogRecord{
			defaultLogEntry: log.NewEntry(logger),
		}

		// Record sys and mod default init
		logRecord.defaultLogEntry = logRecord.defaultLogEntry.
			WithField(RecordSysField, conf.RecordSysName).WithField(RecordModField, conf.RecordModName)
	}
}

// Log file output configuration
func getFileOutput(conf config.Log) *lumberjack.Logger {
	logConfig := &lumberjack.Logger{
		// Log output file path
		Filename: conf.LogFilePath,
		// Maximum log file size, unit: MB
		MaxSize: conf.LogMaxSize,
		// Maximum time interval for keeping expired files, unit: days
		MaxAge: conf.LogMaxAge,
		// Maximum number of expired logs to backups
		MaxBackups: conf.LogMaxBackups,
		// Whether to use local time
		LocalTime: conf.LogLocalTime,
		// Do you need to compress the rolling log, use gzip compression
		Compress: conf.LogCompress,
	}
	return logConfig
}

// Tool method to get log level
func getLogLevel(level string) log.Level {
	switch strings.ToLower(level) {
	case LogDebug:
		return log.DebugLevel
	case LogInfo:
		return log.InfoLevel
	case LogWarn:
		return log.WarnLevel
	case LogError:
		return log.ErrorLevel
	case LogFatal:
		return log.FatalLevel
	}
	return log.DebugLevel
}

func NewContext(ctx context.Context, fields Fields) context.Context {
	return context.WithValue(ctx, logCtxKey, WithContext(ctx).WithFields(log.Fields(fields)))
}

func WithContext(ctx context.Context) *log.Entry {
	if ctx == nil {
		Init()
		return logRecord.defaultLogEntry
	}
	le, ok := ctx.Value(logCtxKey).(*log.Entry)
	if ok {
		return le
	} else {
		return logRecord.defaultLogEntry
	}
}

func GetLogger() *LogRecord {
	if logRecord == nil {
		Init()
	}
	return logRecord
}

func (lr *LogRecord) Record(key string, value interface{}) *LogRecord {
	logRecord.defaultLogEntry = logRecord.defaultLogEntry.WithField(key, value)
	return logRecord
}

func (lr *LogRecord) Records(fields log.Fields) *LogRecord {
	logRecord.defaultLogEntry = logRecord.defaultLogEntry.WithFields(fields)
	return logRecord
}

func (lr *LogRecord) RecordSys(sys string) *LogRecord {
	logRecord.defaultLogEntry = logRecord.defaultLogEntry.WithField(RecordSysField, sys)
	return logRecord
}
func (lr *LogRecord) RecordMod(mod string) *LogRecord {
	logRecord.defaultLogEntry = logRecord.defaultLogEntry.WithField(RecordModField, mod)
	return logRecord
}

func (lr *LogRecord) RecordFun() *log.Entry {
	logEntry := logRecord.defaultLogEntry
	if pc, file, line, ok := runtime.Caller(2); ok {
		ptr := runtime.FuncForPC(pc)
		fmt.Println(ptr.Name(), file, line)
		logEntry = logRecord.defaultLogEntry.WithField(RecordFunField, ptr.Name()).
			WithField(RecordFileField, path.Base(file)).WithField(RecordLineField, line)
	}
	return logEntry
}

func (lr *LogRecord) Debug(args interface{}) {
	lr.RecordFun().Debug(args)
}

func (lr *LogRecord) Debugf(format string, args ...interface{}) {
	lr.RecordFun().Debugf(format, args)
}

func (lr *LogRecord) Debugln(args interface{}) {
	lr.RecordFun().Debugln(args)
}

func (lr *LogRecord) Info(args interface{}) {
	lr.RecordFun().Info(args)
}

func (lr *LogRecord) Infof(format string, args ...interface{}) {
	lr.RecordFun().Infof(format, args)
}

func (lr *LogRecord) Infoln(args interface{}) {
	lr.RecordFun().Infoln(args)
}

func (lr *LogRecord) Warn(args interface{}) {
	lr.RecordFun().Warn(args)
}

func (lr *LogRecord) Warnf(format string, args ...interface{}) {
	lr.RecordFun().Warnf(format, args)
}

func (lr *LogRecord) Warnln(args interface{}) {
	lr.RecordFun().Warnln(args)
}

func (lr *LogRecord) Warning(args interface{}) {
	lr.RecordFun().Warning(args)
}

func (lr *LogRecord) Warningf(format string, args ...interface{}) {
	lr.RecordFun().Warningf(format, args)
}

func (lr *LogRecord) Warningln(args interface{}) {
	lr.RecordFun().Warningln(args)
}

func (lr *LogRecord) Error(args interface{}) {
	lr.RecordFun().Error(args)
}

func (lr *LogRecord) Errorf(format string, args ...interface{}) {
	lr.RecordFun().Errorf(format, args)
}

func (lr *LogRecord) Errorln(args interface{}) {
	lr.RecordFun().Errorln(args)
}

func (lr *LogRecord) Fatal(args interface{}) {
	lr.RecordFun().Fatal(args)
}

func (lr *LogRecord) Fatalf(format string, args ...interface{}) {
	lr.RecordFun().Fatalf(format, args)
}

func (lr *LogRecord) Fatalln(args interface{}) {
	lr.RecordFun().Fatalln(args)
}

// Mit License
// Author: rod6@github

package log6

import (
    "io"
    "log"
    "os"
)

type Level uint8

const (
    DebugLevel Level = iota
    InfoLevel
    WarnLevel
    ErrorLevel
    FatalLevel
    PanicLevel
)

const DefaultLevel Level = InfoLevel

type Logger struct {
    *log.Logger
    level Level
}

func (l *Logger) SetLevel(level Level) {
    l.level = level
}

func (l *Logger) ParseLevel(level string) {
    switch level {
    case "debug":
        l.level = DebugLevel
    case "info":
        l.level = InfoLevel
    case "warn":
        l.level = WarnLevel
    case "error":
        l.level = ErrorLevel
    case "fatal":
        l.level = FatalLevel
    case "panic":
        l.level = PanicLevel
    default:
        l.level = InfoLevel
    }
}

func (l *Logger) GetLevel() string {
    switch l.level {
    case DebugLevel:
        return "debug"
    case InfoLevel:
        return "info"
    case WarnLevel:
        return "warn"
    case ErrorLevel:
        return "error"
    case FatalLevel:
        return "fatal"
    case PanicLevel:
        return "panic"
    }
    return ""
}

func (l *Logger) Debug(format string, v ...interface{}) {
    if l.level > DebugLevel {
        return
    }

    f := "[DEBU] " + format
    l.Printf(f, v...)
}

func (l *Logger) Info(format string, v ...interface{}) {
    if l.level > InfoLevel {
        return
    }

    f := "[INFO] " + format
    l.Printf(f, v...)
}

func (l *Logger) Warn(format string, v ...interface{}) {
    if l.level > WarnLevel {
        return
    }

    f := "[WARN] " + format
    l.Printf(f, v...)
}

func (l *Logger) Error(format string, v ...interface{}) {
    if l.level > ErrorLevel {
        return
    }

    f := "[ERRO] " + format
    l.Printf(f, v...)
}

func (l *Logger) Fatal(format string, v ...interface{}) {
    f := "[FATA] " + format
    l.Fatalf(f, v...)
}

func (l *Logger) Panic(format string, v ...interface{}) {
    f := "[PANI] " + format
    l.Panicf(f, v...)
}

func New(logger *log.Logger, level Level) *Logger {
    return &Logger{logger, level}
}

// Defautl logger & functions
var std = New(log.New(os.Stderr, "◉ ", log.Ldate | log.Ltime | log.Lmicroseconds), InfoLevel)

func SetLevel(level Level) {
    std.SetLevel(level)
}

func ParseLevel(level string) {
    std.ParseLevel(level)
}

func GetLevel() string {
    return std.GetLevel()
}

func SetOutput(w io.Writer) {
    std.SetOutput(w)
}

func Debug(format string, v ...interface{}) {
    std.Debug(format, v...)
}

func Info(format string, v ...interface{}) {
    std.Info(format, v...)
}

func Warn(format string, v ...interface{}) {
    std.Warn(format, v...)
}

func Error(format string, v ...interface{}) {
    std.Error(format, v...)
}

func Fatal(format string, v ...interface{}) {
    std.Fatal(format, v...)
}

func Panic(format string, v ...interface{}) {
    std.Panic(format, v...)
}

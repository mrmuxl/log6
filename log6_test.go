// Mit License
// Author: rod6@github

package log6

import (
    "io/ioutil"
    "os"
    "strings"
    "testing"
)

func TestSetGetLevel(t *testing.T) {
    SetLevel(DebugLevel)
    if std.GetLevel() != "debug" {
        t.Errorf(`Level DebugLevel should be "debug"`)
    }

    SetLevel(InfoLevel)
    if std.GetLevel() != "info" {
        t.Errorf(`Level InfoLevel should be "info"`)
    }

    SetLevel(WarnLevel)
    if std.GetLevel() != "warn" {
        t.Errorf(`Level WarnLevel should be "warn"`)
    }

    SetLevel(ErrorLevel)
    if std.GetLevel() != "error" {
        t.Errorf(`Level ErrorLevel should be "error"`)
    }

    SetLevel(FatalLevel)
    if std.GetLevel() != "fatal" {
        t.Errorf(`Level FatalLevel should be "fatal"`)
    }

    SetLevel(PanicLevel)
    if std.GetLevel() != "panic" {
        t.Errorf(`Level PanicLevel should be "panic"`)
    }
}

func TestParseLevel(t *testing.T) {
    ParseLevel("debug")
    if std.GetLevel() != "debug" {
        t.Errorf(`Level DebugLevel should be "debug"`)
    }

    ParseLevel("info")
    if std.GetLevel() != "info" {
        t.Errorf(`Level InfoLevel should be "info"`)
    }

    ParseLevel("warn")
    if std.GetLevel() != "warn" {
        t.Errorf(`Level WarnLevel should be "warn"`)
    }

    ParseLevel("error")
    if std.GetLevel() != "error" {
        t.Errorf(`Level ErrorLevel should be "error"`)
    }

    ParseLevel("fatal")
    if std.GetLevel() != "fatal" {
        t.Errorf(`Level FatalLevel should be "fatal"`)
    }

    ParseLevel("panic")
    if std.GetLevel() != "panic" {
        t.Errorf(`Level PanicLevel should be "panic"`)
    }

    ParseLevel("wrong")
    if std.GetLevel() != "info" {
        t.Errorf(`Level wrong level input should be "info"`)
    }
}

func TestLogMessage(t *testing.T) {
    f, _ := os.OpenFile("testing.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
    SetOutput(f)
    SetLevel(InfoLevel)

    Debug("Log Debug Message: %v", "debug")
    Info("Log Info Message: %v", "info")
    Warn("Log Warn Message: %v", "warn")
    Error("Log Error Message: %v", "error")

    r, _ := ioutil.ReadFile("testing.log")

    if strings.Contains(string(r), "Log Debug Message") {
        t.Errorf("Should not contain debug information if level = %v", GetLevel())
    }

    if !strings.Contains(string(r), "Log Info Message") {
        t.Errorf("Should contain info information if level = %v", GetLevel())
    }

    if !strings.Contains(string(r), "Log Warn Message") {
        t.Errorf("Should contain warn information if level = %v", GetLevel())
    }

    if !strings.Contains(string(r), "Log Error Message") {
        t.Errorf("Should contain error information if level = %v", GetLevel())
    }
}

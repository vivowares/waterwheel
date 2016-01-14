package waterwheel

import (
	"bytes"
	"regexp"
	"testing"
)

func TestSyncLoggerDebug(t *testing.T) {
	t.Parallel()

	size := 5
	buf := &bytes.Buffer{}
	b := NopCloserBuf{buf}
	w := NewBufferedWriteCloser(size, &b)
	logger := NewSyncLogger(w, SimpleFormatter, "debug")
	logger.Info("info message")
	logger.Debug("debug message")
	logger.Close()

	matchInfo, _ := regexp.MatchString("INFO", string(buf.Bytes()))
	matchDebug, _ := regexp.MatchString("DEBUG", string(buf.Bytes()))
	if !matchInfo || !matchDebug {
		t.Fatalf("logger level filtering doesn't work for DEBUG")
	}
}

func TestSyncLoggerInfo(t *testing.T) {
	t.Parallel()

	size := 5
	buf := &bytes.Buffer{}
	b := NopCloserBuf{buf}
	w := NewBufferedWriteCloser(size, &b)
	logger := NewSyncLogger(w, SimpleFormatter, "info")
	logger.Warn("warn message")
	logger.Info("info message")
	logger.Debug("debug message")
	logger.Close()

	matchInfo, _ := regexp.MatchString("INFO", string(buf.Bytes()))
	matchWarn, _ := regexp.MatchString("WARN", string(buf.Bytes()))
	matchDebug, _ := regexp.MatchString("DEBUG", string(buf.Bytes()))
	if !matchInfo || !matchWarn || matchDebug {
		t.Fatalf("logger level filtering doesn't work for INFO")
	}
}

func TestSyncLoggerWarn(t *testing.T) {
	t.Parallel()

	size := 5
	buf := &bytes.Buffer{}
	b := NopCloserBuf{buf}
	w := NewBufferedWriteCloser(size, &b)
	logger := NewSyncLogger(w, SimpleFormatter, "warn")
	logger.Error("error message")
	logger.Warn("warn message")
	logger.Info("info message")
	logger.Close()

	matchError, _ := regexp.MatchString("ERROR", string(buf.Bytes()))
	matchWarn, _ := regexp.MatchString("WARN", string(buf.Bytes()))
	matchInfo, _ := regexp.MatchString("INFO", string(buf.Bytes()))
	if !matchError || !matchWarn || matchInfo {
		t.Fatalf("logger level filtering doesn't work for WARN")
	}
}

func TestSyncLoggerError(t *testing.T) {
	t.Parallel()

	size := 5
	buf := &bytes.Buffer{}
	b := NopCloserBuf{buf}
	w := NewBufferedWriteCloser(size, &b)
	logger := NewSyncLogger(w, SimpleFormatter, "error")
	logger.Critical("critical message")
	logger.Error("error message")
	logger.Warn("warn message")
	logger.Close()

	matchCritical, _ := regexp.MatchString("CRITICAL", string(buf.Bytes()))
	matchError, _ := regexp.MatchString("ERROR", string(buf.Bytes()))
	matchWarn, _ := regexp.MatchString("WARN", string(buf.Bytes()))
	if !matchCritical || !matchError || matchWarn {
		t.Fatalf("logger level filtering doesn't work for ERROR")
	}
}

func TestSyncLoggerCritical(t *testing.T) {
	t.Parallel()

	size := 5
	buf := &bytes.Buffer{}
	b := NopCloserBuf{buf}
	w := NewBufferedWriteCloser(size, &b)
	logger := NewSyncLogger(w, SimpleFormatter, "critical")
	logger.Critical("critical message")
	logger.Error("error message")
	logger.Close()

	matchCritical, _ := regexp.MatchString("CRITICAL", string(buf.Bytes()))
	matchError, _ := regexp.MatchString("ERROR", string(buf.Bytes()))
	if !matchCritical || matchError {
		t.Fatalf("logger level filtering doesn't work for CRITICAL")
	}
}

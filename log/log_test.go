package log

import (
	"testing"
	"time"
)

func TestInfo(t *testing.T) {
	Info("123")
	Info("", 5, 1.1, "test")
	Info("%d %s", 5, "test")
	InnerLog("123")
}

func TestDebug(t *testing.T) {
	Debug("123")
	Debug("", 5, 1.1, "test", time.Now())
}

func TestError(t *testing.T) {
	Error("123")
	Error("", 5, 1.1, "test", time.Now())
}

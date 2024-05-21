package logger

import (
    "testing"
)

func TestLoggerInitialization(t *testing.T) {
    if InfoLogger == nil {
        t.Error("InfoLogger not initialized")
    }
    if ErrorLogger == nil {
        t.Error("ErrorLogger not initialized")
    }
}

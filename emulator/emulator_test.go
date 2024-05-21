package emulator

import (
    "testing"
    "strings"
)

func TestNewBlock(t *testing.T) {
    _, err := NewBlock(-1, "data")
    if err == nil {
        t.Errorf("Expected error for negative block ID, got nil")
    }

    _, err = NewBlock(1, "")
    if err == nil {
        t.Errorf("Expected error for empty block data, got nil")
    }

    block, err := NewBlock(1, "data")
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }

    if block.ID != 1 || block.Data != "data" {
        t.Errorf("Unexpected block values: %v", block)
    }
}

func TestStart(t *testing.T) {
    // Mock logger to capture log output
    var logs string
    logger := func(msg string) {
        logs += msg + "\n"
    }

    // Assuming you can replace the logger functions or capture output in some way
    InfoLogger := logger
    ErrorLogger := logger

    Start(0, 4, 0, 2, 3)

    if !contains(logs, "Starting emulation for Node 0 of 4") {
        t.Errorf("Missing log entry for starting emulation")
    }
    if !contains(logs, "Emulating blocks from 0 to 2 in mode 3") {
        t.Errorf("Missing log entry for block range and mode")
    }
    if !contains(logs, "Block ID: 0, Data: Data for block 0") {
        t.Errorf("Missing log entry for block 0")
    }
    if !contains(logs, "Block ID: 1, Data: Data for block 1") {
        t.Errorf("Missing log entry for block 1")
    }
    if !contains(logs, "Block ID: 2, Data: Data for block 2") {
        t.Errorf("Missing log entry for block 2")
    }
}

func contains(s, substr string) bool {
    return strings.Contains(s, substr)
}

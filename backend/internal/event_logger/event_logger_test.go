package event_logger

import (
	"testing"
)

func TestCreate(t *testing.T) {
	logger, err := New("test.json")

	if logger == nil {
		t.Fatal("file is nil")
	}

	if err {
		t.Fatal("error when attempting to create event logger")
	}
}

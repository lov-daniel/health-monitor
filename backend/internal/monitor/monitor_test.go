package monitor

import (
	"testing"
)

func TestCreate(t *testing.T) {
	mon, err := New("https://twitch.tv")

	if err {
		t.Errorf("error when attempting to create a monitor")
	}

	if mon.id == 0 {
		t.Errorf("monitor failed to create")
	}
}

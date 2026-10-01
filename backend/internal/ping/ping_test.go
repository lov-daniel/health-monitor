package ping

import (
	"testing"
)

func TestPing(t *testing.T) {
	endpoints := []string{
		"https://www.google.com",
		"https://www.caida.org",
		"https://www.daniellov.com",
		"https://fakeurl.com",
	}

	results, err := CheckAll(endpoints)

	if err {
		t.Errorf(`error when attempting to check all endpoints`)
	}

	if len(results) != 4 {
		t.Errorf(`%d pings returned, %d expected`, len(results), len(endpoints))
	}
}

func TestPingEmpty(t *testing.T) {
	endpoints := []string{}

	results, err := CheckAll(endpoints)

	if err {
		t.Errorf(`error when attempting to check all endpoints: `)
	}

	if len(results) > 0 {
		t.Errorf(`%d pings returned, 0 expected`, len(results))
	}
}

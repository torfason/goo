package assert

import (
	"strings"
	"testing"
)

// NoError asserts that no error was received (that it is nil)
func NoError(t *testing.T, err error, message ...string) {
	t.Helper()
	if err != nil {
		messageOut := strings.Join(message, "\n")
		t.Errorf("%s\nhave: an error (%q)\nwant: no error (<nil>)", messageOut, err)
	}
}

// Error asserts that a non-nil error was received
func Error(t *testing.T, err error, message ...string) {
	t.Helper()
	if err == nil {
		messageOut := strings.Join(message, "\n")
		t.Errorf("%s\nhave: no error (<nil>)\nwant: an error", messageOut)
	}
}

// ErrorIs asserts that a specific error was revived
func ErrorIs(t *testing.T, errHave error, errWant error, message ...string) {
	t.Helper()
	if errHave == nil {
		messageOut := strings.Join(message, "\n")
		t.Errorf("%s\nhave: no error (<nil>)\nwant: an error (%q)", messageOut, errWant)
	} else if errHave != errWant {
		messageOut := strings.Join(message, "\n")
		t.Errorf("%s\nhave: an error: (%q)\nwant: an error (%q)", messageOut, errHave, errWant)
	}
}

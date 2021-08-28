package assert

import (
	"math"
	"reflect"
	"strings"
	"testing"
)

const Eps10 = 0.00000000001

// True asserts that a condition is true
func True(t testing.TB, have bool, message ...string) {
	t.Helper()
	if !have {
		messageOut := strings.Join(message, "\n")
		t.Errorf("%s\nhave: false\nwant: true", messageOut)
	}
}

// False asserts that a condition is true
func False(t testing.TB, have bool, message ...string) {
	t.Helper()
	if have {
		messageOut := strings.Join(message, "\n")
		t.Errorf("%s\nhave: true\nwant: false", messageOut)
	}
}

// EqualString asserts that two strings are equal.
func EqualString(t testing.TB, have, want string, message ...string) {
	t.Helper()
	if have != want {
		messageOut := strings.Join(message, "\n")
		t.Errorf("%s\nhave: %q\nwant: %q", messageOut, have, want)
	}
}

// EqualInt asserts that two integers are equal.
func EqualInt(t testing.TB, have, want int, message ...string) {
	t.Helper()
	if have != want {
		messageOut := strings.Join(message, "\n")
		t.Errorf("%s\nhave: %v\nwant: %v", messageOut, have, want)
	}
}

// EqualBool asserts that two integers are equal.
func EqualBool(t testing.TB, have, want bool, message ...string) {
	t.Helper()
	if have != want {
		messageOut := strings.Join(message, "\n")
		t.Errorf("%s\nhave: %v\nwant: %v", messageOut, have, want)
	}
}

// EqualFloat asserts that two integers are equal.
func EqualFloat64(t testing.TB, have, want float64, epsilon float64, message ...string) {
	t.Helper()
	if epsilon <= 0 {
		t.Errorf("\nepsilon must be greater than zero, but is: %v", epsilon)
	}
	diff := math.Abs(have - want)
	if diff > epsilon {
		messageOut := strings.Join(message, "\n")
		t.Errorf("%s\nhave: %.2f\nwant: %.2f\ndiff: %v", messageOut, have, want, diff)
	}
}

// EqualSliceString asserts that two integer slices are equal.
func EqualSliceString(t testing.TB, have, want []string, message ...string) {
	t.Helper()
	if !reflect.DeepEqual(have, want) {
		messageOut := strings.Join(message, "\n")
		t.Errorf("%s\nhave: %v (len:%v)\nwant: %v (len:%v)",
			messageOut, have, len(have), want, len(want))
	}
}

// EqualSliceInt asserts that two integer slices are equal.
func EqualSliceInt(t testing.TB, have, want []int, message ...string) {
	t.Helper()
	if !reflect.DeepEqual(have, want) {
		messageOut := strings.Join(message, "\n")
		t.Errorf("%s\nhave: %v (len:%v)\nwant: %v (len:%v)",
			messageOut, have, len(have), want, len(want))
	}
}

// EqualMapStringString asserts that two string->string maps are equal.
func EqualMapStringString(t testing.TB, have, want map[string]string, message ...string) {
	t.Helper()
	if !reflect.DeepEqual(have, want) {
		messageOut := strings.Join(message, "\n")
		t.Errorf("%s\nhave: %v\nwant: %v", messageOut, have, want)
	}
}

// EqualMapStringString asserts that two string->string maps are equal.
func EqualMapStringBool(t testing.TB, have, want map[string]bool, message ...string) {
	t.Helper()
	if !reflect.DeepEqual(have, want) {
		messageOut := strings.Join(message, "\n")
		t.Errorf("%s\nhave: %v\nwant: %v", messageOut, have, want)
	}
}

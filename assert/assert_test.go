package assert

import "testing"

func TestAssert(t *testing.T) {

	// Very minimal tests
	EqualString(t, "OK", "OK")
	EqualInt(t, 42, 42)
	EqualBool(t, true, true)
	EqualFloat64(t, 3.14, 3+0.14, Eps10)

}

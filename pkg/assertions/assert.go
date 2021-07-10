package assertions

import "testing"

func AssertEquals(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Fatalf("\nexpected: \n%s\nactual: \n%s", expected, actual)
	}
}

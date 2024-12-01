package is

import (
	"fmt"
	"testing"
)

func NoError(t *testing.T, err error, args ...any) {
	if err != nil {
		t.Fatal(append([]any{err}, args...)...)

		return
	}

	t.Log("no error raised")
}

func Equal[T comparable](t *testing.T, expected, result T, args ...any) {
	if expected != result {
		t.Fatal(append([]any{
			fmt.Sprintf("mismatching values %v and %v", expected, result),
		}, args...)...)

		return
	}

	t.Logf("matching values %v and %v", expected, result)
}

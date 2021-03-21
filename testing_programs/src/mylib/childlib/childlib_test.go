package childlib

import (
	"testing"
)

func Test_MoreBasic(t *testing.T) {
	t.Parallel()
	if 10-5 != 5 {
		t.Error("Failed to subtract")
	}
}

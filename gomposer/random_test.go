package gomposer

import (
	"testing"
)

func Test_GenerateString_ReturnsLength(t *testing.T) {
	n := 10
	s := GenerateRandomString(n)
	if an := len(s); an != n {
		t.Errorf("Expected a length of %v got %v", n, an)
	}
}

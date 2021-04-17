package common

import (
	"testing"
)

func TestPrintArray(t *testing.T) {
	array := []string{"hallo", "test"}
	PrintArray(array)

	PrintArray("hallo", 1, "test")
}

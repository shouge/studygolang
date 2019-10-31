package calc

import "testing"

func TestAdd(t *testing.T) {
	var result int
	result = Add(1, 1)
	if result != 2 {
		t.Error("Expected 2, got", result)
	}
}

func TestSubtract(t *testing.T) {
	var result int
	result = Subtract(10, 5)
	if 5 != result {
		t.Error("Expected 5, got", result)
	}
}

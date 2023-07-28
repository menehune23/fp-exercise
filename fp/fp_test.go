package fp_test

import (
	"testing"

	"fp-exercise/fp"
)

func TestComputeTotal_EmptyList(t *testing.T) {
	expect(t, fp.ComputeTotal(nil), 0)
}

func TestComputeTotal_IgnoresNonIntegers(t *testing.T) {
	expect(t, fp.ComputeTotal([]string{
		"1.0",
		"1",
		"2.5",
		"x",
		"-3",
	}), -2)
}

func expect(t *testing.T, actual int, expected int) {
	t.Helper()

	if actual != expected {
		t.Fatalf("expected %d to be %d", actual, expected)
	}
}

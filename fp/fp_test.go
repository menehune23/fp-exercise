package fp_test

import (
	"reflect"
	"testing"

	"fp-exercise/fp"
)

func TestSimulateTrafficLight(t *testing.T) {
	expectEqual(t,
		fp.SimulateTrafficLight("red", 5),
		[]string{"red", "green", "yellow", "red", "green"},
	)
}

func expectEqual(t *testing.T, actual any, expected any) {
	t.Helper()

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %s to equal %s", actual, expected)
	}
}

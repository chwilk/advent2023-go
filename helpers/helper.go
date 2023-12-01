package helpers

import (
	"reflect"
	"testing"
)

const (
	PartA = iota
	PartB
)

func Check(t testing.TB, expect, got int) {
	t.Helper()
	if expect != got {
		t.Errorf("Expected %d, got %d", expect, got)
	}
}

func CheckBytes(t testing.TB, expect, got []byte) {
	t.Helper()
	if !reflect.DeepEqual(expect, got) {
		t.Errorf("Expected %d, got %d", expect, got)
	}
}

func CheckString(t testing.TB, expect, got string) {
	t.Helper()
	if expect != got {
		t.Errorf("Expected %s, got %s", expect, got)
	}
}

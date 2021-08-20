package nook_test

import (
	"testing"

	"github.com/lindsaygelle/nook"
)

func testResident(t *testing.T, animal nook.Key, r nook.Resident) {
	testResidentSpecial(t, r)
}

func testResidentSpecial(t *testing.T, r nook.Resident) {
	if ok := r.Special; !ok {
		t.Fatalf("%s.Special != true", r.Key)
	}
}

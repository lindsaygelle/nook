package nook_test

import (
	"testing"

	"github.com/lindsaygelle/nook"
)

func TestResidents(t *testing.T) {
	for animal, residents := range residents {
		residents.Each(func(k nook.Key, r nook.Resident) {
			testCharacter(t, animal, r.Character)
			testResident(t, animal, r)
		})
	}
}

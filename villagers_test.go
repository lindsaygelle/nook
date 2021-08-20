package nook_test

import (
	"testing"

	"github.com/lindsaygelle/nook"
)

func TestVillagers(t *testing.T) {
	for animal, villagers := range villagers {
		villagers.Each(func(k nook.Key, v nook.Villager) {
			testCharacter(t, animal, v.Character)
			testVillager(t, v)
		})
	}
}

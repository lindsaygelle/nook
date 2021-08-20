package nook_test

import (
	"reflect"
	"testing"

	"github.com/lindsaygelle/nook"
)

func testVillager(t *testing.T, v nook.Villager) {
	testVillagerPersonality(t, v)
	testVillagerPhrase(t, v)
}

func testVillagerPersonality(t *testing.T, v nook.Villager) {
	if ok := reflect.ValueOf(v.Personality).IsZero(); ok {
		t.Fatalf("%s.Personality is a zero value", v.Key)
	}
}

func testVillagerPhrase(t *testing.T, v nook.Villager) {
	if ok := len(v.Phrase) != 0; !ok {
		t.Fatalf("len(%s.Phrase) == 0", v.Key)
	}
}

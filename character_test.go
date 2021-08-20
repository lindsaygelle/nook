package nook_test

import (
	"reflect"
	"testing"

	"github.com/lindsaygelle/nook"
)

func testCharacter(t *testing.T, animal nook.Key, c nook.Character) {
	if ok := len(c.Key) != 0; !ok {
		t.Fatal(c)
	}
	testCharacterAnimal(t, animal, c)
	testCharacterBirthday(t, c)
	testCharacterGender(t, c)
	testCharacterName(t, c)
}

func testCharacterAnimal(t *testing.T, animal nook.Key, c nook.Character) {
	if ok := c.Animal.Key == animal; !ok {
		t.Fatalf("%s.Animal != animal.%s", c.Key, animal)
	}
}

func testCharacterBirthday(t *testing.T, c nook.Character) {
	if ok := c.Birthday.Ok(); !ok {
		t.Logf("%s.Birthday.Ok() != true", c.Key)
	}
}

func testCharacterGender(t *testing.T, c nook.Character) {
	if ok := reflect.ValueOf(c.Gender).IsZero(); ok {
		t.Fatalf("%s.Gender is a zero value", c.Key)
	}
}

func testCharacterName(t *testing.T, c nook.Character) {
	if ok := len(c.Name) != 0; !ok {
		t.Fatalf("len(%s.Name) == 0", c.Key)
	}
}

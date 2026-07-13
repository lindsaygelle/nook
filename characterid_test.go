package nook_test

import (
	"testing"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/character/cat"
)

func TestNewCharacterID(t *testing.T) {
	id := nook.NewCharacterID(animal.Cat.Key, character.Ankha)
	if id != "Cat:Ankha" {
		t.Fatalf("nook.NewCharacterID(cat, Ankha) = %s", id)
	}

	animalKey, characterKey, ok := id.Parts()
	if !ok {
		t.Fatal("nook.NewCharacterID(cat, Ankha).Parts() did not parse")
	}
	if animalKey != animal.Cat.Key || characterKey != character.Ankha {
		t.Fatalf("nook.NewCharacterID(cat, Ankha).Parts() = %s/%s", animalKey, characterKey)
	}
}

func TestNewCharacterIDRejectsIncompleteValues(t *testing.T) {
	if id := nook.NewCharacterID("", character.Ankha); id.Ok() {
		t.Fatalf("nook.NewCharacterID(blank, Ankha).Ok() = true for %s", id)
	}
	if id := nook.CharacterID("cat"); id.Ok() {
		t.Fatalf("nook.CharacterID(cat).Ok() = true for %s", id)
	}
}

func TestCharacterIDMethod(t *testing.T) {
	if id := cat.Ankha.Character.ID(); id != "Cat:Ankha" {
		t.Fatalf("cat.Ankha.Character.ID() = %s", id)
	}
}

package nook_test

import (
	"reflect"
	"testing"

	"github.com/lindsaygelle/nook"
	catcharacters "github.com/lindsaygelle/nook/character/cat"
	dogcharacters "github.com/lindsaygelle/nook/character/dog"
	squirrelcharacters "github.com/lindsaygelle/nook/character/squirrel"
	"github.com/lindsaygelle/nook/game"
)

func testCharacter(t *testing.T, animal nook.Key, c nook.Character) {
	if ok := len(c.Key) != 0; !ok {
		t.Fatal(c)
	}
	testCharacterAnimal(t, animal, c)
	testCharacterBirthday(t, c)
	testCharacterID(t, c)
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

func testCharacterID(t *testing.T, c nook.Character) {
	id := c.ID()
	if id == "" {
		t.Fatalf("%s.ID() is empty", c.Key)
	}
	expected := nook.Key(string(c.Animal.Key) + ":" + string(c.Key))
	if id != expected {
		t.Fatalf("%s.ID() = %s", c.Key, id)
	}
}

func testCharacterGender(t *testing.T, c nook.Character) {
	if ok := reflect.ValueOf(c.Gender).IsZero(); ok {
		t.Fatalf("%s.Gender is a zero value", c.Key)
	}
	if c.Gender.Key == "" {
		t.Fatalf("%s.Gender.Key is empty", c.Key)
	}
}

func testCharacterName(t *testing.T, c nook.Character) {
	if ok := len(c.Name) != 0; !ok {
		t.Fatalf("len(%s.Name) == 0", c.Key)
	}
}

func TestCharacterAppearsInGame(t *testing.T) {
	if !catcharacters.Ankha.Character.AppearsInGame(game.NewLeaf.Key) {
		t.Fatalf("%s.AppearsInGame(%s) = false", catcharacters.Ankha.Key, game.NewLeaf.Key)
	}
	if catcharacters.Ankha.Character.AppearsInGame("missing") {
		t.Fatalf("%s.AppearsInGame(missing) = true", catcharacters.Ankha.Key)
	}
}

func TestCharacterFirstAndLastGame(t *testing.T) {
	first, ok := catcharacters.Ankha.Character.FirstGame()
	if !ok {
		t.Fatalf("%s.FirstGame() not found", catcharacters.Ankha.Key)
	}
	if first.Key != game.DoubutsuNoMoriPlus.Key {
		t.Fatalf("%s.FirstGame().Key = %s", catcharacters.Ankha.Key, first.Key)
	}

	last, ok := dogcharacters.Isabelle.Character.LastGame()
	if !ok {
		t.Fatalf("%s.LastGame() not found", dogcharacters.Isabelle.Key)
	}
	if last.Key != game.NewHorizons.Key {
		t.Fatalf("%s.LastGame().Key = %s", dogcharacters.Isabelle.Key, last.Key)
	}
}

func TestCharacterGameByKey(t *testing.T) {
	found, ok := dogcharacters.Isabelle.Character.GameByKey(game.NewLeaf.Key)
	if !ok {
		t.Fatalf("%s.GameByKey(%s) not found", dogcharacters.Isabelle.Key, game.NewLeaf.Key)
	}
	if found.Key != game.NewLeaf.Key {
		t.Fatalf("%s.GameByKey(%s).Key = %s", dogcharacters.Isabelle.Key, game.NewLeaf.Key, found.Key)
	}

	if _, ok := dogcharacters.Isabelle.Character.GameByKey("missing"); ok {
		t.Fatalf("%s.GameByKey(missing) unexpectedly found a game", dogcharacters.Isabelle.Key)
	}
}

func TestCharacterGamesByReleaseOrder(t *testing.T) {
	games := catcharacters.Ankha.Character.GamesByReleaseOrder()
	if len(games) != len(catcharacters.Ankha.Character.Games) {
		t.Fatalf("len(%s.GamesByReleaseOrder()) = %d", catcharacters.Ankha.Key, len(games))
	}
	if games[0].Key != game.DoubutsuNoMoriPlus.Key {
		t.Fatalf("%s.GamesByReleaseOrder()[0].Key = %s", catcharacters.Ankha.Key, games[0].Key)
	}
	if games[len(games)-1].Key != game.NewHorizons.Key {
		t.Fatalf("%s.GamesByReleaseOrder()[last].Key = %s", catcharacters.Ankha.Key, games[len(games)-1].Key)
	}

	games[0] = game.NewHorizons

	fresh := catcharacters.Ankha.Character.GamesByReleaseOrder()
	if fresh[0].Key != game.DoubutsuNoMoriPlus.Key {
		t.Fatalf("%s.GamesByReleaseOrder()[0].Key after mutation = %s", catcharacters.Ankha.Key, fresh[0].Key)
	}
}

func TestCharacterGamesWithoutKnownAppearances(t *testing.T) {
	if _, ok := squirrelcharacters.Shaki.Character.FirstGame(); ok {
		t.Fatalf("%s.FirstGame() unexpectedly found a game", squirrelcharacters.Shaki.Key)
	}
	if _, ok := squirrelcharacters.Shaki.Character.LastGame(); ok {
		t.Fatalf("%s.LastGame() unexpectedly found a game", squirrelcharacters.Shaki.Key)
	}
}

package catalog_test

import (
	"testing"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/catalog"
	"github.com/lindsaygelle/nook/character"
	catcharacters "github.com/lindsaygelle/nook/character/cat"
	dogcharacters "github.com/lindsaygelle/nook/character/dog"
	squirrelcharacters "github.com/lindsaygelle/nook/character/squirrel"
	"github.com/lindsaygelle/nook/game"
)

func TestCharacterGamesByID(t *testing.T) {
	games, ok := catalog.CharacterGamesByID(" cat:Ankha ")
	if !ok {
		t.Fatal("catalog.CharacterGamesByID(cat:Ankha) not found")
	}

	expected := []nook.Key{
		game.DoubutsuNoMoriPlus.Key,
		game.AnimalCrossing.Key,
		game.DoubutsuNoMoriEPlus.Key,
		game.CityFolk.Key,
		game.NewLeaf.Key,
		game.NewHorizons.Key,
		game.HappyHomeDesigner.Key,
		game.AmiiboFestival.Key,
		game.PocketCamp.Key,
	}

	if len(games) != len(expected) {
		t.Fatalf("len(catalog.CharacterGamesByID(cat:Ankha)) = %d", len(games))
	}

	for i, want := range expected {
		if games[i].Key != want {
			t.Fatalf("catalog.CharacterGamesByID(cat:Ankha)[%d].Key = %s", i, games[i].Key)
		}
	}
}

func TestCharacterGamesReturnsCopy(t *testing.T) {
	games, ok := catalog.CharacterGamesByID("dog:Isabelle")
	if !ok {
		t.Fatal("catalog.CharacterGamesByID(dog:Isabelle) not found")
	}
	if len(games) == 0 {
		t.Fatal("catalog.CharacterGamesByID(dog:Isabelle) returned no games")
	}

	games[0] = game.DoubutsuNoMori

	fresh, ok := catalog.CharacterGamesByID("dog:Isabelle")
	if !ok {
		t.Fatal("catalog.CharacterGamesByID(dog:Isabelle) not found on second lookup")
	}
	if fresh[0].Key != game.NewLeaf.Key {
		t.Fatalf("catalog.CharacterGamesByID(dog:Isabelle)[0].Key = %s", fresh[0].Key)
	}
}

func TestFirstAndLastAppearanceByID(t *testing.T) {
	first, ok := catalog.FirstAppearanceByID(" dog:Isabelle ")
	if !ok {
		t.Fatal("catalog.FirstAppearanceByID(dog:Isabelle) not found")
	}
	if first.Key != game.NewLeaf.Key {
		t.Fatalf("catalog.FirstAppearanceByID(dog:Isabelle).Key = %s", first.Key)
	}

	last, ok := catalog.LastAppearanceByID(" dog:Isabelle ")
	if !ok {
		t.Fatal("catalog.LastAppearanceByID(dog:Isabelle) not found")
	}
	if last.Key != game.PocketCamp.Key {
		t.Fatalf("catalog.LastAppearanceByID(dog:Isabelle).Key = %s", last.Key)
	}
}

func TestFirstAndLastAppearanceByCharacter(t *testing.T) {
	first, ok := catalog.FirstAppearance(catcharacters.Ankha.Character)
	if !ok {
		t.Fatal("catalog.FirstAppearance(Ankha) not found")
	}
	if first.Key != game.DoubutsuNoMoriPlus.Key {
		t.Fatalf("catalog.FirstAppearance(Ankha).Key = %s", first.Key)
	}

	last, ok := catalog.LastAppearance(dogcharacters.Isabelle.Character)
	if !ok {
		t.Fatal("catalog.LastAppearance(Isabelle) not found")
	}
	if last.Key != game.PocketCamp.Key {
		t.Fatalf("catalog.LastAppearance(Isabelle).Key = %s", last.Key)
	}
}

func TestCharacterGamesHandlesNoKnownAppearances(t *testing.T) {
	games, ok := catalog.CharacterGamesByID(string(squirrelcharacters.Shaki.Character.ID()))
	if !ok {
		t.Fatal("catalog.CharacterGamesByID(Shaki) not found")
	}
	if len(games) != 0 {
		t.Fatalf("len(catalog.CharacterGamesByID(Shaki)) = %d", len(games))
	}

	if _, ok := catalog.FirstAppearance(squirrelcharacters.Shaki.Character); ok {
		t.Fatal("catalog.FirstAppearance(Shaki) unexpectedly found a game")
	}
	if _, ok := catalog.LastAppearance(squirrelcharacters.Shaki.Character); ok {
		t.Fatal("catalog.LastAppearance(Shaki) unexpectedly found a game")
	}
}

func TestCharacterGamesMissingID(t *testing.T) {
	if _, ok := catalog.CharacterGamesByID("missing"); ok {
		t.Fatal("catalog.CharacterGamesByID(missing) unexpectedly found a character")
	}
	if _, ok := catalog.FirstAppearanceByID("missing"); ok {
		t.Fatal("catalog.FirstAppearanceByID(missing) unexpectedly found a game")
	}
	if _, ok := catalog.LastAppearanceByID("missing"); ok {
		t.Fatal("catalog.LastAppearanceByID(missing) unexpectedly found a game")
	}
}

func TestCharacterGamesCoverage(t *testing.T) {
	count := 0

	for _, residents := range catalog.AllResidents {
		for _, resident := range residents {
			count++
			if _, ok := catalog.CharacterGames(resident.Character); !ok {
				t.Fatalf("catalog.CharacterGames(%s) not found", resident.Character.ID())
			}
		}
	}

	for _, villagers := range catalog.AllVillagers {
		for _, villager := range villagers {
			count++
			if _, ok := catalog.CharacterGames(villager.Character); !ok {
				t.Fatalf("catalog.CharacterGames(%s) not found", villager.Character.ID())
			}
		}
	}

	if count == 0 {
		t.Fatal("catalog character coverage count = 0")
	}
}

func TestCharacterGamesUsesComposedID(t *testing.T) {
	if _, ok := catalog.CharacterGamesByID(string(character.Carmen)); ok {
		t.Fatal("catalog.CharacterGamesByID(Carmen) unexpectedly found a character without animal scope")
	}
}

func TestResidentsByGame(t *testing.T) {
	residents := catalog.ResidentsByGame(game.NewLeaf.Key)
	if len(residents) == 0 {
		t.Fatal("catalog.ResidentsByGame(NewLeaf) returned no residents")
	}

	foundIsabelle := false
	for i, resident := range residents {
		if !appearsInGame(resident.Character.Games, game.NewLeaf.Key) {
			t.Fatalf("catalog.ResidentsByGame(NewLeaf)[%d] missing NewLeaf appearance", i)
		}
		if resident.Character.Animal.Key == animal.Dog.Key && resident.Character.Key == character.Isabelle {
			foundIsabelle = true
		}
		if i == 0 {
			continue
		}

		prev := residents[i-1]
		if resident.Character.Animal.Key < prev.Character.Animal.Key {
			t.Fatalf("catalog.ResidentsByGame(NewLeaf)[%d] not sorted by animal key", i)
		}
		if resident.Character.Animal.Key == prev.Character.Animal.Key && resident.Character.Key < prev.Character.Key {
			t.Fatalf("catalog.ResidentsByGame(NewLeaf)[%d] not sorted by character key", i)
		}
	}
	if !foundIsabelle {
		t.Fatal("catalog.ResidentsByGame(NewLeaf) missing Isabelle")
	}
}

func TestResidentsByGameEmptyKey(t *testing.T) {
	residents := catalog.ResidentsByGame("")
	if residents != nil {
		t.Fatalf("catalog.ResidentsByGame(\"\") = %#v", residents)
	}
}

func TestVillagersByGame(t *testing.T) {
	villagers := catalog.VillagersByGame(game.DoubutsuNoMoriPlus.Key)
	if len(villagers) == 0 {
		t.Fatal("catalog.VillagersByGame(DoubutsuNoMoriPlus) returned no villagers")
	}

	foundAnkha := false
	for i, villager := range villagers {
		if !appearsInGame(villager.Character.Games, game.DoubutsuNoMoriPlus.Key) {
			t.Fatalf("catalog.VillagersByGame(DoubutsuNoMoriPlus)[%d] missing DoubutsuNoMoriPlus appearance", i)
		}
		if villager.Character.Animal.Key == animal.Cat.Key && villager.Character.Key == character.Ankha {
			foundAnkha = true
		}
		if i == 0 {
			continue
		}

		prev := villagers[i-1]
		if villager.Character.Animal.Key < prev.Character.Animal.Key {
			t.Fatalf("catalog.VillagersByGame(DoubutsuNoMoriPlus)[%d] not sorted by animal key", i)
		}
		if villager.Character.Animal.Key == prev.Character.Animal.Key && villager.Character.Key < prev.Character.Key {
			t.Fatalf("catalog.VillagersByGame(DoubutsuNoMoriPlus)[%d] not sorted by character key", i)
		}
	}
	if !foundAnkha {
		t.Fatal("catalog.VillagersByGame(DoubutsuNoMoriPlus) missing Ankha")
	}
}

func TestVillagersByGameEmptyKey(t *testing.T) {
	villagers := catalog.VillagersByGame("")
	if villagers != nil {
		t.Fatalf("catalog.VillagersByGame(\"\") = %#v", villagers)
	}
}

func appearsInGame(games []nook.Game, gameKey nook.Key) bool {
	for _, appearance := range games {
		if appearance.Key == gameKey {
			return true
		}
	}
	return false
}

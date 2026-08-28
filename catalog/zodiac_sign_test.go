package catalog_test

import (
	"testing"

	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/catalog"
	"github.com/lindsaygelle/nook/character"
)

func TestResidentsByZodiacSign(t *testing.T) {
	residents := catalog.ResidentsByZodiacSign("Sagittarius")
	if len(residents) == 0 {
		t.Fatal("catalog.ResidentsByZodiacSign(Sagittarius) returned no residents")
	}

	foundIsabelle := false
	for i, resident := range residents {
		if !resident.Character.HasZodiacSign("Sagittarius") {
			t.Fatalf("catalog.ResidentsByZodiacSign(Sagittarius)[%d] missing Sagittarius zodiac sign", i)
		}
		if resident.Character.Animal.Key == animal.Dog.Key && resident.Character.Key == character.Isabelle {
			foundIsabelle = true
		}
		if i == 0 {
			continue
		}

		prev := residents[i-1]
		if resident.Character.Animal.Key < prev.Character.Animal.Key {
			t.Fatalf("catalog.ResidentsByZodiacSign(Sagittarius)[%d] not sorted by animal key", i)
		}
		if resident.Character.Animal.Key == prev.Character.Animal.Key && resident.Character.Key < prev.Character.Key {
			t.Fatalf("catalog.ResidentsByZodiacSign(Sagittarius)[%d] not sorted by character key", i)
		}
	}
	if !foundIsabelle {
		t.Fatal("catalog.ResidentsByZodiacSign(Sagittarius) missing Isabelle")
	}
}

func TestResidentsByZodiacSignEmptyKey(t *testing.T) {
	residents := catalog.ResidentsByZodiacSign("")
	if residents != nil {
		t.Fatalf("catalog.ResidentsByZodiacSign(\"\") = %#v", residents)
	}
}

func TestVillagersByZodiacSign(t *testing.T) {
	villagers := catalog.VillagersByZodiacSign("Virgo")
	if len(villagers) == 0 {
		t.Fatal("catalog.VillagersByZodiacSign(Virgo) returned no villagers")
	}

	foundAnkha := false
	for i, villager := range villagers {
		if !villager.Character.HasZodiacSign("Virgo") {
			t.Fatalf("catalog.VillagersByZodiacSign(Virgo)[%d] missing Virgo zodiac sign", i)
		}
		if villager.Character.Animal.Key == animal.Cat.Key && villager.Character.Key == character.Ankha {
			foundAnkha = true
		}
		if i == 0 {
			continue
		}

		prev := villagers[i-1]
		if villager.Character.Animal.Key < prev.Character.Animal.Key {
			t.Fatalf("catalog.VillagersByZodiacSign(Virgo)[%d] not sorted by animal key", i)
		}
		if villager.Character.Animal.Key == prev.Character.Animal.Key && villager.Character.Key < prev.Character.Key {
			t.Fatalf("catalog.VillagersByZodiacSign(Virgo)[%d] not sorted by character key", i)
		}
	}
	if !foundAnkha {
		t.Fatal("catalog.VillagersByZodiacSign(Virgo) missing Ankha")
	}
}

func TestVillagersByZodiacSignEmptyKey(t *testing.T) {
	villagers := catalog.VillagersByZodiacSign("")
	if villagers != nil {
		t.Fatalf("catalog.VillagersByZodiacSign(\"\") = %#v", villagers)
	}
}

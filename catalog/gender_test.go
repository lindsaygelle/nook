package catalog_test

import (
	"testing"

	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/catalog"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
)

func TestResidentsByGender(t *testing.T) {
	residents := catalog.ResidentsByGender(gender.Female.Key)
	if len(residents) == 0 {
		t.Fatal("catalog.ResidentsByGender(Female) returned no residents")
	}

	foundIsabelle := false
	for i, resident := range residents {
		if resident.Character.Gender.Key != gender.Female.Key {
			t.Fatalf("catalog.ResidentsByGender(Female)[%d].Gender.Key = %s", i, resident.Character.Gender.Key)
		}
		if resident.Character.Animal.Key == animal.Dog.Key && resident.Character.Key == character.Isabelle {
			foundIsabelle = true
		}
		if i == 0 {
			continue
		}

		prev := residents[i-1]
		if resident.Character.Animal.Key < prev.Character.Animal.Key {
			t.Fatalf("catalog.ResidentsByGender(Female)[%d] not sorted by animal key", i)
		}
		if resident.Character.Animal.Key == prev.Character.Animal.Key && resident.Character.Key < prev.Character.Key {
			t.Fatalf("catalog.ResidentsByGender(Female)[%d] not sorted by character key", i)
		}
	}
	if !foundIsabelle {
		t.Fatal("catalog.ResidentsByGender(Female) missing Isabelle")
	}
}

func TestResidentsByGenderEmptyKey(t *testing.T) {
	residents := catalog.ResidentsByGender("")
	if residents != nil {
		t.Fatalf("catalog.ResidentsByGender(\"\") = %#v", residents)
	}
}

func TestVillagersByGender(t *testing.T) {
	villagers := catalog.VillagersByGender(gender.Male.Key)
	if len(villagers) == 0 {
		t.Fatal("catalog.VillagersByGender(Male) returned no villagers")
	}

	foundAlfonso := false
	for i, villager := range villagers {
		if villager.Character.Gender.Key != gender.Male.Key {
			t.Fatalf("catalog.VillagersByGender(Male)[%d].Gender.Key = %s", i, villager.Character.Gender.Key)
		}
		if villager.Character.Animal.Key == animal.Alligator.Key && villager.Character.Key == character.Alfonso {
			foundAlfonso = true
		}
		if i == 0 {
			continue
		}

		prev := villagers[i-1]
		if villager.Character.Animal.Key < prev.Character.Animal.Key {
			t.Fatalf("catalog.VillagersByGender(Male)[%d] not sorted by animal key", i)
		}
		if villager.Character.Animal.Key == prev.Character.Animal.Key && villager.Character.Key < prev.Character.Key {
			t.Fatalf("catalog.VillagersByGender(Male)[%d] not sorted by character key", i)
		}
	}
	if !foundAlfonso {
		t.Fatal("catalog.VillagersByGender(Male) missing Alfonso")
	}
}

func TestVillagersByGenderEmptyKey(t *testing.T) {
	villagers := catalog.VillagersByGender("")
	if villagers != nil {
		t.Fatalf("catalog.VillagersByGender(\"\") = %#v", villagers)
	}
}

package catalog_test

import (
	"testing"

	"github.com/lindsaygelle/nook/catalog"
	"github.com/lindsaygelle/nook/character"
	"golang.org/x/text/language"
)

func TestNormalizeLookupValue(t *testing.T) {
	got := catalog.NormalizeLookupValue("  ToM NoOk  ")
	if got != "tom nook" {
		t.Fatalf("catalog.NormalizeLookupValue(...) = %q", got)
	}
}

func TestVillagersByNameFindsDuplicateNamesDeterministically(t *testing.T) {
	villagers := catalog.VillagersByName(language.AmericanEnglish, "  carmen ")
	if len(villagers) != 2 {
		t.Fatalf("len(catalog.VillagersByName(en-US, carmen)) = %d", len(villagers))
	}
	if villagers[0].Animal.Key != "Mouse" || villagers[0].Key != character.Carmen {
		t.Fatalf("catalog.VillagersByName(en-US, carmen)[0] = %s/%s", villagers[0].Animal.Key, villagers[0].Key)
	}
	if villagers[1].Animal.Key != "Rabbit" || villagers[1].Key != character.Carmen {
		t.Fatalf("catalog.VillagersByName(en-US, carmen)[1] = %s/%s", villagers[1].Animal.Key, villagers[1].Key)
	}
}

func TestVillagersByNameMatchesLocaleExactlyAfterNormalization(t *testing.T) {
	villagers := catalog.VillagersByName(language.French, " zoé ")
	if len(villagers) != 2 {
		t.Fatalf("len(catalog.VillagersByName(fr, zoé)) = %d", len(villagers))
	}
	if villagers[0].Animal.Key != "Anteater" || villagers[0].Key != character.Zoe {
		t.Fatalf("catalog.VillagersByName(fr, zoé)[0] = %s/%s", villagers[0].Animal.Key, villagers[0].Key)
	}
	if villagers[1].Animal.Key != "Rabbit" || villagers[1].Key != character.Carmen {
		t.Fatalf("catalog.VillagersByName(fr, zoé)[1] = %s/%s", villagers[1].Animal.Key, villagers[1].Key)
	}

	villagers = catalog.VillagersByName(language.French, "zoe")
	if len(villagers) != 0 {
		t.Fatalf("len(catalog.VillagersByName(fr, zoe)) = %d", len(villagers))
	}
}

func TestResidentsByName(t *testing.T) {
	residents := catalog.ResidentsByName(language.AmericanEnglish, " tom nook ")
	if len(residents) != 1 {
		t.Fatalf("len(catalog.ResidentsByName(en-US, tom nook)) = %d", len(residents))
	}
	if residents[0].Key != character.TomNook {
		t.Fatalf("catalog.ResidentsByName(en-US, tom nook)[0].Key = %s", residents[0].Key)
	}
}

func TestNameLookupBlankQuery(t *testing.T) {
	if residents := catalog.ResidentsByName(language.AmericanEnglish, "   "); len(residents) != 0 {
		t.Fatalf("len(catalog.ResidentsByName(en-US, blank)) = %d", len(residents))
	}
	if villagers := catalog.VillagersByName(language.AmericanEnglish, ""); len(villagers) != 0 {
		t.Fatalf("len(catalog.VillagersByName(en-US, blank)) = %d", len(villagers))
	}
}

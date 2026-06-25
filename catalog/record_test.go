package catalog_test

import (
	"testing"

	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/catalog"
	"github.com/lindsaygelle/nook/character/cat"
	"github.com/lindsaygelle/nook/character/mouse"
	"github.com/lindsaygelle/nook/character/raccoon"
	"golang.org/x/text/language"
)

func TestCharacterRecordOf(t *testing.T) {
	record := catalog.CharacterRecordOf(cat.Ankha.Character)

	if record.AnimalKey != string(animal.Cat.Key) {
		t.Fatalf("catalog.CharacterRecordOf(cat.Ankha).AnimalKey = %s", record.AnimalKey)
	}
	if record.Key != "Ankha" {
		t.Fatalf("catalog.CharacterRecordOf(cat.Ankha).Key = %s", record.Key)
	}
	if record.Birthday.Day != 22 || record.Birthday.Month != 9 {
		t.Fatalf("catalog.CharacterRecordOf(cat.Ankha).Birthday = %#v", record.Birthday)
	}
	if record.Name[language.AmericanEnglish.String()] != "Ankha" {
		t.Fatalf("catalog.CharacterRecordOf(cat.Ankha).Name[en-US] = %s", record.Name[language.AmericanEnglish.String()])
	}
}

func TestVillagerRecordOf(t *testing.T) {
	record := catalog.VillagerRecordOf(cat.Ankha)

	if record.Special {
		t.Fatal("catalog.VillagerRecordOf(cat.Ankha).Special unexpectedly true")
	}
	if record.Personality[language.AmericanEnglish.String()] != "Snooty" {
		t.Fatalf("catalog.VillagerRecordOf(cat.Ankha).Personality[en-US] = %s", record.Personality[language.AmericanEnglish.String()])
	}
	if record.Phrase[language.AmericanEnglish.String()] != "me meow" {
		t.Fatalf("catalog.VillagerRecordOf(cat.Ankha).Phrase[en-US] = %s", record.Phrase[language.AmericanEnglish.String()])
	}
}

func TestResidentRecordOf(t *testing.T) {
	record := catalog.ResidentRecordOf(raccoon.TomNook)

	if !record.Special {
		t.Fatal("catalog.ResidentRecordOf(raccoon.TomNook).Special unexpectedly false")
	}
	if record.Code != "rcn/rco" {
		t.Fatalf("catalog.ResidentRecordOf(raccoon.TomNook).Code = %s", record.Code)
	}
}

func TestLocalizedValuesOfOmitsEmptyValues(t *testing.T) {
	values := catalog.LocalizedValuesOf(mouse.Carmen.Name)

	if values[language.AmericanEnglish.String()] != "Carmen" {
		t.Fatalf("catalog.LocalizedValuesOf(mouse.Carmen.Name)[en-US] = %s", values[language.AmericanEnglish.String()])
	}
	if _, ok := values[language.CanadianFrench.String()]; ok {
		t.Fatal("catalog.LocalizedValuesOf(mouse.Carmen.Name) unexpectedly included an empty value")
	}
}

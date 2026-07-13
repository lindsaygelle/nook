package catalog_test

import (
	"testing"
	"time"

	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/catalog"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/character/cat"
	"github.com/lindsaygelle/nook/character/dog"
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

func TestResidentRecordByCode(t *testing.T) {
	record, ok := catalog.ResidentRecordByCode(" rcn/rco ")
	if !ok {
		t.Fatal("catalog.ResidentRecordByCode(rcn/rco) not found")
	}
	if record.Key != string(character.TomNook) {
		t.Fatalf("catalog.ResidentRecordByCode(rcn/rco).Key = %s", record.Key)
	}
	if !record.Special {
		t.Fatal("catalog.ResidentRecordByCode(rcn/rco).Special unexpectedly false")
	}
}

func TestResidentRecordsByAnimal(t *testing.T) {
	records, ok := catalog.ResidentRecordsByAnimal(animal.Dog.Key)
	if !ok {
		t.Fatalf("catalog.ResidentRecordsByAnimal(%s) not found", animal.Dog.Key)
	}
	if len(records) != len(dog.Residents) {
		t.Fatalf("catalog.ResidentRecordsByAnimal(%s) length = %d", animal.Dog.Key, len(records))
	}
	if records[0].Key != string(character.Booker) {
		t.Fatalf("catalog.ResidentRecordsByAnimal(%s)[0].Key = %s", animal.Dog.Key, records[0].Key)
	}
	if records[len(records)-1].Key != string(character.Serena) {
		t.Fatalf("catalog.ResidentRecordsByAnimal(%s)[last].Key = %s", animal.Dog.Key, records[len(records)-1].Key)
	}
}

func TestResidentRecordsByBirthday(t *testing.T) {
	records := catalog.ResidentRecordsByBirthday(time.December, 20)
	if len(records) != 2 {
		t.Fatalf("len(catalog.ResidentRecordsByBirthday(December, 20)) = %d", len(records))
	}
	if records[0].Key != string(character.Digby) {
		t.Fatalf("catalog.ResidentRecordsByBirthday(December, 20)[0].Key = %s", records[0].Key)
	}
	if records[1].Key != string(character.Isabelle) {
		t.Fatalf("catalog.ResidentRecordsByBirthday(December, 20)[1].Key = %s", records[1].Key)
	}
}

func TestResidentRecordsByBirthMonth(t *testing.T) {
	records := catalog.ResidentRecordsByBirthMonth(time.December)
	if len(records) != 6 {
		t.Fatalf("len(catalog.ResidentRecordsByBirthMonth(December)) = %d", len(records))
	}
	if records[0].Key != string(character.Chip) {
		t.Fatalf("catalog.ResidentRecordsByBirthMonth(December)[0].Key = %s", records[0].Key)
	}
	if records[len(records)-1].Key != string(character.Tortimer) {
		t.Fatalf("catalog.ResidentRecordsByBirthMonth(December)[last].Key = %s", records[len(records)-1].Key)
	}
	if records[0].Birthday.Month != uint8(time.December) {
		t.Fatalf("catalog.ResidentRecordsByBirthMonth(December)[0].Birthday.Month = %d", records[0].Birthday.Month)
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

func TestVillagerRecordByCode(t *testing.T) {
	record, ok := catalog.VillagerRecordByCode(" DUK13 ")
	if !ok {
		t.Fatal("catalog.VillagerRecordByCode(duk13) not found")
	}
	if record.Key != string(character.Ketchup) {
		t.Fatalf("catalog.VillagerRecordByCode(duk13).Key = %s", record.Key)
	}
	if record.Phrase[language.AmericanEnglish.String()] != "bitty" {
		t.Fatalf("catalog.VillagerRecordByCode(duk13).Phrase[en-US] = %s", record.Phrase[language.AmericanEnglish.String()])
	}
}

func TestVillagerRecordsByAnimal(t *testing.T) {
	records, ok := catalog.VillagerRecordsByAnimal(animal.Mouse.Key)
	if !ok {
		t.Fatalf("catalog.VillagerRecordsByAnimal(%s) not found", animal.Mouse.Key)
	}
	if len(records) != len(mouse.Villagers) {
		t.Fatalf("catalog.VillagerRecordsByAnimal(%s) length = %d", animal.Mouse.Key, len(records))
	}
	if records[0].Key != string(character.Anicotti) {
		t.Fatalf("catalog.VillagerRecordsByAnimal(%s)[0].Key = %s", animal.Mouse.Key, records[0].Key)
	}
	if records[len(records)-1].Key != string(character.Samson) {
		t.Fatalf("catalog.VillagerRecordsByAnimal(%s)[last].Key = %s", animal.Mouse.Key, records[len(records)-1].Key)
	}
	if records[0].AnimalKey != string(animal.Mouse.Key) {
		t.Fatalf("catalog.VillagerRecordsByAnimal(%s)[0].AnimalKey = %s", animal.Mouse.Key, records[0].AnimalKey)
	}
}

func TestVillagerRecordsByBirthday(t *testing.T) {
	records := catalog.VillagerRecordsByBirthday(time.April, 23)
	if len(records) != 2 {
		t.Fatalf("len(catalog.VillagerRecordsByBirthday(April, 23)) = %d", len(records))
	}
	if records[0].Key != string(character.Miranda) {
		t.Fatalf("catalog.VillagerRecordsByBirthday(April, 23)[0].Key = %s", records[0].Key)
	}
	if records[1].Key != string(character.Tammi) {
		t.Fatalf("catalog.VillagerRecordsByBirthday(April, 23)[1].Key = %s", records[1].Key)
	}
	if records[0].Birthday.Day != 23 || records[0].Birthday.Month != uint8(time.April) {
		t.Fatalf("catalog.VillagerRecordsByBirthday(April, 23)[0].Birthday = %#v", records[0].Birthday)
	}
}

func TestVillagerRecordsByPersonality(t *testing.T) {
	records := catalog.VillagerRecordsByPersonality(language.AmericanEnglish, " snooty ")
	if len(records) != 66 {
		t.Fatalf("len(catalog.VillagerRecordsByPersonality(en-US, snooty)) = %d", len(records))
	}
	if records[0].Key != string(character.Alli) {
		t.Fatalf("catalog.VillagerRecordsByPersonality(en-US, snooty)[0].Key = %s", records[0].Key)
	}
	if records[len(records)-1].Key != string(character.Whitney) {
		t.Fatalf("catalog.VillagerRecordsByPersonality(en-US, snooty)[last].Key = %s", records[len(records)-1].Key)
	}
	if records[0].Personality[language.AmericanEnglish.String()] != "Snooty" {
		t.Fatalf("catalog.VillagerRecordsByPersonality(en-US, snooty)[0].Personality[en-US] = %s", records[0].Personality[language.AmericanEnglish.String()])
	}
}

func TestVillagerRecordsByBirthMonth(t *testing.T) {
	records := catalog.VillagerRecordsByBirthMonth(time.September)
	if len(records) != 32 {
		t.Fatalf("len(catalog.VillagerRecordsByBirthMonth(September)) = %d", len(records))
	}
	if records[0].Key != string(character.Beardo) {
		t.Fatalf("catalog.VillagerRecordsByBirthMonth(September)[0].Key = %s", records[0].Key)
	}
	if records[len(records)-1].Key != string(character.Whitney) {
		t.Fatalf("catalog.VillagerRecordsByBirthMonth(September)[last].Key = %s", records[len(records)-1].Key)
	}
	if records[0].Birthday.Month != uint8(time.September) {
		t.Fatalf("catalog.VillagerRecordsByBirthMonth(September)[0].Birthday.Month = %d", records[0].Birthday.Month)
	}
}

func TestRecordHelpersMissingAnimalBucket(t *testing.T) {
	if _, ok := catalog.ResidentRecordsByAnimal("missing"); ok {
		t.Fatal("catalog.ResidentRecordsByAnimal(missing) unexpectedly found a bucket")
	}
	if _, ok := catalog.VillagerRecordsByAnimal("missing"); ok {
		t.Fatal("catalog.VillagerRecordsByAnimal(missing) unexpectedly found a bucket")
	}
	if _, ok := catalog.ResidentRecordByCode(""); ok {
		t.Fatal("catalog.ResidentRecordByCode(blank) unexpectedly found a character")
	}
	if _, ok := catalog.VillagerRecordByCode("missing"); ok {
		t.Fatal("catalog.VillagerRecordByCode(missing) unexpectedly found a character")
	}
	if records := catalog.ResidentRecordsByBirthday(0, 0); len(records) != 0 {
		t.Fatalf("len(catalog.ResidentRecordsByBirthday(0, 0)) = %d", len(records))
	}
	if records := catalog.VillagerRecordsByBirthday(time.March, 0); len(records) != 0 {
		t.Fatalf("len(catalog.VillagerRecordsByBirthday(March, 0)) = %d", len(records))
	}
	if records := catalog.ResidentRecordsByBirthMonth(0); len(records) != 0 {
		t.Fatalf("len(catalog.ResidentRecordsByBirthMonth(0)) = %d", len(records))
	}
	if records := catalog.VillagerRecordsByBirthMonth(0); len(records) != 0 {
		t.Fatalf("len(catalog.VillagerRecordsByBirthMonth(0)) = %d", len(records))
	}
	if records := catalog.VillagerRecordsByPersonality(language.AmericanEnglish, ""); len(records) != 0 {
		t.Fatalf("len(catalog.VillagerRecordsByPersonality(en-US, blank)) = %d", len(records))
	}
	if records := catalog.VillagerRecordsByPersonality(language.French, "snooty"); len(records) != 0 {
		t.Fatalf("len(catalog.VillagerRecordsByPersonality(fr, snooty)) = %d", len(records))
	}
}

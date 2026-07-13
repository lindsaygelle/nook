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
	"github.com/lindsaygelle/nook/game"
	"github.com/lindsaygelle/nook/gender"
	"golang.org/x/text/language"
)

func TestGameRecordOf(t *testing.T) {
	record := catalog.GameRecordOf(game.NewLeaf)

	if record.Key != string(game.NewLeaf.Key) {
		t.Fatalf("catalog.GameRecordOf(game.NewLeaf).Key = %s", record.Key)
	}
	if record.Name[language.AmericanEnglish.String()] != "Animal Crossing: New Leaf" {
		t.Fatalf("catalog.GameRecordOf(game.NewLeaf).Name[en-US] = %s", record.Name[language.AmericanEnglish.String()])
	}
}

func TestCharacterRecordOf(t *testing.T) {
	record := catalog.CharacterRecordOf(cat.Ankha.Character)

	if record.AnimalKey != string(animal.Cat.Key) {
		t.Fatalf("catalog.CharacterRecordOf(cat.Ankha).AnimalKey = %s", record.AnimalKey)
	}
	if record.ID != "Cat:Ankha" {
		t.Fatalf("catalog.CharacterRecordOf(cat.Ankha).ID = %s", record.ID)
	}
	if record.Key != "Ankha" {
		t.Fatalf("catalog.CharacterRecordOf(cat.Ankha).Key = %s", record.Key)
	}
	if record.GenderKey != "Female" {
		t.Fatalf("catalog.CharacterRecordOf(cat.Ankha).GenderKey = %s", record.GenderKey)
	}
	if record.Birthday.Day != 22 || record.Birthday.Month != 9 {
		t.Fatalf("catalog.CharacterRecordOf(cat.Ankha).Birthday = %#v", record.Birthday)
	}
	if record.Name[language.AmericanEnglish.String()] != "Ankha" {
		t.Fatalf("catalog.CharacterRecordOf(cat.Ankha).Name[en-US] = %s", record.Name[language.AmericanEnglish.String()])
	}
	if len(record.Games) != 9 {
		t.Fatalf("len(catalog.CharacterRecordOf(cat.Ankha).Games) = %d", len(record.Games))
	}
	if record.Games[0].Key != string(game.DoubutsuNoMoriPlus.Key) {
		t.Fatalf("catalog.CharacterRecordOf(cat.Ankha).Games[0].Key = %s", record.Games[0].Key)
	}
	if record.Games[len(record.Games)-1].Key != string(game.PocketCamp.Key) {
		t.Fatalf("catalog.CharacterRecordOf(cat.Ankha).Games[last].Key = %s", record.Games[len(record.Games)-1].Key)
	}
}

func TestVillagerRecordOf(t *testing.T) {
	record := catalog.VillagerRecordOf(cat.Ankha)

	if record.Special {
		t.Fatal("catalog.VillagerRecordOf(cat.Ankha).Special unexpectedly true")
	}
	if record.PersonalityKey != "Snooty" {
		t.Fatalf("catalog.VillagerRecordOf(cat.Ankha).PersonalityKey = %s", record.PersonalityKey)
	}
	if record.Personality[language.AmericanEnglish.String()] != "Snooty" {
		t.Fatalf("catalog.VillagerRecordOf(cat.Ankha).Personality[en-US] = %s", record.Personality[language.AmericanEnglish.String()])
	}
	if record.Phrase[language.AmericanEnglish.String()] != "me meow" {
		t.Fatalf("catalog.VillagerRecordOf(cat.Ankha).Phrase[en-US] = %s", record.Phrase[language.AmericanEnglish.String()])
	}
	if len(record.Games) != 9 {
		t.Fatalf("len(catalog.VillagerRecordOf(cat.Ankha).Games) = %d", len(record.Games))
	}
}

func TestResidentRecordOf(t *testing.T) {
	record := catalog.ResidentRecordOf(raccoon.TomNook)

	if !record.Special {
		t.Fatal("catalog.ResidentRecordOf(raccoon.TomNook).Special unexpectedly false")
	}
	if record.ID != "Raccoon:TomNook" {
		t.Fatalf("catalog.ResidentRecordOf(raccoon.TomNook).ID = %s", record.ID)
	}
	if record.Code != "rcn/rco" {
		t.Fatalf("catalog.ResidentRecordOf(raccoon.TomNook).Code = %s", record.Code)
	}
	if len(record.Games) == 0 {
		t.Fatal("catalog.ResidentRecordOf(raccoon.TomNook).Games returned no games")
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

func TestResidentRecordByID(t *testing.T) {
	record, ok := catalog.ResidentRecordByID("raccoon:TomNook")
	if !ok {
		t.Fatal("catalog.ResidentRecordByID(raccoon:TomNook) not found")
	}
	if record.Key != string(character.TomNook) {
		t.Fatalf("catalog.ResidentRecordByID(raccoon:TomNook).Key = %s", record.Key)
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

func TestResidentRecordsByGame(t *testing.T) {
	records := catalog.ResidentRecordsByGame(game.NewLeaf.Key)
	if len(records) == 0 {
		t.Fatal("catalog.ResidentRecordsByGame(NewLeaf) returned no records")
	}

	foundTomNook := false
	for i, record := range records {
		if record.ID == "Raccoon:TomNook" {
			foundTomNook = true
		}
		if i == 0 {
			continue
		}

		prev := records[i-1]
		if record.AnimalKey < prev.AnimalKey {
			t.Fatalf("catalog.ResidentRecordsByGame(NewLeaf)[%d] not sorted by animal key", i)
		}
		if record.AnimalKey == prev.AnimalKey && record.Key < prev.Key {
			t.Fatalf("catalog.ResidentRecordsByGame(NewLeaf)[%d] not sorted by character key", i)
		}
	}
	if !foundTomNook {
		t.Fatal("catalog.ResidentRecordsByGame(NewLeaf) missing Tom Nook")
	}
}

func TestResidentRecordsByGender(t *testing.T) {
	records := catalog.ResidentRecordsByGender(gender.Female.Key)
	if len(records) == 0 {
		t.Fatal("catalog.ResidentRecordsByGender(Female) returned no records")
	}

	foundCeleste := false
	for i, record := range records {
		if record.GenderKey != string(gender.Female.Key) {
			t.Fatalf("catalog.ResidentRecordsByGender(Female)[%d].GenderKey = %s", i, record.GenderKey)
		}
		if record.ID == "Owl:Celeste" {
			foundCeleste = true
		}
		if i == 0 {
			continue
		}

		prev := records[i-1]
		if record.AnimalKey < prev.AnimalKey {
			t.Fatalf("catalog.ResidentRecordsByGender(Female)[%d] not sorted by animal key", i)
		}
		if record.AnimalKey == prev.AnimalKey && record.Key < prev.Key {
			t.Fatalf("catalog.ResidentRecordsByGender(Female)[%d] not sorted by character key", i)
		}
	}
	if !foundCeleste {
		t.Fatal("catalog.ResidentRecordsByGender(Female) missing Celeste")
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

func TestVillagerRecordByID(t *testing.T) {
	record, ok := catalog.VillagerRecordByID("duck:Ketchup")
	if !ok {
		t.Fatal("catalog.VillagerRecordByID(duck:Ketchup) not found")
	}
	if record.Key != string(character.Ketchup) {
		t.Fatalf("catalog.VillagerRecordByID(duck:Ketchup).Key = %s", record.Key)
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

func TestVillagerRecordsByGame(t *testing.T) {
	records := catalog.VillagerRecordsByGame(game.DoubutsuNoMoriPlus.Key)
	if len(records) == 0 {
		t.Fatal("catalog.VillagerRecordsByGame(DoubutsuNoMoriPlus) returned no records")
	}

	foundAnkha := false
	for i, record := range records {
		if record.ID == "Cat:Ankha" {
			foundAnkha = true
		}
		if i == 0 {
			continue
		}

		prev := records[i-1]
		if record.AnimalKey < prev.AnimalKey {
			t.Fatalf("catalog.VillagerRecordsByGame(DoubutsuNoMoriPlus)[%d] not sorted by animal key", i)
		}
		if record.AnimalKey == prev.AnimalKey && record.Key < prev.Key {
			t.Fatalf("catalog.VillagerRecordsByGame(DoubutsuNoMoriPlus)[%d] not sorted by character key", i)
		}
	}
	if !foundAnkha {
		t.Fatal("catalog.VillagerRecordsByGame(DoubutsuNoMoriPlus) missing Ankha")
	}
}

func TestVillagerRecordsByGender(t *testing.T) {
	records := catalog.VillagerRecordsByGender(gender.Male.Key)
	if len(records) == 0 {
		t.Fatal("catalog.VillagerRecordsByGender(Male) returned no records")
	}

	foundMarshal := false
	for i, record := range records {
		if record.GenderKey != string(gender.Male.Key) {
			t.Fatalf("catalog.VillagerRecordsByGender(Male)[%d].GenderKey = %s", i, record.GenderKey)
		}
		if record.ID == "Squirrel:Marshal" {
			foundMarshal = true
		}
		if i == 0 {
			continue
		}

		prev := records[i-1]
		if record.AnimalKey < prev.AnimalKey {
			t.Fatalf("catalog.VillagerRecordsByGender(Male)[%d] not sorted by animal key", i)
		}
		if record.AnimalKey == prev.AnimalKey && record.Key < prev.Key {
			t.Fatalf("catalog.VillagerRecordsByGender(Male)[%d] not sorted by character key", i)
		}
	}
	if !foundMarshal {
		t.Fatal("catalog.VillagerRecordsByGender(Male) missing Marshal")
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
	if _, ok := catalog.ResidentRecordByID(""); ok {
		t.Fatal("catalog.ResidentRecordByID(blank) unexpectedly found a character")
	}
	if _, ok := catalog.VillagerRecordByCode("missing"); ok {
		t.Fatal("catalog.VillagerRecordByCode(missing) unexpectedly found a character")
	}
	if _, ok := catalog.VillagerRecordByID("missing"); ok {
		t.Fatal("catalog.VillagerRecordByID(missing) unexpectedly found a character")
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
	if records := catalog.ResidentRecordsByGame(""); len(records) != 0 {
		t.Fatalf("len(catalog.ResidentRecordsByGame(\"\")) = %d", len(records))
	}
	if records := catalog.ResidentRecordsByGender(""); len(records) != 0 {
		t.Fatalf("len(catalog.ResidentRecordsByGender(\"\")) = %d", len(records))
	}
	if records := catalog.VillagerRecordsByBirthMonth(0); len(records) != 0 {
		t.Fatalf("len(catalog.VillagerRecordsByBirthMonth(0)) = %d", len(records))
	}
	if records := catalog.VillagerRecordsByGame(""); len(records) != 0 {
		t.Fatalf("len(catalog.VillagerRecordsByGame(\"\")) = %d", len(records))
	}
	if records := catalog.VillagerRecordsByGender(""); len(records) != 0 {
		t.Fatalf("len(catalog.VillagerRecordsByGender(\"\")) = %d", len(records))
	}
	if records := catalog.VillagerRecordsByPersonality(language.AmericanEnglish, ""); len(records) != 0 {
		t.Fatalf("len(catalog.VillagerRecordsByPersonality(en-US, blank)) = %d", len(records))
	}
	if records := catalog.VillagerRecordsByPersonality(language.French, "snooty"); len(records) != 0 {
		t.Fatalf("len(catalog.VillagerRecordsByPersonality(fr, snooty)) = %d", len(records))
	}
}

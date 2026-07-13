package catalog_test

import (
	"strings"
	"testing"
	"time"

	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/catalog"
	"github.com/lindsaygelle/nook/character"
	dogcharacters "github.com/lindsaygelle/nook/character/dog"
	duckcharacters "github.com/lindsaygelle/nook/character/duck"
	mousecharacters "github.com/lindsaygelle/nook/character/mouse"
	rabbitcharacters "github.com/lindsaygelle/nook/character/rabbit"
	raccooncharacters "github.com/lindsaygelle/nook/character/raccoon"
	"golang.org/x/text/language"
)

func TestVillagersByAnimal(t *testing.T) {
	villagers, ok := catalog.VillagersByAnimal(animal.Rabbit.Key)
	if !ok {
		t.Fatalf("catalog.VillagersByAnimal(%s) not found", animal.Rabbit.Key)
	}
	if len(villagers) != len(rabbitcharacters.Villagers) {
		t.Fatalf("catalog.VillagersByAnimal(%s) length = %d", animal.Rabbit.Key, len(villagers))
	}
}

func TestResidentListByAnimalSorted(t *testing.T) {
	residents, ok := catalog.ResidentListByAnimal(animal.Dog.Key)
	if !ok {
		t.Fatalf("catalog.ResidentListByAnimal(%s) not found", animal.Dog.Key)
	}
	if len(residents) != len(dogcharacters.Residents) {
		t.Fatalf("catalog.ResidentListByAnimal(%s) length = %d", animal.Dog.Key, len(residents))
	}
	if residents[0].Key != character.Booker {
		t.Fatalf("catalog.ResidentListByAnimal(%s)[0].Key = %s", animal.Dog.Key, residents[0].Key)
	}
	if residents[len(residents)-1].Key != character.Serena {
		t.Fatalf("catalog.ResidentListByAnimal(%s)[last].Key = %s", animal.Dog.Key, residents[len(residents)-1].Key)
	}
	for i := 1; i < len(residents); i++ {
		if strings.Compare(string(residents[i-1].Key), string(residents[i].Key)) > 0 {
			t.Fatalf("catalog.ResidentListByAnimal(%s) not sorted at %s > %s", animal.Dog.Key, residents[i-1].Key, residents[i].Key)
		}
	}
}

func TestResidentByKey(t *testing.T) {
	resident, ok := catalog.ResidentByKey(animal.Raccoon.Key, character.TomNook)
	if !ok {
		t.Fatalf("catalog.ResidentByKey(%s, %s) not found", animal.Raccoon.Key, character.TomNook)
	}
	if resident.Key != raccooncharacters.TomNook.Key {
		t.Fatalf("catalog.ResidentByKey(%s, %s).Key = %s", animal.Raccoon.Key, character.TomNook, resident.Key)
	}
}

func TestResidentByKeySupportsMixedAnimalBuckets(t *testing.T) {
	resident, ok := catalog.ResidentByKey(animal.Dog.Key, character.KKSlider)
	if !ok {
		t.Fatalf("catalog.ResidentByKey(%s, %s) not found", animal.Dog.Key, character.KKSlider)
	}
	if resident.Key != character.KKSlider {
		t.Fatalf("catalog.ResidentByKey(%s, %s).Key = %s", animal.Dog.Key, character.KKSlider, resident.Key)
	}
}

func TestResidentByCode(t *testing.T) {
	resident, ok := catalog.ResidentByCode(" RCN/RCO ")
	if !ok {
		t.Fatal("catalog.ResidentByCode(rcn/rco) not found")
	}
	if resident.Key != raccooncharacters.TomNook.Key {
		t.Fatalf("catalog.ResidentByCode(rcn/rco).Key = %s", resident.Key)
	}
}

func TestResidentsByBirthdaySorted(t *testing.T) {
	residents := catalog.ResidentsByBirthday(time.December, 20)
	if len(residents) != 2 {
		t.Fatalf("len(catalog.ResidentsByBirthday(December, 20)) = %d", len(residents))
	}
	if residents[0].Animal.Key != animal.Dog.Key || residents[0].Key != character.Digby {
		t.Fatalf("catalog.ResidentsByBirthday(December, 20)[0] = %s/%s", residents[0].Animal.Key, residents[0].Key)
	}
	if residents[1].Animal.Key != animal.Dog.Key || residents[1].Key != character.Isabelle {
		t.Fatalf("catalog.ResidentsByBirthday(December, 20)[1] = %s/%s", residents[1].Animal.Key, residents[1].Key)
	}
}

func TestResidentsByBirthMonthSorted(t *testing.T) {
	residents := catalog.ResidentsByBirthMonth(time.December)
	if len(residents) != 6 {
		t.Fatalf("len(catalog.ResidentsByBirthMonth(December)) = %d", len(residents))
	}
	if residents[0].Animal.Key != animal.Beaver.Key || residents[0].Key != character.Chip {
		t.Fatalf("catalog.ResidentsByBirthMonth(December)[0] = %s/%s", residents[0].Animal.Key, residents[0].Key)
	}
	if residents[len(residents)-1].Animal.Key != animal.Tortoise.Key || residents[len(residents)-1].Key != character.Tortimer {
		t.Fatalf("catalog.ResidentsByBirthMonth(December)[last] = %s/%s", residents[len(residents)-1].Animal.Key, residents[len(residents)-1].Key)
	}
}

func TestVillagerByKeyUsesAnimalScope(t *testing.T) {
	rabbitCarmen, ok := catalog.VillagerByKey(animal.Rabbit.Key, character.Carmen)
	if !ok {
		t.Fatalf("catalog.VillagerByKey(%s, %s) not found", animal.Rabbit.Key, character.Carmen)
	}
	if rabbitCarmen.Code.Value != rabbitcharacters.Carmen.Code.Value {
		t.Fatalf("catalog.VillagerByKey(%s, %s).Code = %s", animal.Rabbit.Key, character.Carmen, rabbitCarmen.Code.Value)
	}

	mouseCarmen, ok := catalog.VillagerByKey(animal.Mouse.Key, character.Carmen)
	if !ok {
		t.Fatalf("catalog.VillagerByKey(%s, %s) not found", animal.Mouse.Key, character.Carmen)
	}
	if mouseCarmen.Code.Value != mousecharacters.Carmen.Code.Value {
		t.Fatalf("catalog.VillagerByKey(%s, %s).Code = %s", animal.Mouse.Key, character.Carmen, mouseCarmen.Code.Value)
	}

	if _, ok := catalog.VillagerByKey(animal.Cat.Key, character.Carmen); ok {
		t.Fatalf("catalog.VillagerByKey(%s, %s) unexpectedly found a villager", animal.Cat.Key, character.Carmen)
	}
}

func TestVillagerByCode(t *testing.T) {
	villager, ok := catalog.VillagerByCode(" duk13 ")
	if !ok {
		t.Fatal("catalog.VillagerByCode(duk13) not found")
	}
	if villager.Key != duckcharacters.Ketchup.Key {
		t.Fatalf("catalog.VillagerByCode(duk13).Key = %s", villager.Key)
	}
}

func TestVillagersByBirthdayFindsDuplicateDatesDeterministically(t *testing.T) {
	villagers := catalog.VillagersByBirthday(time.April, 23)
	if len(villagers) != 2 {
		t.Fatalf("len(catalog.VillagersByBirthday(April, 23)) = %d", len(villagers))
	}
	if villagers[0].Animal.Key != animal.Duck.Key || villagers[0].Key != character.Miranda {
		t.Fatalf("catalog.VillagersByBirthday(April, 23)[0] = %s/%s", villagers[0].Animal.Key, villagers[0].Key)
	}
	if villagers[1].Animal.Key != animal.Monkey.Key || villagers[1].Key != character.Tammi {
		t.Fatalf("catalog.VillagersByBirthday(April, 23)[1] = %s/%s", villagers[1].Animal.Key, villagers[1].Key)
	}
	for i := 1; i < len(villagers); i++ {
		leftAnimal := string(villagers[i-1].Animal.Key)
		rightAnimal := string(villagers[i].Animal.Key)
		if leftAnimal > rightAnimal || (leftAnimal == rightAnimal && string(villagers[i-1].Key) > string(villagers[i].Key)) {
			t.Fatalf("catalog.VillagersByBirthday(April, 23) not sorted at %s/%s > %s/%s", villagers[i-1].Animal.Key, villagers[i-1].Key, villagers[i].Animal.Key, villagers[i].Key)
		}
	}
}

func TestVillagersByBirthMonthSorted(t *testing.T) {
	villagers := catalog.VillagersByBirthMonth(time.September)
	if len(villagers) != 32 {
		t.Fatalf("len(catalog.VillagersByBirthMonth(September)) = %d", len(villagers))
	}
	if villagers[0].Animal.Key != animal.Bear.Key || villagers[0].Key != character.Beardo {
		t.Fatalf("catalog.VillagersByBirthMonth(September)[0] = %s/%s", villagers[0].Animal.Key, villagers[0].Key)
	}
	if villagers[len(villagers)-1].Animal.Key != animal.Wolf.Key || villagers[len(villagers)-1].Key != character.Whitney {
		t.Fatalf("catalog.VillagersByBirthMonth(September)[last] = %s/%s", villagers[len(villagers)-1].Animal.Key, villagers[len(villagers)-1].Key)
	}
}

func TestVillagersByPersonalityDeterministic(t *testing.T) {
	villagers := catalog.VillagersByPersonality(language.AmericanEnglish, " snooty ")
	if len(villagers) != 66 {
		t.Fatalf("len(catalog.VillagersByPersonality(en-US, snooty)) = %d", len(villagers))
	}
	if villagers[0].Animal.Key != animal.Alligator.Key || villagers[0].Key != character.Alli {
		t.Fatalf("catalog.VillagersByPersonality(en-US, snooty)[0] = %s/%s", villagers[0].Animal.Key, villagers[0].Key)
	}
	if villagers[len(villagers)-1].Animal.Key != animal.Wolf.Key || villagers[len(villagers)-1].Key != character.Whitney {
		t.Fatalf("catalog.VillagersByPersonality(en-US, snooty)[last] = %s/%s", villagers[len(villagers)-1].Animal.Key, villagers[len(villagers)-1].Key)
	}
}

func TestVillagerListByAnimalSorted(t *testing.T) {
	villagers, ok := catalog.VillagerListByAnimal(animal.Mouse.Key)
	if !ok {
		t.Fatalf("catalog.VillagerListByAnimal(%s) not found", animal.Mouse.Key)
	}
	if len(villagers) != len(mousecharacters.Villagers) {
		t.Fatalf("catalog.VillagerListByAnimal(%s) length = %d", animal.Mouse.Key, len(villagers))
	}
	if villagers[0].Key != character.Anicotti {
		t.Fatalf("catalog.VillagerListByAnimal(%s)[0].Key = %s", animal.Mouse.Key, villagers[0].Key)
	}
	if villagers[len(villagers)-1].Key != character.Samson {
		t.Fatalf("catalog.VillagerListByAnimal(%s)[last].Key = %s", animal.Mouse.Key, villagers[len(villagers)-1].Key)
	}
	for i := 1; i < len(villagers); i++ {
		if strings.Compare(string(villagers[i-1].Key), string(villagers[i].Key)) > 0 {
			t.Fatalf("catalog.VillagerListByAnimal(%s) not sorted at %s > %s", animal.Mouse.Key, villagers[i-1].Key, villagers[i].Key)
		}
	}
}

func TestMissingAnimalBucket(t *testing.T) {
	if _, ok := catalog.VillagersByAnimal("missing"); ok {
		t.Fatal("catalog.VillagersByAnimal(missing) unexpectedly found a bucket")
	}
	if _, ok := catalog.ResidentsByAnimal("missing"); ok {
		t.Fatal("catalog.ResidentsByAnimal(missing) unexpectedly found a bucket")
	}
	if _, ok := catalog.VillagerListByAnimal("missing"); ok {
		t.Fatal("catalog.VillagerListByAnimal(missing) unexpectedly found a bucket")
	}
	if _, ok := catalog.ResidentListByAnimal("missing"); ok {
		t.Fatal("catalog.ResidentListByAnimal(missing) unexpectedly found a bucket")
	}
	if _, ok := catalog.ResidentByCode(""); ok {
		t.Fatal("catalog.ResidentByCode(blank) unexpectedly found a character")
	}
	if _, ok := catalog.VillagerByCode("missing"); ok {
		t.Fatal("catalog.VillagerByCode(missing) unexpectedly found a character")
	}
	if residents := catalog.ResidentsByBirthday(0, 0); len(residents) != 0 {
		t.Fatalf("len(catalog.ResidentsByBirthday(0, 0)) = %d", len(residents))
	}
	if villagers := catalog.VillagersByBirthday(time.March, 0); len(villagers) != 0 {
		t.Fatalf("len(catalog.VillagersByBirthday(March, 0)) = %d", len(villagers))
	}
	if residents := catalog.ResidentsByBirthMonth(0); len(residents) != 0 {
		t.Fatalf("len(catalog.ResidentsByBirthMonth(0)) = %d", len(residents))
	}
	if villagers := catalog.VillagersByBirthMonth(0); len(villagers) != 0 {
		t.Fatalf("len(catalog.VillagersByBirthMonth(0)) = %d", len(villagers))
	}
	if villagers := catalog.VillagersByPersonality(language.AmericanEnglish, ""); len(villagers) != 0 {
		t.Fatalf("len(catalog.VillagersByPersonality(en-US, blank)) = %d", len(villagers))
	}
	if villagers := catalog.VillagersByPersonality(language.French, "snooty"); len(villagers) != 0 {
		t.Fatalf("len(catalog.VillagersByPersonality(fr, snooty)) = %d", len(villagers))
	}
}

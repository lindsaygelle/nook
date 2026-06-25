package catalog_test

import (
	"testing"

	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/catalog"
	"github.com/lindsaygelle/nook/character"
	mousecharacters "github.com/lindsaygelle/nook/character/mouse"
	rabbitcharacters "github.com/lindsaygelle/nook/character/rabbit"
	raccooncharacters "github.com/lindsaygelle/nook/character/raccoon"
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

func TestResidentByKey(t *testing.T) {
	resident, ok := catalog.ResidentByKey(animal.Raccoon.Key, character.TomNook)
	if !ok {
		t.Fatalf("catalog.ResidentByKey(%s, %s) not found", animal.Raccoon.Key, character.TomNook)
	}
	if resident.Key != raccooncharacters.TomNook.Key {
		t.Fatalf("catalog.ResidentByKey(%s, %s).Key = %s", animal.Raccoon.Key, character.TomNook, resident.Key)
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

func TestMissingAnimalBucket(t *testing.T) {
	if _, ok := catalog.VillagersByAnimal("missing"); ok {
		t.Fatal("catalog.VillagersByAnimal(missing) unexpectedly found a bucket")
	}
	if _, ok := catalog.ResidentsByAnimal("missing"); ok {
		t.Fatal("catalog.ResidentsByAnimal(missing) unexpectedly found a bucket")
	}
}

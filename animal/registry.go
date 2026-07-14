package animal

import "github.com/lindsaygelle/nook"

var (
	animals = []nook.Animal{
		Alligator,
		Alpaca,
		Anteater,
		Axolotl,
		Bear,
		Bearcub,
		Beaver,
		Bird,
		Boar,
		Bull,
		Camel,
		Cat,
		Chameleon,
		Chicken,
		Cow,
		Deer,
		Dodo,
		Dog,
		Duck,
		Eagle,
		Elephant,
		Fox,
		FrillneckedLizard,
		Frog,
		FurSeal,
		Giraffe,
		Goat,
		Gorilla,
		Gyroid,
		Hamster,
		Hedgehog,
		Hippo,
		Horse,
		Kangaroo,
		Koala,
		Lion,
		Mole,
		Monkey,
		Mouse,
		Octopus,
		Ostrich,
		Otter,
		Owl,
		Panther,
		Peacock,
		Pelican,
		Penguin,
		Pig,
		Pigeon,
		Rabbit,
		Raccoon,
		Reindeer,
		Rhinoceros,
		Seagull,
		Sheep,
		Skunk,
		Sloth,
		Squirrel,
		Tapir,
		Tiger,
		Tortoise,
		Turkey,
		Turtle,
		Walrus,
		Wolf,
	}
	animalsByKey = func() map[nook.Key]nook.Animal {
		index := make(map[nook.Key]nook.Animal, len(animals))
		for _, animal := range animals {
			index[animal.Key] = animal
		}
		return index
	}()
)

// ByKey returns the canonical animal with the provided key.
func ByKey(key nook.Key) (nook.Animal, bool) {
	animal, ok := animalsByKey[key]
	return animal, ok
}

// List returns all canonical animals in deterministic key order.
func List() []nook.Animal {
	return append([]nook.Animal(nil), animals...)
}

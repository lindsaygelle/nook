package nook_test

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character/alligator"
	"github.com/lindsaygelle/nook/character/alpaca"
	"github.com/lindsaygelle/nook/character/anteater"
	"github.com/lindsaygelle/nook/character/axolotl"
	"github.com/lindsaygelle/nook/character/bear"
	"github.com/lindsaygelle/nook/character/bearcub"
	"github.com/lindsaygelle/nook/character/beaver"
	"github.com/lindsaygelle/nook/character/bird"
	"github.com/lindsaygelle/nook/character/boar"
	"github.com/lindsaygelle/nook/character/bull"
	"github.com/lindsaygelle/nook/character/camel"
	"github.com/lindsaygelle/nook/character/cat"
	"github.com/lindsaygelle/nook/character/chameleon"
	"github.com/lindsaygelle/nook/character/chicken"
	"github.com/lindsaygelle/nook/character/cow"
	"github.com/lindsaygelle/nook/character/deer"
	"github.com/lindsaygelle/nook/character/dodo"
	"github.com/lindsaygelle/nook/character/dog"
	"github.com/lindsaygelle/nook/character/duck"
	"github.com/lindsaygelle/nook/character/eagle"
	"github.com/lindsaygelle/nook/character/elephant"
	"github.com/lindsaygelle/nook/character/fox"
	"github.com/lindsaygelle/nook/character/frillneckedlizard"
	"github.com/lindsaygelle/nook/character/frog"
	"github.com/lindsaygelle/nook/character/furseal"
	"github.com/lindsaygelle/nook/character/giraffe"
	"github.com/lindsaygelle/nook/character/goat"
	"github.com/lindsaygelle/nook/character/gorilla"
	"github.com/lindsaygelle/nook/character/gyroid"
	"github.com/lindsaygelle/nook/character/hamster"
	"github.com/lindsaygelle/nook/character/hedgehog"
	"github.com/lindsaygelle/nook/character/hippo"
	"github.com/lindsaygelle/nook/character/horse"
	"github.com/lindsaygelle/nook/character/kangaroo"
	"github.com/lindsaygelle/nook/character/koala"
	"github.com/lindsaygelle/nook/character/lion"
	"github.com/lindsaygelle/nook/character/mole"
	"github.com/lindsaygelle/nook/character/monkey"
	"github.com/lindsaygelle/nook/character/mouse"
	"github.com/lindsaygelle/nook/character/octopus"
	"github.com/lindsaygelle/nook/character/ostrich"
	"github.com/lindsaygelle/nook/character/otter"
	"github.com/lindsaygelle/nook/character/owl"
	"github.com/lindsaygelle/nook/character/panther"
	"github.com/lindsaygelle/nook/character/peacock"
	"github.com/lindsaygelle/nook/character/pelican"
	"github.com/lindsaygelle/nook/character/penguin"
	"github.com/lindsaygelle/nook/character/pig"
	"github.com/lindsaygelle/nook/character/pigeon"
	"github.com/lindsaygelle/nook/character/rabbit"
	"github.com/lindsaygelle/nook/character/raccoon"
	"github.com/lindsaygelle/nook/character/reindeer"
	"github.com/lindsaygelle/nook/character/rhinoceros"
	"github.com/lindsaygelle/nook/character/seagull"
	"github.com/lindsaygelle/nook/character/sheep"
	"github.com/lindsaygelle/nook/character/skunk"
	"github.com/lindsaygelle/nook/character/sloth"
	"github.com/lindsaygelle/nook/character/squirrel"
	"github.com/lindsaygelle/nook/character/tapir"
	"github.com/lindsaygelle/nook/character/tiger"
	"github.com/lindsaygelle/nook/character/tortoise"
	"github.com/lindsaygelle/nook/character/turkey"
	"github.com/lindsaygelle/nook/character/turtle"
	"github.com/lindsaygelle/nook/character/walrus"
	"github.com/lindsaygelle/nook/character/wolf"
)

var (
	residents = map[nook.Key]nook.Residents{
		animal.Alpaca.Key:            alpaca.Residents,
		animal.Axolotl.Key:           axolotl.Residents,
		animal.Beaver.Key:            beaver.Residents,
		animal.Boar.Key:              boar.Residents,
		animal.Camel.Key:             camel.Residents,
		animal.Chameleon.Key:         chameleon.Residents,
		animal.Dodo.Key:              dodo.Residents,
		animal.Fox.Key:               fox.Residents,
		animal.FrillneckedLizard.Key: frillneckedlizard.Residents,
		animal.FurSeal.Key:           furseal.Residents,
		animal.Giraffe.Key:           giraffe.Residents,
		animal.Gyroid.Key:            gyroid.Residents,
		animal.Hedgehog.Key:          hedgehog.Residents,
		animal.Mole.Key:              mole.Residents,
		animal.Otter.Key:             otter.Residents,
		animal.Owl.Key:               owl.Residents,
		animal.Panther.Key:           panther.Residents,
		animal.Peacock.Key:           peacock.Residents,
		animal.Pelican.Key:           pelican.Residents,
		animal.Pigeon.Key:            pigeon.Residents,
		animal.Raccoon.Key:           raccoon.Residents,
		animal.Reindeer.Key:          reindeer.Residents,
		animal.Seagull.Key:           seagull.Residents,
		animal.Skunk.Key:             skunk.Residents,
		animal.Sloth.Key:             sloth.Residents,
		animal.Tapir.Key:             tapir.Residents,
		animal.Tortoise.Key:          tortoise.Residents,
		animal.Turkey.Key:            turkey.Residents,
		animal.Turtle.Key:            turtle.Residents,
		animal.Walrus.Key:            walrus.Residents}
)

var (
	villagers = map[nook.Key]nook.Villagers{
		animal.Alligator.Key:  alligator.Villagers,
		animal.Anteater.Key:   anteater.Villagers,
		animal.Bear.Key:       bear.Villagers,
		animal.Bearcub.Key:    bearcub.Villagers,
		animal.Bird.Key:       bird.Villagers,
		animal.Bull.Key:       bull.Villagers,
		animal.Cat.Key:        cat.Villagers,
		animal.Chicken.Key:    chicken.Villagers,
		animal.Cow.Key:        cow.Villagers,
		animal.Deer.Key:       deer.Villagers,
		animal.Dog.Key:        dog.Villagers,
		animal.Duck.Key:       duck.Villagers,
		animal.Eagle.Key:      eagle.Villagers,
		animal.Elephant.Key:   elephant.Villagers,
		animal.Frog.Key:       frog.Villagers,
		animal.Goat.Key:       goat.Villagers,
		animal.Gorilla.Key:    gorilla.Villagers,
		animal.Hamster.Key:    hamster.Villagers,
		animal.Hippo.Key:      hippo.Villagers,
		animal.Horse.Key:      horse.Villagers,
		animal.Kangaroo.Key:   kangaroo.Villagers,
		animal.Koala.Key:      koala.Villagers,
		animal.Lion.Key:       lion.Villagers,
		animal.Monkey.Key:     monkey.Villagers,
		animal.Mouse.Key:      mouse.Villagers,
		animal.Octopus.Key:    octopus.Villagers,
		animal.Ostrich.Key:    ostrich.Villagers,
		animal.Penguin.Key:    penguin.Villagers,
		animal.Pig.Key:        pig.Villagers,
		animal.Rabbit.Key:     rabbit.Villagers,
		animal.Rhinoceros.Key: rhinoceros.Villagers,
		animal.Sheep.Key:      sheep.Villagers,
		animal.Squirrel.Key:   squirrel.Villagers,
		animal.Tiger.Key:      tiger.Villagers,
		animal.Wolf.Key:       wolf.Villagers}
)

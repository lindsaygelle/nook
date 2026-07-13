package catalog

import (
	"slices"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/game"
)

var gameReleaseOrder = map[nook.Key]int{
	game.DoubutsuNoMori.Key:      0,
	game.DoubutsuNoMoriPlus.Key:  1,
	game.AnimalCrossing.Key:      2,
	game.DoubutsuNoMoriEPlus.Key: 3,
	game.DongwuSenlin.Key:        4,
	game.WildWorld.Key:           5,
	game.CityFolk.Key:            6,
	game.NewLeaf.Key:             7,
	game.NewHorizons.Key:         8,
	game.HappyHomeDesigner.Key:   9,
	game.AmiiboFestival.Key:      10,
	game.PocketCamp.Key:          11,
}

// CharacterGames returns a character's game appearance history in release
// order.
func CharacterGames(character nook.Character) ([]nook.Game, bool) {
	if character.ID() == "" || character.Games == nil {
		return nil, false
	}
	return sortedCharacterGames(character.Games), true
}

// CharacterGamesByID returns a character's game appearance history using an
// exact global character identifier match after normalization.
func CharacterGamesByID(id string) ([]nook.Game, bool) {
	if resident, ok := ResidentByID(id); ok {
		return CharacterGames(resident.Character)
	}
	if villager, ok := VillagerByID(id); ok {
		return CharacterGames(villager.Character)
	}
	return nil, false
}

// FirstAppearance returns the earliest known game appearance for a character.
func FirstAppearance(character nook.Character) (nook.Game, bool) {
	return FirstAppearanceByID(string(character.ID()))
}

// FirstAppearanceByID returns the earliest known game appearance using an
// exact global character identifier match after normalization.
func FirstAppearanceByID(id string) (nook.Game, bool) {
	games, ok := CharacterGamesByID(id)
	if !ok || len(games) == 0 {
		return nook.Game{}, false
	}
	return games[0], true
}

// LastAppearance returns the latest known game appearance for a character.
func LastAppearance(character nook.Character) (nook.Game, bool) {
	return LastAppearanceByID(string(character.ID()))
}

// LastAppearanceByID returns the latest known game appearance using an exact
// global character identifier match after normalization.
func LastAppearanceByID(id string) (nook.Game, bool) {
	games, ok := CharacterGamesByID(id)
	if !ok || len(games) == 0 {
		return nook.Game{}, false
	}
	return games[len(games)-1], true
}

func sortedCharacterGames(games []nook.Game) []nook.Game {
	out := append([]nook.Game(nil), games...)
	slices.SortFunc(out, compareGamesByReleaseOrder)
	return out
}

func compareGamesByReleaseOrder(a, b nook.Game) int {
	left, ok := gameReleaseOrder[a.Key]
	if !ok {
		left = len(gameReleaseOrder)
	}
	right, ok := gameReleaseOrder[b.Key]
	if !ok {
		right = len(gameReleaseOrder)
	}
	switch {
	case left < right:
		return -1
	case left > right:
		return 1
	case a.Key < b.Key:
		return -1
	case a.Key > b.Key:
		return 1
	default:
		return 0
	}
}

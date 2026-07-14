package game

import "github.com/lindsaygelle/nook"

var (
	// games contains the canonical games in deterministic series order.
	games = []nook.Game{
		DoubutsuNoMori,
		DoubutsuNoMoriPlus,
		AnimalCrossing,
		DoubutsuNoMoriEPlus,
		DongwuSenlin,
		WildWorld,
		CityFolk,
		NewLeaf,
		HappyHomeDesigner,
		AmiiboFestival,
		PocketCamp,
		NewHorizons,
	}
)

var (
	// gamesByKey contains canonical games indexed by key.
	gamesByKey = func() map[nook.Key]nook.Game {
		index := make(map[nook.Key]nook.Game, len(games))
		for _, game := range games {
			index[game.Key] = game
		}
		return index
	}()
)

// ByKey returns the canonical game with the provided key.
func ByKey(key nook.Key) (nook.Game, bool) {
	game, ok := gamesByKey[key]
	return game, ok
}

// List returns all canonical games in deterministic series order.
func List() []nook.Game {
	return append([]nook.Game(nil), games...)
}

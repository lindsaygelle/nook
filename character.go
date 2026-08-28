package nook

import "slices"

const characterIDSeparator = ":"

func compareGameCategoriesByKey(a, b GameCategory) int {
	switch {
	case a.Key < b.Key:
		return -1
	case a.Key > b.Key:
		return 1
	default:
		return 0
	}
}

func compareGamesByReleaseOrder(a, b Game) int {
	switch {
	case a.ReleaseOrder == 0 && b.ReleaseOrder != 0:
		return 1
	case a.ReleaseOrder != 0 && b.ReleaseOrder == 0:
		return -1
	case a.ReleaseOrder < b.ReleaseOrder:
		return -1
	case a.ReleaseOrder > b.ReleaseOrder:
		return 1
	}

	switch {
	case a.Key < b.Key:
		return -1
	case a.Key > b.Key:
		return 1
	default:
		return 0
	}
}

func comparePlatformsByKey(a, b Platform) int {
	switch {
	case a.Key < b.Key:
		return -1
	case a.Key > b.Key:
		return 1
	default:
		return 0
	}
}

func compareRegionsByKey(a, b Region) int {
	switch {
	case a.Key < b.Key:
		return -1
	case a.Key > b.Key:
		return 1
	default:
		return 0
	}
}

// Character is a composite type that combines various attributes of an Animal Crossing character.
type Character struct {
	// Animal represents the animal type of the character.
	Animal

	// Birthday contains the birthday information of the character.
	Birthday Birthday

	// Code is a unique identifier for the character.
	Code Code

	// Games contains the character's game appearances.
	Games []Game

	// Gender represents the gender of the character.
	Gender Gender

	// Key is the character's canonical package key.
	// Keys are not guaranteed to be globally unique across animal types.
	Key Key

	// Name contains the names of the character in different languages.
	Name Languages

	// Special indicates whether the character is special or has a unique role.
	Special bool
}

// AppearsInGame reports whether the character appears in the provided game.
func (c Character) AppearsInGame(gameKey Key) bool {
	if gameKey == "" {
		return false
	}

	for _, game := range c.Games {
		if game.Key == gameKey {
			return true
		}
	}

	return false
}

// FirstGame returns the earliest known game appearance for the character.
func (c Character) FirstGame() (Game, bool) {
	games := c.GamesByReleaseOrder()
	if len(games) == 0 {
		return Game{}, false
	}
	return games[0], true
}

// FirstReleaseDate returns the earliest known release date for the character.
func (c Character) FirstReleaseDate() (ReleaseDate, bool) {
	game, ok := c.FirstGame()
	if !ok {
		return ReleaseDate{}, false
	}
	return game.FirstReleaseDate()
}

// GameByKey returns the character's game appearance for the provided game key.
func (c Character) GameByKey(gameKey Key) (Game, bool) {
	if gameKey == "" {
		return Game{}, false
	}

	for _, game := range c.Games {
		if game.Key == gameKey {
			return game, true
		}
	}

	return Game{}, false
}

// HasZodiacSign reports whether the character has the provided zodiac sign.
func (c Character) HasZodiacSign(zodiacSignKey Key) bool {
	if zodiacSignKey == "" {
		return false
	}

	key, ok := c.ZodiacSignKey()
	return ok && key == zodiacSignKey
}

// GamesByCategory returns the character's game appearances within the provided
// category sorted into deterministic release order.
func (c Character) GamesByCategory(categoryKey Key) []Game {
	if categoryKey == "" {
		return nil
	}

	games := make([]Game, 0)
	for _, game := range c.GamesByReleaseOrder() {
		if !game.HasCategory(categoryKey) {
			continue
		}
		games = append(games, game)
	}

	return games
}

// GamesByPlatform returns the character's game appearances on the provided
// platform sorted into deterministic release order.
func (c Character) GamesByPlatform(platformKey Key) []Game {
	if platformKey == "" {
		return nil
	}

	games := make([]Game, 0)
	for _, game := range c.GamesByReleaseOrder() {
		if !game.OnPlatform(platformKey) {
			continue
		}
		games = append(games, game)
	}

	return games
}

// GamesByRegion returns the character's game appearances with a release in the
// provided region sorted into deterministic release order.
func (c Character) GamesByRegion(regionKey Key) []Game {
	if regionKey == "" {
		return nil
	}

	games := make([]Game, 0)
	for _, game := range c.GamesByReleaseOrder() {
		if _, ok := game.ReleaseDateByRegion(regionKey); !ok {
			continue
		}
		games = append(games, game)
	}

	return games
}

// GameCategories returns the unique game categories represented across the
// character's known game appearances in deterministic key order.
func (c Character) GameCategories() []GameCategory {
	index := make(map[Key]GameCategory, len(c.Games))

	for _, game := range c.Games {
		if game.Category.Key == "" {
			continue
		}
		index[game.Category.Key] = game.Category
	}

	categories := make([]GameCategory, 0, len(index))
	for _, category := range index {
		categories = append(categories, category)
	}
	slices.SortFunc(categories, compareGameCategoriesByKey)
	return categories
}

// GamePlatforms returns the unique platforms represented across the
// character's known game appearances in deterministic key order.
func (c Character) GamePlatforms() []Platform {
	index := make(map[Key]Platform)

	for _, game := range c.Games {
		for _, platform := range game.Platforms {
			if platform.Key == "" {
				continue
			}
			index[platform.Key] = platform
		}
	}

	platforms := make([]Platform, 0, len(index))
	for _, platform := range index {
		platforms = append(platforms, platform)
	}
	slices.SortFunc(platforms, comparePlatformsByKey)
	return platforms
}

// GameRegions returns the unique release regions represented across the
// character's known game appearances in deterministic key order.
func (c Character) GameRegions() []Region {
	index := make(map[Key]Region)

	for _, game := range c.Games {
		for _, releaseDate := range game.ReleaseDates {
			if releaseDate.Region.Key == "" {
				continue
			}
			index[releaseDate.Region.Key] = releaseDate.Region
		}
	}

	regions := make([]Region, 0, len(index))
	for _, region := range index {
		regions = append(regions, region)
	}
	slices.SortFunc(regions, compareRegionsByKey)
	return regions
}

// GamesOnPlatform returns a boolean indicating whether the character appears on
// the provided platform in any known game appearance.
func (c Character) GamesOnPlatform(platformKey Key) bool {
	if platformKey == "" {
		return false
	}

	for _, game := range c.Games {
		if game.OnPlatform(platformKey) {
			return true
		}
	}

	return false
}

// GamesReleasedInRegion returns a boolean indicating whether the character
// appears in any game that released in the provided region.
func (c Character) GamesReleasedInRegion(regionKey Key) bool {
	if regionKey == "" {
		return false
	}

	for _, game := range c.Games {
		if _, ok := game.ReleaseDateByRegion(regionKey); ok {
			return true
		}
	}

	return false
}

// GamesByReleaseOrder returns the character's game appearances sorted into
// deterministic release order.
func (c Character) GamesByReleaseOrder() []Game {
	games := append([]Game(nil), c.Games...)
	slices.SortFunc(games, compareGamesByReleaseOrder)
	return games
}

// ID returns a globally unique identifier composed from the character's animal
// key and character key.
func (c Character) ID() Key {
	if c.Animal.Key == "" || c.Key == "" {
		return ""
	}
	return Key(string(c.Animal.Key) + characterIDSeparator + string(c.Key))
}

// LastGame returns the latest known game appearance for the character.
func (c Character) LastGame() (Game, bool) {
	games := c.GamesByReleaseOrder()
	if len(games) == 0 {
		return Game{}, false
	}
	return games[len(games)-1], true
}

// LastReleaseDate returns the latest known release date for the character.
func (c Character) LastReleaseDate() (ReleaseDate, bool) {
	game, ok := c.LastGame()
	if !ok {
		return ReleaseDate{}, false
	}
	return game.LastReleaseDate()
}

// ZodiacSignKey returns the zodiac sign key derived from the character
// birthday.
func (c Character) ZodiacSignKey() (Key, bool) {
	return c.Birthday.ZodiacSignKey()
}

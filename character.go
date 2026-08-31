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

func firstGameFromGames(games []Game) (Game, bool) {
	if len(games) == 0 {
		return Game{}, false
	}

	return games[0], true
}

func firstReleaseDateFromGames(games []Game) (ReleaseDate, bool) {
	var first ReleaseDate
	found := false

	for _, game := range games {
		releaseDate, ok := game.FirstReleaseDate()
		if !ok || !releaseDate.Ok() {
			continue
		}

		if !found || releaseDate.Before(first) {
			first = releaseDate
			found = true
		}
	}

	return first, found
}

func lastGameFromGames(games []Game) (Game, bool) {
	if len(games) == 0 {
		return Game{}, false
	}

	return games[len(games)-1], true
}

func lastReleaseDateFromGames(games []Game) (ReleaseDate, bool) {
	var last ReleaseDate
	found := false

	for _, game := range games {
		releaseDate, ok := game.LastReleaseDate()
		if !ok || !releaseDate.Ok() {
			continue
		}

		if !found || releaseDate.After(last) {
			last = releaseDate
			found = true
		}
	}

	return last, found
}

func releaseYearsFromGames(games []Game) []uint16 {
	index := make(map[uint16]struct{})

	for _, game := range games {
		for _, releaseDate := range game.ReleaseDates {
			if releaseDate.Year == 0 {
				continue
			}

			index[releaseDate.Year] = struct{}{}
		}
	}

	years := make([]uint16, 0, len(index))
	for year := range index {
		years = append(years, year)
	}

	slices.Sort(years)
	return years
}

func releaseYearsFromGamesByRegion(games []Game, regionKey Key) []uint16 {
	index := make(map[uint16]struct{})

	for _, game := range games {
		releaseDate, ok := game.ReleaseDateByRegion(regionKey)
		if !ok || releaseDate.Year == 0 {
			continue
		}

		index[releaseDate.Year] = struct{}{}
	}

	years := make([]uint16, 0, len(index))
	for year := range index {
		years = append(years, year)
	}

	slices.Sort(years)
	return years
}

func gamesByReleaseYear(games []Game, year uint16) []Game {
	filtered := make([]Game, 0)

	for _, game := range games {
		for _, releaseDate := range game.ReleaseDates {
			if releaseDate.Year != year {
				continue
			}

			filtered = append(filtered, game)
			break
		}
	}

	return filtered
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
	return firstGameFromGames(c.GamesByReleaseOrder())
}

// FirstGameByCategory returns the earliest known game appearance for the
// character within the provided game category.
func (c Character) FirstGameByCategory(categoryKey Key) (Game, bool) {
	return firstGameFromGames(c.GamesByCategory(categoryKey))
}

// FirstGameByPlatform returns the earliest known game appearance for the
// character on the provided platform.
func (c Character) FirstGameByPlatform(platformKey Key) (Game, bool) {
	return firstGameFromGames(c.GamesByPlatform(platformKey))
}

// FirstGameByRegion returns the earliest known game appearance for the
// character in the provided release region.
func (c Character) FirstGameByRegion(regionKey Key) (Game, bool) {
	return firstGameFromGames(c.GamesByRegion(regionKey))
}

// FirstReleaseDate returns the earliest known release date for the character.
func (c Character) FirstReleaseDate() (ReleaseDate, bool) {
	return firstReleaseDateFromGames(c.Games)
}

// FirstReleaseDateByCategory returns the earliest known release date for the
// character within the provided game category.
func (c Character) FirstReleaseDateByCategory(categoryKey Key) (ReleaseDate, bool) {
	if categoryKey == "" {
		return ReleaseDate{}, false
	}

	return firstReleaseDateFromGames(c.GamesByCategory(categoryKey))
}

// FirstReleaseDateByPlatform returns the earliest known release date for the
// character on the provided platform.
func (c Character) FirstReleaseDateByPlatform(platformKey Key) (ReleaseDate, bool) {
	if platformKey == "" {
		return ReleaseDate{}, false
	}

	return firstReleaseDateFromGames(c.GamesByPlatform(platformKey))
}

// FirstReleaseDateByRegion returns the earliest known release date for the
// character in the provided release region.
func (c Character) FirstReleaseDateByRegion(regionKey Key) (ReleaseDate, bool) {
	if regionKey == "" {
		return ReleaseDate{}, false
	}

	var first ReleaseDate
	found := false

	for _, game := range c.Games {
		releaseDate, ok := game.ReleaseDateByRegion(regionKey)
		if !ok || !releaseDate.Ok() {
			continue
		}

		if !found || releaseDate.Before(first) {
			first = releaseDate
			found = true
		}
	}

	return first, found
}

// GameCategoryCount returns the number of unique game categories represented
// across the character's known game appearances.
func (c Character) GameCategoryCount() int {
	return len(c.GameCategories())
}

// GameCount returns the number of known game appearances for the character.
func (c Character) GameCount() int {
	return len(c.Games)
}

// GameCountByCategory returns the number of known game appearances for the
// character within the provided category.
func (c Character) GameCountByCategory(categoryKey Key) int {
	return len(c.GamesByCategory(categoryKey))
}

// GameCountByPlatform returns the number of known game appearances for the
// character on the provided platform.
func (c Character) GameCountByPlatform(platformKey Key) int {
	return len(c.GamesByPlatform(platformKey))
}

// GameCountByRegion returns the number of known game appearances for the
// character in games released within the provided region.
func (c Character) GameCountByRegion(regionKey Key) int {
	return len(c.GamesByRegion(regionKey))
}

// GameCountByReleaseYear returns the number of known game appearances for the
// character with a release in the provided year.
func (c Character) GameCountByReleaseYear(year uint16) int {
	return len(c.GamesByReleaseYear(year))
}

// ReleaseYears returns the unique release years represented across the
// character's known regional release history in chronological order.
func (c Character) ReleaseYears() []uint16 {
	return releaseYearsFromGames(c.Games)
}

// ReleaseYearsByCategory returns the unique release years represented across
// the character's known releases within the provided category in chronological
// order.
func (c Character) ReleaseYearsByCategory(categoryKey Key) []uint16 {
	if categoryKey == "" {
		return nil
	}

	return releaseYearsFromGames(c.GamesByCategory(categoryKey))
}

// ReleaseYearsByPlatform returns the unique release years represented across
// the character's known releases on the provided platform in chronological
// order.
func (c Character) ReleaseYearsByPlatform(platformKey Key) []uint16 {
	if platformKey == "" {
		return nil
	}

	return releaseYearsFromGames(c.GamesByPlatform(platformKey))
}

// ReleaseYearsByRegion returns the unique release years represented across the
// character's known releases in the provided region in chronological order.
func (c Character) ReleaseYearsByRegion(regionKey Key) []uint16 {
	if regionKey == "" {
		return nil
	}

	return releaseYearsFromGamesByRegion(c.GamesByRegion(regionKey), regionKey)
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

// GamesByReleaseYear returns the character's game appearances with a release
// in the provided year sorted into deterministic release order.
func (c Character) GamesByReleaseYear(year uint16) []Game {
	if year == 0 {
		return nil
	}

	return gamesByReleaseYear(c.GamesByReleaseOrder(), year)
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

// GamePlatformCount returns the number of unique platforms represented across
// the character's known game appearances.
func (c Character) GamePlatformCount() int {
	return len(c.GamePlatforms())
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

// GameRegionCount returns the number of unique release regions represented
// across the character's known game appearances.
func (c Character) GameRegionCount() int {
	return len(c.GameRegions())
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

// GamesReleasedInYear returns a boolean indicating whether the character
// appears in any game that released in the provided year.
func (c Character) GamesReleasedInYear(year uint16) bool {
	if year == 0 {
		return false
	}

	return len(c.GamesByReleaseYear(year)) != 0
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
	return lastGameFromGames(c.GamesByReleaseOrder())
}

// LastGameByCategory returns the latest known game appearance for the
// character within the provided game category.
func (c Character) LastGameByCategory(categoryKey Key) (Game, bool) {
	return lastGameFromGames(c.GamesByCategory(categoryKey))
}

// LastGameByPlatform returns the latest known game appearance for the
// character on the provided platform.
func (c Character) LastGameByPlatform(platformKey Key) (Game, bool) {
	return lastGameFromGames(c.GamesByPlatform(platformKey))
}

// LastGameByRegion returns the latest known game appearance for the character
// in the provided release region.
func (c Character) LastGameByRegion(regionKey Key) (Game, bool) {
	return lastGameFromGames(c.GamesByRegion(regionKey))
}

// LastReleaseDate returns the latest known release date for the character.
func (c Character) LastReleaseDate() (ReleaseDate, bool) {
	return lastReleaseDateFromGames(c.Games)
}

// LastReleaseDateByCategory returns the latest known release date for the
// character within the provided game category.
func (c Character) LastReleaseDateByCategory(categoryKey Key) (ReleaseDate, bool) {
	if categoryKey == "" {
		return ReleaseDate{}, false
	}

	return lastReleaseDateFromGames(c.GamesByCategory(categoryKey))
}

// LastReleaseDateByPlatform returns the latest known release date for the
// character on the provided platform.
func (c Character) LastReleaseDateByPlatform(platformKey Key) (ReleaseDate, bool) {
	if platformKey == "" {
		return ReleaseDate{}, false
	}

	return lastReleaseDateFromGames(c.GamesByPlatform(platformKey))
}

// LastReleaseDateByRegion returns the latest known release date for the
// character in the provided release region.
func (c Character) LastReleaseDateByRegion(regionKey Key) (ReleaseDate, bool) {
	if regionKey == "" {
		return ReleaseDate{}, false
	}

	var last ReleaseDate
	found := false

	for _, game := range c.Games {
		releaseDate, ok := game.ReleaseDateByRegion(regionKey)
		if !ok || !releaseDate.Ok() {
			continue
		}

		if !found || releaseDate.After(last) {
			last = releaseDate
			found = true
		}
	}

	return last, found
}

// ZodiacSignKey returns the zodiac sign key derived from the character
// birthday.
func (c Character) ZodiacSignKey() (Key, bool) {
	return c.Birthday.ZodiacSignKey()
}

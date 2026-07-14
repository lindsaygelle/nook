package nook

// Game represents a game in the Animal Crossing series.
type Game struct {
	// Key is the language-agnostic key of the game.
	Key Key

	// Name contains the localized names of the game.
	Name Languages

	// Platforms contains the game's known release platforms in deterministic
	// key order.
	Platforms []Platform

	// ReleaseDates contains the game's known regional release history in
	// chronological order.
	ReleaseDates []ReleaseDate

	// ReleaseOrder is the series chronology position of the game.
	ReleaseOrder uint8
}

// FirstReleaseDate returns the earliest known release date for the game.
func (g Game) FirstReleaseDate() (ReleaseDate, bool) {
	if len(g.ReleaseDates) == 0 {
		return ReleaseDate{}, false
	}
	return g.ReleaseDates[0], true
}

// LastReleaseDate returns the latest known release date for the game.
func (g Game) LastReleaseDate() (ReleaseDate, bool) {
	if len(g.ReleaseDates) == 0 {
		return ReleaseDate{}, false
	}
	return g.ReleaseDates[len(g.ReleaseDates)-1], true
}

// OnPlatform reports whether the game released on the provided platform.
func (g Game) OnPlatform(platformKey Key) bool {
	_, ok := g.PlatformByKey(platformKey)
	return ok
}

// PlatformByKey returns the game's platform with the provided key.
func (g Game) PlatformByKey(platformKey Key) (Platform, bool) {
	if platformKey == "" {
		return Platform{}, false
	}

	for _, platform := range g.Platforms {
		if platform.Key == platformKey {
			return platform, true
		}
	}

	return Platform{}, false
}

// ReleaseDateByRegion returns the game's release date for the provided region.
func (g Game) ReleaseDateByRegion(regionKey Key) (ReleaseDate, bool) {
	if regionKey == "" {
		return ReleaseDate{}, false
	}

	for _, releaseDate := range g.ReleaseDates {
		if releaseDate.Region.Key == regionKey {
			return releaseDate, true
		}
	}

	return ReleaseDate{}, false
}

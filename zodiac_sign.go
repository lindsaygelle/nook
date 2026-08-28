package nook

// ZodiacSign represents a canonical zodiac sign derived from a character
// birthday.
type ZodiacSign struct {
	// Key is the language-agnostic key of the zodiac sign.
	Key Key

	// Name contains the localized names of the zodiac sign.
	Name Languages
}

package zodiac

import "github.com/lindsaygelle/nook"

var (
	// zodiacSigns contains canonical zodiac signs in deterministic key order.
	zodiacSigns = []nook.ZodiacSign{
		Aquarius,
		Aries,
		Cancer,
		Capricorn,
		Gemini,
		Leo,
		Libra,
		Pisces,
		Sagittarius,
		Scorpio,
		Taurus,
		Virgo}
)

var (
	// zodiacSignsByKey contains canonical zodiac signs indexed by key.
	zodiacSignsByKey = func() map[nook.Key]nook.ZodiacSign {
		index := make(map[nook.Key]nook.ZodiacSign, len(zodiacSigns))
		for _, zodiacSign := range zodiacSigns {
			index[zodiacSign.Key] = zodiacSign
		}
		return index
	}()
)

// ByKey returns the canonical zodiac sign with the provided key.
func ByKey(key nook.Key) (nook.ZodiacSign, bool) {
	zodiacSign, ok := zodiacSignsByKey[key]
	return zodiacSign, ok
}

// List returns all canonical zodiac signs in deterministic key order.
func List() []nook.ZodiacSign {
	return append([]nook.ZodiacSign(nil), zodiacSigns...)
}

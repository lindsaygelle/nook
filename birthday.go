package nook

import (
	"time"
)

const (
	capricornStartDay   = 22
	aquariusStartDay    = 20
	piscesStartDay      = 19
	ariesStartDay       = 21
	taurusStartDay      = 20
	geminiStartDay      = 21
	cancerStartDay      = 21
	leoStartDay         = 23
	virgoStartDay       = 23
	libraStartDay       = 23
	scorpioStartDay     = 23
	sagittariusStartDay = 22
)

const (
	capricornKey   = Key("Capricorn")
	aquariusKey    = Key("Aquarius")
	piscesKey      = Key("Pisces")
	ariesKey       = Key("Aries")
	taurusKey      = Key("Taurus")
	geminiKey      = Key("Gemini")
	cancerKey      = Key("Cancer")
	leoKey         = Key("Leo")
	virgoKey       = Key("Virgo")
	libraKey       = Key("Libra")
	scorpioKey     = Key("Scorpio")
	sagittariusKey = Key("Sagittarius")
)

// Birthday represents the birthday information of an Animal Crossing character.
type Birthday struct {
	// Day is the day of the month when the character's birthday falls.
	Day uint8

	// Month is the month of the year when the character's birthday falls.
	Month time.Month
}

// Ok returns a boolean indicating whether the Birthday information is complete and valid.
// It checks if both the Day and Month fields have been set.
func (b Birthday) Ok() bool {
	return (b.Day != 0) && (b.Month != 0)
}

// ZodiacSignKey returns the zodiac sign key derived from the birthday.
func (b Birthday) ZodiacSignKey() (Key, bool) {
	if !b.Ok() {
		return "", false
	}

	switch b.Month {
	case time.January:
		if b.Day >= aquariusStartDay {
			return aquariusKey, true
		}
		return capricornKey, true
	case time.February:
		if b.Day >= piscesStartDay {
			return piscesKey, true
		}
		return aquariusKey, true
	case time.March:
		if b.Day >= ariesStartDay {
			return ariesKey, true
		}
		return piscesKey, true
	case time.April:
		if b.Day >= taurusStartDay {
			return taurusKey, true
		}
		return ariesKey, true
	case time.May:
		if b.Day >= geminiStartDay {
			return geminiKey, true
		}
		return taurusKey, true
	case time.June:
		if b.Day >= cancerStartDay {
			return cancerKey, true
		}
		return geminiKey, true
	case time.July:
		if b.Day >= leoStartDay {
			return leoKey, true
		}
		return cancerKey, true
	case time.August:
		if b.Day >= virgoStartDay {
			return virgoKey, true
		}
		return leoKey, true
	case time.September:
		if b.Day >= libraStartDay {
			return libraKey, true
		}
		return virgoKey, true
	case time.October:
		if b.Day >= scorpioStartDay {
			return scorpioKey, true
		}
		return libraKey, true
	case time.November:
		if b.Day >= sagittariusStartDay {
			return sagittariusKey, true
		}
		return scorpioKey, true
	case time.December:
		if b.Day >= capricornStartDay {
			return capricornKey, true
		}
		return sagittariusKey, true
	default:
		return "", false
	}
}

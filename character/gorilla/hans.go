package gorilla

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	// hansBirthday represents hans birthday.
	hansBirthday = nook.Birthday{
		Day:   5,
		Month: time.December}
)

var (
	// hansCode represents hans code.
	hansCode = nook.Code{
		Value: "gor10"}
)

var (
	// hansAmericanEnglishName represents hans american english name.
	hansAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Hans"}

	// hansCanadianFrenchName represents hans canadian french name.
	hansCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Loran"}

	// hansDutchName represents hans dutch name.
	hansDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Hans"}

	// hansFrenchName represents hans french name.
	hansFrenchName = nook.Name{
		Language: language.French,
		Value:    "Loran"}

	// hansGermanName represents hans german name.
	hansGermanName = nook.Name{
		Language: language.German,
		Value:    "Hans"}

	// hansItalianName represents hans italian name.
	hansItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Grigilio"}

	// hansJapaneseName represents hans japanese name.
	hansJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "スナイル"}

	// hansKoreanName represents hans korean name.
	hansKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "스나일"}

	// hansLatinAmericanSpanishName represents hans latin american spanish name.
	hansLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Hans"}

	// hansRussianName represents hans russian name.
	hansRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Ганс"}

	// hansSimplifiedChineseName represents hans simplified chinese name.
	hansSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "史奈鲁"}

	// hansSpanishName represents hans spanish name.
	hansSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Hans"}

	// hansTraditionalChineseName represents hans traditional chinese name.
	hansTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "史奈魯"}
)

var (
	// hansName represents hans name.
	hansName = nook.Languages{
		language.AmericanEnglish:      hansAmericanEnglishName,
		language.CanadianFrench:       hansCanadianFrenchName,
		language.Dutch:                hansDutchName,
		language.French:               hansFrenchName,
		language.German:               hansGermanName,
		language.Italian:              hansItalianName,
		language.Japanese:             hansJapaneseName,
		language.Korean:               hansKoreanName,
		language.LatinAmericanSpanish: hansLatinAmericanSpanishName,
		language.Russian:              hansRussianName,
		language.SimplifiedChinese:    hansSimplifiedChineseName,
		language.Spanish:              hansSpanishName,
		language.TraditionalChinese:   hansTraditionalChineseName}
)

var (
	// hansCharacter represents hans character.
	hansCharacter = nook.Character{
		Animal:   animal.Gorilla,
		Birthday: hansBirthday,
		Code:     hansCode,
		Key:      character.Hans,
		Gender:   gender.Male,
		Name:     hansName,
		Special:  false}
)

var (
	// hansAmericanEnglishPhrase represents hans american english phrase.
	hansAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "groovy"}

	// hansCanadianFrenchPhrase represents hans canadian french phrase.
	hansCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tût tût"}

	// hansDutchPhrase represents hans dutch phrase.
	hansDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "groovy"}

	// hansFrenchPhrase represents hans french phrase.
	hansFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tût tût"}

	// hansGermanPhrase represents hans german phrase.
	hansGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kongkong"}

	// hansItalianPhrase represents hans italian phrase.
	hansItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "entonces"}

	// hansJapanesePhrase represents hans japanese phrase.
	hansJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "いえてる"}

	// hansKoreanPhrase represents hans korean phrase.
	hansKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "알만하다"}

	// hansLatinAmericanSpanishPhrase represents hans latin american spanish phrase.
	hansLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cungagún"}

	// hansRussianPhrase represents hans russian phrase.
	hansRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "респект"}

	// hansSimplifiedChinesePhrase represents hans simplified chinese phrase.
	hansSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "同意"}

	// hansSpanishPhrase represents hans spanish phrase.
	hansSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cungagún"}

	// hansTraditionalChinesePhrase represents hans traditional chinese phrase.
	hansTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "同意"}
)

var (
	// hansPhrase represents hans phrase.
	hansPhrase = nook.Languages{
		language.AmericanEnglish:      hansAmericanEnglishPhrase,
		language.CanadianFrench:       hansCanadianFrenchPhrase,
		language.Dutch:                hansDutchPhrase,
		language.French:               hansFrenchPhrase,
		language.German:               hansGermanPhrase,
		language.Italian:              hansItalianPhrase,
		language.Japanese:             hansJapanesePhrase,
		language.Korean:               hansKoreanPhrase,
		language.LatinAmericanSpanish: hansLatinAmericanSpanishPhrase,
		language.Russian:              hansRussianPhrase,
		language.SimplifiedChinese:    hansSimplifiedChinesePhrase,
		language.Spanish:              hansSpanishPhrase,
		language.TraditionalChinese:   hansTraditionalChinesePhrase}
)

var (
	// Hans represents hans.
	Hans = nook.Villager{
		Character:   hansCharacter,
		Personality: personality.Smug,
		Phrase:      hansPhrase}
)

package duck

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
	// mollyBirthday represents molly birthday.
	mollyBirthday = nook.Birthday{
		Day:   7,
		Month: time.March}
)

var (
	// mollyCode represents molly code.
	mollyCode = nook.Code{
		Value: "duk16"}
)

var (
	// mollyAmericanEnglishName represents molly american english name.
	mollyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Molly"}

	// mollyCanadianFrenchName represents molly canadian french name.
	mollyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Molly"}

	// mollyDutchName represents molly dutch name.
	mollyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Molly"}

	// mollyFrenchName represents molly french name.
	mollyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Molly"}

	// mollyGermanName represents molly german name.
	mollyGermanName = nook.Name{
		Language: language.German,
		Value:    "Monika"}

	// mollyItalianName represents molly italian name.
	mollyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Molly"}

	// mollyJapaneseName represents molly japanese name.
	mollyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "カモミ"}

	// mollyKoreanName represents molly korean name.
	mollyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "귀오미"}

	// mollyLatinAmericanSpanishName represents molly latin american spanish name.
	mollyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Deira"}

	// mollyRussianName represents molly russian name.
	mollyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Молли"}

	// mollySimplifiedChineseName represents molly simplified chinese name.
	mollySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "亚美"}

	// mollySpanishName represents molly spanish name.
	mollySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Deira"}

	// mollyTraditionalChineseName represents molly traditional chinese name.
	mollyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "亞美"}
)

var (
	// mollyName represents molly name.
	mollyName = nook.Languages{
		language.AmericanEnglish:      mollyAmericanEnglishName,
		language.CanadianFrench:       mollyCanadianFrenchName,
		language.Dutch:                mollyDutchName,
		language.French:               mollyFrenchName,
		language.German:               mollyGermanName,
		language.Italian:              mollyItalianName,
		language.Japanese:             mollyJapaneseName,
		language.Korean:               mollyKoreanName,
		language.LatinAmericanSpanish: mollyLatinAmericanSpanishName,
		language.Russian:              mollyRussianName,
		language.SimplifiedChinese:    mollySimplifiedChineseName,
		language.Spanish:              mollySpanishName,
		language.TraditionalChinese:   mollyTraditionalChineseName}
)

var (
	// mollyCharacter represents molly character.
	mollyCharacter = nook.Character{
		Animal:   animal.Duck,
		Birthday: mollyBirthday,
		Code:     mollyCode,
		Key:      character.Molly,
		Gender:   gender.Female,
		Name:     mollyName,
		Special:  false}
)

var (
	// mollyAmericanEnglishPhrase represents molly american english phrase.
	mollyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "quackidee"}

	// mollyCanadianFrenchPhrase represents molly canadian french phrase.
	mollyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "canénette"}

	// mollyDutchPhrase represents molly dutch phrase.
	mollyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "topper"}

	// mollyFrenchPhrase represents molly french phrase.
	mollyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "canénette"}

	// mollyGermanPhrase represents molly german phrase.
	mollyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "krümel"}

	// mollyItalianPhrase represents molly italian phrase.
	mollyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "coin coin"}

	// mollyJapanesePhrase represents molly japanese phrase.
	mollyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "そうかも"}

	// mollyKoreanPhrase represents molly korean phrase.
	mollyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그럴지도"}

	// mollyLatinAmericanSpanishPhrase represents molly latin american spanish phrase.
	mollyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cuáquidi"}

	// mollyRussianPhrase represents molly russian phrase.
	mollyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "крякульки"}

	// mollySimplifiedChinesePhrase represents molly simplified chinese phrase.
	mollySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "说不定鸭"}

	// mollySpanishPhrase represents molly spanish phrase.
	mollySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "deverdá"}

	// mollyTraditionalChinesePhrase represents molly traditional chinese phrase.
	mollyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "說不定鴨"}
)

var (
	// mollyPhrase represents molly phrase.
	mollyPhrase = nook.Languages{
		language.AmericanEnglish:      mollyAmericanEnglishPhrase,
		language.CanadianFrench:       mollyCanadianFrenchPhrase,
		language.Dutch:                mollyDutchPhrase,
		language.French:               mollyFrenchPhrase,
		language.German:               mollyGermanPhrase,
		language.Italian:              mollyItalianPhrase,
		language.Japanese:             mollyJapanesePhrase,
		language.Korean:               mollyKoreanPhrase,
		language.LatinAmericanSpanish: mollyLatinAmericanSpanishPhrase,
		language.Russian:              mollyRussianPhrase,
		language.SimplifiedChinese:    mollySimplifiedChinesePhrase,
		language.Spanish:              mollySpanishPhrase,
		language.TraditionalChinese:   mollyTraditionalChinesePhrase}
)

var (
	// Molly represents molly.
	Molly = nook.Villager{
		Character:   mollyCharacter,
		Personality: personality.Normal,
		Phrase:      mollyPhrase}
)

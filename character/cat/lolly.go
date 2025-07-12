package cat

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
	// lollyBirthday represents lolly birthday.
	lollyBirthday = nook.Birthday{
		Day:   27,
		Month: time.March}
)

var (
	// lollyCode represents lolly code.
	lollyCode = nook.Code{
		Value: "cat18"}
)

var (
	// lollyAmericanEnglishName represents lolly american english name.
	lollyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Lolly"}

	// lollyCanadianFrenchName represents lolly canadian french name.
	lollyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Linette"}

	// lollyDutchName represents lolly dutch name.
	lollyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Lolly"}

	// lollyFrenchName represents lolly french name.
	lollyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Linette"}

	// lollyGermanName represents lolly german name.
	lollyGermanName = nook.Name{
		Language: language.German,
		Value:    "Feline"}

	// lollyItalianName represents lolly italian name.
	lollyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Maty"}

	// lollyJapaneseName represents lolly japanese name.
	lollyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ラムネ"}

	// lollyKoreanName represents lolly korean name.
	lollyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "사이다"}

	// lollyLatinAmericanSpanishName represents lolly latin american spanish name.
	lollyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Feli"}

	// lollyRussianName represents lolly russian name.
	lollyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лолли"}

	// lollySimplifiedChineseName represents lolly simplified chinese name.
	lollySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "柠檬娜"}

	// lollySpanishName represents lolly spanish name.
	lollySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Feli"}

	// lollyTraditionalChineseName represents lolly traditional chinese name.
	lollyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "檸檬娜"}
)

var (
	// lollyName represents lolly name.
	lollyName = nook.Languages{
		language.AmericanEnglish:      lollyAmericanEnglishName,
		language.CanadianFrench:       lollyCanadianFrenchName,
		language.Dutch:                lollyDutchName,
		language.French:               lollyFrenchName,
		language.German:               lollyGermanName,
		language.Italian:              lollyItalianName,
		language.Japanese:             lollyJapaneseName,
		language.Korean:               lollyKoreanName,
		language.LatinAmericanSpanish: lollyLatinAmericanSpanishName,
		language.Russian:              lollyRussianName,
		language.SimplifiedChinese:    lollySimplifiedChineseName,
		language.Spanish:              lollySpanishName,
		language.TraditionalChinese:   lollyTraditionalChineseName}
)

var (
	// lollyCharacter represents lolly character.
	lollyCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: lollyBirthday,
		Code:     lollyCode,
		Key:      character.Lolly,
		Gender:   gender.Female,
		Name:     lollyName,
		Special:  false}
)

var (
	// lollyAmericanEnglishPhrase represents lolly american english phrase.
	lollyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bonbon"}

	// lollyCanadianFrenchPhrase represents lolly canadian french phrase.
	lollyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ron-ron"}

	// lollyDutchPhrase represents lolly dutch phrase.
	lollyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "zuurtje"}

	// lollyFrenchPhrase represents lolly french phrase.
	lollyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ron-ron"}

	// lollyGermanPhrase represents lolly german phrase.
	lollyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kratz"}

	// lollyItalianPhrase represents lolly italian phrase.
	lollyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "purrrr"}

	// lollyJapanesePhrase represents lolly japanese phrase.
	lollyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "あのね"}

	// lollyKoreanPhrase represents lolly korean phrase.
	lollyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "퐁퐁"}

	// lollyLatinAmericanSpanishPhrase represents lolly latin american spanish phrase.
	lollyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "purrrr"}

	// lollyRussianPhrase represents lolly russian phrase.
	lollyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "конфетка"}

	// lollySimplifiedChinesePhrase represents lolly simplified chinese phrase.
	lollySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "娜个"}

	// lollySpanishPhrase represents lolly spanish phrase.
	lollySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "purrrr"}

	// lollyTraditionalChinesePhrase represents lolly traditional chinese phrase.
	lollyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "娜個"}
)

var (
	// lollyPhrase represents lolly phrase.
	lollyPhrase = nook.Languages{
		language.AmericanEnglish:      lollyAmericanEnglishPhrase,
		language.CanadianFrench:       lollyCanadianFrenchPhrase,
		language.Dutch:                lollyDutchPhrase,
		language.French:               lollyFrenchPhrase,
		language.German:               lollyGermanPhrase,
		language.Italian:              lollyItalianPhrase,
		language.Japanese:             lollyJapanesePhrase,
		language.Korean:               lollyKoreanPhrase,
		language.LatinAmericanSpanish: lollyLatinAmericanSpanishPhrase,
		language.Russian:              lollyRussianPhrase,
		language.SimplifiedChinese:    lollySimplifiedChinesePhrase,
		language.Spanish:              lollySpanishPhrase,
		language.TraditionalChinese:   lollyTraditionalChinesePhrase}
)

var (
	// Lolly represents lolly.
	Lolly = nook.Villager{
		Character:   lollyCharacter,
		Personality: personality.Normal,
		Phrase:      lollyPhrase}
)

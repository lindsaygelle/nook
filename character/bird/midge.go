package bird

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
	// midgeBirthday represents midge birthday.
	midgeBirthday = nook.Birthday{
		Day:   12,
		Month: time.March}
)

var (
	// midgeCode represents midge code.
	midgeCode = nook.Code{
		Value: "brd08"}
)

var (
	// midgeAmericanEnglishName represents midge american english name.
	midgeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Midge"}

	// midgeCanadianFrenchName represents midge canadian french name.
	midgeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Michèle"}

	// midgeDutchName represents midge dutch name.
	midgeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Midge"}

	// midgeFrenchName represents midge french name.
	midgeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Michèle"}

	// midgeGermanName represents midge german name.
	midgeGermanName = nook.Name{
		Language: language.German,
		Value:    "Anna"}

	// midgeItalianName represents midge italian name.
	midgeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Minù"}

	// midgeJapaneseName represents midge japanese name.
	midgeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "うずまき"}

	// midgeKoreanName represents midge korean name.
	midgeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "핑글이"}

	// midgeLatinAmericanSpanishName represents midge latin american spanish name.
	midgeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Paloma"}

	// midgeRussianName represents midge russian name.
	midgeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Мидж"}

	// midgeSimplifiedChineseName represents midge simplified chinese name.
	midgeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小酒窝"}

	// midgeSpanishName represents midge spanish name.
	midgeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Paloma"}

	// midgeTraditionalChineseName represents midge traditional chinese name.
	midgeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小酒窩"}
)

var (
	// midgeName represents midge name.
	midgeName = nook.Languages{
		language.AmericanEnglish:      midgeAmericanEnglishName,
		language.CanadianFrench:       midgeCanadianFrenchName,
		language.Dutch:                midgeDutchName,
		language.French:               midgeFrenchName,
		language.German:               midgeGermanName,
		language.Italian:              midgeItalianName,
		language.Japanese:             midgeJapaneseName,
		language.Korean:               midgeKoreanName,
		language.LatinAmericanSpanish: midgeLatinAmericanSpanishName,
		language.Russian:              midgeRussianName,
		language.SimplifiedChinese:    midgeSimplifiedChineseName,
		language.Spanish:              midgeSpanishName,
		language.TraditionalChinese:   midgeTraditionalChineseName}
)

var (
	// midgeCharacter represents midge character.
	midgeCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: midgeBirthday,
		Code:     midgeCode,
		Key:      character.Midge,
		Gender:   gender.Female,
		Name:     midgeName,
		Special:  false}
)

var (
	// midgeAmericanEnglishPhrase represents midge american english phrase.
	midgeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "tweedledee"}

	// midgeCanadianFrenchPhrase represents midge canadian french phrase.
	midgeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "loulou"}

	// midgeDutchPhrase represents midge dutch phrase.
	midgeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kwetter"}

	// midgeFrenchPhrase represents midge french phrase.
	midgeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "loulou"}

	// midgeGermanPhrase represents midge german phrase.
	midgeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "twitiriti"}

	// midgeItalianPhrase represents midge italian phrase.
	midgeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "firulì"}

	// midgeJapanesePhrase represents midge japanese phrase.
	midgeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "キョン"}

	// midgeKoreanPhrase represents midge korean phrase.
	midgeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "핑그르르"}

	// midgeLatinAmericanSpanishPhrase represents midge latin american spanish phrase.
	midgeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "arrurrú"}

	// midgeRussianPhrase represents midge russian phrase.
	midgeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "тви-тви"}

	// midgeSimplifiedChinesePhrase represents midge simplified chinese phrase.
	midgeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "脸红"}

	// midgeSpanishPhrase represents midge spanish phrase.
	midgeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "arrurrú"}

	// midgeTraditionalChinesePhrase represents midge traditional chinese phrase.
	midgeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "臉紅"}
)

var (
	// midgePhrase represents midge phrase.
	midgePhrase = nook.Languages{
		language.AmericanEnglish:      midgeAmericanEnglishPhrase,
		language.CanadianFrench:       midgeCanadianFrenchPhrase,
		language.Dutch:                midgeDutchPhrase,
		language.French:               midgeFrenchPhrase,
		language.German:               midgeGermanPhrase,
		language.Italian:              midgeItalianPhrase,
		language.Japanese:             midgeJapanesePhrase,
		language.Korean:               midgeKoreanPhrase,
		language.LatinAmericanSpanish: midgeLatinAmericanSpanishPhrase,
		language.Russian:              midgeRussianPhrase,
		language.SimplifiedChinese:    midgeSimplifiedChinesePhrase,
		language.Spanish:              midgeSpanishPhrase,
		language.TraditionalChinese:   midgeTraditionalChinesePhrase}
)

var (
	// Midge represents midge.
	Midge = nook.Villager{
		Character:   midgeCharacter,
		Personality: personality.Normal,
		Phrase:      midgePhrase}
)

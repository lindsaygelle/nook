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
	// mallaryBirthday represents mallary birthday.
	mallaryBirthday = nook.Birthday{
		Day:   17,
		Month: time.November}
)

var (
	// mallaryCode represents mallary code.
	mallaryCode = nook.Code{
		Value: "duk06"}
)

var (
	// mallaryAmericanEnglishName represents mallary american english name.
	mallaryAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Mallary"}

	// mallaryCanadianFrenchName represents mallary canadian french name.
	mallaryCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Mallory"}

	// mallaryDutchName represents mallary dutch name.
	mallaryDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Mallary"}

	// mallaryFrenchName represents mallary french name.
	mallaryFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mallory"}

	// mallaryGermanName represents mallary german name.
	mallaryGermanName = nook.Name{
		Language: language.German,
		Value:    "Marina"}

	// mallaryItalianName represents mallary italian name.
	mallaryItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Sofia"}

	// mallaryJapaneseName represents mallary japanese name.
	mallaryJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "スミモモ"}

	// mallaryKoreanName represents mallary korean name.
	mallaryKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "스미모"}

	// mallaryLatinAmericanSpanishName represents mallary latin american spanish name.
	mallaryLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Mirta"}

	// mallaryRussianName represents mallary russian name.
	mallaryRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Мэллари"}

	// mallarySimplifiedChineseName represents mallary simplified chinese name.
	mallarySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "苏米"}

	// mallarySpanishName represents mallary spanish name.
	mallarySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Mirta"}

	// mallaryTraditionalChineseName represents mallary traditional chinese name.
	mallaryTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "蘇米"}
)

var (
	// mallaryName represents mallary name.
	mallaryName = nook.Languages{
		language.AmericanEnglish:      mallaryAmericanEnglishName,
		language.CanadianFrench:       mallaryCanadianFrenchName,
		language.Dutch:                mallaryDutchName,
		language.French:               mallaryFrenchName,
		language.German:               mallaryGermanName,
		language.Italian:              mallaryItalianName,
		language.Japanese:             mallaryJapaneseName,
		language.Korean:               mallaryKoreanName,
		language.LatinAmericanSpanish: mallaryLatinAmericanSpanishName,
		language.Russian:              mallaryRussianName,
		language.SimplifiedChinese:    mallarySimplifiedChineseName,
		language.Spanish:              mallarySpanishName,
		language.TraditionalChinese:   mallaryTraditionalChineseName}
)

var (
	// mallaryCharacter represents mallary character.
	mallaryCharacter = nook.Character{
		Animal:   animal.Duck,
		Birthday: mallaryBirthday,
		Code:     mallaryCode,
		Key:      character.Mallary,
		Gender:   gender.Female,
		Name:     mallaryName,
		Special:  false}
)

var (
	// mallaryAmericanEnglishPhrase represents mallary american english phrase.
	mallaryAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "quackpth"}

	// mallaryCanadianFrenchPhrase represents mallary canadian french phrase.
	mallaryCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "coinpff"}

	// mallaryDutchPhrase represents mallary dutch phrase.
	mallaryDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kwaks"}

	// mallaryFrenchPhrase represents mallary french phrase.
	mallaryFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "coinpff"}

	// mallaryGermanPhrase represents mallary german phrase.
	mallaryGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "quackpss"}

	// mallaryItalianPhrase represents mallary italian phrase.
	mallaryItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "quackplà"}

	// mallaryJapanesePhrase represents mallary japanese phrase.
	mallaryJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ヨネ"}

	// mallaryKoreanPhrase represents mallary korean phrase.
	mallaryKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "모모"}

	// mallaryLatinAmericanSpanishPhrase represents mallary latin american spanish phrase.
	mallaryLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "recuac"}

	// mallaryRussianPhrase represents mallary russian phrase.
	mallaryRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кряк-фи"}

	// mallarySimplifiedChinesePhrase represents mallary simplified chinese phrase.
	mallarySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "八十八"}

	// mallarySpanishPhrase represents mallary spanish phrase.
	mallarySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "recuac"}

	// mallaryTraditionalChinesePhrase represents mallary traditional chinese phrase.
	mallaryTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "八十八"}
)

var (
	// mallaryPhrase represents mallary phrase.
	mallaryPhrase = nook.Languages{
		language.AmericanEnglish:      mallaryAmericanEnglishPhrase,
		language.CanadianFrench:       mallaryCanadianFrenchPhrase,
		language.Dutch:                mallaryDutchPhrase,
		language.French:               mallaryFrenchPhrase,
		language.German:               mallaryGermanPhrase,
		language.Italian:              mallaryItalianPhrase,
		language.Japanese:             mallaryJapanesePhrase,
		language.Korean:               mallaryKoreanPhrase,
		language.LatinAmericanSpanish: mallaryLatinAmericanSpanishPhrase,
		language.Russian:              mallaryRussianPhrase,
		language.SimplifiedChinese:    mallarySimplifiedChinesePhrase,
		language.Spanish:              mallarySpanishPhrase,
		language.TraditionalChinese:   mallaryTraditionalChinesePhrase}
)

var (
	// Mallary represents mallary.
	Mallary = nook.Villager{
		Character:   mallaryCharacter,
		Personality: personality.Snooty,
		Phrase:      mallaryPhrase}
)

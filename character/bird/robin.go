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
	// robinBirthday represents robin birthday.
	robinBirthday = nook.Birthday{
		Day:   4,
		Month: time.December}
)

var (
	// robinCode represents robin code.
	robinCode = nook.Code{
		Value: "brd01"}
)

var (
	// robinAmericanEnglishName represents robin american english name.
	robinAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Robin"}

	// robinCanadianFrenchName represents robin canadian french name.
	robinCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Robie"}

	// robinDutchName represents robin dutch name.
	robinDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Robin"}

	// robinFrenchName represents robin french name.
	robinFrenchName = nook.Name{
		Language: language.French,
		Value:    "Robie"}

	// robinGermanName represents robin german name.
	robinGermanName = nook.Name{
		Language: language.German,
		Value:    "Jule"}

	// robinItalianName represents robin italian name.
	robinItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rossana"}

	// robinJapaneseName represents robin japanese name.
	robinJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "パーチク"}

	// robinKoreanName represents robin korean name.
	robinKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "파틱"}

	// robinLatinAmericanSpanishName represents robin latin american spanish name.
	robinLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Aria"}

	// robinRussianName represents robin russian name.
	robinRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Робин"}

	// robinSimplifiedChineseName represents robin simplified chinese name.
	robinSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喳喳"}

	// robinSpanishName represents robin spanish name.
	robinSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Aria"}

	// robinTraditionalChineseName represents robin traditional chinese name.
	robinTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "喳喳"}
)

var (
	// robinName represents robin name.
	robinName = nook.Languages{
		language.AmericanEnglish:      robinAmericanEnglishName,
		language.CanadianFrench:       robinCanadianFrenchName,
		language.Dutch:                robinDutchName,
		language.French:               robinFrenchName,
		language.German:               robinGermanName,
		language.Italian:              robinItalianName,
		language.Japanese:             robinJapaneseName,
		language.Korean:               robinKoreanName,
		language.LatinAmericanSpanish: robinLatinAmericanSpanishName,
		language.Russian:              robinRussianName,
		language.SimplifiedChinese:    robinSimplifiedChineseName,
		language.Spanish:              robinSpanishName,
		language.TraditionalChinese:   robinTraditionalChineseName}
)

var (
	// robinCharacter represents robin character.
	robinCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: robinBirthday,
		Code:     robinCode,
		Key:      character.Robin,
		Gender:   gender.Female,
		Name:     robinName,
		Special:  false}
)

var (
	// robinAmericanEnglishPhrase represents robin american english phrase.
	robinAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "la-di-da"}

	// robinCanadianFrenchPhrase represents robin canadian french phrase.
	robinCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ta-di-da"}

	// robinDutchPhrase represents robin dutch phrase.
	robinDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "lalala"}

	// robinFrenchPhrase represents robin french phrase.
	robinFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ta-di-da"}

	// robinGermanPhrase represents robin german phrase.
	robinGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "ladida"}

	// robinItalianPhrase represents robin italian phrase.
	robinItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "yuppidù"}

	// robinJapanesePhrase represents robin japanese phrase.
	robinJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "さ"}

	// robinKoreanPhrase represents robin korean phrase.
	robinKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "글쎄"}

	// robinLatinAmericanSpanishPhrase represents robin latin american spanish phrase.
	robinLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "larará"}

	// robinRussianPhrase represents robin russian phrase.
	robinRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ля-ля-ля"}

	// robinSimplifiedChinesePhrase represents robin simplified chinese phrase.
	robinSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喳"}

	// robinSpanishPhrase represents robin spanish phrase.
	robinSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "picatoste"}

	// robinTraditionalChinesePhrase represents robin traditional chinese phrase.
	robinTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "喳"}
)

var (
	// robinPhrase represents robin phrase.
	robinPhrase = nook.Languages{
		language.AmericanEnglish:      robinAmericanEnglishPhrase,
		language.CanadianFrench:       robinCanadianFrenchPhrase,
		language.Dutch:                robinDutchPhrase,
		language.French:               robinFrenchPhrase,
		language.German:               robinGermanPhrase,
		language.Italian:              robinItalianPhrase,
		language.Japanese:             robinJapanesePhrase,
		language.Korean:               robinKoreanPhrase,
		language.LatinAmericanSpanish: robinLatinAmericanSpanishPhrase,
		language.Russian:              robinRussianPhrase,
		language.SimplifiedChinese:    robinSimplifiedChinesePhrase,
		language.Spanish:              robinSpanishPhrase,
		language.TraditionalChinese:   robinTraditionalChinesePhrase}
)

var (
	// Robin represents robin.
	Robin = nook.Villager{
		Character:   robinCharacter,
		Personality: personality.Snooty,
		Phrase:      robinPhrase}
)

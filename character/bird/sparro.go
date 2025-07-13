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
	// sparroBirthday represents sparro birthday.
	sparroBirthday = nook.Birthday{
		Day:   20,
		Month: time.November}
)

var (
	// sparroCode represents sparro code.
	sparroCode = nook.Code{
		Value: "brd18"}
)

var (
	// sparroAmericanEnglishName represents sparro american english name.
	sparroAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sparro"}

	// sparroCanadianFrenchName represents sparro canadian french name.
	sparroCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Darius"}

	// sparroDutchName represents sparro dutch name.
	sparroDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sparro"}

	// sparroFrenchName represents sparro french name.
	sparroFrenchName = nook.Name{
		Language: language.French,
		Value:    "Darius"}

	// sparroGermanName represents sparro german name.
	sparroGermanName = nook.Name{
		Language: language.German,
		Value:    "Nestor"}

	// sparroItalianName represents sparro italian name.
	sparroItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Piumardo"}

	// sparroJapaneseName represents sparro japanese name.
	sparroJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ちゅんのすけ"}

	// sparroKoreanName represents sparro korean name.
	sparroKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "춘섭"}

	// sparroLatinAmericanSpanishName represents sparro latin american spanish name.
	sparroLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Jaime"}

	// sparroRussianName represents sparro russian name.
	sparroRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Спарро"}

	// sparroSimplifiedChineseName represents sparro simplified chinese name.
	sparroSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "周之翔"}

	// sparroSpanishName represents sparro spanish name.
	sparroSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Jaime"}

	// sparroTraditionalChineseName represents sparro traditional chinese name.
	sparroTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "周之翔"}
)

var (
	// sparroName represents sparro name.
	sparroName = nook.Languages{
		language.AmericanEnglish:      sparroAmericanEnglishName,
		language.CanadianFrench:       sparroCanadianFrenchName,
		language.Dutch:                sparroDutchName,
		language.French:               sparroFrenchName,
		language.German:               sparroGermanName,
		language.Italian:              sparroItalianName,
		language.Japanese:             sparroJapaneseName,
		language.Korean:               sparroKoreanName,
		language.LatinAmericanSpanish: sparroLatinAmericanSpanishName,
		language.Russian:              sparroRussianName,
		language.SimplifiedChinese:    sparroSimplifiedChineseName,
		language.Spanish:              sparroSpanishName,
		language.TraditionalChinese:   sparroTraditionalChineseName}
)

var (
	// sparroCharacter represents sparro character.
	sparroCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: sparroBirthday,
		Code:     sparroCode,
		Key:      character.Sparro,
		Gender:   gender.Male,
		Name:     sparroName,
		Special:  false}
)

var (
	// sparroAmericanEnglishPhrase represents sparro american english phrase.
	sparroAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "like whoa"}

	// sparroCanadianFrenchPhrase represents sparro canadian french phrase.
	sparroCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "piaaaaf"}

	// sparroDutchPhrase represents sparro dutch phrase.
	sparroDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "wreed"}

	// sparroFrenchPhrase represents sparro french phrase.
	sparroFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "piaaaaf"}

	// sparroGermanPhrase represents sparro german phrase.
	sparroGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "birtz"}

	// sparroItalianPhrase represents sparro italian phrase.
	sparroItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "adiós"}

	// sparroJapanesePhrase represents sparro japanese phrase.
	sparroJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だチュン"}

	// sparroKoreanPhrase represents sparro korean phrase.
	sparroKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "섭섭"}

	// sparroLatinAmericanSpanishPhrase represents sparro latin american spanish phrase.
	sparroLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "chirpi"}

	// sparroRussianPhrase represents sparro russian phrase.
	sparroRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "так-так"}

	// sparroSimplifiedChinesePhrase represents sparro simplified chinese phrase.
	sparroSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啁啁"}

	// sparroSpanishPhrase represents sparro spanish phrase.
	sparroSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "gorrión"}

	// sparroTraditionalChinesePhrase represents sparro traditional chinese phrase.
	sparroTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "啁啁"}
)

var (
	// sparroPhrase represents sparro phrase.
	sparroPhrase = nook.Languages{
		language.AmericanEnglish:      sparroAmericanEnglishPhrase,
		language.CanadianFrench:       sparroCanadianFrenchPhrase,
		language.Dutch:                sparroDutchPhrase,
		language.French:               sparroFrenchPhrase,
		language.German:               sparroGermanPhrase,
		language.Italian:              sparroItalianPhrase,
		language.Japanese:             sparroJapanesePhrase,
		language.Korean:               sparroKoreanPhrase,
		language.LatinAmericanSpanish: sparroLatinAmericanSpanishPhrase,
		language.Russian:              sparroRussianPhrase,
		language.SimplifiedChinese:    sparroSimplifiedChinesePhrase,
		language.Spanish:              sparroSpanishPhrase,
		language.TraditionalChinese:   sparroTraditionalChinesePhrase}
)

var (
	// Sparro represents sparro.
	Sparro = nook.Villager{
		Character:   sparroCharacter,
		Personality: personality.Jock,
		Phrase:      sparroPhrase}
)

package koala

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
	// sydneyBirthday represents sydney birthday.
	sydneyBirthday = nook.Birthday{
		Day:   21,
		Month: time.June}
)

var (
	// sydneyCode represents sydney code.
	sydneyCode = nook.Code{
		Value: "kal03"}
)

var (
	// sydneyAmericanEnglishName represents sydney american english name.
	sydneyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sydney"}

	// sydneyCanadianFrenchName represents sydney canadian french name.
	sydneyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Koaline"}

	// sydneyDutchName represents sydney dutch name.
	sydneyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sydney"}

	// sydneyFrenchName represents sydney french name.
	sydneyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Koaline"}

	// sydneyGermanName represents sydney german name.
	sydneyGermanName = nook.Name{
		Language: language.German,
		Value:    "Silke"}

	// sydneyItalianName represents sydney italian name.
	sydneyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Sidney"}

	// sydneyJapaneseName represents sydney japanese name.
	sydneyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "シドニー"}

	// sydneyKoreanName represents sydney korean name.
	sydneyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "시드니"}

	// sydneyLatinAmericanSpanishName represents sydney latin american spanish name.
	sydneyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Sidney"}

	// sydneyRussianName represents sydney russian name.
	sydneyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Сидни"}

	// sydneySimplifiedChineseName represents sydney simplified chinese name.
	sydneySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "思颖"}

	// sydneySpanishName represents sydney spanish name.
	sydneySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Sidney"}

	// sydneyTraditionalChineseName represents sydney traditional chinese name.
	sydneyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "思穎"}
)

var (
	// sydneyName represents sydney name.
	sydneyName = nook.Languages{
		language.AmericanEnglish:      sydneyAmericanEnglishName,
		language.CanadianFrench:       sydneyCanadianFrenchName,
		language.Dutch:                sydneyDutchName,
		language.French:               sydneyFrenchName,
		language.German:               sydneyGermanName,
		language.Italian:              sydneyItalianName,
		language.Japanese:             sydneyJapaneseName,
		language.Korean:               sydneyKoreanName,
		language.LatinAmericanSpanish: sydneyLatinAmericanSpanishName,
		language.Russian:              sydneyRussianName,
		language.SimplifiedChinese:    sydneySimplifiedChineseName,
		language.Spanish:              sydneySpanishName,
		language.TraditionalChinese:   sydneyTraditionalChineseName}
)

var (
	// sydneyCharacter represents sydney character.
	sydneyCharacter = nook.Character{
		Animal:   animal.Koala,
		Birthday: sydneyBirthday,
		Code:     sydneyCode,
		Key:      character.Sydney,
		Gender:   gender.Female,
		Name:     sydneyName,
		Special:  false}
)

var (
	// sydneyAmericanEnglishPhrase represents sydney american english phrase.
	sydneyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sunshine"}

	// sydneyCanadianFrenchPhrase represents sydney canadian french phrase.
	sydneyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mon trésor"}

	// sydneyDutchPhrase represents sydney dutch phrase.
	sydneyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "zonnetje"}

	// sydneyFrenchPhrase represents sydney french phrase.
	sydneyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mon trésor"}

	// sydneyGermanPhrase represents sydney german phrase.
	sydneyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "sternchen"}

	// sydneyItalianPhrase represents sydney italian phrase.
	sydneyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "calipto"}

	// sydneyJapanesePhrase represents sydney japanese phrase.
	sydneyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だコアラ"}

	// sydneyKoreanPhrase represents sydney korean phrase.
	sydneyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "발그레"}

	// sydneyLatinAmericanSpanishPhrase represents sydney latin american spanish phrase.
	sydneyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "calipto"}

	// sydneyRussianPhrase represents sydney russian phrase.
	sydneyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "солнышко"}

	// sydneySimplifiedChinesePhrase represents sydney simplified chinese phrase.
	sydneySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "无尾熊"}

	// sydneySpanishPhrase represents sydney spanish phrase.
	sydneySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "súper"}

	// sydneyTraditionalChinesePhrase represents sydney traditional chinese phrase.
	sydneyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "無尾熊"}
)

var (
	// sydneyPhrase represents sydney phrase.
	sydneyPhrase = nook.Languages{
		language.AmericanEnglish:      sydneyAmericanEnglishPhrase,
		language.CanadianFrench:       sydneyCanadianFrenchPhrase,
		language.Dutch:                sydneyDutchPhrase,
		language.French:               sydneyFrenchPhrase,
		language.German:               sydneyGermanPhrase,
		language.Italian:              sydneyItalianPhrase,
		language.Japanese:             sydneyJapanesePhrase,
		language.Korean:               sydneyKoreanPhrase,
		language.LatinAmericanSpanish: sydneyLatinAmericanSpanishPhrase,
		language.Russian:              sydneyRussianPhrase,
		language.SimplifiedChinese:    sydneySimplifiedChinesePhrase,
		language.Spanish:              sydneySpanishPhrase,
		language.TraditionalChinese:   sydneyTraditionalChinesePhrase}
)

var (
	// Sydney represents sydney.
	Sydney = nook.Villager{
		Character:   sydneyCharacter,
		Personality: personality.Normal,
		Phrase:      sydneyPhrase}
)

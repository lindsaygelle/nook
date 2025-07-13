package bearcub

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
	// aisleBirthday represents aisle birthday.
	aisleBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// aisleCode represents aisle code.
	aisleCode = nook.Code{
		Value: ""}
)

var (
	// aisleAmericanEnglishName represents aisle american english name.
	aisleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Aisle"}

	// aisleCanadianFrenchName represents aisle canadian french name.
	aisleCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// aisleDutchName represents aisle dutch name.
	aisleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// aisleFrenchName represents aisle french name.
	aisleFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	// aisleGermanName represents aisle german name.
	aisleGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	// aisleItalianName represents aisle italian name.
	aisleItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// aisleJapaneseName represents aisle japanese name.
	aisleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アイル"}

	// aisleKoreanName represents aisle korean name.
	aisleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// aisleLatinAmericanSpanishName represents aisle latin american spanish name.
	aisleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// aisleRussianName represents aisle russian name.
	aisleRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// aisleSimplifiedChineseName represents aisle simplified chinese name.
	aisleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// aisleSpanishName represents aisle spanish name.
	aisleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// aisleTraditionalChineseName represents aisle traditional chinese name.
	aisleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// aisleName represents aisle name.
	aisleName = nook.Languages{
		language.AmericanEnglish:      aisleAmericanEnglishName,
		language.CanadianFrench:       aisleCanadianFrenchName,
		language.Dutch:                aisleDutchName,
		language.French:               aisleFrenchName,
		language.German:               aisleGermanName,
		language.Italian:              aisleItalianName,
		language.Japanese:             aisleJapaneseName,
		language.Korean:               aisleKoreanName,
		language.LatinAmericanSpanish: aisleLatinAmericanSpanishName,
		language.Russian:              aisleRussianName,
		language.SimplifiedChinese:    aisleSimplifiedChineseName,
		language.Spanish:              aisleSpanishName,
		language.TraditionalChinese:   aisleTraditionalChineseName}
)

var (
	// aisleCharacter represents aisle character.
	aisleCharacter = nook.Character{
		Animal:   animal.Bearcub,
		Birthday: aisleBirthday,
		Code:     aisleCode,
		Key:      character.Aisle,
		Gender:   gender.Male,
		Name:     aisleName,
		Special:  false}
)

var (
	// aisleAmericanEnglishPhrase represents aisle american english phrase.
	aisleAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "あーあ"}

	// aisleCanadianFrenchPhrase represents aisle canadian french phrase.
	aisleCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// aisleDutchPhrase represents aisle dutch phrase.
	aisleDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// aisleFrenchPhrase represents aisle french phrase.
	aisleFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	// aisleGermanPhrase represents aisle german phrase.
	aisleGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	// aisleItalianPhrase represents aisle italian phrase.
	aisleItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// aisleJapanesePhrase represents aisle japanese phrase.
	aisleJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "あーあ"}

	// aisleKoreanPhrase represents aisle korean phrase.
	aisleKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// aisleLatinAmericanSpanishPhrase represents aisle latin american spanish phrase.
	aisleLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// aisleRussianPhrase represents aisle russian phrase.
	aisleRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// aisleSimplifiedChinesePhrase represents aisle simplified chinese phrase.
	aisleSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// aisleSpanishPhrase represents aisle spanish phrase.
	aisleSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// aisleTraditionalChinesePhrase represents aisle traditional chinese phrase.
	aisleTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// aislePhrase represents aisle phrase.
	aislePhrase = nook.Languages{
		language.AmericanEnglish:      aisleAmericanEnglishPhrase,
		language.CanadianFrench:       aisleCanadianFrenchPhrase,
		language.Dutch:                aisleDutchPhrase,
		language.French:               aisleFrenchPhrase,
		language.German:               aisleGermanPhrase,
		language.Italian:              aisleItalianPhrase,
		language.Japanese:             aisleJapanesePhrase,
		language.Korean:               aisleKoreanPhrase,
		language.LatinAmericanSpanish: aisleLatinAmericanSpanishPhrase,
		language.Russian:              aisleRussianPhrase,
		language.SimplifiedChinese:    aisleSimplifiedChinesePhrase,
		language.Spanish:              aisleSpanishPhrase,
		language.TraditionalChinese:   aisleTraditionalChinesePhrase}
)

var (
	// Aisle represents aisle.
	Aisle = nook.Villager{
		Character:   aisleCharacter,
		Personality: personality.Lazy,
		Phrase:      aislePhrase}
)

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
	// huggyBirthday represents huggy birthday.
	huggyBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// huggyCode represents huggy code.
	huggyCode = nook.Code{
		Value: ""}
)

var (
	// huggyAmericanEnglishName represents huggy american english name.
	huggyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Huggy"}

	// huggyCanadianFrenchName represents huggy canadian french name.
	huggyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Tina"}

	// huggyDutchName represents huggy dutch name.
	huggyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// huggyFrenchName represents huggy french name.
	huggyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tina"}

	// huggyGermanName represents huggy german name.
	huggyGermanName = nook.Name{
		Language: language.German,
		Value:    "Heike"}

	// huggyItalianName represents huggy italian name.
	huggyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Colette"}

	// huggyJapaneseName represents huggy japanese name.
	huggyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ダッコ"}

	// huggyKoreanName represents huggy korean name.
	huggyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// huggyLatinAmericanSpanishName represents huggy latin american spanish name.
	huggyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Mimí"}

	// huggyRussianName represents huggy russian name.
	huggyRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// huggySimplifiedChineseName represents huggy simplified chinese name.
	huggySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "包包"}

	// huggySpanishName represents huggy spanish name.
	huggySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Mimí"}

	// huggyTraditionalChineseName represents huggy traditional chinese name.
	huggyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// huggyName represents huggy name.
	huggyName = nook.Languages{
		language.AmericanEnglish:      huggyAmericanEnglishName,
		language.CanadianFrench:       huggyCanadianFrenchName,
		language.Dutch:                huggyDutchName,
		language.French:               huggyFrenchName,
		language.German:               huggyGermanName,
		language.Italian:              huggyItalianName,
		language.Japanese:             huggyJapaneseName,
		language.Korean:               huggyKoreanName,
		language.LatinAmericanSpanish: huggyLatinAmericanSpanishName,
		language.Russian:              huggyRussianName,
		language.SimplifiedChinese:    huggySimplifiedChineseName,
		language.Spanish:              huggySpanishName,
		language.TraditionalChinese:   huggyTraditionalChineseName}
)

var (
	// huggyCharacter represents huggy character.
	huggyCharacter = nook.Character{
		Animal:   animal.Koala,
		Birthday: huggyBirthday,
		Code:     huggyCode,
		Key:      character.Huggy,
		Gender:   gender.Female,
		Name:     huggyName,
		Special:  false}
)

var (
	// huggyAmericanEnglishPhrase represents huggy american english phrase.
	huggyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bear"}

	// huggyCanadianFrenchPhrase represents huggy canadian french phrase.
	huggyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "nounours"}

	// huggyDutchPhrase represents huggy dutch phrase.
	huggyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// huggyFrenchPhrase represents huggy french phrase.
	huggyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "nounours"}

	// huggyGermanPhrase represents huggy german phrase.
	huggyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bärchen"}

	// huggyItalianPhrase represents huggy italian phrase.
	huggyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "coalilà"}

	// huggyJapanesePhrase represents huggy japanese phrase.
	huggyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "エヘ"}

	// huggyKoreanPhrase represents huggy korean phrase.
	huggyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// huggyLatinAmericanSpanishPhrase represents huggy latin american spanish phrase.
	huggyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "osi-cosi"}

	// huggyRussianPhrase represents huggy russian phrase.
	huggyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// huggySimplifiedChinesePhrase represents huggy simplified chinese phrase.
	huggySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啀"}

	// huggySpanishPhrase represents huggy spanish phrase.
	huggySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "osi-cosi"}

	// huggyTraditionalChinesePhrase represents huggy traditional chinese phrase.
	huggyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// huggyPhrase represents huggy phrase.
	huggyPhrase = nook.Languages{
		language.AmericanEnglish:      huggyAmericanEnglishPhrase,
		language.CanadianFrench:       huggyCanadianFrenchPhrase,
		language.Dutch:                huggyDutchPhrase,
		language.French:               huggyFrenchPhrase,
		language.German:               huggyGermanPhrase,
		language.Italian:              huggyItalianPhrase,
		language.Japanese:             huggyJapanesePhrase,
		language.Korean:               huggyKoreanPhrase,
		language.LatinAmericanSpanish: huggyLatinAmericanSpanishPhrase,
		language.Russian:              huggyRussianPhrase,
		language.SimplifiedChinese:    huggySimplifiedChinesePhrase,
		language.Spanish:              huggySpanishPhrase,
		language.TraditionalChinese:   huggyTraditionalChinesePhrase}
)

var (
	// Huggy represents huggy.
	Huggy = nook.Villager{
		Character:   huggyCharacter,
		Personality: personality.Peppy,
		Phrase:      huggyPhrase}
)

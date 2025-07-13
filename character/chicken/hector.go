package chicken

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
	// hectorBirthday represents hector birthday.
	hectorBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// hectorCode represents hector code.
	hectorCode = nook.Code{
		Value: ""}
)

var (
	// hectorAmericanEnglishName represents hector american english name.
	hectorAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Hector"}

	// hectorCanadianFrenchName represents hector canadian french name.
	hectorCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// hectorDutchName represents hector dutch name.
	hectorDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// hectorFrenchName represents hector french name.
	hectorFrenchName = nook.Name{
		Language: language.French,
		Value:    "Hector"}

	// hectorGermanName represents hector german name.
	hectorGermanName = nook.Name{
		Language: language.German,
		Value:    "Hektor"}

	// hectorItalianName represents hector italian name.
	hectorItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pettore"}

	// hectorJapaneseName represents hector japanese name.
	hectorJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ピロシ"}

	// hectorKoreanName represents hector korean name.
	hectorKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// hectorLatinAmericanSpanishName represents hector latin american spanish name.
	hectorLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// hectorRussianName represents hector russian name.
	hectorRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// hectorSimplifiedChineseName represents hector simplified chinese name.
	hectorSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "博博"}

	// hectorSpanishName represents hector spanish name.
	hectorSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Héctor"}

	// hectorTraditionalChineseName represents hector traditional chinese name.
	hectorTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// hectorName represents hector name.
	hectorName = nook.Languages{
		language.AmericanEnglish:      hectorAmericanEnglishName,
		language.CanadianFrench:       hectorCanadianFrenchName,
		language.Dutch:                hectorDutchName,
		language.French:               hectorFrenchName,
		language.German:               hectorGermanName,
		language.Italian:              hectorItalianName,
		language.Japanese:             hectorJapaneseName,
		language.Korean:               hectorKoreanName,
		language.LatinAmericanSpanish: hectorLatinAmericanSpanishName,
		language.Russian:              hectorRussianName,
		language.SimplifiedChinese:    hectorSimplifiedChineseName,
		language.Spanish:              hectorSpanishName,
		language.TraditionalChinese:   hectorTraditionalChineseName}
)

var (
	// hectorCharacter represents hector character.
	hectorCharacter = nook.Character{
		Animal:   animal.Chicken,
		Birthday: hectorBirthday,
		Code:     hectorCode,
		Key:      character.Hector,
		Gender:   gender.Male,
		Name:     hectorName,
		Special:  false}
)

var (
	// hectorAmericanEnglishPhrase represents hector american english phrase.
	hectorAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "coo-HAH"}

	// hectorCanadianFrenchPhrase represents hector canadian french phrase.
	hectorCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// hectorDutchPhrase represents hector dutch phrase.
	hectorDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// hectorFrenchPhrase represents hector french phrase.
	hectorFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mon coco"}

	// hectorGermanPhrase represents hector german phrase.
	hectorGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kikiriii"}

	// hectorItalianPhrase represents hector italian phrase.
	hectorItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "chirichì"}

	// hectorJapanesePhrase represents hector japanese phrase.
	hectorJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なのだ"}

	// hectorKoreanPhrase represents hector korean phrase.
	hectorKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// hectorLatinAmericanSpanishPhrase represents hector latin american spanish phrase.
	hectorLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// hectorRussianPhrase represents hector russian phrase.
	hectorRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// hectorSimplifiedChinesePhrase represents hector simplified chinese phrase.
	hectorSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "叮叮"}

	// hectorSpanishPhrase represents hector spanish phrase.
	hectorSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "yuu-ju"}

	// hectorTraditionalChinesePhrase represents hector traditional chinese phrase.
	hectorTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// hectorPhrase represents hector phrase.
	hectorPhrase = nook.Languages{
		language.AmericanEnglish:      hectorAmericanEnglishPhrase,
		language.CanadianFrench:       hectorCanadianFrenchPhrase,
		language.Dutch:                hectorDutchPhrase,
		language.French:               hectorFrenchPhrase,
		language.German:               hectorGermanPhrase,
		language.Italian:              hectorItalianPhrase,
		language.Japanese:             hectorJapanesePhrase,
		language.Korean:               hectorKoreanPhrase,
		language.LatinAmericanSpanish: hectorLatinAmericanSpanishPhrase,
		language.Russian:              hectorRussianPhrase,
		language.SimplifiedChinese:    hectorSimplifiedChinesePhrase,
		language.Spanish:              hectorSpanishPhrase,
		language.TraditionalChinese:   hectorTraditionalChinesePhrase}
)

var (
	// Hector represents hector.
	Hector = nook.Villager{
		Character:   hectorCharacter,
		Personality: personality.Jock,
		Phrase:      hectorPhrase}
)

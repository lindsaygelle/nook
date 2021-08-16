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
	hectorBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	hectorCode = nook.Code{
		Value: ""}
)

var (
	hectorAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Hector"}

	hectorCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	hectorDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	hectorFrenchName = nook.Name{
		Language: language.French,
		Value:    "Hector"}

	hectorGermanName = nook.Name{
		Language: language.German,
		Value:    "Hektor"}

	hectorItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pettore"}

	hectorJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ピロシ"}

	hectorKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	hectorLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	hectorRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	hectorSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "博博"}

	hectorSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Héctor"}

	hectorTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
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
	hectorCharacter = nook.Character{
		Animal:   animal.Chicken,
		Birthday: hectorBirthday,
		Code:     hectorCode,
		Key:      character.Hector,
		Gender:   gender.Male,
		Name:     hectorName}
)

var (
	hectorAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "coo-HAH"}

	hectorCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	hectorDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	hectorFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mon coco"}

	hectorGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kikiriii"}

	hectorItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "chirichì"}

	hectorJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なのだ"}

	hectorKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	hectorLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	hectorRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	hectorSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "叮叮"}

	hectorSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "yuu-ju"}

	hectorTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	hectorPhrase = nook.Languages{
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
	Hector = nook.Villager{
		Character:   hectorCharacter,
		Personality: personality.Jock,
		Phrase:      hectorPhrase}
)

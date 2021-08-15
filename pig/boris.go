package pig

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	borisBirthday = nook.Birthday{
		Day:   6,
		Month: time.November}
)

var (
	borisCode = nook.Code{
		Value: "pig09"}
)

var (
	borisAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Boris"}

	borisCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Borisminot"}

	borisDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Borisoink"}

	borisFrenchName = nook.Name{
		Language: language.French,
		Value:    "Borisminot"}

	borisGermanName = nook.Name{
		Language: language.German,
		Value:    "Bollequiek"}

	borisItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Borissgrunf"}

	borisJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ダリーブヒッ"}

	borisKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "보리쌀"}

	borisLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Borisoink"}

	borisRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Борисхрю-хряк"}

	borisSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "达利噗嘻"}

	borisSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Borisoink"}

	borisTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "達利噗嘻"}
)

var (
	borisName = nook.Languages{
		language.AmericanEnglish:      borisAmericanEnglishName,
		language.CanadianFrench:       borisCanadianFrenchName,
		language.Dutch:                borisDutchName,
		language.French:               borisFrenchName,
		language.German:               borisGermanName,
		language.Italian:              borisItalianName,
		language.Japanese:             borisJapaneseName,
		language.Korean:               borisKoreanName,
		language.LatinAmericanSpanish: borisLatinAmericanSpanishName,
		language.Russian:              borisRussianName,
		language.SimplifiedChinese:    borisSimplifiedChineseName,
		language.Spanish:              borisSpanishName,
		language.TraditionalChinese:   borisTraditionalChineseName}
)

var (
	borisCharacter = nook.Character{
		Animal:   Pig,
		Birthday: borisBirthday,
		Code:     borisCode,
		Gender:   nook.Male,
		Name:     borisName}
)

var (
	borisAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "schnort"}

	borisCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "minot"}

	borisDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "oink"}

	borisFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "minot"}

	borisGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "quiek"}

	borisItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sgrunf"}

	borisJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ブヒッ"}

	borisKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "쌀"}

	borisLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "oink"}

	borisRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хрю-хряк"}

	borisSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "噗嘻"}

	borisSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "oink"}

	borisTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "噗嘻"}
)

var (
	borisPhrase = nook.Languages{
		language.AmericanEnglish:      borisAmericanEnglishName,
		language.CanadianFrench:       borisCanadianFrenchName,
		language.Dutch:                borisDutchName,
		language.French:               borisFrenchName,
		language.German:               borisGermanName,
		language.Italian:              borisItalianName,
		language.Japanese:             borisJapaneseName,
		language.Korean:               borisKoreanName,
		language.LatinAmericanSpanish: borisLatinAmericanSpanishName,
		language.Russian:              borisRussianName,
		language.SimplifiedChinese:    borisSimplifiedChineseName,
		language.Spanish:              borisSpanishName,
		language.TraditionalChinese:   borisTraditionalChineseName}
)

var (
	Boris = nook.Villager{
		Character:   borisCharacter,
		Personality: nook.Cranky,
		Phrase:      borisPhrase}
)

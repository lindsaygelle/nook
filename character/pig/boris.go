package pig

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
		Value:    "Boris"}

	borisDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Boris"}

	borisFrenchName = nook.Name{
		Language: language.French,
		Value:    "Boris"}

	borisGermanName = nook.Name{
		Language: language.German,
		Value:    "Bolle"}

	borisItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Boris"}

	borisJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ダリー"}

	borisKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "보리"}

	borisLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Boris"}

	borisRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Борис"}

	borisSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "达利"}

	borisSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Boris"}

	borisTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "達利"}
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
		Animal:   animal.Pig,
		Birthday: borisBirthday,
		Code:     borisCode,
		Key:      character.Boris,
		Gender:   gender.Male,
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
		language.AmericanEnglish:      borisAmericanEnglishPhrase,
		language.CanadianFrench:       borisCanadianFrenchPhrase,
		language.Dutch:                borisDutchPhrase,
		language.French:               borisFrenchPhrase,
		language.German:               borisGermanPhrase,
		language.Italian:              borisItalianPhrase,
		language.Japanese:             borisJapanesePhrase,
		language.Korean:               borisKoreanPhrase,
		language.LatinAmericanSpanish: borisLatinAmericanSpanishPhrase,
		language.Russian:              borisRussianPhrase,
		language.SimplifiedChinese:    borisSimplifiedChinesePhrase,
		language.Spanish:              borisSpanishPhrase,
		language.TraditionalChinese:   borisTraditionalChinesePhrase}
)

var (
	Boris = nook.Villager{
		Character:   borisCharacter,
		Personality: personality.Cranky,
		Phrase:      borisPhrase}
)

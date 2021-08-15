package chicken

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	benedictBirthday = nook.Birthday{
		Day:   10,
		Month: time.October}
)

var (
	benedictCode = nook.Code{
		Value: "chn01"}
)

var (
	benedictAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Benedict"}

	benedictCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Dimitrioh oh"}

	benedictDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Benedicteitje"}

	benedictFrenchName = nook.Name{
		Language: language.French,
		Value:    "Dimitrioh oh"}

	benedictGermanName = nook.Name{
		Language: language.German,
		Value:    "Benediktputputtt"}

	benedictItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Biagiouh-hoo"}

	benedictJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ぺしみちウヒョー"}

	benedictKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "페니실린우힛"}

	benedictLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Benitococorico"}

	benedictRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бенедикто-ко-ко"}

	benedictSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "沛希好棒"}

	benedictSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Benitococorico"}

	benedictTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "沛希好棒"}
)

var (
	benedictName = nook.Languages{
		language.AmericanEnglish:      benedictAmericanEnglishName,
		language.CanadianFrench:       benedictCanadianFrenchName,
		language.Dutch:                benedictDutchName,
		language.French:               benedictFrenchName,
		language.German:               benedictGermanName,
		language.Italian:              benedictItalianName,
		language.Japanese:             benedictJapaneseName,
		language.Korean:               benedictKoreanName,
		language.LatinAmericanSpanish: benedictLatinAmericanSpanishName,
		language.Russian:              benedictRussianName,
		language.SimplifiedChinese:    benedictSimplifiedChineseName,
		language.Spanish:              benedictSpanishName,
		language.TraditionalChinese:   benedictTraditionalChineseName}
)

var (
	benedictCharacter = nook.Character{
		Animal:   Chicken,
		Birthday: benedictBirthday,
		Code:     benedictCode,
		Gender:   nook.Male,
		Name:     benedictName}
)

var (
	benedictAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "uh-hoo"}

	benedictCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "oh oh"}

	benedictDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "eitje"}

	benedictFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "oh oh"}

	benedictGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "putputtt"}

	benedictItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "uh-hoo"}

	benedictJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ウヒョー"}

	benedictKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "우힛"}

	benedictLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cocorico"}

	benedictRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "о-ко-ко"}

	benedictSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "好棒"}

	benedictSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cocorico"}

	benedictTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "好棒"}
)

var (
	benedictPhrase = nook.Languages{
		language.AmericanEnglish:      benedictAmericanEnglishName,
		language.CanadianFrench:       benedictCanadianFrenchName,
		language.Dutch:                benedictDutchName,
		language.French:               benedictFrenchName,
		language.German:               benedictGermanName,
		language.Italian:              benedictItalianName,
		language.Japanese:             benedictJapaneseName,
		language.Korean:               benedictKoreanName,
		language.LatinAmericanSpanish: benedictLatinAmericanSpanishName,
		language.Russian:              benedictRussianName,
		language.SimplifiedChinese:    benedictSimplifiedChineseName,
		language.Spanish:              benedictSpanishName,
		language.TraditionalChinese:   benedictTraditionalChineseName}
)

var (
	Benedict = nook.Villager{
		Character:   benedictCharacter,
		Personality: nook.Lazy,
		Phrase:      benedictPhrase}
)

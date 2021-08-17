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
		Value:    "Dimitri"}

	benedictDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Benedict"}

	benedictFrenchName = nook.Name{
		Language: language.French,
		Value:    "Dimitri"}

	benedictGermanName = nook.Name{
		Language: language.German,
		Value:    "Benedikt"}

	benedictItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Biagio"}

	benedictJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ぺしみち"}

	benedictKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "페니실린"}

	benedictLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Benito"}

	benedictRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бенедикт"}

	benedictSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "沛希"}

	benedictSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Benito"}

	benedictTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "沛希"}
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
		Animal:   animal.Chicken,
		Birthday: benedictBirthday,
		Code:     benedictCode,
		Key:      character.Benedict,
		Gender:   gender.Male,
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
		language.AmericanEnglish:      benedictAmericanEnglishPhrase,
		language.CanadianFrench:       benedictCanadianFrenchPhrase,
		language.Dutch:                benedictDutchPhrase,
		language.French:               benedictFrenchPhrase,
		language.German:               benedictGermanPhrase,
		language.Italian:              benedictItalianPhrase,
		language.Japanese:             benedictJapanesePhrase,
		language.Korean:               benedictKoreanPhrase,
		language.LatinAmericanSpanish: benedictLatinAmericanSpanishPhrase,
		language.Russian:              benedictRussianPhrase,
		language.SimplifiedChinese:    benedictSimplifiedChinesePhrase,
		language.Spanish:              benedictSpanishPhrase,
		language.TraditionalChinese:   benedictTraditionalChinesePhrase}
)

var (
	Benedict = nook.Villager{
		Character:   benedictCharacter,
		Personality: personality.Lazy,
		Phrase:      benedictPhrase}
)

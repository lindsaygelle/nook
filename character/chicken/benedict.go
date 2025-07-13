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
	// benedictBirthday represents benedict birthday.
	benedictBirthday = nook.Birthday{
		Day:   10,
		Month: time.October}
)

var (
	// benedictCode represents benedict code.
	benedictCode = nook.Code{
		Value: "chn01"}
)

var (
	// benedictAmericanEnglishName represents benedict american english name.
	benedictAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Benedict"}

	// benedictCanadianFrenchName represents benedict canadian french name.
	benedictCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Dimitri"}

	// benedictDutchName represents benedict dutch name.
	benedictDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Benedict"}

	// benedictFrenchName represents benedict french name.
	benedictFrenchName = nook.Name{
		Language: language.French,
		Value:    "Dimitri"}

	// benedictGermanName represents benedict german name.
	benedictGermanName = nook.Name{
		Language: language.German,
		Value:    "Benedikt"}

	// benedictItalianName represents benedict italian name.
	benedictItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Biagio"}

	// benedictJapaneseName represents benedict japanese name.
	benedictJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ぺしみち"}

	// benedictKoreanName represents benedict korean name.
	benedictKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "페니실린"}

	// benedictLatinAmericanSpanishName represents benedict latin american spanish name.
	benedictLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Benito"}

	// benedictRussianName represents benedict russian name.
	benedictRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бенедикт"}

	// benedictSimplifiedChineseName represents benedict simplified chinese name.
	benedictSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "沛希"}

	// benedictSpanishName represents benedict spanish name.
	benedictSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Benito"}

	// benedictTraditionalChineseName represents benedict traditional chinese name.
	benedictTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "沛希"}
)

var (
	// benedictName represents benedict name.
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
	// benedictCharacter represents benedict character.
	benedictCharacter = nook.Character{
		Animal:   animal.Chicken,
		Birthday: benedictBirthday,
		Code:     benedictCode,
		Key:      character.Benedict,
		Gender:   gender.Male,
		Name:     benedictName,
		Special:  false}
)

var (
	// benedictAmericanEnglishPhrase represents benedict american english phrase.
	benedictAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "uh-hoo"}

	// benedictCanadianFrenchPhrase represents benedict canadian french phrase.
	benedictCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "oh oh"}

	// benedictDutchPhrase represents benedict dutch phrase.
	benedictDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "eitje"}

	// benedictFrenchPhrase represents benedict french phrase.
	benedictFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "oh oh"}

	// benedictGermanPhrase represents benedict german phrase.
	benedictGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "putputtt"}

	// benedictItalianPhrase represents benedict italian phrase.
	benedictItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "uh-hoo"}

	// benedictJapanesePhrase represents benedict japanese phrase.
	benedictJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ウヒョー"}

	// benedictKoreanPhrase represents benedict korean phrase.
	benedictKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "우힛"}

	// benedictLatinAmericanSpanishPhrase represents benedict latin american spanish phrase.
	benedictLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cocorico"}

	// benedictRussianPhrase represents benedict russian phrase.
	benedictRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "о-ко-ко"}

	// benedictSimplifiedChinesePhrase represents benedict simplified chinese phrase.
	benedictSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "好棒"}

	// benedictSpanishPhrase represents benedict spanish phrase.
	benedictSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cocorico"}

	// benedictTraditionalChinesePhrase represents benedict traditional chinese phrase.
	benedictTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "好棒"}
)

var (
	// benedictPhrase represents benedict phrase.
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
	// Benedict represents benedict.
	Benedict = nook.Villager{
		Character:   benedictCharacter,
		Personality: personality.Lazy,
		Phrase:      benedictPhrase}
)

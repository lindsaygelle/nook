package rabbit

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
	cocoBirthday = nook.Birthday{
		Day:   1,
		Month: time.March}
)

var (
	cocoCode = nook.Code{
		Value: "rbt02"}
)

var (
	cocoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Coco"}

	cocoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Coco"}

	cocoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Coco"}

	cocoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Coco"}

	cocoGermanName = nook.Name{
		Language: language.German,
		Value:    "Koko"}

	cocoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Coco"}

	cocoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "やよい"}

	cocoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "이요"}

	cocoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cocoloca"}

	cocoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Коко"}

	cocoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "仰韶"}

	cocoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cocoloca"}

	cocoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "仰韶"}
)

var (
	cocoName = nook.Languages{
		language.AmericanEnglish:      cocoAmericanEnglishName,
		language.CanadianFrench:       cocoCanadianFrenchName,
		language.Dutch:                cocoDutchName,
		language.French:               cocoFrenchName,
		language.German:               cocoGermanName,
		language.Italian:              cocoItalianName,
		language.Japanese:             cocoJapaneseName,
		language.Korean:               cocoKoreanName,
		language.LatinAmericanSpanish: cocoLatinAmericanSpanishName,
		language.Russian:              cocoRussianName,
		language.SimplifiedChinese:    cocoSimplifiedChineseName,
		language.Spanish:              cocoSpanishName,
		language.TraditionalChinese:   cocoTraditionalChineseName}
)

var (
	cocoCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: cocoBirthday,
		Code:     cocoCode,
		Key:      character.Coco,
		Gender:   gender.Female,
		Name:     cocoName,
		Special:  false}
)

var (
	cocoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "doyoing"}

	cocoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "touloulou"}

	cocoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "bojoing"}

	cocoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "touloulou"}

	cocoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hoppsa"}

	cocoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "evviva"}

	cocoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "はにょ"}

	cocoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "삐용"}

	cocoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "toin-toin"}

	cocoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "прыг-скок"}

	cocoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "土俑"}

	cocoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "toin-toin"}

	cocoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "土俑"}
)

var (
	cocoPhrase = nook.Languages{
		language.AmericanEnglish:      cocoAmericanEnglishPhrase,
		language.CanadianFrench:       cocoCanadianFrenchPhrase,
		language.Dutch:                cocoDutchPhrase,
		language.French:               cocoFrenchPhrase,
		language.German:               cocoGermanPhrase,
		language.Italian:              cocoItalianPhrase,
		language.Japanese:             cocoJapanesePhrase,
		language.Korean:               cocoKoreanPhrase,
		language.LatinAmericanSpanish: cocoLatinAmericanSpanishPhrase,
		language.Russian:              cocoRussianPhrase,
		language.SimplifiedChinese:    cocoSimplifiedChinesePhrase,
		language.Spanish:              cocoSpanishPhrase,
		language.TraditionalChinese:   cocoTraditionalChinesePhrase}
)

var (
	Coco = nook.Villager{
		Character:   cocoCharacter,
		Personality: personality.Normal,
		Phrase:      cocoPhrase}
)

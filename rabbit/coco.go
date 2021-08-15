package rabbit

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Cocotouloulou"}

	cocoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cocobojoing"}

	cocoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Cocotouloulou"}

	cocoGermanName = nook.Name{
		Language: language.German,
		Value:    "Kokohoppsa"}

	cocoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cocoevviva"}

	cocoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "やよいはにょ"}

	cocoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "이요삐용"}

	cocoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cocolocatoin-toin"}

	cocoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кокопрыг-скок"}

	cocoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "仰韶土俑"}

	cocoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cocolocatoin-toin"}

	cocoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "仰韶土俑"}
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
		Animal:   Rabbit,
		Birthday: cocoBirthday,
		Code:     cocoCode,
		Gender:   nook.Female,
		Name:     cocoName}
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
	Coco = nook.Villager{
		Character:   cocoCharacter,
		Personality: nook.Normal,
		Phrase:      cocoPhrase}
)

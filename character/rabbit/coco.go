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
	// cocoBirthday represents Coco's birthday.
	cocoBirthday = nook.Birthday{
		Day:   1,
		Month: time.March}
)

var (
	// cocoCode represents Coco's unique code ("rbt02").
	cocoCode = nook.Code{
		Value: "rbt02"}
)

var (
	// cocoAmericanEnglishName represents Coco's name in American English.
	cocoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Coco"}

	// cocoCanadianFrenchName represents Coco's name in Canadian French.
	cocoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Coco"}

	// cocoDutchName represents Coco's name in Dutch.
	cocoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Coco"}

	// cocoFrenchName represents Coco's name in French.
	cocoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Coco"}

	// cocoGermanName represents Coco's name in German.
	cocoGermanName = nook.Name{
		Language: language.German,
		Value:    "Koko"}

	// cocoItalianName represents Coco's name in Italian.
	cocoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Coco"}

	// cocoJapaneseName represents Coco's name in Japanese.
	cocoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "やよい"}

	// cocoKoreanName represents Coco's name in Korean.
	cocoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "이요"}

	// cocoLatinAmericanSpanishName represents Coco's name in Latin American Spanish.
	cocoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cocoloca"}

	// cocoRussianName represents Coco's name in Russian.
	cocoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Коко"}

	// cocoSimplifiedChineseName represents Coco's name in Simplified Chinese.
	cocoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "仰韶"}

	// cocoSpanishName represents Coco's name in Spanish.
	cocoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cocoloca"}

	// cocoTraditionalChineseName represents Coco's name in Traditional Chinese.
	cocoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "仰韶"}
)

var (
	// cocoName represents Coco's name in different languages.
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
	// cocoCharacter represents Coco's character information.
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
	// cocoAmericanEnglishPhrase represents Coco's phrase in American English.
	cocoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "doyoing"}

	// cocoCanadianFrenchPhrase represents Coco's phrase in Canadian French.
	cocoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "touloulou"}

	// cocoDutchPhrase represents Coco's phrase in Dutch.
	cocoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "bojoing"}

	// cocoFrenchPhrase represents Coco's phrase in French.
	cocoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "touloulou"}

	// cocoGermanPhrase represents Coco's phrase in German.
	cocoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hoppsa"}

	// cocoItalianPhrase represents Coco's phrase in Italian.
	cocoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "evviva"}

	// cocoJapanesePhrase represents Coco's phrase in Japanese.
	cocoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "はにょ"}

	// cocoKoreanPhrase represents Coco's phrase in Korean.
	cocoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "삐용"}

	// cocoLatinAmericanSpanishPhrase represents Coco's phrase in Latin American Spanish.
	cocoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "toin-toin"}

	// cocoRussianPhrase represents Coco's phrase in Russian.
	cocoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "прыг-скок"}

	// cocoSimplifiedChinesePhrase represents Coco's phrase in Simplified Chinese.
	cocoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "土俑"}

	// cocoSpanishPhrase represents Coco's phrase in Spanish.
	cocoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "toin-toin"}

	// cocoTraditionalChinesePhrase represents Coco's phrase in Traditional Chinese.
	cocoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "土俑"}
)

var (
	// cocoPhrase represents Coco's phrase in different languages.
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
	// Coco represents the character Coco with her complete information.
	Coco = nook.Villager{
		Character:   cocoCharacter,
		Personality: personality.Normal,
		Phrase:      cocoPhrase}
)

package dog

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
	macBirthday = nook.Birthday{
		Day:   11,
		Month: time.November}
)

var (
	macCode = nook.Code{
		Value: "dog14"}
)

var (
	macAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Mac"}

	macCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Brutus"}

	macDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Mac"}

	macFrenchName = nook.Name{
		Language: language.French,
		Value:    "Brutus"}

	macGermanName = nook.Name{
		Language: language.German,
		Value:    "Wuffi"}

	macItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Teo"}

	macJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チャンプ"}

	macKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "챔프"}

	macLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pit"}

	macRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Мак"}

	macSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "金牌"}

	macSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pit"}

	macTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "金牌"}
)

var (
	macName = nook.Languages{
		language.AmericanEnglish:      macAmericanEnglishName,
		language.CanadianFrench:       macCanadianFrenchName,
		language.Dutch:                macDutchName,
		language.French:               macFrenchName,
		language.German:               macGermanName,
		language.Italian:              macItalianName,
		language.Japanese:             macJapaneseName,
		language.Korean:               macKoreanName,
		language.LatinAmericanSpanish: macLatinAmericanSpanishName,
		language.Russian:              macRussianName,
		language.SimplifiedChinese:    macSimplifiedChineseName,
		language.Spanish:              macSpanishName,
		language.TraditionalChinese:   macTraditionalChineseName}
)

var (
	macCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: macBirthday,
		Code:     macCode,
		Key:      character.Mac,
		Gender:   gender.Male,
		Name:     macName}
)

var (
	macAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "woo woof"}

	macCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "babines"}

	macDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "woewoef"}

	macFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "babines"}

	macGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bellbell"}

	macItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "arfidenti"}

	macJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "モグ"}

	macKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "우물우물"}

	macLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "frusky"}

	macRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ау-у-гав"}

	macSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "摘"}

	macSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "frusky"}

	macTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "摘"}
)

var (
	macPhrase = nook.Languages{
		language.AmericanEnglish:      macAmericanEnglishName,
		language.CanadianFrench:       macCanadianFrenchName,
		language.Dutch:                macDutchName,
		language.French:               macFrenchName,
		language.German:               macGermanName,
		language.Italian:              macItalianName,
		language.Japanese:             macJapaneseName,
		language.Korean:               macKoreanName,
		language.LatinAmericanSpanish: macLatinAmericanSpanishName,
		language.Russian:              macRussianName,
		language.SimplifiedChinese:    macSimplifiedChineseName,
		language.Spanish:              macSpanishName,
		language.TraditionalChinese:   macTraditionalChineseName}
)

var (
	Mac = nook.Villager{
		Character:   macCharacter,
		Personality: personality.Jock,
		Phrase:      macPhrase}
)

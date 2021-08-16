package anteater

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
	snootyBirthday = nook.Birthday{
		Day:   24,
		Month: time.October}
)

var (
	snootyCode = nook.Code{
		Value: "ant06"}
)

var (
	snootyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Snooty"}

	snootyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Tarina"}

	snootyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Snooty"}

	snootyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tarina"}

	snootyGermanName = nook.Name{
		Language: language.German,
		Value:    "Hilde"}

	snootyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lucrezia"}

	snootyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "こまち"}

	snootyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "스누티"}

	snootyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rosanari"}

	snootyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Снути"}

	snootySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "贵妃"}

	snootySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rosanari"}

	snootyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "貴妃"}
)

var (
	snootyName = nook.Languages{
		language.AmericanEnglish:      snootyAmericanEnglishName,
		language.CanadianFrench:       snootyCanadianFrenchName,
		language.Dutch:                snootyDutchName,
		language.French:               snootyFrenchName,
		language.German:               snootyGermanName,
		language.Italian:              snootyItalianName,
		language.Japanese:             snootyJapaneseName,
		language.Korean:               snootyKoreanName,
		language.LatinAmericanSpanish: snootyLatinAmericanSpanishName,
		language.Russian:              snootyRussianName,
		language.SimplifiedChinese:    snootySimplifiedChineseName,
		language.Spanish:              snootySpanishName,
		language.TraditionalChinese:   snootyTraditionalChineseName}
)

var (
	snootyCharacter = nook.Character{
		Animal:   animal.Anteater,
		Birthday: snootyBirthday,
		Code:     snootyCode,
		Key:      character.Snooty,
		Gender:   gender.Female,
		Name:     snootyName}
)

var (
	snootyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sniffff"}

	snootyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "snif snif"}

	snootyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snuifff"}

	snootyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "snif snif"}

	snootyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnuff"}

	snootyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sniffff"}

	snootyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "たもれ"}

	snootyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "추릅"}

	snootyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "sniff"}

	snootyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "шмыг"}

	snootySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "麻烦了"}

	snootySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "sniff"}

	snootyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "麻煩了"}
)

var (
	snootyPhrase = nook.Languages{
		language.AmericanEnglish:      snootyAmericanEnglishName,
		language.CanadianFrench:       snootyCanadianFrenchName,
		language.Dutch:                snootyDutchName,
		language.French:               snootyFrenchName,
		language.German:               snootyGermanName,
		language.Italian:              snootyItalianName,
		language.Japanese:             snootyJapaneseName,
		language.Korean:               snootyKoreanName,
		language.LatinAmericanSpanish: snootyLatinAmericanSpanishName,
		language.Russian:              snootyRussianName,
		language.SimplifiedChinese:    snootySimplifiedChineseName,
		language.Spanish:              snootySpanishName,
		language.TraditionalChinese:   snootyTraditionalChineseName}
)

var (
	Snooty = nook.Villager{
		Character:   snootyCharacter,
		Personality: personality.Snooty,
		Phrase:      snootyPhrase}
)

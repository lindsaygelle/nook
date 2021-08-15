package anteater

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	luluBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	luluCode = nook.Code{
		Value: ""}
)

var (
	luluAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Lulu"}

	luluCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	luluDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	luluFrenchName = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	luluGermanName = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	luluItalianName = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	luluJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ルル"}

	luluKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	luluLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	luluRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	luluSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	luluSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	luluTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	luluName = nook.Languages{
		language.AmericanEnglish:      luluAmericanEnglishName,
		language.CanadianFrench:       luluCanadianFrenchName,
		language.Dutch:                luluDutchName,
		language.French:               luluFrenchName,
		language.German:               luluGermanName,
		language.Italian:              luluItalianName,
		language.Japanese:             luluJapaneseName,
		language.Korean:               luluKoreanName,
		language.LatinAmericanSpanish: luluLatinAmericanSpanishName,
		language.Russian:              luluRussianName,
		language.SimplifiedChinese:    luluSimplifiedChineseName,
		language.Spanish:              luluSpanishName,
		language.TraditionalChinese:   luluTraditionalChineseName}
)

var (
	luluCharacter = nook.Character{
		Animal:   Anteater,
		Birthday: luluBirthday,
		Code:     luluCode,
		Gender:   nook.Female,
		Name:     luluName}
)

var (
	luluAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	luluCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	luluDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	luluFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	luluGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	luluItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	luluJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "あらあら"}

	luluKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	luluLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	luluRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	luluSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	luluSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	luluTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	luluPhrase = nook.Languages{
		language.AmericanEnglish:      luluAmericanEnglishName,
		language.CanadianFrench:       luluCanadianFrenchName,
		language.Dutch:                luluDutchName,
		language.French:               luluFrenchName,
		language.German:               luluGermanName,
		language.Italian:              luluItalianName,
		language.Japanese:             luluJapaneseName,
		language.Korean:               luluKoreanName,
		language.LatinAmericanSpanish: luluLatinAmericanSpanishName,
		language.Russian:              luluRussianName,
		language.SimplifiedChinese:    luluSimplifiedChineseName,
		language.Spanish:              luluSpanishName,
		language.TraditionalChinese:   luluTraditionalChineseName}
)

var (
	Lulu = nook.Villager{
		Character:   luluCharacter,
		Personality: nook.Snooty,
		Phrase:      luluPhrase}
)

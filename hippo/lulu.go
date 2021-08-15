package hippo

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
		Value:    "Lulu"}

	luluGermanName = nook.Name{
		Language: language.German,
		Value:    "Lulu"}

	luluItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lulù"}

	luluJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "エルニーニョ"}

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
		Value:    "萍萍"}

	luluSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lulù"}

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
		Animal:   Hippo,
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
		Value:    "mon bichon"}

	luluGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "gääähn"}

	luluItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "hippurrà"}

	luluJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ニョニョ"}

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
		Value:    "妞妞"}

	luluSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "lilalá"}

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
		Personality: nook.Peppy,
		Phrase:      luluPhrase}
)

package dog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	walkerBirthday = nook.Birthday{
		Day:   10,
		Month: time.June}
)

var (
	walkerCode = nook.Code{
		Value: "dog06"}
)

var (
	walkerAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Walker"}

	walkerCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Georgeouh"}

	walkerDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Walkerwaf"}

	walkerFrenchName = nook.Name{
		Language: language.French,
		Value:    "Georgeouh"}

	walkerGermanName = nook.Name{
		Language: language.German,
		Value:    "Fidowafff"}

	walkerItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Walterwuff"}

	walkerJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ベンバウ"}

	walkerKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "벤컹컹"}

	walkerLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Mikiguuu-ah"}

	walkerRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Уокерау-у"}

	walkerSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿笨咆"}

	walkerSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Mikihuesitos"}

	walkerTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿笨咆"}
)

var (
	walkerName = nook.Languages{
		language.AmericanEnglish:      walkerAmericanEnglishName,
		language.CanadianFrench:       walkerCanadianFrenchName,
		language.Dutch:                walkerDutchName,
		language.French:               walkerFrenchName,
		language.German:               walkerGermanName,
		language.Italian:              walkerItalianName,
		language.Japanese:             walkerJapaneseName,
		language.Korean:               walkerKoreanName,
		language.LatinAmericanSpanish: walkerLatinAmericanSpanishName,
		language.Russian:              walkerRussianName,
		language.SimplifiedChinese:    walkerSimplifiedChineseName,
		language.Spanish:              walkerSpanishName,
		language.TraditionalChinese:   walkerTraditionalChineseName}
)

var (
	walkerCharacter = nook.Character{
		Animal:   Dog,
		Birthday: walkerBirthday,
		Code:     walkerCode,
		Gender:   nook.Male,
		Name:     walkerName}
)

var (
	walkerAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "wuh"}

	walkerCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ouh"}

	walkerDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "waf"}

	walkerFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ouh"}

	walkerGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "wafff"}

	walkerItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "wuff"}

	walkerJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "バウ"}

	walkerKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "컹컹"}

	walkerLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "guuu-ah"}

	walkerRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ау-у"}

	walkerSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咆"}

	walkerSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "huesitos"}

	walkerTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咆"}
)

var (
	walkerPhrase = nook.Languages{
		language.AmericanEnglish:      walkerAmericanEnglishName,
		language.CanadianFrench:       walkerCanadianFrenchName,
		language.Dutch:                walkerDutchName,
		language.French:               walkerFrenchName,
		language.German:               walkerGermanName,
		language.Italian:              walkerItalianName,
		language.Japanese:             walkerJapaneseName,
		language.Korean:               walkerKoreanName,
		language.LatinAmericanSpanish: walkerLatinAmericanSpanishName,
		language.Russian:              walkerRussianName,
		language.SimplifiedChinese:    walkerSimplifiedChineseName,
		language.Spanish:              walkerSpanishName,
		language.TraditionalChinese:   walkerTraditionalChineseName}
)

var (
	Walker = nook.Villager{
		Character:   walkerCharacter,
		Personality: nook.Lazy,
		Phrase:      walkerPhrase}
)

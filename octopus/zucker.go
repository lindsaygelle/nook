package octopus

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	zuckerBirthday = nook.Birthday{
		Day:   8,
		Month: time.March}
)

var (
	zuckerCode = nook.Code{
		Value: "ocp02"}
)

var (
	zuckerAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Zucker"}

	zuckerCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Marvinsploush"}

	zuckerDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Zuckerbloep"}

	zuckerFrenchName = nook.Name{
		Language: language.French,
		Value:    "Marvinsploush"}

	zuckerGermanName = nook.Name{
		Language: language.German,
		Value:    "Ottokarsaug"}

	zuckerItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Poldoblub"}

	zuckerJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "タコヤせやねん"}

	zuckerKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "탁호약히"}

	zuckerLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Paulinochurubú"}

	zuckerRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Цукербуль"}

	zuckerSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "章丸丸认同"}

	zuckerSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Paulinochurubú"}

	zuckerTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "章丸丸認同"}
)

var (
	zuckerName = nook.Languages{
		language.AmericanEnglish:      zuckerAmericanEnglishName,
		language.CanadianFrench:       zuckerCanadianFrenchName,
		language.Dutch:                zuckerDutchName,
		language.French:               zuckerFrenchName,
		language.German:               zuckerGermanName,
		language.Italian:              zuckerItalianName,
		language.Japanese:             zuckerJapaneseName,
		language.Korean:               zuckerKoreanName,
		language.LatinAmericanSpanish: zuckerLatinAmericanSpanishName,
		language.Russian:              zuckerRussianName,
		language.SimplifiedChinese:    zuckerSimplifiedChineseName,
		language.Spanish:              zuckerSpanishName,
		language.TraditionalChinese:   zuckerTraditionalChineseName}
)

var (
	zuckerCharacter = nook.Character{
		Animal:   Octopus,
		Birthday: zuckerBirthday,
		Code:     zuckerCode,
		Gender:   nook.Male,
		Name:     zuckerName}
)

var (
	zuckerAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bloop"}

	zuckerCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "sploush"}

	zuckerDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "bloep"}

	zuckerFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "sploush"}

	zuckerGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "saug"}

	zuckerItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "blub"}

	zuckerJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "せやねん"}

	zuckerKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "약히"}

	zuckerLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "churubú"}

	zuckerRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "буль"}

	zuckerSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "认同"}

	zuckerSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "churubú"}

	zuckerTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "認同"}
)

var (
	zuckerPhrase = nook.Languages{
		language.AmericanEnglish:      zuckerAmericanEnglishName,
		language.CanadianFrench:       zuckerCanadianFrenchName,
		language.Dutch:                zuckerDutchName,
		language.French:               zuckerFrenchName,
		language.German:               zuckerGermanName,
		language.Italian:              zuckerItalianName,
		language.Japanese:             zuckerJapaneseName,
		language.Korean:               zuckerKoreanName,
		language.LatinAmericanSpanish: zuckerLatinAmericanSpanishName,
		language.Russian:              zuckerRussianName,
		language.SimplifiedChinese:    zuckerSimplifiedChineseName,
		language.Spanish:              zuckerSpanishName,
		language.TraditionalChinese:   zuckerTraditionalChineseName}
)

var (
	Zucker = nook.Villager{
		Character:   zuckerCharacter,
		Personality: nook.Lazy,
		Phrase:      zuckerPhrase}
)

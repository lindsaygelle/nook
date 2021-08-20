package octopus

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
		Value:    "Marvin"}

	zuckerDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Zucker"}

	zuckerFrenchName = nook.Name{
		Language: language.French,
		Value:    "Marvin"}

	zuckerGermanName = nook.Name{
		Language: language.German,
		Value:    "Ottokar"}

	zuckerItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Poldo"}

	zuckerJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "タコヤ"}

	zuckerKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "탁호"}

	zuckerLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Paulino"}

	zuckerRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Цукер"}

	zuckerSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "章丸丸"}

	zuckerSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Paulino"}

	zuckerTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "章丸丸"}
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
		Animal:   animal.Octopus,
		Birthday: zuckerBirthday,
		Code:     zuckerCode,
		Key:      character.Zucker,
		Gender:   gender.Male,
		Name:     zuckerName,
		Special:  false}
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
		language.AmericanEnglish:      zuckerAmericanEnglishPhrase,
		language.CanadianFrench:       zuckerCanadianFrenchPhrase,
		language.Dutch:                zuckerDutchPhrase,
		language.French:               zuckerFrenchPhrase,
		language.German:               zuckerGermanPhrase,
		language.Italian:              zuckerItalianPhrase,
		language.Japanese:             zuckerJapanesePhrase,
		language.Korean:               zuckerKoreanPhrase,
		language.LatinAmericanSpanish: zuckerLatinAmericanSpanishPhrase,
		language.Russian:              zuckerRussianPhrase,
		language.SimplifiedChinese:    zuckerSimplifiedChinesePhrase,
		language.Spanish:              zuckerSpanishPhrase,
		language.TraditionalChinese:   zuckerTraditionalChinesePhrase}
)

var (
	Zucker = nook.Villager{
		Character:   zuckerCharacter,
		Personality: personality.Lazy,
		Phrase:      zuckerPhrase}
)

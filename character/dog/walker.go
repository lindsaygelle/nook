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
		Value:    "George"}

	walkerDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Walker"}

	walkerFrenchName = nook.Name{
		Language: language.French,
		Value:    "George"}

	walkerGermanName = nook.Name{
		Language: language.German,
		Value:    "Fido"}

	walkerItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Walter"}

	walkerJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ベン"}

	walkerKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "벤"}

	walkerLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Miki"}

	walkerRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Уокер"}

	walkerSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿笨"}

	walkerSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Miki"}

	walkerTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿笨"}
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
		Animal:   animal.Dog,
		Birthday: walkerBirthday,
		Code:     walkerCode,
		Key:      character.Walker,
		Gender:   gender.Male,
		Name:     walkerName,
		Special:  false}
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
		language.AmericanEnglish:      walkerAmericanEnglishPhrase,
		language.CanadianFrench:       walkerCanadianFrenchPhrase,
		language.Dutch:                walkerDutchPhrase,
		language.French:               walkerFrenchPhrase,
		language.German:               walkerGermanPhrase,
		language.Italian:              walkerItalianPhrase,
		language.Japanese:             walkerJapanesePhrase,
		language.Korean:               walkerKoreanPhrase,
		language.LatinAmericanSpanish: walkerLatinAmericanSpanishPhrase,
		language.Russian:              walkerRussianPhrase,
		language.SimplifiedChinese:    walkerSimplifiedChinesePhrase,
		language.Spanish:              walkerSpanishPhrase,
		language.TraditionalChinese:   walkerTraditionalChinesePhrase}
)

var (
	Walker = nook.Villager{
		Character:   walkerCharacter,
		Personality: personality.Lazy,
		Phrase:      walkerPhrase}
)

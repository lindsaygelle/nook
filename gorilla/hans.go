package gorilla

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	hansBirthday = nook.Birthday{
		Day:   5,
		Month: time.December}
)

var (
	hansCode = nook.Code{
		Value: "gor10"}
)

var (
	hansAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Hans"}

	hansCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Lorantût tût"}

	hansDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Hansgroovy"}

	hansFrenchName = nook.Name{
		Language: language.French,
		Value:    "Lorantût tût"}

	hansGermanName = nook.Name{
		Language: language.German,
		Value:    "Hanskongkong"}

	hansItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Grigilioentonces"}

	hansJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "スナイルいえてる"}

	hansKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "스나일알만하다"}

	hansLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Hanscungagún"}

	hansRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Гансреспект"}

	hansSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "史奈鲁同意"}

	hansSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Hanscungagún"}

	hansTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "史奈魯同意"}
)

var (
	hansName = nook.Languages{
		language.AmericanEnglish:      hansAmericanEnglishName,
		language.CanadianFrench:       hansCanadianFrenchName,
		language.Dutch:                hansDutchName,
		language.French:               hansFrenchName,
		language.German:               hansGermanName,
		language.Italian:              hansItalianName,
		language.Japanese:             hansJapaneseName,
		language.Korean:               hansKoreanName,
		language.LatinAmericanSpanish: hansLatinAmericanSpanishName,
		language.Russian:              hansRussianName,
		language.SimplifiedChinese:    hansSimplifiedChineseName,
		language.Spanish:              hansSpanishName,
		language.TraditionalChinese:   hansTraditionalChineseName}
)

var (
	hansCharacter = nook.Character{
		Animal:   Gorilla,
		Birthday: hansBirthday,
		Code:     hansCode,
		Gender:   nook.Male,
		Name:     hansName}
)

var (
	hansAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "groovy"}

	hansCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tût tût"}

	hansDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "groovy"}

	hansFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tût tût"}

	hansGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kongkong"}

	hansItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "entonces"}

	hansJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "いえてる"}

	hansKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "알만하다"}

	hansLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cungagún"}

	hansRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "респект"}

	hansSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "同意"}

	hansSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cungagún"}

	hansTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "同意"}
)

var (
	hansPhrase = nook.Languages{
		language.AmericanEnglish:      hansAmericanEnglishName,
		language.CanadianFrench:       hansCanadianFrenchName,
		language.Dutch:                hansDutchName,
		language.French:               hansFrenchName,
		language.German:               hansGermanName,
		language.Italian:              hansItalianName,
		language.Japanese:             hansJapaneseName,
		language.Korean:               hansKoreanName,
		language.LatinAmericanSpanish: hansLatinAmericanSpanishName,
		language.Russian:              hansRussianName,
		language.SimplifiedChinese:    hansSimplifiedChineseName,
		language.Spanish:              hansSpanishName,
		language.TraditionalChinese:   hansTraditionalChineseName}
)

var (
	Hans = nook.Villager{
		Character:   hansCharacter,
		Personality: nook.Smug,
		Phrase:      hansPhrase}
)

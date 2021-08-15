package bearcub

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	oliveBirthday = nook.Birthday{
		Day:   12,
		Month: time.July}
)

var (
	oliveCode = nook.Code{
		Value: "cbr09"}
)

var (
	oliveAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Olive"}

	oliveCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Grisatrognon"}

	oliveDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Olivehoningkelkje"}

	oliveFrenchName = nook.Name{
		Language: language.French,
		Value:    "Grisatrognon"}

	oliveGermanName = nook.Name{
		Language: language.German,
		Value:    "Lindahonigtopf"}

	oliveItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Oliviafagiolino"}

	oliveJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ピッコロマグ"}

	oliveKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "올리브오잉"}

	oliveLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Osalinapanipof"}

	oliveRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Оливмилашка"}

	oliveSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "毕克萝马克杯"}

	oliveSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Osalinapanalito"}

	oliveTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "畢克蘿馬克杯"}
)

var (
	oliveName = nook.Languages{
		language.AmericanEnglish:      oliveAmericanEnglishName,
		language.CanadianFrench:       oliveCanadianFrenchName,
		language.Dutch:                oliveDutchName,
		language.French:               oliveFrenchName,
		language.German:               oliveGermanName,
		language.Italian:              oliveItalianName,
		language.Japanese:             oliveJapaneseName,
		language.Korean:               oliveKoreanName,
		language.LatinAmericanSpanish: oliveLatinAmericanSpanishName,
		language.Russian:              oliveRussianName,
		language.SimplifiedChinese:    oliveSimplifiedChineseName,
		language.Spanish:              oliveSpanishName,
		language.TraditionalChinese:   oliveTraditionalChineseName}
)

var (
	oliveCharacter = nook.Character{
		Animal:   Bearcub,
		Birthday: oliveBirthday,
		Code:     oliveCode,
		Gender:   nook.Female,
		Name:     oliveName}
)

var (
	oliveAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sweet pea"}

	oliveCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "trognon"}

	oliveDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "honingkelkje"}

	oliveFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "trognon"}

	oliveGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "honigtopf"}

	oliveItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "fagiolino"}

	oliveJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "マグ"}

	oliveKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "오잉"}

	oliveLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "panipof"}

	oliveRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "милашка"}

	oliveSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "马克杯"}

	oliveSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "panalito"}

	oliveTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "馬克杯"}
)

var (
	olivePhrase = nook.Languages{
		language.AmericanEnglish:      oliveAmericanEnglishName,
		language.CanadianFrench:       oliveCanadianFrenchName,
		language.Dutch:                oliveDutchName,
		language.French:               oliveFrenchName,
		language.German:               oliveGermanName,
		language.Italian:              oliveItalianName,
		language.Japanese:             oliveJapaneseName,
		language.Korean:               oliveKoreanName,
		language.LatinAmericanSpanish: oliveLatinAmericanSpanishName,
		language.Russian:              oliveRussianName,
		language.SimplifiedChinese:    oliveSimplifiedChineseName,
		language.Spanish:              oliveSpanishName,
		language.TraditionalChinese:   oliveTraditionalChineseName}
)

var (
	Olive = nook.Villager{
		Character:   oliveCharacter,
		Personality: nook.Normal,
		Phrase:      olivePhrase}
)

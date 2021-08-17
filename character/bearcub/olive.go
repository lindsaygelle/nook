package bearcub

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
		Value:    "Grisa"}

	oliveDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Olive"}

	oliveFrenchName = nook.Name{
		Language: language.French,
		Value:    "Grisa"}

	oliveGermanName = nook.Name{
		Language: language.German,
		Value:    "Linda"}

	oliveItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Olivia"}

	oliveJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ピッコロ"}

	oliveKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "올리브"}

	oliveLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Osalina"}

	oliveRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Олив"}

	oliveSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "毕克萝"}

	oliveSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Osalina"}

	oliveTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "畢克蘿"}
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
		Animal:   animal.Bearcub,
		Birthday: oliveBirthday,
		Code:     oliveCode,
		Key:      character.Olive,
		Gender:   gender.Female,
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
		language.AmericanEnglish:      oliveAmericanEnglishPhrase,
		language.CanadianFrench:       oliveCanadianFrenchPhrase,
		language.Dutch:                oliveDutchPhrase,
		language.French:               oliveFrenchPhrase,
		language.German:               oliveGermanPhrase,
		language.Italian:              oliveItalianPhrase,
		language.Japanese:             oliveJapanesePhrase,
		language.Korean:               oliveKoreanPhrase,
		language.LatinAmericanSpanish: oliveLatinAmericanSpanishPhrase,
		language.Russian:              oliveRussianPhrase,
		language.SimplifiedChinese:    oliveSimplifiedChinesePhrase,
		language.Spanish:              oliveSpanishPhrase,
		language.TraditionalChinese:   oliveTraditionalChinesePhrase}
)

var (
	Olive = nook.Villager{
		Character:   oliveCharacter,
		Personality: personality.Normal,
		Phrase:      olivePhrase}
)

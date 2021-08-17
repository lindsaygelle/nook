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
	pudgeBirthday = nook.Birthday{
		Day:   11,
		Month: time.June}
)

var (
	pudgeCode = nook.Code{
		Value: "cbr03"}
)

var (
	pudgeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pudge"}

	pudgeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gradub"}

	pudgeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pudge"}

	pudgeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gradub"}

	pudgeGermanName = nook.Name{
		Language: language.German,
		Value:    "Bertram"}

	pudgeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Tombolo"}

	pudgeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "きんぞう"}

	pudgeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "우띠"}

	pudgeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bollito"}

	pudgeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Падж"}

	pudgeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿钦"}

	pudgeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bollito"}

	pudgeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿欽"}
)

var (
	pudgeName = nook.Languages{
		language.AmericanEnglish:      pudgeAmericanEnglishName,
		language.CanadianFrench:       pudgeCanadianFrenchName,
		language.Dutch:                pudgeDutchName,
		language.French:               pudgeFrenchName,
		language.German:               pudgeGermanName,
		language.Italian:              pudgeItalianName,
		language.Japanese:             pudgeJapaneseName,
		language.Korean:               pudgeKoreanName,
		language.LatinAmericanSpanish: pudgeLatinAmericanSpanishName,
		language.Russian:              pudgeRussianName,
		language.SimplifiedChinese:    pudgeSimplifiedChineseName,
		language.Spanish:              pudgeSpanishName,
		language.TraditionalChinese:   pudgeTraditionalChineseName}
)

var (
	pudgeCharacter = nook.Character{
		Animal:   animal.Bearcub,
		Birthday: pudgeBirthday,
		Code:     pudgeCode,
		Key:      character.Pudge,
		Gender:   gender.Male,
		Name:     pudgeName}
)

var (
	pudgeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "pudgy"}

	pudgeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "louloute"}

	pudgeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "pudding"}

	pudgeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "louloute"}

	pudgeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "griesgram"}

	pudgeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "pulcino"}

	pudgeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "んもう"}

	pudgeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아이참"}

	pudgeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "bruah"}

	pudgeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "карапуз"}

	pudgeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "真是的"}

	pudgeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "rebollos"}

	pudgeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "真是的"}
)

var (
	pudgePhrase = nook.Languages{
		language.AmericanEnglish:      pudgeAmericanEnglishPhrase,
		language.CanadianFrench:       pudgeCanadianFrenchPhrase,
		language.Dutch:                pudgeDutchPhrase,
		language.French:               pudgeFrenchPhrase,
		language.German:               pudgeGermanPhrase,
		language.Italian:              pudgeItalianPhrase,
		language.Japanese:             pudgeJapanesePhrase,
		language.Korean:               pudgeKoreanPhrase,
		language.LatinAmericanSpanish: pudgeLatinAmericanSpanishPhrase,
		language.Russian:              pudgeRussianPhrase,
		language.SimplifiedChinese:    pudgeSimplifiedChinesePhrase,
		language.Spanish:              pudgeSpanishPhrase,
		language.TraditionalChinese:   pudgeTraditionalChinesePhrase}
)

var (
	Pudge = nook.Villager{
		Character:   pudgeCharacter,
		Personality: personality.Lazy,
		Phrase:      pudgePhrase}
)

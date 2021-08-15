package bearcub

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Gradublouloute"}

	pudgeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pudgepudding"}

	pudgeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gradublouloute"}

	pudgeGermanName = nook.Name{
		Language: language.German,
		Value:    "Bertramgriesgram"}

	pudgeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Tombolopulcino"}

	pudgeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "きんぞうんもう"}

	pudgeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "우띠아이참"}

	pudgeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bollitobruah"}

	pudgeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Паджкарапуз"}

	pudgeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿钦真是的"}

	pudgeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bollitorebollos"}

	pudgeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿欽真是的"}
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
		Animal:   Bearcub,
		Birthday: pudgeBirthday,
		Code:     pudgeCode,
		Gender:   nook.Male,
		Name:     pudgeName}
)

var (
	pudgeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "pudgygolly"}

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
	Pudge = nook.Villager{
		Character:   pudgeCharacter,
		Personality: nook.Lazy,
		Phrase:      pudgePhrase}
)

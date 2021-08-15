package elephant

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	bigtopBirthday = nook.Birthday{
		Day:   3,
		Month: time.October}
)

var (
	bigtopCode = nook.Code{
		Value: "elp02"}
)

var (
	bigtopAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Big Top"}

	bigtopCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Miles"}

	bigtopDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Big Top"}

	bigtopFrenchName = nook.Name{
		Language: language.French,
		Value:    "Miles"}

	bigtopGermanName = nook.Name{
		Language: language.German,
		Value:    "Benni"}

	bigtopItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Grande C"}

	bigtopJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "３ごう"}

	bigtopKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "3호"}

	bigtopLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Déivid"}

	bigtopRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Биг-Топ"}

	bigtopSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿三"}

	bigtopSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Déivid"}

	bigtopTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿三"}
)

var (
	bigtopName = nook.Languages{
		language.AmericanEnglish:      bigtopAmericanEnglishName,
		language.CanadianFrench:       bigtopCanadianFrenchName,
		language.Dutch:                bigtopDutchName,
		language.French:               bigtopFrenchName,
		language.German:               bigtopGermanName,
		language.Italian:              bigtopItalianName,
		language.Japanese:             bigtopJapaneseName,
		language.Korean:               bigtopKoreanName,
		language.LatinAmericanSpanish: bigtopLatinAmericanSpanishName,
		language.Russian:              bigtopRussianName,
		language.SimplifiedChinese:    bigtopSimplifiedChineseName,
		language.Spanish:              bigtopSpanishName,
		language.TraditionalChinese:   bigtopTraditionalChineseName}
)

var (
	bigtopCharacter = nook.Character{
		Animal:   Elephant,
		Birthday: bigtopBirthday,
		Code:     bigtopCode,
		Gender:   nook.Male,
		Name:     bigtopName}
)

var (
	bigtopAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "villain"}

	bigtopCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "vipère"}

	bigtopDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "boef"}

	bigtopFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "vipère"}

	bigtopGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "truuu"}

	bigtopItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "furfante"}

	bigtopJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "うりゃー"}

	bigtopKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "으랏차"}

	bigtopLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "duuud"}

	bigtopRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хитрюга"}

	bigtopSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哈啊"}

	bigtopSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "mutante"}

	bigtopTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哈啊"}
)

var (
	bigtopPhrase = nook.Languages{
		language.AmericanEnglish:      bigtopAmericanEnglishName,
		language.CanadianFrench:       bigtopCanadianFrenchName,
		language.Dutch:                bigtopDutchName,
		language.French:               bigtopFrenchName,
		language.German:               bigtopGermanName,
		language.Italian:              bigtopItalianName,
		language.Japanese:             bigtopJapaneseName,
		language.Korean:               bigtopKoreanName,
		language.LatinAmericanSpanish: bigtopLatinAmericanSpanishName,
		language.Russian:              bigtopRussianName,
		language.SimplifiedChinese:    bigtopSimplifiedChineseName,
		language.Spanish:              bigtopSpanishName,
		language.TraditionalChinese:   bigtopTraditionalChineseName}
)

var (
	BigTop = nook.Villager{
		Character:   bigtopCharacter,
		Personality: nook.Lazy,
		Phrase:      bigtopPhrase}
)

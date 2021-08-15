package eagle

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	apolloBirthday = nook.Birthday{
		Day:   4,
		Month: time.July}
)

var (
	apolloCode = nook.Code{
		Value: "pbr00"}
)

var (
	apolloAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Apollo"}

	apolloCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Apollontrraaaa"}

	apolloDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Apollopoeh"}

	apolloFrenchName = nook.Name{
		Language: language.French,
		Value:    "Apollontrraaaa"}

	apolloGermanName = nook.Name{
		Language: language.German,
		Value:    "Apolloahhh"}

	apolloItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Apollopah"}

	apolloJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アポロだワイ"}

	apolloKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아폴로다꿇어"}

	apolloLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Apolorapahhh"}

	apolloRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Аполлопф-ф"}

	apolloSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿波罗唔咿"}

	apolloSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Apolorapaz"}

	apolloTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿波羅唔咿"}
)

var (
	apolloName = nook.Languages{
		language.AmericanEnglish:      apolloAmericanEnglishName,
		language.CanadianFrench:       apolloCanadianFrenchName,
		language.Dutch:                apolloDutchName,
		language.French:               apolloFrenchName,
		language.German:               apolloGermanName,
		language.Italian:              apolloItalianName,
		language.Japanese:             apolloJapaneseName,
		language.Korean:               apolloKoreanName,
		language.LatinAmericanSpanish: apolloLatinAmericanSpanishName,
		language.Russian:              apolloRussianName,
		language.SimplifiedChinese:    apolloSimplifiedChineseName,
		language.Spanish:              apolloSpanishName,
		language.TraditionalChinese:   apolloTraditionalChineseName}
)

var (
	apolloCharacter = nook.Character{
		Animal:   Eagle,
		Birthday: apolloBirthday,
		Code:     apolloCode,
		Gender:   nook.Male,
		Name:     apolloName}
)

var (
	apolloAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "pah"}

	apolloCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "trraaaa"}

	apolloDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "poeh"}

	apolloFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "trraaaa"}

	apolloGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "ahhh"}

	apolloItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "pah"}

	apolloJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だワイ"}

	apolloKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "다꿇어"}

	apolloLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "rapahhh"}

	apolloRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "пф-ф"}

	apolloSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "唔咿"}

	apolloSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "rapaz"}

	apolloTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "唔咿"}
)

var (
	apolloPhrase = nook.Languages{
		language.AmericanEnglish:      apolloAmericanEnglishName,
		language.CanadianFrench:       apolloCanadianFrenchName,
		language.Dutch:                apolloDutchName,
		language.French:               apolloFrenchName,
		language.German:               apolloGermanName,
		language.Italian:              apolloItalianName,
		language.Japanese:             apolloJapaneseName,
		language.Korean:               apolloKoreanName,
		language.LatinAmericanSpanish: apolloLatinAmericanSpanishName,
		language.Russian:              apolloRussianName,
		language.SimplifiedChinese:    apolloSimplifiedChineseName,
		language.Spanish:              apolloSpanishName,
		language.TraditionalChinese:   apolloTraditionalChineseName}
)

var (
	Apollo = nook.Villager{
		Character:   apolloCharacter,
		Personality: nook.Cranky,
		Phrase:      apolloPhrase}
)

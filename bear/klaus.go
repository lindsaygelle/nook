package bear

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	klausBirthday = nook.Birthday{
		Day:   31,
		Month: time.March}
)

var (
	klausCode = nook.Code{
		Value: "bea14"}
)

var (
	klausAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Klaus"}

	klausCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Klausachtung"}

	klausDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Klausstrudel"}

	klausFrenchName = nook.Name{
		Language: language.French,
		Value:    "Klausachtung"}

	klausGermanName = nook.Name{
		Language: language.German,
		Value:    "Grischagruaaa"}

	klausItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Macistebailar"}

	klausJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クマロスオパー"}

	klausKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "곰도로스라저"}

	klausLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Gruñertoachurf"}

	klausRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Клаусштрудель"}

	klausSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "熊战士Over"}

	klausSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Gruñertoricitos"}

	klausTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "熊戰士Over"}
)

var (
	klausName = nook.Languages{
		language.AmericanEnglish:      klausAmericanEnglishName,
		language.CanadianFrench:       klausCanadianFrenchName,
		language.Dutch:                klausDutchName,
		language.French:               klausFrenchName,
		language.German:               klausGermanName,
		language.Italian:              klausItalianName,
		language.Japanese:             klausJapaneseName,
		language.Korean:               klausKoreanName,
		language.LatinAmericanSpanish: klausLatinAmericanSpanishName,
		language.Russian:              klausRussianName,
		language.SimplifiedChinese:    klausSimplifiedChineseName,
		language.Spanish:              klausSpanishName,
		language.TraditionalChinese:   klausTraditionalChineseName}
)

var (
	klausCharacter = nook.Character{
		Animal:   Bear,
		Birthday: klausBirthday,
		Code:     klausCode,
		Gender:   nook.Male,
		Name:     klausName}
)

var (
	klausAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "strudel"}

	klausCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "achtung"}

	klausDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "strudel"}

	klausFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "achtung"}

	klausGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "gruaaa"}

	klausItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bailar"}

	klausJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "オパー"}

	klausKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "라저"}

	klausLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "achurf"}

	klausRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "штрудель"}

	klausSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Over"}

	klausSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ricitos"}

	klausTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Over"}
)

var (
	klausPhrase = nook.Languages{
		language.AmericanEnglish:      klausAmericanEnglishName,
		language.CanadianFrench:       klausCanadianFrenchName,
		language.Dutch:                klausDutchName,
		language.French:               klausFrenchName,
		language.German:               klausGermanName,
		language.Italian:              klausItalianName,
		language.Japanese:             klausJapaneseName,
		language.Korean:               klausKoreanName,
		language.LatinAmericanSpanish: klausLatinAmericanSpanishName,
		language.Russian:              klausRussianName,
		language.SimplifiedChinese:    klausSimplifiedChineseName,
		language.Spanish:              klausSpanishName,
		language.TraditionalChinese:   klausTraditionalChineseName}
)

var (
	Klaus = nook.Villager{
		Character:   klausCharacter,
		Personality: nook.Smug,
		Phrase:      klausPhrase}
)

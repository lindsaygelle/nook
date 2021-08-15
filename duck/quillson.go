package duck

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	quillsonBirthday = nook.Birthday{
		Day:   22,
		Month: time.December}
)

var (
	quillsonCode = nook.Code{
		Value: "duk17"}
)

var (
	quillsonAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Quillson"}

	quillsonCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Narcissearrêêêête"}

	quillsonDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Quillsonsmient"}

	quillsonFrenchName = nook.Name{
		Language: language.French,
		Value:    "Narcissearrêêêête"}

	quillsonGermanName = nook.Name{
		Language: language.German,
		Value:    "Quentinpluster"}

	quillsonItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Verdoniocua cua"}

	quillsonJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "タックンマジかよ"}

	quillsonKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "덕근리얼리"}

	quillsonLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cuáltercuacoac"}

	quillsonRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Квилсоншу-утка"}

	quillsonSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "何童说真的"}

	quillsonSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cuálterverdoncho"}

	quillsonTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "何童說真的"}
)

var (
	quillsonName = nook.Languages{
		language.AmericanEnglish:      quillsonAmericanEnglishName,
		language.CanadianFrench:       quillsonCanadianFrenchName,
		language.Dutch:                quillsonDutchName,
		language.French:               quillsonFrenchName,
		language.German:               quillsonGermanName,
		language.Italian:              quillsonItalianName,
		language.Japanese:             quillsonJapaneseName,
		language.Korean:               quillsonKoreanName,
		language.LatinAmericanSpanish: quillsonLatinAmericanSpanishName,
		language.Russian:              quillsonRussianName,
		language.SimplifiedChinese:    quillsonSimplifiedChineseName,
		language.Spanish:              quillsonSpanishName,
		language.TraditionalChinese:   quillsonTraditionalChineseName}
)

var (
	quillsonCharacter = nook.Character{
		Animal:   Duck,
		Birthday: quillsonBirthday,
		Code:     quillsonCode,
		Gender:   nook.Male,
		Name:     quillsonName}
)

var (
	quillsonAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ridukulous"}

	quillsonCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "arrêêêête"}

	quillsonDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "smient"}

	quillsonFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "arrêêêête"}

	quillsonGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "pluster"}

	quillsonItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cua cua"}

	quillsonJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "マジかよ"}

	quillsonKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "리얼리"}

	quillsonLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cuacoac"}

	quillsonRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "шу-утка"}

	quillsonSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "说真的"}

	quillsonSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "verdoncho"}

	quillsonTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "說真的"}
)

var (
	quillsonPhrase = nook.Languages{
		language.AmericanEnglish:      quillsonAmericanEnglishName,
		language.CanadianFrench:       quillsonCanadianFrenchName,
		language.Dutch:                quillsonDutchName,
		language.French:               quillsonFrenchName,
		language.German:               quillsonGermanName,
		language.Italian:              quillsonItalianName,
		language.Japanese:             quillsonJapaneseName,
		language.Korean:               quillsonKoreanName,
		language.LatinAmericanSpanish: quillsonLatinAmericanSpanishName,
		language.Russian:              quillsonRussianName,
		language.SimplifiedChinese:    quillsonSimplifiedChineseName,
		language.Spanish:              quillsonSpanishName,
		language.TraditionalChinese:   quillsonTraditionalChineseName}
)

var (
	Quillson = nook.Villager{
		Character:   quillsonCharacter,
		Personality: nook.Smug,
		Phrase:      quillsonPhrase}
)

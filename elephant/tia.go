package elephant

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	tiaBirthday = nook.Birthday{
		Day:   18,
		Month: time.November}
)

var (
	tiaCode = nook.Code{
		Value: "elp10"}
)

var (
	tiaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tia"}

	tiaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Fannyénoooorme"}

	tiaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tiakopje thee"}

	tiaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Fannyénoooorme"}

	tiaGermanName = nook.Name{
		Language: language.German,
		Value:    "Eleonorehuiuiui"}

	tiaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Fannyparbleu"}

	tiaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ティーナホッ"}

	tiaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "티나따끈따끈"}

	tiaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Meraldasnuuuf"}

	tiaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Тиачайку-у"}

	tiaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "茉莉呼"}

	tiaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Meraldasnuuuf"}

	tiaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "茉莉呼"}
)

var (
	tiaName = nook.Languages{
		language.AmericanEnglish:      tiaAmericanEnglishName,
		language.CanadianFrench:       tiaCanadianFrenchName,
		language.Dutch:                tiaDutchName,
		language.French:               tiaFrenchName,
		language.German:               tiaGermanName,
		language.Italian:              tiaItalianName,
		language.Japanese:             tiaJapaneseName,
		language.Korean:               tiaKoreanName,
		language.LatinAmericanSpanish: tiaLatinAmericanSpanishName,
		language.Russian:              tiaRussianName,
		language.SimplifiedChinese:    tiaSimplifiedChineseName,
		language.Spanish:              tiaSpanishName,
		language.TraditionalChinese:   tiaTraditionalChineseName}
)

var (
	tiaCharacter = nook.Character{
		Animal:   Elephant,
		Birthday: tiaBirthday,
		Code:     tiaCode,
		Gender:   nook.Female,
		Name:     tiaName}
)

var (
	tiaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "teacup"}

	tiaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "énoooorme"}

	tiaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kopje thee"}

	tiaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "énoooorme"}

	tiaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "huiuiui"}

	tiaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "parbleu"}

	tiaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ホッ"}

	tiaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "따끈따끈"}

	tiaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "snuuuf"}

	tiaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "чайку-у"}

	tiaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "呼"}

	tiaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "snuuuf"}

	tiaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "呼"}
)

var (
	tiaPhrase = nook.Languages{
		language.AmericanEnglish:      tiaAmericanEnglishName,
		language.CanadianFrench:       tiaCanadianFrenchName,
		language.Dutch:                tiaDutchName,
		language.French:               tiaFrenchName,
		language.German:               tiaGermanName,
		language.Italian:              tiaItalianName,
		language.Japanese:             tiaJapaneseName,
		language.Korean:               tiaKoreanName,
		language.LatinAmericanSpanish: tiaLatinAmericanSpanishName,
		language.Russian:              tiaRussianName,
		language.SimplifiedChinese:    tiaSimplifiedChineseName,
		language.Spanish:              tiaSpanishName,
		language.TraditionalChinese:   tiaTraditionalChineseName}
)

var (
	Tia = nook.Villager{
		Character:   tiaCharacter,
		Personality: nook.Normal,
		Phrase:      tiaPhrase}
)

package elephant

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
		Value:    "Fanny"}

	tiaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tia"}

	tiaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Fanny"}

	tiaGermanName = nook.Name{
		Language: language.German,
		Value:    "Eleonore"}

	tiaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Fanny"}

	tiaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ティーナ"}

	tiaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "티나"}

	tiaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Meralda"}

	tiaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Тиа"}

	tiaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "茉莉"}

	tiaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Meralda"}

	tiaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "茉莉"}
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
		Animal:   animal.Elephant,
		Birthday: tiaBirthday,
		Code:     tiaCode,
		Key:      character.Tia,
		Gender:   gender.Female,
		Name:     tiaName,
		Special:  false}
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
		language.AmericanEnglish:      tiaAmericanEnglishPhrase,
		language.CanadianFrench:       tiaCanadianFrenchPhrase,
		language.Dutch:                tiaDutchPhrase,
		language.French:               tiaFrenchPhrase,
		language.German:               tiaGermanPhrase,
		language.Italian:              tiaItalianPhrase,
		language.Japanese:             tiaJapanesePhrase,
		language.Korean:               tiaKoreanPhrase,
		language.LatinAmericanSpanish: tiaLatinAmericanSpanishPhrase,
		language.Russian:              tiaRussianPhrase,
		language.SimplifiedChinese:    tiaSimplifiedChinesePhrase,
		language.Spanish:              tiaSpanishPhrase,
		language.TraditionalChinese:   tiaTraditionalChinesePhrase}
)

var (
	Tia = nook.Villager{
		Character:   tiaCharacter,
		Personality: personality.Normal,
		Phrase:      tiaPhrase}
)

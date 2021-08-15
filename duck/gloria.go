package duck

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	gloriaBirthday = nook.Birthday{
		Day:   12,
		Month: time.August}
)

var (
	gloriaCode = nook.Code{
		Value: "duk15"}
)

var (
	gloriaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Gloria"}

	gloriaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Déborahkwakos"}

	gloriaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Gloriakwebbel"}

	gloriaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Déborahkwakos"}

	gloriaGermanName = nook.Name{
		Language: language.German,
		Value:    "Gustavianerv"}

	gloriaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pappypaperonz"}

	gloriaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "スワンソンぎゃくに"}

	gloriaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마릴린유～후"}

	gloriaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Jarreapicodoro"}

	gloriaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Глорияне кряк ли"}

	gloriaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "素返真相反"}

	gloriaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Jarreapicodoro"}

	gloriaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "素返真相反"}
)

var (
	gloriaName = nook.Languages{
		language.AmericanEnglish:      gloriaAmericanEnglishName,
		language.CanadianFrench:       gloriaCanadianFrenchName,
		language.Dutch:                gloriaDutchName,
		language.French:               gloriaFrenchName,
		language.German:               gloriaGermanName,
		language.Italian:              gloriaItalianName,
		language.Japanese:             gloriaJapaneseName,
		language.Korean:               gloriaKoreanName,
		language.LatinAmericanSpanish: gloriaLatinAmericanSpanishName,
		language.Russian:              gloriaRussianName,
		language.SimplifiedChinese:    gloriaSimplifiedChineseName,
		language.Spanish:              gloriaSpanishName,
		language.TraditionalChinese:   gloriaTraditionalChineseName}
)

var (
	gloriaCharacter = nook.Character{
		Animal:   Duck,
		Birthday: gloriaBirthday,
		Code:     gloriaCode,
		Gender:   nook.Female,
		Name:     gloriaName}
)

var (
	gloriaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "quacker"}

	gloriaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "kwakos"}

	gloriaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kwebbel"}

	gloriaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "kwakos"}

	gloriaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "nerv"}

	gloriaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "paperonz"}

	gloriaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ぎゃくに"}

	gloriaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "유～후"}

	gloriaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "picodoro"}

	gloriaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "не кряк ли"}

	gloriaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "相反"}

	gloriaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "picodoro"}

	gloriaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "相反"}
)

var (
	gloriaPhrase = nook.Languages{
		language.AmericanEnglish:      gloriaAmericanEnglishName,
		language.CanadianFrench:       gloriaCanadianFrenchName,
		language.Dutch:                gloriaDutchName,
		language.French:               gloriaFrenchName,
		language.German:               gloriaGermanName,
		language.Italian:              gloriaItalianName,
		language.Japanese:             gloriaJapaneseName,
		language.Korean:               gloriaKoreanName,
		language.LatinAmericanSpanish: gloriaLatinAmericanSpanishName,
		language.Russian:              gloriaRussianName,
		language.SimplifiedChinese:    gloriaSimplifiedChineseName,
		language.Spanish:              gloriaSpanishName,
		language.TraditionalChinese:   gloriaTraditionalChineseName}
)

var (
	Gloria = nook.Villager{
		Character:   gloriaCharacter,
		Personality: nook.Snooty,
		Phrase:      gloriaPhrase}
)

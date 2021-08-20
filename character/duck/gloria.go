package duck

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
		Value:    "Déborah"}

	gloriaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Gloria"}

	gloriaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Déborah"}

	gloriaGermanName = nook.Name{
		Language: language.German,
		Value:    "Gustavia"}

	gloriaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pappy"}

	gloriaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "スワンソン"}

	gloriaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마릴린"}

	gloriaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Jarrea"}

	gloriaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Глория"}

	gloriaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "素返真"}

	gloriaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Jarrea"}

	gloriaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "素返真"}
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
		Animal:   animal.Duck,
		Birthday: gloriaBirthday,
		Code:     gloriaCode,
		Key:      character.Gloria,
		Gender:   gender.Female,
		Name:     gloriaName,
		Special:  false}
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
		language.AmericanEnglish:      gloriaAmericanEnglishPhrase,
		language.CanadianFrench:       gloriaCanadianFrenchPhrase,
		language.Dutch:                gloriaDutchPhrase,
		language.French:               gloriaFrenchPhrase,
		language.German:               gloriaGermanPhrase,
		language.Italian:              gloriaItalianPhrase,
		language.Japanese:             gloriaJapanesePhrase,
		language.Korean:               gloriaKoreanPhrase,
		language.LatinAmericanSpanish: gloriaLatinAmericanSpanishPhrase,
		language.Russian:              gloriaRussianPhrase,
		language.SimplifiedChinese:    gloriaSimplifiedChinesePhrase,
		language.Spanish:              gloriaSpanishPhrase,
		language.TraditionalChinese:   gloriaTraditionalChinesePhrase}
)

var (
	Gloria = nook.Villager{
		Character:   gloriaCharacter,
		Personality: personality.Snooty,
		Phrase:      gloriaPhrase}
)

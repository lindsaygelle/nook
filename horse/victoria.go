package horse

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	victoriaBirthday = nook.Birthday{
		Day:   11,
		Month: time.July}
)

var (
	victoriaCode = nook.Code{
		Value: "hrs01"}
)

var (
	victoriaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Victoria"}

	victoriaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Victoria"}

	victoriaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Victoria"}

	victoriaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Victoria"}

	victoriaGermanName = nook.Name{
		Language: language.German,
		Value:    "Emma"}

	victoriaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Vittoria"}

	victoriaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "セントアロー"}

	victoriaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "센트엘로"}

	victoriaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Victoria"}

	victoriaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Виктория"}

	victoriaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "圣萝"}

	victoriaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Victoria"}

	victoriaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "聖蘿"}
)

var (
	victoriaName = nook.Languages{
		language.AmericanEnglish:      victoriaAmericanEnglishName,
		language.CanadianFrench:       victoriaCanadianFrenchName,
		language.Dutch:                victoriaDutchName,
		language.French:               victoriaFrenchName,
		language.German:               victoriaGermanName,
		language.Italian:              victoriaItalianName,
		language.Japanese:             victoriaJapaneseName,
		language.Korean:               victoriaKoreanName,
		language.LatinAmericanSpanish: victoriaLatinAmericanSpanishName,
		language.Russian:              victoriaRussianName,
		language.SimplifiedChinese:    victoriaSimplifiedChineseName,
		language.Spanish:              victoriaSpanishName,
		language.TraditionalChinese:   victoriaTraditionalChineseName}
)

var (
	victoriaCharacter = nook.Character{
		Animal:   Horse,
		Birthday: victoriaBirthday,
		Code:     victoriaCode,
		Gender:   nook.Female,
		Name:     victoriaName}
)

var (
	victoriaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sugar cube"}

	victoriaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "susucre"}

	victoriaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "klontje"}

	victoriaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "susucre"}

	victoriaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "gimazuka"}

	victoriaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "zolletta"}

	victoriaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "いくわよ"}

	victoriaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "달리자"}

	victoriaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "clic-cloc"}

	victoriaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "сахарок"}

	victoriaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "走哦"}

	victoriaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "terroncito"}

	victoriaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "走囉"}
)

var (
	victoriaPhrase = nook.Languages{
		language.AmericanEnglish:      victoriaAmericanEnglishName,
		language.CanadianFrench:       victoriaCanadianFrenchName,
		language.Dutch:                victoriaDutchName,
		language.French:               victoriaFrenchName,
		language.German:               victoriaGermanName,
		language.Italian:              victoriaItalianName,
		language.Japanese:             victoriaJapaneseName,
		language.Korean:               victoriaKoreanName,
		language.LatinAmericanSpanish: victoriaLatinAmericanSpanishName,
		language.Russian:              victoriaRussianName,
		language.SimplifiedChinese:    victoriaSimplifiedChineseName,
		language.Spanish:              victoriaSpanishName,
		language.TraditionalChinese:   victoriaTraditionalChineseName}
)

var (
	Victoria = nook.Villager{
		Character:   victoriaCharacter,
		Personality: nook.Peppy,
		Phrase:      victoriaPhrase}
)

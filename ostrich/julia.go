package ostrich

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	juliaBirthday = nook.Birthday{
		Day:   31,
		Month: time.July}
)

var (
	juliaCode = nook.Code{
		Value: "ost05"}
)

var (
	juliaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Julia"}

	juliaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Julietrutruche"}

	juliaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Juliaschatje"}

	juliaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Julietrutruche"}

	juliaGermanName = nook.Name{
		Language: language.German,
		Value:    "Juliapüh"}

	juliaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Giuliablaaah"}

	juliaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジュリアやだわね"}

	juliaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "줄리아어머몰라"}

	juliaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Juliachurri"}

	juliaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Джулиядарлинг"}

	juliaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "朱莉亚吼唷"}

	juliaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Juliachurri"}

	juliaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "朱莉亞吼唷"}
)

var (
	juliaName = nook.Languages{
		language.AmericanEnglish:      juliaAmericanEnglishName,
		language.CanadianFrench:       juliaCanadianFrenchName,
		language.Dutch:                juliaDutchName,
		language.French:               juliaFrenchName,
		language.German:               juliaGermanName,
		language.Italian:              juliaItalianName,
		language.Japanese:             juliaJapaneseName,
		language.Korean:               juliaKoreanName,
		language.LatinAmericanSpanish: juliaLatinAmericanSpanishName,
		language.Russian:              juliaRussianName,
		language.SimplifiedChinese:    juliaSimplifiedChineseName,
		language.Spanish:              juliaSpanishName,
		language.TraditionalChinese:   juliaTraditionalChineseName}
)

var (
	juliaCharacter = nook.Character{
		Animal:   Ostrich,
		Birthday: juliaBirthday,
		Code:     juliaCode,
		Gender:   nook.Female,
		Name:     juliaName}
)

var (
	juliaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "dahling"}

	juliaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "trutruche"}

	juliaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "schatje"}

	juliaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "trutruche"}

	juliaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "püh"}

	juliaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "blaaah"}

	juliaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "やだわね"}

	juliaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "어머몰라"}

	juliaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "churri"}

	juliaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "дарлинг"}

	juliaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "吼唷"}

	juliaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "churri"}

	juliaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "吼唷"}
)

var (
	juliaPhrase = nook.Languages{
		language.AmericanEnglish:      juliaAmericanEnglishName,
		language.CanadianFrench:       juliaCanadianFrenchName,
		language.Dutch:                juliaDutchName,
		language.French:               juliaFrenchName,
		language.German:               juliaGermanName,
		language.Italian:              juliaItalianName,
		language.Japanese:             juliaJapaneseName,
		language.Korean:               juliaKoreanName,
		language.LatinAmericanSpanish: juliaLatinAmericanSpanishName,
		language.Russian:              juliaRussianName,
		language.SimplifiedChinese:    juliaSimplifiedChineseName,
		language.Spanish:              juliaSpanishName,
		language.TraditionalChinese:   juliaTraditionalChineseName}
)

var (
	Julia = nook.Villager{
		Character:   juliaCharacter,
		Personality: nook.Snooty,
		Phrase:      juliaPhrase}
)

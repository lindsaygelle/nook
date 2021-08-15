package penguin

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	frigaBirthday = nook.Birthday{
		Day:   16,
		Month: time.October}
)

var (
	frigaCode = nook.Code{
		Value: "pgn04"}
)

var (
	frigaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Friga"}

	frigaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Frigabourfff"}

	frigaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Frigabrrrhmpf"}

	frigaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Frigabourfff"}

	frigaGermanName = nook.Name{
		Language: language.German,
		Value:    "Friedagrrmpf"}

	frigaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Fridabrrrumf"}

	frigaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "サブリナツルルン"}

	frigaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "사브리나으쓱"}

	frigaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Fridaurfff"}

	frigaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Фригабрр-хм"}

	frigaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "谢宾娜滑溜溜"}

	frigaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Fridaheladito"}

	frigaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "謝賓娜滑溜溜"}
)

var (
	frigaName = nook.Languages{
		language.AmericanEnglish:      frigaAmericanEnglishName,
		language.CanadianFrench:       frigaCanadianFrenchName,
		language.Dutch:                frigaDutchName,
		language.French:               frigaFrenchName,
		language.German:               frigaGermanName,
		language.Italian:              frigaItalianName,
		language.Japanese:             frigaJapaneseName,
		language.Korean:               frigaKoreanName,
		language.LatinAmericanSpanish: frigaLatinAmericanSpanishName,
		language.Russian:              frigaRussianName,
		language.SimplifiedChinese:    frigaSimplifiedChineseName,
		language.Spanish:              frigaSpanishName,
		language.TraditionalChinese:   frigaTraditionalChineseName}
)

var (
	frigaCharacter = nook.Character{
		Animal:   Penguin,
		Birthday: frigaBirthday,
		Code:     frigaCode,
		Gender:   nook.Female,
		Name:     frigaName}
)

var (
	frigaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "brrrmph"}

	frigaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bourfff"}

	frigaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "brrrhmpf"}

	frigaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bourfff"}

	frigaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "grrmpf"}

	frigaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "brrrumf"}

	frigaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ツルルン"}

	frigaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "으쓱"}

	frigaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "urfff"}

	frigaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "брр-хм"}

	frigaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "滑溜溜"}

	frigaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "heladito"}

	frigaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "滑溜溜"}
)

var (
	frigaPhrase = nook.Languages{
		language.AmericanEnglish:      frigaAmericanEnglishName,
		language.CanadianFrench:       frigaCanadianFrenchName,
		language.Dutch:                frigaDutchName,
		language.French:               frigaFrenchName,
		language.German:               frigaGermanName,
		language.Italian:              frigaItalianName,
		language.Japanese:             frigaJapaneseName,
		language.Korean:               frigaKoreanName,
		language.LatinAmericanSpanish: frigaLatinAmericanSpanishName,
		language.Russian:              frigaRussianName,
		language.SimplifiedChinese:    frigaSimplifiedChineseName,
		language.Spanish:              frigaSpanishName,
		language.TraditionalChinese:   frigaTraditionalChineseName}
)

var (
	Friga = nook.Villager{
		Character:   frigaCharacter,
		Personality: nook.Snooty,
		Phrase:      frigaPhrase}
)

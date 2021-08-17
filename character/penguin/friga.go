package penguin

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
		Value:    "Friga"}

	frigaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Friga"}

	frigaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Friga"}

	frigaGermanName = nook.Name{
		Language: language.German,
		Value:    "Frieda"}

	frigaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Frida"}

	frigaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "サブリナ"}

	frigaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "사브리나"}

	frigaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Frida"}

	frigaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Фрига"}

	frigaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "谢宾娜"}

	frigaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Frida"}

	frigaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "謝賓娜"}
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
		Animal:   animal.Penguin,
		Birthday: frigaBirthday,
		Code:     frigaCode,
		Key:      character.Friga,
		Gender:   gender.Female,
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
		language.AmericanEnglish:      frigaAmericanEnglishPhrase,
		language.CanadianFrench:       frigaCanadianFrenchPhrase,
		language.Dutch:                frigaDutchPhrase,
		language.French:               frigaFrenchPhrase,
		language.German:               frigaGermanPhrase,
		language.Italian:              frigaItalianPhrase,
		language.Japanese:             frigaJapanesePhrase,
		language.Korean:               frigaKoreanPhrase,
		language.LatinAmericanSpanish: frigaLatinAmericanSpanishPhrase,
		language.Russian:              frigaRussianPhrase,
		language.SimplifiedChinese:    frigaSimplifiedChinesePhrase,
		language.Spanish:              frigaSpanishPhrase,
		language.TraditionalChinese:   frigaTraditionalChinesePhrase}
)

var (
	Friga = nook.Villager{
		Character:   frigaCharacter,
		Personality: personality.Snooty,
		Phrase:      frigaPhrase}
)

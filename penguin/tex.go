package penguin

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	texBirthday = nook.Birthday{
		Day:   6,
		Month: time.October}
)

var (
	texCode = nook.Code{
		Value: "pgn12"}
)

var (
	texAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tex"}

	texCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Émilien"}

	texDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tex"}

	texFrenchName = nook.Name{
		Language: language.French,
		Value:    "Émilien"}

	texGermanName = nook.Name{
		Language: language.German,
		Value:    "Matze"}

	texItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Freddy"}

	texJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ボルト"}

	texKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "볼트"}

	texLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Tomeo"}

	texRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Текс"}

	texSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "伏特"}

	texSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tomeo"}

	texTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "伏特"}
)

var (
	texName = nook.Languages{
		language.AmericanEnglish:      texAmericanEnglishName,
		language.CanadianFrench:       texCanadianFrenchName,
		language.Dutch:                texDutchName,
		language.French:               texFrenchName,
		language.German:               texGermanName,
		language.Italian:              texItalianName,
		language.Japanese:             texJapaneseName,
		language.Korean:               texKoreanName,
		language.LatinAmericanSpanish: texLatinAmericanSpanishName,
		language.Russian:              texRussianName,
		language.SimplifiedChinese:    texSimplifiedChineseName,
		language.Spanish:              texSpanishName,
		language.TraditionalChinese:   texTraditionalChineseName}
)

var (
	texCharacter = nook.Character{
		Animal:   Penguin,
		Birthday: texBirthday,
		Code:     texCode,
		Gender:   nook.Male,
		Name:     texName}
)

var (
	texAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "picante"}

	texCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "gla gla"}

	texDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "picante"}

	texFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "gla gla"}

	texGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "watschl"}

	texItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "frio"}

	texJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ベイベッ"}

	texKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "베이베"}

	texLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "siclaro"}

	texRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "огонек"}

	texSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "宝宝"}

	texSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "siclaro"}

	texTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "寶寶"}
)

var (
	texPhrase = nook.Languages{
		language.AmericanEnglish:      texAmericanEnglishName,
		language.CanadianFrench:       texCanadianFrenchName,
		language.Dutch:                texDutchName,
		language.French:               texFrenchName,
		language.German:               texGermanName,
		language.Italian:              texItalianName,
		language.Japanese:             texJapaneseName,
		language.Korean:               texKoreanName,
		language.LatinAmericanSpanish: texLatinAmericanSpanishName,
		language.Russian:              texRussianName,
		language.SimplifiedChinese:    texSimplifiedChineseName,
		language.Spanish:              texSpanishName,
		language.TraditionalChinese:   texTraditionalChineseName}
)

var (
	Tex = nook.Villager{
		Character:   texCharacter,
		Personality: nook.Smug,
		Phrase:      texPhrase}
)

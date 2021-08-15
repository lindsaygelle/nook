package squirrel

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	sylvanaBirthday = nook.Birthday{
		Day:   22,
		Month: time.October}
)

var (
	sylvanaCode = nook.Code{
		Value: "squ14"}
)

var (
	sylvanaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sylvana"}

	sylvanaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Mouniagrignote"}

	sylvanaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sylvanajammie"}

	sylvanaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mouniagrignote"}

	sylvanaGermanName = nook.Name{
		Language: language.German,
		Value:    "Marenjammi"}

	sylvanaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Silvanasgranocc"}

	sylvanaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "もんぺひゅん"}

	sylvanaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "실바나퓨우"}

	sylvanaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Silvanayupip"}

	sylvanaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Сильванахрум"}

	sylvanaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "孟珮咻"}

	sylvanaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Silvanacaldito"}

	sylvanaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "孟珮咻"}
)

var (
	sylvanaName = nook.Languages{
		language.AmericanEnglish:      sylvanaAmericanEnglishName,
		language.CanadianFrench:       sylvanaCanadianFrenchName,
		language.Dutch:                sylvanaDutchName,
		language.French:               sylvanaFrenchName,
		language.German:               sylvanaGermanName,
		language.Italian:              sylvanaItalianName,
		language.Japanese:             sylvanaJapaneseName,
		language.Korean:               sylvanaKoreanName,
		language.LatinAmericanSpanish: sylvanaLatinAmericanSpanishName,
		language.Russian:              sylvanaRussianName,
		language.SimplifiedChinese:    sylvanaSimplifiedChineseName,
		language.Spanish:              sylvanaSpanishName,
		language.TraditionalChinese:   sylvanaTraditionalChineseName}
)

var (
	sylvanaCharacter = nook.Character{
		Animal:   Squirrel,
		Birthday: sylvanaBirthday,
		Code:     sylvanaCode,
		Gender:   nook.Female,
		Name:     sylvanaName}
)

var (
	sylvanaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "hubbub"}

	sylvanaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "grignote"}

	sylvanaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "jammie"}

	sylvanaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "grignote"}

	sylvanaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "jammi"}

	sylvanaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sgranocc"}

	sylvanaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ひゅん"}

	sylvanaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "퓨우"}

	sylvanaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "yupip"}

	sylvanaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хрум"}

	sylvanaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咻"}

	sylvanaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "caldito"}

	sylvanaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咻"}
)

var (
	sylvanaPhrase = nook.Languages{
		language.AmericanEnglish:      sylvanaAmericanEnglishName,
		language.CanadianFrench:       sylvanaCanadianFrenchName,
		language.Dutch:                sylvanaDutchName,
		language.French:               sylvanaFrenchName,
		language.German:               sylvanaGermanName,
		language.Italian:              sylvanaItalianName,
		language.Japanese:             sylvanaJapaneseName,
		language.Korean:               sylvanaKoreanName,
		language.LatinAmericanSpanish: sylvanaLatinAmericanSpanishName,
		language.Russian:              sylvanaRussianName,
		language.SimplifiedChinese:    sylvanaSimplifiedChineseName,
		language.Spanish:              sylvanaSpanishName,
		language.TraditionalChinese:   sylvanaTraditionalChineseName}
)

var (
	Sylvana = nook.Villager{
		Character:   sylvanaCharacter,
		Personality: nook.Normal,
		Phrase:      sylvanaPhrase}
)

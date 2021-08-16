package squirrel

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
		Value:    "Mounia"}

	sylvanaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sylvana"}

	sylvanaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mounia"}

	sylvanaGermanName = nook.Name{
		Language: language.German,
		Value:    "Maren"}

	sylvanaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Silvana"}

	sylvanaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "もんぺ"}

	sylvanaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "실바나"}

	sylvanaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Silvana"}

	sylvanaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Сильвана"}

	sylvanaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "孟珮"}

	sylvanaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Silvana"}

	sylvanaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "孟珮"}
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
		Animal:   animal.Squirrel,
		Birthday: sylvanaBirthday,
		Code:     sylvanaCode,
		Key:      character.Sylvana,
		Gender:   gender.Female,
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
		Personality: personality.Normal,
		Phrase:      sylvanaPhrase}
)

package rabbit

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	rubyBirthday = nook.Birthday{
		Day:   25,
		Month: time.December}
)

var (
	rubyCode = nook.Code{
		Value: "rbt09"}
)

var (
	rubyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Ruby"}

	rubyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rubisn'oreille"}

	rubyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rubyhangoor"}

	rubyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rubisn'oreille"}

	rubyGermanName = nook.Name{
		Language: language.German,
		Value:    "Rubinapaffpaff"}

	rubyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rubinaorecchine"}

	rubyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ルナフツーに"}

	rubyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "루나보통이지"}

	rubyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rubíorejí"}

	rubyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рубиушки"}

	rubySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "月兔普通"}

	rubySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rubíorejitas"}

	rubyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "月兔普通"}
)

var (
	rubyName = nook.Languages{
		language.AmericanEnglish:      rubyAmericanEnglishName,
		language.CanadianFrench:       rubyCanadianFrenchName,
		language.Dutch:                rubyDutchName,
		language.French:               rubyFrenchName,
		language.German:               rubyGermanName,
		language.Italian:              rubyItalianName,
		language.Japanese:             rubyJapaneseName,
		language.Korean:               rubyKoreanName,
		language.LatinAmericanSpanish: rubyLatinAmericanSpanishName,
		language.Russian:              rubyRussianName,
		language.SimplifiedChinese:    rubySimplifiedChineseName,
		language.Spanish:              rubySpanishName,
		language.TraditionalChinese:   rubyTraditionalChineseName}
)

var (
	rubyCharacter = nook.Character{
		Animal:   Rabbit,
		Birthday: rubyBirthday,
		Code:     rubyCode,
		Gender:   nook.Female,
		Name:     rubyName}
)

var (
	rubyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "li'l ears"}

	rubyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "n'oreille"}

	rubyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hangoor"}

	rubyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "n'oreille"}

	rubyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "paffpaff"}

	rubyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "orecchine"}

	rubyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "フツーに"}

	rubyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "보통이지"}

	rubyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "orejí"}

	rubyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ушки"}

	rubySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "普通"}

	rubySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "orejitas"}

	rubyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "普通"}
)

var (
	rubyPhrase = nook.Languages{
		language.AmericanEnglish:      rubyAmericanEnglishName,
		language.CanadianFrench:       rubyCanadianFrenchName,
		language.Dutch:                rubyDutchName,
		language.French:               rubyFrenchName,
		language.German:               rubyGermanName,
		language.Italian:              rubyItalianName,
		language.Japanese:             rubyJapaneseName,
		language.Korean:               rubyKoreanName,
		language.LatinAmericanSpanish: rubyLatinAmericanSpanishName,
		language.Russian:              rubyRussianName,
		language.SimplifiedChinese:    rubySimplifiedChineseName,
		language.Spanish:              rubySpanishName,
		language.TraditionalChinese:   rubyTraditionalChineseName}
)

var (
	Ruby = nook.Villager{
		Character:   rubyCharacter,
		Personality: nook.Peppy,
		Phrase:      rubyPhrase}
)

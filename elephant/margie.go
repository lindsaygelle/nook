package elephant

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	margieBirthday = nook.Birthday{
		Day:   28,
		Month: time.January}
)

var (
	margieCode = nook.Code{
		Value: "elp04"}
)

var (
	margieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Margie"}

	margieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Maguytoutie"}

	margieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Margiepwèèpie"}

	margieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Maguytoutie"}

	margieGermanName = nook.Name{
		Language: language.German,
		Value:    "Elisatrampl"}

	margieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Mariannatootie"}

	margieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "サリーシャララ"}

	margieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "샐리샬랄라"}

	margieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ritafiuf-fiuf"}

	margieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Марджитру-у-ти"}

	margieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "莎莉莎啦啦"}

	margieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ritaflaqui"}

	margieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "莎莉莎啦啦"}
)

var (
	margieName = nook.Languages{
		language.AmericanEnglish:      margieAmericanEnglishName,
		language.CanadianFrench:       margieCanadianFrenchName,
		language.Dutch:                margieDutchName,
		language.French:               margieFrenchName,
		language.German:               margieGermanName,
		language.Italian:              margieItalianName,
		language.Japanese:             margieJapaneseName,
		language.Korean:               margieKoreanName,
		language.LatinAmericanSpanish: margieLatinAmericanSpanishName,
		language.Russian:              margieRussianName,
		language.SimplifiedChinese:    margieSimplifiedChineseName,
		language.Spanish:              margieSpanishName,
		language.TraditionalChinese:   margieTraditionalChineseName}
)

var (
	margieCharacter = nook.Character{
		Animal:   Elephant,
		Birthday: margieBirthday,
		Code:     margieCode,
		Gender:   nook.Female,
		Name:     margieName}
)

var (
	margieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "tootie"}

	margieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "toutie"}

	margieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "pwèèpie"}

	margieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "toutie"}

	margieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "trampl"}

	margieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "tootie"}

	margieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "シャララ"}

	margieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "샬랄라"}

	margieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "fiuf-fiuf"}

	margieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "тру-у-ти"}

	margieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "莎啦啦"}

	margieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "flaqui"}

	margieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "莎啦啦"}
)

var (
	margiePhrase = nook.Languages{
		language.AmericanEnglish:      margieAmericanEnglishName,
		language.CanadianFrench:       margieCanadianFrenchName,
		language.Dutch:                margieDutchName,
		language.French:               margieFrenchName,
		language.German:               margieGermanName,
		language.Italian:              margieItalianName,
		language.Japanese:             margieJapaneseName,
		language.Korean:               margieKoreanName,
		language.LatinAmericanSpanish: margieLatinAmericanSpanishName,
		language.Russian:              margieRussianName,
		language.SimplifiedChinese:    margieSimplifiedChineseName,
		language.Spanish:              margieSpanishName,
		language.TraditionalChinese:   margieTraditionalChineseName}
)

var (
	Margie = nook.Villager{
		Character:   margieCharacter,
		Personality: nook.Normal,
		Phrase:      margiePhrase}
)

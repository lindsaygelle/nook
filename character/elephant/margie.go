package elephant

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
		Value:    "Maguy"}

	margieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Margie"}

	margieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Maguy"}

	margieGermanName = nook.Name{
		Language: language.German,
		Value:    "Elisa"}

	margieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Marianna"}

	margieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "サリー"}

	margieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "샐리"}

	margieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rita"}

	margieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Марджи"}

	margieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "莎莉"}

	margieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rita"}

	margieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "莎莉"}
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
		Animal:   animal.Elephant,
		Birthday: margieBirthday,
		Code:     margieCode,
		Key:      character.Margie,
		Gender:   gender.Female,
		Name:     margieName,
		Special:  false}
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
		language.AmericanEnglish:      margieAmericanEnglishPhrase,
		language.CanadianFrench:       margieCanadianFrenchPhrase,
		language.Dutch:                margieDutchPhrase,
		language.French:               margieFrenchPhrase,
		language.German:               margieGermanPhrase,
		language.Italian:              margieItalianPhrase,
		language.Japanese:             margieJapanesePhrase,
		language.Korean:               margieKoreanPhrase,
		language.LatinAmericanSpanish: margieLatinAmericanSpanishPhrase,
		language.Russian:              margieRussianPhrase,
		language.SimplifiedChinese:    margieSimplifiedChinesePhrase,
		language.Spanish:              margieSpanishPhrase,
		language.TraditionalChinese:   margieTraditionalChinesePhrase}
)

var (
	Margie = nook.Villager{
		Character:   margieCharacter,
		Personality: personality.Normal,
		Phrase:      margiePhrase}
)

package bird

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	twiggyBirthday = nook.Birthday{
		Day:   13,
		Month: time.July}
)

var (
	twiggyCode = nook.Code{
		Value: "brd03"}
)

var (
	twiggyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Twiggy"}

	twiggyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Titicoolos"}

	twiggyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Twiggyfwieeet"}

	twiggyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Titicoolos"}

	twiggyGermanName = nook.Name{
		Language: language.German,
		Value:    "Twiggyzwirtschi"}

	twiggyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Titticiiip"}

	twiggyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ピーチクッピ"}

	twiggyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "핀틱크"}

	twiggyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Titítirití"}

	twiggyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Твиггитю-вить"}

	twiggySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "叽叽叽"}

	twiggySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Titítirití"}

	twiggyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘰嘰嘰"}
)

var (
	twiggyName = nook.Languages{
		language.AmericanEnglish:      twiggyAmericanEnglishName,
		language.CanadianFrench:       twiggyCanadianFrenchName,
		language.Dutch:                twiggyDutchName,
		language.French:               twiggyFrenchName,
		language.German:               twiggyGermanName,
		language.Italian:              twiggyItalianName,
		language.Japanese:             twiggyJapaneseName,
		language.Korean:               twiggyKoreanName,
		language.LatinAmericanSpanish: twiggyLatinAmericanSpanishName,
		language.Russian:              twiggyRussianName,
		language.SimplifiedChinese:    twiggySimplifiedChineseName,
		language.Spanish:              twiggySpanishName,
		language.TraditionalChinese:   twiggyTraditionalChineseName}
)

var (
	twiggyCharacter = nook.Character{
		Animal:   Bird,
		Birthday: twiggyBirthday,
		Code:     twiggyCode,
		Gender:   nook.Female,
		Name:     twiggyName}
)

var (
	twiggyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cheepers"}

	twiggyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "coolos"}

	twiggyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "fwieeet"}

	twiggyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "coolos"}

	twiggyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "zwirtschi"}

	twiggyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ciiip"}

	twiggyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ッピ"}

	twiggyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "크"}

	twiggyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "tirití"}

	twiggyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "тю-вить"}

	twiggySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "叽"}

	twiggySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tirití"}

	twiggyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘰"}
)

var (
	twiggyPhrase = nook.Languages{
		language.AmericanEnglish:      twiggyAmericanEnglishName,
		language.CanadianFrench:       twiggyCanadianFrenchName,
		language.Dutch:                twiggyDutchName,
		language.French:               twiggyFrenchName,
		language.German:               twiggyGermanName,
		language.Italian:              twiggyItalianName,
		language.Japanese:             twiggyJapaneseName,
		language.Korean:               twiggyKoreanName,
		language.LatinAmericanSpanish: twiggyLatinAmericanSpanishName,
		language.Russian:              twiggyRussianName,
		language.SimplifiedChinese:    twiggySimplifiedChineseName,
		language.Spanish:              twiggySpanishName,
		language.TraditionalChinese:   twiggyTraditionalChineseName}
)

var (
	Twiggy = nook.Villager{
		Character:   twiggyCharacter,
		Personality: nook.Peppy,
		Phrase:      twiggyPhrase}
)

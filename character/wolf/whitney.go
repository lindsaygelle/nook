package wolf

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
	whitneyBirthday = nook.Birthday{
		Day:   17,
		Month: time.September}
)

var (
	whitneyCode = nook.Code{
		Value: "wol03"}
)

var (
	whitneyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Whitney"}

	whitneyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Blanche"}

	whitneyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Whitney"}

	whitneyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Blanche"}

	whitneyGermanName = nook.Name{
		Language: language.German,
		Value:    "Lupa"}

	whitneyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Bianca"}

	whitneyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ビアンカ"}

	whitneyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "비앙카"}

	whitneyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lupe"}

	whitneyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Уитни"}

	whitneySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "毕安卡"}

	whitneySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lupe"}

	whitneyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "畢安卡"}
)

var (
	whitneyName = nook.Languages{
		language.AmericanEnglish:      whitneyAmericanEnglishName,
		language.CanadianFrench:       whitneyCanadianFrenchName,
		language.Dutch:                whitneyDutchName,
		language.French:               whitneyFrenchName,
		language.German:               whitneyGermanName,
		language.Italian:              whitneyItalianName,
		language.Japanese:             whitneyJapaneseName,
		language.Korean:               whitneyKoreanName,
		language.LatinAmericanSpanish: whitneyLatinAmericanSpanishName,
		language.Russian:              whitneyRussianName,
		language.SimplifiedChinese:    whitneySimplifiedChineseName,
		language.Spanish:              whitneySpanishName,
		language.TraditionalChinese:   whitneyTraditionalChineseName}
)

var (
	whitneyCharacter = nook.Character{
		Animal:   animal.Wolf,
		Birthday: whitneyBirthday,
		Code:     whitneyCode,
		Key:      character.Whitney,
		Gender:   gender.Female,
		Name:     whitneyName}
)

var (
	whitneyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "snappy"}

	whitneyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bing bang"}

	whitneyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hap"}

	whitneyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "sbam"}

	whitneyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "jaulll"}

	whitneyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "snappy"}

	whitneyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ステキね"}

	whitneyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "멋져"}

	whitneyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "auf-auf"}

	whitneyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "цап"}

	whitneySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "漂亮"}

	whitneySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "auf-auf"}

	whitneyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "漂亮"}
)

var (
	whitneyPhrase = nook.Languages{
		language.AmericanEnglish:      whitneyAmericanEnglishName,
		language.CanadianFrench:       whitneyCanadianFrenchName,
		language.Dutch:                whitneyDutchName,
		language.French:               whitneyFrenchName,
		language.German:               whitneyGermanName,
		language.Italian:              whitneyItalianName,
		language.Japanese:             whitneyJapaneseName,
		language.Korean:               whitneyKoreanName,
		language.LatinAmericanSpanish: whitneyLatinAmericanSpanishName,
		language.Russian:              whitneyRussianName,
		language.SimplifiedChinese:    whitneySimplifiedChineseName,
		language.Spanish:              whitneySpanishName,
		language.TraditionalChinese:   whitneyTraditionalChineseName}
)

var (
	Whitney = nook.Villager{
		Character:   whitneyCharacter,
		Personality: personality.Snooty,
		Phrase:      whitneyPhrase}
)

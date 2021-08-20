package bird

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
	shoukichiBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	shoukichiCode = nook.Code{
		Value: ""}
)

var (
	shoukichiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Shoukichi"}

	shoukichiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	shoukichiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	shoukichiFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	shoukichiGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	shoukichiItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	shoukichiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "しょうきち"}

	shoukichiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	shoukichiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	shoukichiRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	shoukichiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	shoukichiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	shoukichiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	shoukichiName = nook.Languages{
		language.AmericanEnglish:      shoukichiAmericanEnglishName,
		language.CanadianFrench:       shoukichiCanadianFrenchName,
		language.Dutch:                shoukichiDutchName,
		language.French:               shoukichiFrenchName,
		language.German:               shoukichiGermanName,
		language.Italian:              shoukichiItalianName,
		language.Japanese:             shoukichiJapaneseName,
		language.Korean:               shoukichiKoreanName,
		language.LatinAmericanSpanish: shoukichiLatinAmericanSpanishName,
		language.Russian:              shoukichiRussianName,
		language.SimplifiedChinese:    shoukichiSimplifiedChineseName,
		language.Spanish:              shoukichiSpanishName,
		language.TraditionalChinese:   shoukichiTraditionalChineseName}
)

var (
	shoukichiCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: shoukichiBirthday,
		Code:     shoukichiCode,
		Key:      character.Shoukichi,
		Gender:   gender.Male,
		Name:     shoukichiName,
		Special:  false}
)

var (
	shoukichiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ズン"}

	shoukichiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	shoukichiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	shoukichiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	shoukichiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	shoukichiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	shoukichiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ズン"}

	shoukichiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	shoukichiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	shoukichiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	shoukichiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	shoukichiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	shoukichiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	shoukichiPhrase = nook.Languages{
		language.AmericanEnglish:      shoukichiAmericanEnglishPhrase,
		language.CanadianFrench:       shoukichiCanadianFrenchPhrase,
		language.Dutch:                shoukichiDutchPhrase,
		language.French:               shoukichiFrenchPhrase,
		language.German:               shoukichiGermanPhrase,
		language.Italian:              shoukichiItalianPhrase,
		language.Japanese:             shoukichiJapanesePhrase,
		language.Korean:               shoukichiKoreanPhrase,
		language.LatinAmericanSpanish: shoukichiLatinAmericanSpanishPhrase,
		language.Russian:              shoukichiRussianPhrase,
		language.SimplifiedChinese:    shoukichiSimplifiedChinesePhrase,
		language.Spanish:              shoukichiSpanishPhrase,
		language.TraditionalChinese:   shoukichiTraditionalChinesePhrase}
)

var (
	Shoukichi = nook.Villager{
		Character:   shoukichiCharacter,
		Personality: personality.Jock,
		Phrase:      shoukichiPhrase}
)

package bird

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "N/A"}

	shoukichiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	shoukichiFrenchName = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	shoukichiGermanName = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	shoukichiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	shoukichiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "しょうきち"}

	shoukichiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	shoukichiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	shoukichiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	shoukichiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	shoukichiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	shoukichiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
		Animal:   Bird,
		Birthday: shoukichiBirthday,
		Code:     shoukichiCode,
		Gender:   nook.Male,
		Name:     shoukichiName}
)

var (
	shoukichiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	shoukichiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	shoukichiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	shoukichiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	shoukichiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	shoukichiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	shoukichiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ズン"}

	shoukichiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	shoukichiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	shoukichiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	shoukichiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	shoukichiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	shoukichiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	shoukichiPhrase = nook.Languages{
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
	Shoukichi = nook.Villager{
		Character:   shoukichiCharacter,
		Personality: nook.Jock,
		Phrase:      shoukichiPhrase}
)

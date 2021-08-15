package chicken

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	pluckyBirthday = nook.Birthday{
		Day:   12,
		Month: time.October}
)

var (
	pluckyCode = nook.Code{
		Value: "chn10"}
)

var (
	pluckyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Plucky"}

	pluckyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Poulettegaline"}

	pluckyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pluckykippetje"}

	pluckyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Poulettegaline"}

	pluckyGermanName = nook.Name{
		Language: language.German,
		Value:    "Mareilehühnermist"}

	pluckyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Coccazampoli"}

	pluckyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "パタヤどうだい"}

	pluckyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "파타야그게어때"}

	pluckyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Herminiacorocó"}

	pluckyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Плакищип-щип"}

	pluckySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "帕塔雅泰泰"}

	pluckySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Herminiatesorito"}

	pluckyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "帕塔雅泰泰"}
)

var (
	pluckyName = nook.Languages{
		language.AmericanEnglish:      pluckyAmericanEnglishName,
		language.CanadianFrench:       pluckyCanadianFrenchName,
		language.Dutch:                pluckyDutchName,
		language.French:               pluckyFrenchName,
		language.German:               pluckyGermanName,
		language.Italian:              pluckyItalianName,
		language.Japanese:             pluckyJapaneseName,
		language.Korean:               pluckyKoreanName,
		language.LatinAmericanSpanish: pluckyLatinAmericanSpanishName,
		language.Russian:              pluckyRussianName,
		language.SimplifiedChinese:    pluckySimplifiedChineseName,
		language.Spanish:              pluckySpanishName,
		language.TraditionalChinese:   pluckyTraditionalChineseName}
)

var (
	pluckyCharacter = nook.Character{
		Animal:   Chicken,
		Birthday: pluckyBirthday,
		Code:     pluckyCode,
		Gender:   nook.Female,
		Name:     pluckyName}
)

var (
	pluckyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "chicky-poo"}

	pluckyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "galine"}

	pluckyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kippetje"}

	pluckyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "galine"}

	pluckyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hühnermist"}

	pluckyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "zampoli"}

	pluckyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "どうだい"}

	pluckyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그게어때"}

	pluckyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "corocó"}

	pluckyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "щип-щип"}

	pluckySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "泰泰"}

	pluckySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tesorito"}

	pluckyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "泰泰"}
)

var (
	pluckyPhrase = nook.Languages{
		language.AmericanEnglish:      pluckyAmericanEnglishName,
		language.CanadianFrench:       pluckyCanadianFrenchName,
		language.Dutch:                pluckyDutchName,
		language.French:               pluckyFrenchName,
		language.German:               pluckyGermanName,
		language.Italian:              pluckyItalianName,
		language.Japanese:             pluckyJapaneseName,
		language.Korean:               pluckyKoreanName,
		language.LatinAmericanSpanish: pluckyLatinAmericanSpanishName,
		language.Russian:              pluckyRussianName,
		language.SimplifiedChinese:    pluckySimplifiedChineseName,
		language.Spanish:              pluckySpanishName,
		language.TraditionalChinese:   pluckyTraditionalChineseName}
)

var (
	Plucky = nook.Villager{
		Character:   pluckyCharacter,
		Personality: nook.BigSister,
		Phrase:      pluckyPhrase}
)

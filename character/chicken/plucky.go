package chicken

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
		Value:    "Poulette"}

	pluckyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Plucky"}

	pluckyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Poulette"}

	pluckyGermanName = nook.Name{
		Language: language.German,
		Value:    "Mareile"}

	pluckyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cocca"}

	pluckyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "パタヤ"}

	pluckyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "파타야"}

	pluckyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Herminia"}

	pluckyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Плаки"}

	pluckySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "帕塔雅"}

	pluckySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Herminia"}

	pluckyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "帕塔雅"}
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
		Animal:   animal.Chicken,
		Birthday: pluckyBirthday,
		Code:     pluckyCode,
		Key:      character.Plucky,
		Gender:   gender.Female,
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
		language.AmericanEnglish:      pluckyAmericanEnglishPhrase,
		language.CanadianFrench:       pluckyCanadianFrenchPhrase,
		language.Dutch:                pluckyDutchPhrase,
		language.French:               pluckyFrenchPhrase,
		language.German:               pluckyGermanPhrase,
		language.Italian:              pluckyItalianPhrase,
		language.Japanese:             pluckyJapanesePhrase,
		language.Korean:               pluckyKoreanPhrase,
		language.LatinAmericanSpanish: pluckyLatinAmericanSpanishPhrase,
		language.Russian:              pluckyRussianPhrase,
		language.SimplifiedChinese:    pluckySimplifiedChinesePhrase,
		language.Spanish:              pluckySpanishPhrase,
		language.TraditionalChinese:   pluckyTraditionalChinesePhrase}
)

var (
	Plucky = nook.Villager{
		Character:   pluckyCharacter,
		Personality: personality.BigSister,
		Phrase:      pluckyPhrase}
)

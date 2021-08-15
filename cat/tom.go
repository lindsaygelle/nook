package cat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	tomBirthday = nook.Birthday{
		Day:   10,
		Month: time.December}
)

var (
	tomCode = nook.Code{
		Value: "cat15"}
)

var (
	tomAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tom"}

	tomCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Tom"}

	tomDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tom"}

	tomFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tom"}

	tomGermanName = nook.Name{
		Language: language.German,
		Value:    "Timo"}

	tomItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Isidoro"}

	tomJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "バンタム"}

	tomKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "밴덤"}

	tomLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Zapirón"}

	tomRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Том"}

	tomSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿邦"}

	tomSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Zapirón"}

	tomTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿邦"}
)

var (
	tomName = nook.Languages{
		language.AmericanEnglish:      tomAmericanEnglishName,
		language.CanadianFrench:       tomCanadianFrenchName,
		language.Dutch:                tomDutchName,
		language.French:               tomFrenchName,
		language.German:               tomGermanName,
		language.Italian:              tomItalianName,
		language.Japanese:             tomJapaneseName,
		language.Korean:               tomKoreanName,
		language.LatinAmericanSpanish: tomLatinAmericanSpanishName,
		language.Russian:              tomRussianName,
		language.SimplifiedChinese:    tomSimplifiedChineseName,
		language.Spanish:              tomSpanishName,
		language.TraditionalChinese:   tomTraditionalChineseName}
)

var (
	tomCharacter = nook.Character{
		Animal:   Cat,
		Birthday: tomBirthday,
		Code:     tomCode,
		Gender:   nook.Male,
		Name:     tomName}
)

var (
	tomAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "me-YOWZA"}

	tomCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mistigri"}

	tomDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "miauw moe"}

	tomFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mistigri"}

	tomGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "fauch"}

	tomItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "mi-IAO"}

	tomJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ナックル"}

	tomKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "쳇"}

	tomLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "miooou"}

	tomRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "мя-я-яоза"}

	tomSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "切"}

	tomSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "miooou"}

	tomTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "關節"}
)

var (
	tomPhrase = nook.Languages{
		language.AmericanEnglish:      tomAmericanEnglishName,
		language.CanadianFrench:       tomCanadianFrenchName,
		language.Dutch:                tomDutchName,
		language.French:               tomFrenchName,
		language.German:               tomGermanName,
		language.Italian:              tomItalianName,
		language.Japanese:             tomJapaneseName,
		language.Korean:               tomKoreanName,
		language.LatinAmericanSpanish: tomLatinAmericanSpanishName,
		language.Russian:              tomRussianName,
		language.SimplifiedChinese:    tomSimplifiedChineseName,
		language.Spanish:              tomSpanishName,
		language.TraditionalChinese:   tomTraditionalChineseName}
)

var (
	Tom = nook.Villager{
		Character:   tomCharacter,
		Personality: nook.Cranky,
		Phrase:      tomPhrase}
)

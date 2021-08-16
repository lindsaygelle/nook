package frog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	tadBirthday = nook.Birthday{
		Day:   3,
		Month: time.August}
)

var (
	tadCode = nook.Code{
		Value: "flg09"}
)

var (
	tadAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tad"}

	tadCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rénato"}

	tadDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tad"}

	tadFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rénato"}

	tadGermanName = nook.Name{
		Language: language.German,
		Value:    "Paul"}

	tadItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Achille"}

	tadJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "タンボ"}

	tadKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "텀보"}

	tadLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Saltonio"}

	tadRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Тэд"}

	tadSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿田"}

	tadSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Saltonio"}

	tadTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿田"}
)

var (
	tadName = nook.Languages{
		language.AmericanEnglish:      tadAmericanEnglishName,
		language.CanadianFrench:       tadCanadianFrenchName,
		language.Dutch:                tadDutchName,
		language.French:               tadFrenchName,
		language.German:               tadGermanName,
		language.Italian:              tadItalianName,
		language.Japanese:             tadJapaneseName,
		language.Korean:               tadKoreanName,
		language.LatinAmericanSpanish: tadLatinAmericanSpanishName,
		language.Russian:              tadRussianName,
		language.SimplifiedChinese:    tadSimplifiedChineseName,
		language.Spanish:              tadSpanishName,
		language.TraditionalChinese:   tadTraditionalChineseName}
)

var (
	tadCharacter = nook.Character{
		Animal:   animal.Frog,
		Birthday: tadBirthday,
		Code:     tadCode,
		Gender:   gender.Male,
		Name:     tadName}
)

var (
	tadAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sluuuurp"}

	tadCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "nénuf"}

	tadDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "sluuuurp"}

	tadFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "nénuf"}

	tadGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schlürpf"}

	tadItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sluuuurp"}

	tadJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だよん"}

	tadKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "흐압"}

	tadLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "zump"}

	tadRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "жабррр"}

	tadSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "蝌蚪"}

	tadSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "zump"}

	tadTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "蝌蚪"}
)

var (
	tadPhrase = nook.Languages{
		language.AmericanEnglish:      tadAmericanEnglishName,
		language.CanadianFrench:       tadCanadianFrenchName,
		language.Dutch:                tadDutchName,
		language.French:               tadFrenchName,
		language.German:               tadGermanName,
		language.Italian:              tadItalianName,
		language.Japanese:             tadJapaneseName,
		language.Korean:               tadKoreanName,
		language.LatinAmericanSpanish: tadLatinAmericanSpanishName,
		language.Russian:              tadRussianName,
		language.SimplifiedChinese:    tadSimplifiedChineseName,
		language.Spanish:              tadSpanishName,
		language.TraditionalChinese:   tadTraditionalChineseName}
)

var (
	Tad = nook.Villager{
		Character:   tadCharacter,
		Personality: personality.Jock,
		Phrase:      tadPhrase}
)

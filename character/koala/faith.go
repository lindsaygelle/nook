package koala

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
	faithBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	faithCode = nook.Code{
		Value: ""}
)

var (
	faithAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Faith"}

	faithCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	faithDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	faithFrenchName = nook.Name{
		Language: language.French,
		Value:    "Kolette"}

	faithGermanName = nook.Name{
		Language: language.German,
		Value:    "Finchen"}

	faithItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Calipso"}

	faithJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マーチ"}

	faithKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	faithLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	faithRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	faithSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	faithSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Felicia"}

	faithTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	faithName = nook.Languages{
		language.AmericanEnglish:      faithAmericanEnglishName,
		language.CanadianFrench:       faithCanadianFrenchName,
		language.Dutch:                faithDutchName,
		language.French:               faithFrenchName,
		language.German:               faithGermanName,
		language.Italian:              faithItalianName,
		language.Japanese:             faithJapaneseName,
		language.Korean:               faithKoreanName,
		language.LatinAmericanSpanish: faithLatinAmericanSpanishName,
		language.Russian:              faithRussianName,
		language.SimplifiedChinese:    faithSimplifiedChineseName,
		language.Spanish:              faithSpanishName,
		language.TraditionalChinese:   faithTraditionalChineseName}
)

var (
	faithCharacter = nook.Character{
		Animal:   animal.Koala,
		Birthday: faithBirthday,
		Code:     faithCode,
		Key:      character.Faith,
		Gender:   gender.Female,
		Name:     faithName}
)

var (
	faithAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "aloha"}

	faithCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	faithDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	faithFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "aloha"}

	faithGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "aloha"}

	faithItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "aloha"}

	faithJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "サクッ"}

	faithKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	faithLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	faithRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	faithSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	faithSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "aloha"}

	faithTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	faithPhrase = nook.Languages{
		language.AmericanEnglish:      faithAmericanEnglishName,
		language.CanadianFrench:       faithCanadianFrenchName,
		language.Dutch:                faithDutchName,
		language.French:               faithFrenchName,
		language.German:               faithGermanName,
		language.Italian:              faithItalianName,
		language.Japanese:             faithJapaneseName,
		language.Korean:               faithKoreanName,
		language.LatinAmericanSpanish: faithLatinAmericanSpanishName,
		language.Russian:              faithRussianName,
		language.SimplifiedChinese:    faithSimplifiedChineseName,
		language.Spanish:              faithSpanishName,
		language.TraditionalChinese:   faithTraditionalChineseName}
)

var (
	Faith = nook.Villager{
		Character:   faithCharacter,
		Personality: personality.Normal,
		Phrase:      faithPhrase}
)

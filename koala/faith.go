package koala

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "N/A"}

	faithDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

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
		Value:    "N/A"}

	faithLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	faithRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	faithSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	faithSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Felicia"}

	faithTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
		Animal:   Koala,
		Birthday: faithBirthday,
		Code:     faithCode,
		Gender:   nook.Female,
		Name:     faithName}
)

var (
	faithAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	faithCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	faithDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

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
		Value:    "N/A"}

	faithLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	faithRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	faithSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	faithSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "aloha"}

	faithTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
		Personality: nook.Normal,
		Phrase:      faithPhrase}
)

package gorilla

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	cesarBirthday = nook.Birthday{
		Day:   6,
		Month: time.September}
)

var (
	cesarCode = nook.Code{
		Value: "gor00"}
)

var (
	cesarAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cesar"}

	cesarCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Césarbrutus"}

	cesarDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cesarhoogheid"}

	cesarFrenchName = nook.Name{
		Language: language.French,
		Value:    "Césarbrutus"}

	cesarGermanName = nook.Name{
		Language: language.German,
		Value:    "Alfredohoheit"}

	cesarItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cesaredado"}

	cesarJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アランウホウホ"}

	cesarKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "앨런우갸우갸"}

	cesarLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Césarejeeem"}

	cesarRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Цезарьвысочество"}

	cesarSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿朗拍胸脯"}

	cesarSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Césarmonigote"}

	cesarTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿朗拍胸脯"}
)

var (
	cesarName = nook.Languages{
		language.AmericanEnglish:      cesarAmericanEnglishName,
		language.CanadianFrench:       cesarCanadianFrenchName,
		language.Dutch:                cesarDutchName,
		language.French:               cesarFrenchName,
		language.German:               cesarGermanName,
		language.Italian:              cesarItalianName,
		language.Japanese:             cesarJapaneseName,
		language.Korean:               cesarKoreanName,
		language.LatinAmericanSpanish: cesarLatinAmericanSpanishName,
		language.Russian:              cesarRussianName,
		language.SimplifiedChinese:    cesarSimplifiedChineseName,
		language.Spanish:              cesarSpanishName,
		language.TraditionalChinese:   cesarTraditionalChineseName}
)

var (
	cesarCharacter = nook.Character{
		Animal:   Gorilla,
		Birthday: cesarBirthday,
		Code:     cesarCode,
		Gender:   nook.Male,
		Name:     cesarName}
)

var (
	cesarAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "highness"}

	cesarCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "brutus"}

	cesarDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hoogheid"}

	cesarFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "brutus"}

	cesarGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hoheit"}

	cesarItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "dado"}

	cesarJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ウホウホ"}

	cesarKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "우갸우갸"}

	cesarLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ejeeem"}

	cesarRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "высочество"}

	cesarSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "拍胸脯"}

	cesarSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "monigote"}

	cesarTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "拍胸脯"}
)

var (
	cesarPhrase = nook.Languages{
		language.AmericanEnglish:      cesarAmericanEnglishName,
		language.CanadianFrench:       cesarCanadianFrenchName,
		language.Dutch:                cesarDutchName,
		language.French:               cesarFrenchName,
		language.German:               cesarGermanName,
		language.Italian:              cesarItalianName,
		language.Japanese:             cesarJapaneseName,
		language.Korean:               cesarKoreanName,
		language.LatinAmericanSpanish: cesarLatinAmericanSpanishName,
		language.Russian:              cesarRussianName,
		language.SimplifiedChinese:    cesarSimplifiedChineseName,
		language.Spanish:              cesarSpanishName,
		language.TraditionalChinese:   cesarTraditionalChineseName}
)

var (
	Cesar = nook.Villager{
		Character:   cesarCharacter,
		Personality: nook.Cranky,
		Phrase:      cesarPhrase}
)

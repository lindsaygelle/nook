package bird

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	madamrosaBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	madamrosaCode = nook.Code{
		Value: ""}
)

var (
	madamrosaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Madam Rosa"}

	madamrosaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	madamrosaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	madamrosaFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	madamrosaGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	madamrosaItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	madamrosaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マダムローザ"}

	madamrosaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	madamrosaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	madamrosaRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	madamrosaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	madamrosaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	madamrosaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	madamrosaName = nook.Languages{
		language.AmericanEnglish:      madamrosaAmericanEnglishName,
		language.CanadianFrench:       madamrosaCanadianFrenchName,
		language.Dutch:                madamrosaDutchName,
		language.French:               madamrosaFrenchName,
		language.German:               madamrosaGermanName,
		language.Italian:              madamrosaItalianName,
		language.Japanese:             madamrosaJapaneseName,
		language.Korean:               madamrosaKoreanName,
		language.LatinAmericanSpanish: madamrosaLatinAmericanSpanishName,
		language.Russian:              madamrosaRussianName,
		language.SimplifiedChinese:    madamrosaSimplifiedChineseName,
		language.Spanish:              madamrosaSpanishName,
		language.TraditionalChinese:   madamrosaTraditionalChineseName}
)

var (
	madamrosaCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: madamrosaBirthday,
		Code:     madamrosaCode,
		Gender:   gender.Female,
		Name:     madamrosaName}
)

var (
	madamrosaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ほほほ"}

	madamrosaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	madamrosaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	madamrosaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	madamrosaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	madamrosaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	madamrosaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ほほほ"}

	madamrosaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	madamrosaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	madamrosaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	madamrosaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	madamrosaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	madamrosaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	madamrosaPhrase = nook.Languages{
		language.AmericanEnglish:      madamrosaAmericanEnglishName,
		language.CanadianFrench:       madamrosaCanadianFrenchName,
		language.Dutch:                madamrosaDutchName,
		language.French:               madamrosaFrenchName,
		language.German:               madamrosaGermanName,
		language.Italian:              madamrosaItalianName,
		language.Japanese:             madamrosaJapaneseName,
		language.Korean:               madamrosaKoreanName,
		language.LatinAmericanSpanish: madamrosaLatinAmericanSpanishName,
		language.Russian:              madamrosaRussianName,
		language.SimplifiedChinese:    madamrosaSimplifiedChineseName,
		language.Spanish:              madamrosaSpanishName,
		language.TraditionalChinese:   madamrosaTraditionalChineseName}
)

var (
	MadamRosa = nook.Villager{
		Character:   madamrosaCharacter,
		Personality: personality.Snooty,
		Phrase:      madamrosaPhrase}
)

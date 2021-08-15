package bird

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "マダムローザほほほ"}

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
		Animal:   Bird,
		Birthday: madamrosaBirthday,
		Code:     madamrosaCode,
		Gender:   nook.Female,
		Name:     madamrosaName}
)

var (
	madamrosaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	madamrosaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	madamrosaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	madamrosaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	madamrosaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	madamrosaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	madamrosaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ほほほ"}

	madamrosaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	madamrosaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	madamrosaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	madamrosaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	madamrosaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	madamrosaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
		Personality: nook.Snooty,
		Phrase:      madamrosaPhrase}
)

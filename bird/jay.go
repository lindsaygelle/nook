package bird

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	jayBirthday = nook.Birthday{
		Day:   17,
		Month: time.July}
)

var (
	jayCode = nook.Code{
		Value: "brd00"}
)

var (
	jayAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Jay"}

	jayCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gérardmaaaaaais"}

	jayDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Jayheeeeeyy"}

	jayFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gérardmaaaaaais"}

	jayGermanName = nook.Name{
		Language: language.German,
		Value:    "Janhiiiijjj"}

	jayItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Jetcipinci"}

	jayJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ツバクロでおます"}

	jayKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "참돌이그쵸"}

	jayLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Jairoeeey"}

	jayRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Джейхей-хей"}

	jaySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿燕呢喃"}

	jaySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Jotaheeey"}

	jayTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿燕呢喃"}
)

var (
	jayName = nook.Languages{
		language.AmericanEnglish:      jayAmericanEnglishName,
		language.CanadianFrench:       jayCanadianFrenchName,
		language.Dutch:                jayDutchName,
		language.French:               jayFrenchName,
		language.German:               jayGermanName,
		language.Italian:              jayItalianName,
		language.Japanese:             jayJapaneseName,
		language.Korean:               jayKoreanName,
		language.LatinAmericanSpanish: jayLatinAmericanSpanishName,
		language.Russian:              jayRussianName,
		language.SimplifiedChinese:    jaySimplifiedChineseName,
		language.Spanish:              jaySpanishName,
		language.TraditionalChinese:   jayTraditionalChineseName}
)

var (
	jayCharacter = nook.Character{
		Animal:   Bird,
		Birthday: jayBirthday,
		Code:     jayCode,
		Gender:   nook.Male,
		Name:     jayName}
)

var (
	jayAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "heeeeeyy"}

	jayCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "maaaaaais"}

	jayDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "heeeeeyy"}

	jayFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "maaaaaais"}

	jayGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hiiiijjj"}

	jayItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cipinci"}

	jayJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でおます"}

	jayKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그쵸"}

	jayLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "eeey"}

	jayRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хей-хей"}

	jaySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "呢喃"}

	jaySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "heeey"}

	jayTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "呢喃"}
)

var (
	jayPhrase = nook.Languages{
		language.AmericanEnglish:      jayAmericanEnglishName,
		language.CanadianFrench:       jayCanadianFrenchName,
		language.Dutch:                jayDutchName,
		language.French:               jayFrenchName,
		language.German:               jayGermanName,
		language.Italian:              jayItalianName,
		language.Japanese:             jayJapaneseName,
		language.Korean:               jayKoreanName,
		language.LatinAmericanSpanish: jayLatinAmericanSpanishName,
		language.Russian:              jayRussianName,
		language.SimplifiedChinese:    jaySimplifiedChineseName,
		language.Spanish:              jaySpanishName,
		language.TraditionalChinese:   jayTraditionalChineseName}
)

var (
	Jay = nook.Villager{
		Character:   jayCharacter,
		Personality: nook.Jock,
		Phrase:      jayPhrase}
)

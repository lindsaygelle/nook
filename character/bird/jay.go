package bird

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
		Value:    "Gérard"}

	jayDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Jay"}

	jayFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gérard"}

	jayGermanName = nook.Name{
		Language: language.German,
		Value:    "Jan"}

	jayItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Jet"}

	jayJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ツバクロ"}

	jayKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "참돌이"}

	jayLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Jairo"}

	jayRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Джей"}

	jaySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿燕"}

	jaySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Jota"}

	jayTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿燕"}
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
		Animal:   animal.Bird,
		Birthday: jayBirthday,
		Code:     jayCode,
		Key:      character.Jay,
		Gender:   gender.Male,
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
		Personality: personality.Jock,
		Phrase:      jayPhrase}
)

package horse

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	elmerBirthday = nook.Birthday{
		Day:   5,
		Month: time.October}
)

var (
	elmerCode = nook.Code{
		Value: "hrs03"}
)

var (
	elmerAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Elmer"}

	elmerCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Martin"}

	elmerDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Elmer"}

	elmerFrenchName = nook.Name{
		Language: language.French,
		Value:    "Martin"}

	elmerGermanName = nook.Name{
		Language: language.German,
		Value:    "Emil"}

	elmerItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Oreste"}

	elmerJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "サブレ"}

	elmerKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "샤브렌"}

	elmerLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Jacinto"}

	elmerRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Элмер"}

	elmerSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "酥饼"}

	elmerSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Jacinto"}

	elmerTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "酥餅"}
)

var (
	elmerName = nook.Languages{
		language.AmericanEnglish:      elmerAmericanEnglishName,
		language.CanadianFrench:       elmerCanadianFrenchName,
		language.Dutch:                elmerDutchName,
		language.French:               elmerFrenchName,
		language.German:               elmerGermanName,
		language.Italian:              elmerItalianName,
		language.Japanese:             elmerJapaneseName,
		language.Korean:               elmerKoreanName,
		language.LatinAmericanSpanish: elmerLatinAmericanSpanishName,
		language.Russian:              elmerRussianName,
		language.SimplifiedChinese:    elmerSimplifiedChineseName,
		language.Spanish:              elmerSpanishName,
		language.TraditionalChinese:   elmerTraditionalChineseName}
)

var (
	elmerCharacter = nook.Character{
		Animal:   animal.Horse,
		Birthday: elmerBirthday,
		Code:     elmerCode,
		Gender:   gender.Male,
		Name:     elmerName}
)

var (
	elmerAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "tenderfoot"}

	elmerCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mouuuuiii"}

	elmerDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "jonkie"}

	elmerFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mouuuuiii"}

	elmerGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "anfänger"}

	elmerItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "pinopino"}

	elmerJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だヒヒン"}

	elmerKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "히히힝"}

	elmerLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "hop-uah"}

	elmerRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "жеребенок"}

	elmerSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "酥酥"}

	elmerSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pezuñitas"}

	elmerTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "酥酥"}
)

var (
	elmerPhrase = nook.Languages{
		language.AmericanEnglish:      elmerAmericanEnglishName,
		language.CanadianFrench:       elmerCanadianFrenchName,
		language.Dutch:                elmerDutchName,
		language.French:               elmerFrenchName,
		language.German:               elmerGermanName,
		language.Italian:              elmerItalianName,
		language.Japanese:             elmerJapaneseName,
		language.Korean:               elmerKoreanName,
		language.LatinAmericanSpanish: elmerLatinAmericanSpanishName,
		language.Russian:              elmerRussianName,
		language.SimplifiedChinese:    elmerSimplifiedChineseName,
		language.Spanish:              elmerSpanishName,
		language.TraditionalChinese:   elmerTraditionalChineseName}
)

var (
	Elmer = nook.Villager{
		Character:   elmerCharacter,
		Personality: personality.Lazy,
		Phrase:      elmerPhrase}
)

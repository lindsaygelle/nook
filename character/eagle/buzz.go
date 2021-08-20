package eagle

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
	buzzBirthday = nook.Birthday{
		Day:   7,
		Month: time.December}
)

var (
	buzzCode = nook.Code{
		Value: "pbr03"}
)

var (
	buzzAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Buzz"}

	buzzCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Phébus"}

	buzzDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Buzz"}

	buzzFrenchName = nook.Name{
		Language: language.French,
		Value:    "Phébus"}

	buzzGermanName = nook.Name{
		Language: language.German,
		Value:    "Balduin"}

	buzzItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Golia"}

	buzzJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ひでよし"}

	buzzKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "근엄"}

	buzzLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Nabar"}

	buzzRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Базз"}

	buzzSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "继光"}

	buzzSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Nabar"}

	buzzTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "繼光"}
)

var (
	buzzName = nook.Languages{
		language.AmericanEnglish:      buzzAmericanEnglishName,
		language.CanadianFrench:       buzzCanadianFrenchName,
		language.Dutch:                buzzDutchName,
		language.French:               buzzFrenchName,
		language.German:               buzzGermanName,
		language.Italian:              buzzItalianName,
		language.Japanese:             buzzJapaneseName,
		language.Korean:               buzzKoreanName,
		language.LatinAmericanSpanish: buzzLatinAmericanSpanishName,
		language.Russian:              buzzRussianName,
		language.SimplifiedChinese:    buzzSimplifiedChineseName,
		language.Spanish:              buzzSpanishName,
		language.TraditionalChinese:   buzzTraditionalChineseName}
)

var (
	buzzCharacter = nook.Character{
		Animal:   animal.Eagle,
		Birthday: buzzBirthday,
		Code:     buzzCode,
		Key:      character.Buzz,
		Gender:   gender.Male,
		Name:     buzzName,
		Special:  false}
)

var (
	buzzAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "captain"}

	buzzCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "flap flap"}

	buzzDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kaptein"}

	buzzFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "flap flap"}

	buzzGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "käpten"}

	buzzItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "capo"}

	buzzJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ッキーッ"}

	buzzKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "짜란～"}

	buzzLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "capííí"}

	buzzRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "капитан"}

	buzzSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "继继"}

	buzzSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cadete"}

	buzzTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "繼繼"}
)

var (
	buzzPhrase = nook.Languages{
		language.AmericanEnglish:      buzzAmericanEnglishPhrase,
		language.CanadianFrench:       buzzCanadianFrenchPhrase,
		language.Dutch:                buzzDutchPhrase,
		language.French:               buzzFrenchPhrase,
		language.German:               buzzGermanPhrase,
		language.Italian:              buzzItalianPhrase,
		language.Japanese:             buzzJapanesePhrase,
		language.Korean:               buzzKoreanPhrase,
		language.LatinAmericanSpanish: buzzLatinAmericanSpanishPhrase,
		language.Russian:              buzzRussianPhrase,
		language.SimplifiedChinese:    buzzSimplifiedChinesePhrase,
		language.Spanish:              buzzSpanishPhrase,
		language.TraditionalChinese:   buzzTraditionalChinesePhrase}
)

var (
	Buzz = nook.Villager{
		Character:   buzzCharacter,
		Personality: personality.Cranky,
		Phrase:      buzzPhrase}
)

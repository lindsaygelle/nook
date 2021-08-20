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
	jacobBirthday = nook.Birthday{
		Day:   24,
		Month: time.August}
)

var (
	jacobCode = nook.Code{
		Value: "brd11"}
)

var (
	jacobAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Jacob Jakey"}

	jacobCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Jacob"}

	jacobDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Jacob"}

	jacobFrenchName = nook.Name{
		Language: language.French,
		Value:    "Jacob"}

	jacobGermanName = nook.Name{
		Language: language.German,
		Value:    "Jesko"}

	jacobItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Valerio"}

	jacobJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジャコテン"}

	jacobKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "야곱"}

	jacobLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Repollo"}

	jacobRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Джейкоб"}

	jacobSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "余板"}

	jacobSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Repollo"}

	jacobTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "余板"}
)

var (
	jacobName = nook.Languages{
		language.AmericanEnglish:      jacobAmericanEnglishName,
		language.CanadianFrench:       jacobCanadianFrenchName,
		language.Dutch:                jacobDutchName,
		language.French:               jacobFrenchName,
		language.German:               jacobGermanName,
		language.Italian:              jacobItalianName,
		language.Japanese:             jacobJapaneseName,
		language.Korean:               jacobKoreanName,
		language.LatinAmericanSpanish: jacobLatinAmericanSpanishName,
		language.Russian:              jacobRussianName,
		language.SimplifiedChinese:    jacobSimplifiedChineseName,
		language.Spanish:              jacobSpanishName,
		language.TraditionalChinese:   jacobTraditionalChineseName}
)

var (
	jacobCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: jacobBirthday,
		Code:     jacobCode,
		Key:      character.Jacob,
		Gender:   gender.Male,
		Name:     jacobName,
		Special:  false}
)

var (
	jacobAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ya feel"}

	jacobCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "asticô"}

	jacobDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snappie"}

	jacobFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "asticô"}

	jacobGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "tirilü"}

	jacobItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ciiipito"}

	jacobJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "っつーの"}

	jacobKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그거야"}

	jacobLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "chiurp"}

	jacobRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "чуррп"}

	jacobSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "话说"}

	jacobSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pelusilla"}

	jacobTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "話說"}
)

var (
	jacobPhrase = nook.Languages{
		language.AmericanEnglish:      jacobAmericanEnglishPhrase,
		language.CanadianFrench:       jacobCanadianFrenchPhrase,
		language.Dutch:                jacobDutchPhrase,
		language.French:               jacobFrenchPhrase,
		language.German:               jacobGermanPhrase,
		language.Italian:              jacobItalianPhrase,
		language.Japanese:             jacobJapanesePhrase,
		language.Korean:               jacobKoreanPhrase,
		language.LatinAmericanSpanish: jacobLatinAmericanSpanishPhrase,
		language.Russian:              jacobRussianPhrase,
		language.SimplifiedChinese:    jacobSimplifiedChinesePhrase,
		language.Spanish:              jacobSpanishPhrase,
		language.TraditionalChinese:   jacobTraditionalChinesePhrase}
)

var (
	Jacob = nook.Villager{
		Character:   jacobCharacter,
		Personality: personality.Lazy,
		Phrase:      jacobPhrase}
)

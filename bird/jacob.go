package bird

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Jacobasticô"}

	jacobDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Jacobsnappie"}

	jacobFrenchName = nook.Name{
		Language: language.French,
		Value:    "Jacobasticô"}

	jacobGermanName = nook.Name{
		Language: language.German,
		Value:    "Jeskotirilü"}

	jacobItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Valeriociiipito"}

	jacobJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジャコテンっつーの"}

	jacobKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "야곱그거야"}

	jacobLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Repollochiurp"}

	jacobRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Джейкобчуррп"}

	jacobSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "余板话说"}

	jacobSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Repollopelusilla"}

	jacobTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "余板話說"}
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
		Animal:   Bird,
		Birthday: jacobBirthday,
		Code:     jacobCode,
		Gender:   nook.Male,
		Name:     jacobName}
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
	Jacob = nook.Villager{
		Character:   jacobCharacter,
		Personality: nook.Lazy,
		Phrase:      jacobPhrase}
)

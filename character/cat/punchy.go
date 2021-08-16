package cat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	punchyBirthday = nook.Birthday{
		Day:   11,
		Month: time.April}
)

var (
	punchyCode = nook.Code{
		Value: "cat06"}
)

var (
	punchyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Punchy"}

	punchyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Cédric"}

	punchyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Punchy"}

	punchyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Cédric"}

	punchyGermanName = nook.Name{
		Language: language.German,
		Value:    "Julian"}

	punchyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Felix"}

	punchyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ビンタ"}

	punchyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "빙티"}

	punchyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Félix"}

	punchyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Панчи"}

	punchySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "尔光"}

	punchySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Félix"}

	punchyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "爾光"}
)

var (
	punchyName = nook.Languages{
		language.AmericanEnglish:      punchyAmericanEnglishName,
		language.CanadianFrench:       punchyCanadianFrenchName,
		language.Dutch:                punchyDutchName,
		language.French:               punchyFrenchName,
		language.German:               punchyGermanName,
		language.Italian:              punchyItalianName,
		language.Japanese:             punchyJapaneseName,
		language.Korean:               punchyKoreanName,
		language.LatinAmericanSpanish: punchyLatinAmericanSpanishName,
		language.Russian:              punchyRussianName,
		language.SimplifiedChinese:    punchySimplifiedChineseName,
		language.Spanish:              punchySpanishName,
		language.TraditionalChinese:   punchyTraditionalChineseName}
)

var (
	punchyCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: punchyBirthday,
		Code:     punchyCode,
		Gender:   gender.Male,
		Name:     punchyName}
)

var (
	punchyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "mrmpht"}

	punchyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mimine"}

	punchyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "plof"}

	punchyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "blébléblé"}

	punchyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mrmpft"}

	punchyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "maramao"}

	punchyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だのら"}

	punchyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "노라줘"}

	punchyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "gruuuah"}

	punchyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "отдыхаем"}

	punchySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "晃晃"}

	punchySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "arenque"}

	punchyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "晃晃"}
)

var (
	punchyPhrase = nook.Languages{
		language.AmericanEnglish:      punchyAmericanEnglishName,
		language.CanadianFrench:       punchyCanadianFrenchName,
		language.Dutch:                punchyDutchName,
		language.French:               punchyFrenchName,
		language.German:               punchyGermanName,
		language.Italian:              punchyItalianName,
		language.Japanese:             punchyJapaneseName,
		language.Korean:               punchyKoreanName,
		language.LatinAmericanSpanish: punchyLatinAmericanSpanishName,
		language.Russian:              punchyRussianName,
		language.SimplifiedChinese:    punchySimplifiedChineseName,
		language.Spanish:              punchySpanishName,
		language.TraditionalChinese:   punchyTraditionalChineseName}
)

var (
	Punchy = nook.Villager{
		Character:   punchyCharacter,
		Personality: personality.Lazy,
		Phrase:      punchyPhrase}
)

package pig

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
	lucyBirthday = nook.Birthday{
		Day:   2,
		Month: time.June}
)

var (
	lucyCode = nook.Code{
		Value: "pig04"}
)

var (
	lucyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Lucy"}

	lucyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Lucie"}

	lucyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Lucy"}

	lucyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Lucie"}

	lucyGermanName = nook.Name{
		Language: language.German,
		Value:    "Larissa"}

	lucyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lucy"}

	lucyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ルーシー"}

	lucyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "루시"}

	lucyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Aurelia"}

	lucyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Люси"}

	lucySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "露西"}

	lucySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Aurelia"}

	lucyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "露西"}
)

var (
	lucyName = nook.Languages{
		language.AmericanEnglish:      lucyAmericanEnglishName,
		language.CanadianFrench:       lucyCanadianFrenchName,
		language.Dutch:                lucyDutchName,
		language.French:               lucyFrenchName,
		language.German:               lucyGermanName,
		language.Italian:              lucyItalianName,
		language.Japanese:             lucyJapaneseName,
		language.Korean:               lucyKoreanName,
		language.LatinAmericanSpanish: lucyLatinAmericanSpanishName,
		language.Russian:              lucyRussianName,
		language.SimplifiedChinese:    lucySimplifiedChineseName,
		language.Spanish:              lucySpanishName,
		language.TraditionalChinese:   lucyTraditionalChineseName}
)

var (
	lucyCharacter = nook.Character{
		Animal:   animal.Pig,
		Birthday: lucyBirthday,
		Code:     lucyCode,
		Key:      character.Lucy,
		Gender:   gender.Female,
		Name:     lucyName}
)

var (
	lucyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "snoooink"}

	lucyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "groingroin"}

	lucyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "knorrie"}

	lucyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "groingroin"}

	lucyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnoink"}

	lucyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "snoink"}

	lucyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "よぅ"}

	lucyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "예예"}

	lucyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cochín"}

	lucyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хрю-ю-ю"}

	lucySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "唷"}

	lucySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cochín"}

	lucyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "唷"}
)

var (
	lucyPhrase = nook.Languages{
		language.AmericanEnglish:      lucyAmericanEnglishName,
		language.CanadianFrench:       lucyCanadianFrenchName,
		language.Dutch:                lucyDutchName,
		language.French:               lucyFrenchName,
		language.German:               lucyGermanName,
		language.Italian:              lucyItalianName,
		language.Japanese:             lucyJapaneseName,
		language.Korean:               lucyKoreanName,
		language.LatinAmericanSpanish: lucyLatinAmericanSpanishName,
		language.Russian:              lucyRussianName,
		language.SimplifiedChinese:    lucySimplifiedChineseName,
		language.Spanish:              lucySpanishName,
		language.TraditionalChinese:   lucyTraditionalChineseName}
)

var (
	Lucy = nook.Villager{
		Character:   lucyCharacter,
		Personality: personality.Normal,
		Phrase:      lucyPhrase}
)

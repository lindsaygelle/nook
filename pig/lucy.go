package pig

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Luciegroingroin"}

	lucyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Lucyknorrie"}

	lucyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Luciegroingroin"}

	lucyGermanName = nook.Name{
		Language: language.German,
		Value:    "Larissaschnoink"}

	lucyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lucysnoink"}

	lucyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ルーシーよぅ"}

	lucyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "루시예예"}

	lucyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Aureliacochín"}

	lucyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Люсихрю-ю-ю"}

	lucySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "露西唷"}

	lucySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Aureliacochín"}

	lucyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "露西唷"}
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
		Animal:   Pig,
		Birthday: lucyBirthday,
		Code:     lucyCode,
		Gender:   nook.Female,
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
		Personality: nook.Normal,
		Phrase:      lucyPhrase}
)

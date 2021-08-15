package eagle

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	frankBirthday = nook.Birthday{
		Day:   30,
		Month: time.July}
)

var (
	frankCode = nook.Code{
		Value: "pbr06"}
)

var (
	frankAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Frank"}

	frankCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Greggaemaaaarre"}

	frankDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Frankpletter"}

	frankFrenchName = nook.Name{
		Language: language.French,
		Value:    "Greggaemaaaarre"}

	frankGermanName = nook.Name{
		Language: language.German,
		Value:    "Arthurkriiieck"}

	frankItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Frankrapaciao"}

	frankJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハルクダンケ"}

	frankKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "헐크크헐"}

	frankLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Aquilinovulanico"}

	frankRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Франккрасава"}

	frankSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "浩克多谢"}

	frankSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Aquilinovulanico"}

	frankTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "浩克多謝"}
)

var (
	frankName = nook.Languages{
		language.AmericanEnglish:      frankAmericanEnglishName,
		language.CanadianFrench:       frankCanadianFrenchName,
		language.Dutch:                frankDutchName,
		language.French:               frankFrenchName,
		language.German:               frankGermanName,
		language.Italian:              frankItalianName,
		language.Japanese:             frankJapaneseName,
		language.Korean:               frankKoreanName,
		language.LatinAmericanSpanish: frankLatinAmericanSpanishName,
		language.Russian:              frankRussianName,
		language.SimplifiedChinese:    frankSimplifiedChineseName,
		language.Spanish:              frankSpanishName,
		language.TraditionalChinese:   frankTraditionalChineseName}
)

var (
	frankCharacter = nook.Character{
		Animal:   Eagle,
		Birthday: frankBirthday,
		Code:     frankCode,
		Gender:   nook.Male,
		Name:     frankName}
)

var (
	frankAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "crushy"}

	frankCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "maaaarre"}

	frankDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "pletter"}

	frankFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "maaaarre"}

	frankGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kriiieck"}

	frankItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "rapaciao"}

	frankJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ダンケ"}

	frankKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "크헐"}

	frankLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "vulanico"}

	frankRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "красава"}

	frankSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "多谢"}

	frankSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "vulanico"}

	frankTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "多謝"}
)

var (
	frankPhrase = nook.Languages{
		language.AmericanEnglish:      frankAmericanEnglishName,
		language.CanadianFrench:       frankCanadianFrenchName,
		language.Dutch:                frankDutchName,
		language.French:               frankFrenchName,
		language.German:               frankGermanName,
		language.Italian:              frankItalianName,
		language.Japanese:             frankJapaneseName,
		language.Korean:               frankKoreanName,
		language.LatinAmericanSpanish: frankLatinAmericanSpanishName,
		language.Russian:              frankRussianName,
		language.SimplifiedChinese:    frankSimplifiedChineseName,
		language.Spanish:              frankSpanishName,
		language.TraditionalChinese:   frankTraditionalChineseName}
)

var (
	Frank = nook.Villager{
		Character:   frankCharacter,
		Personality: nook.Cranky,
		Phrase:      frankPhrase}
)

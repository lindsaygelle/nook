package hamster

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	grahamBirthday = nook.Birthday{
		Day:   20,
		Month: time.June}
)

var (
	grahamCode = nook.Code{
		Value: "ham02"}
)

var (
	grahamAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Graham"}

	grahamCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Grahamau pif"}

	grahamDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Grahaminderdaad"}

	grahamFrenchName = nook.Name{
		Language: language.French,
		Value:    "Grahamau pif"}

	grahamGermanName = nook.Name{
		Language: language.German,
		Value:    "Güntherhähäm"}

	grahamItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Evaristopaella"}

	grahamJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "グラハムですぞ"}

	grahamKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "글라햄그렇다고"}

	grahamLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Roelioñisqui"}

	grahamRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Грэмрезонно"}

	grahamSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "麦麦就是这样"}

	grahamSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Roelioñisqui"}

	grahamTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "麥麥就是這樣"}
)

var (
	grahamName = nook.Languages{
		language.AmericanEnglish:      grahamAmericanEnglishName,
		language.CanadianFrench:       grahamCanadianFrenchName,
		language.Dutch:                grahamDutchName,
		language.French:               grahamFrenchName,
		language.German:               grahamGermanName,
		language.Italian:              grahamItalianName,
		language.Japanese:             grahamJapaneseName,
		language.Korean:               grahamKoreanName,
		language.LatinAmericanSpanish: grahamLatinAmericanSpanishName,
		language.Russian:              grahamRussianName,
		language.SimplifiedChinese:    grahamSimplifiedChineseName,
		language.Spanish:              grahamSpanishName,
		language.TraditionalChinese:   grahamTraditionalChineseName}
)

var (
	grahamCharacter = nook.Character{
		Animal:   Hamster,
		Birthday: grahamBirthday,
		Code:     grahamCode,
		Gender:   nook.Male,
		Name:     grahamName}
)

var (
	grahamAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "indeed"}

	grahamCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "au pif"}

	grahamDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "inderdaad"}

	grahamFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "au pif"}

	grahamGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hähäm"}

	grahamItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "paella"}

	grahamJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ですぞ"}

	grahamKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그렇다고"}

	grahamLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ñisqui"}

	grahamRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "резонно"}

	grahamSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "就是这样"}

	grahamSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ñisqui"}

	grahamTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "就是這樣"}
)

var (
	grahamPhrase = nook.Languages{
		language.AmericanEnglish:      grahamAmericanEnglishName,
		language.CanadianFrench:       grahamCanadianFrenchName,
		language.Dutch:                grahamDutchName,
		language.French:               grahamFrenchName,
		language.German:               grahamGermanName,
		language.Italian:              grahamItalianName,
		language.Japanese:             grahamJapaneseName,
		language.Korean:               grahamKoreanName,
		language.LatinAmericanSpanish: grahamLatinAmericanSpanishName,
		language.Russian:              grahamRussianName,
		language.SimplifiedChinese:    grahamSimplifiedChineseName,
		language.Spanish:              grahamSpanishName,
		language.TraditionalChinese:   grahamTraditionalChineseName}
)

var (
	Graham = nook.Villager{
		Character:   grahamCharacter,
		Personality: nook.Smug,
		Phrase:      grahamPhrase}
)

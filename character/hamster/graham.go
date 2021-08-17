package hamster

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
		Value:    "Graham"}

	grahamDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Graham"}

	grahamFrenchName = nook.Name{
		Language: language.French,
		Value:    "Graham"}

	grahamGermanName = nook.Name{
		Language: language.German,
		Value:    "Günther"}

	grahamItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Evaristo"}

	grahamJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "グラハム"}

	grahamKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "글라햄"}

	grahamLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Roelio"}

	grahamRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Грэм"}

	grahamSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "麦麦"}

	grahamSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Roelio"}

	grahamTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "麥麥"}
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
		Animal:   animal.Hamster,
		Birthday: grahamBirthday,
		Code:     grahamCode,
		Key:      character.Graham,
		Gender:   gender.Male,
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
		language.AmericanEnglish:      grahamAmericanEnglishPhrase,
		language.CanadianFrench:       grahamCanadianFrenchPhrase,
		language.Dutch:                grahamDutchPhrase,
		language.French:               grahamFrenchPhrase,
		language.German:               grahamGermanPhrase,
		language.Italian:              grahamItalianPhrase,
		language.Japanese:             grahamJapanesePhrase,
		language.Korean:               grahamKoreanPhrase,
		language.LatinAmericanSpanish: grahamLatinAmericanSpanishPhrase,
		language.Russian:              grahamRussianPhrase,
		language.SimplifiedChinese:    grahamSimplifiedChinesePhrase,
		language.Spanish:              grahamSpanishPhrase,
		language.TraditionalChinese:   grahamTraditionalChinesePhrase}
)

var (
	Graham = nook.Villager{
		Character:   grahamCharacter,
		Personality: personality.Smug,
		Phrase:      grahamPhrase}
)

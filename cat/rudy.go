package cat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	rudyBirthday = nook.Birthday{
		Day:   20,
		Month: time.December}
)

var (
	rudyCode = nook.Code{
		Value: "cat20"}
)

var (
	rudyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rudy"}

	rudyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rougepifmouah"}

	rudyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rudyrobijn"}

	rudyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rougepifalatienne"}

	rudyGermanName = nook.Name{
		Language: language.German,
		Value:    "Heinzkatzaaa"}

	rudyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gomitoloron ron"}

	rudyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チャスとかナ"}

	rudyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "찰스그러거나"}

	rudyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rufinocachusco"}

	rudyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рудишмяк"}

	rudySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "茶茶之类的"}

	rudySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rufinocachusco"}

	rudyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "茶茶之類的"}
)

var (
	rudyName = nook.Languages{
		language.AmericanEnglish:      rudyAmericanEnglishName,
		language.CanadianFrench:       rudyCanadianFrenchName,
		language.Dutch:                rudyDutchName,
		language.French:               rudyFrenchName,
		language.German:               rudyGermanName,
		language.Italian:              rudyItalianName,
		language.Japanese:             rudyJapaneseName,
		language.Korean:               rudyKoreanName,
		language.LatinAmericanSpanish: rudyLatinAmericanSpanishName,
		language.Russian:              rudyRussianName,
		language.SimplifiedChinese:    rudySimplifiedChineseName,
		language.Spanish:              rudySpanishName,
		language.TraditionalChinese:   rudyTraditionalChineseName}
)

var (
	rudyCharacter = nook.Character{
		Animal:   Cat,
		Birthday: rudyBirthday,
		Code:     rudyCode,
		Gender:   nook.Male,
		Name:     rudyName}
)

var (
	rudyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "mush"}

	rudyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mouah"}

	rudyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "robijn"}

	rudyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "alatienne"}

	rudyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "katzaaa"}

	rudyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ron ron"}

	rudyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "とかナ"}

	rudyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그러거나"}

	rudyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cachusco"}

	rudyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "шмяк"}

	rudySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "之类的"}

	rudySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cachusco"}

	rudyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "之類的"}
)

var (
	rudyPhrase = nook.Languages{
		language.AmericanEnglish:      rudyAmericanEnglishName,
		language.CanadianFrench:       rudyCanadianFrenchName,
		language.Dutch:                rudyDutchName,
		language.French:               rudyFrenchName,
		language.German:               rudyGermanName,
		language.Italian:              rudyItalianName,
		language.Japanese:             rudyJapaneseName,
		language.Korean:               rudyKoreanName,
		language.LatinAmericanSpanish: rudyLatinAmericanSpanishName,
		language.Russian:              rudyRussianName,
		language.SimplifiedChinese:    rudySimplifiedChineseName,
		language.Spanish:              rudySpanishName,
		language.TraditionalChinese:   rudyTraditionalChineseName}
)

var (
	Rudy = nook.Villager{
		Character:   rudyCharacter,
		Personality: nook.Jock,
		Phrase:      rudyPhrase}
)

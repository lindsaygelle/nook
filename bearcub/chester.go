package bearcub

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	chesterBirthday = nook.Birthday{
		Day:   6,
		Month: time.August}
)

var (
	chesterCode = nook.Code{
		Value: "cbr15"}
)

var (
	chesterAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Chester"}

	chesterCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Placidepla-pla"}

	chesterDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Chesterbamboe"}

	chesterFrenchName = nook.Name{
		Language: language.French,
		Value:    "Placidepla-pla"}

	chesterGermanName = nook.Name{
		Language: language.German,
		Value:    "Eduardbaaambus"}

	chesterItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Clementebrupp"}

	chesterJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "パンタだバンブ"}

	chesterKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "팬타부끄"}

	chesterLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Osunioahivá"}

	chesterRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Честербамбук"}

	chesterSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "胖达竹子"}

	chesterSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Osunioahivá"}

	chesterTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "胖達竹子"}
)

var (
	chesterName = nook.Languages{
		language.AmericanEnglish:      chesterAmericanEnglishName,
		language.CanadianFrench:       chesterCanadianFrenchName,
		language.Dutch:                chesterDutchName,
		language.French:               chesterFrenchName,
		language.German:               chesterGermanName,
		language.Italian:              chesterItalianName,
		language.Japanese:             chesterJapaneseName,
		language.Korean:               chesterKoreanName,
		language.LatinAmericanSpanish: chesterLatinAmericanSpanishName,
		language.Russian:              chesterRussianName,
		language.SimplifiedChinese:    chesterSimplifiedChineseName,
		language.Spanish:              chesterSpanishName,
		language.TraditionalChinese:   chesterTraditionalChineseName}
)

var (
	chesterCharacter = nook.Character{
		Animal:   Bearcub,
		Birthday: chesterBirthday,
		Code:     chesterCode,
		Gender:   nook.Male,
		Name:     chesterName}
)

var (
	chesterAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "rookie"}

	chesterCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pla-pla"}

	chesterDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "bamboe"}

	chesterFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pla-pla"}

	chesterGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "baaambus"}

	chesterItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "brupp"}

	chesterJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だバンブ"}

	chesterKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "부끄"}

	chesterLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ahivá"}

	chesterRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бамбук"}

	chesterSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "竹子"}

	chesterSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ahivá"}

	chesterTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "竹子"}
)

var (
	chesterPhrase = nook.Languages{
		language.AmericanEnglish:      chesterAmericanEnglishName,
		language.CanadianFrench:       chesterCanadianFrenchName,
		language.Dutch:                chesterDutchName,
		language.French:               chesterFrenchName,
		language.German:               chesterGermanName,
		language.Italian:              chesterItalianName,
		language.Japanese:             chesterJapaneseName,
		language.Korean:               chesterKoreanName,
		language.LatinAmericanSpanish: chesterLatinAmericanSpanishName,
		language.Russian:              chesterRussianName,
		language.SimplifiedChinese:    chesterSimplifiedChineseName,
		language.Spanish:              chesterSpanishName,
		language.TraditionalChinese:   chesterTraditionalChineseName}
)

var (
	Chester = nook.Villager{
		Character:   chesterCharacter,
		Personality: nook.Lazy,
		Phrase:      chesterPhrase}
)

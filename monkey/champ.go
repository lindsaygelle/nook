package monkey

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	champBirthday = nook.Birthday{
		Day:   4,
		Month: time.June}
)

var (
	champCode = nook.Code{
		Value: "mnk00"}
)

var (
	champAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Champ"}

	champCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ThibauttchouTCHOU"}

	champDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	champFrenchName = nook.Name{
		Language: language.French,
		Value:    "ThibauttchouTCHOU"}

	champGermanName = nook.Name{
		Language: language.German,
		Value:    "Tschitau-u-u"}

	champItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Champchoo CHOO"}

	champJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "さるおウッキー"}

	champKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "몽돌이야아아"}

	champLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Micouh-ah-ah"}

	champRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	champSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	champSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Micouh-ah-ah"}

	champTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	champName = nook.Languages{
		language.AmericanEnglish:      champAmericanEnglishName,
		language.CanadianFrench:       champCanadianFrenchName,
		language.Dutch:                champDutchName,
		language.French:               champFrenchName,
		language.German:               champGermanName,
		language.Italian:              champItalianName,
		language.Japanese:             champJapaneseName,
		language.Korean:               champKoreanName,
		language.LatinAmericanSpanish: champLatinAmericanSpanishName,
		language.Russian:              champRussianName,
		language.SimplifiedChinese:    champSimplifiedChineseName,
		language.Spanish:              champSpanishName,
		language.TraditionalChinese:   champTraditionalChineseName}
)

var (
	champCharacter = nook.Character{
		Animal:   Monkey,
		Birthday: champBirthday,
		Code:     champCode,
		Gender:   nook.Male,
		Name:     champName}
)

var (
	champAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "choo CHOO"}

	champCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tchouTCHOU"}

	champDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	champFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tchouTCHOU"}

	champGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "u-u-u"}

	champItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "choo CHOO"}

	champJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ウッキー"}

	champKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "야아아"}

	champLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "uh-ah-ah"}

	champRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	champSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	champSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "uh-ah-ah"}

	champTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	champPhrase = nook.Languages{
		language.AmericanEnglish:      champAmericanEnglishName,
		language.CanadianFrench:       champCanadianFrenchName,
		language.Dutch:                champDutchName,
		language.French:               champFrenchName,
		language.German:               champGermanName,
		language.Italian:              champItalianName,
		language.Japanese:             champJapaneseName,
		language.Korean:               champKoreanName,
		language.LatinAmericanSpanish: champLatinAmericanSpanishName,
		language.Russian:              champRussianName,
		language.SimplifiedChinese:    champSimplifiedChineseName,
		language.Spanish:              champSpanishName,
		language.TraditionalChinese:   champTraditionalChineseName}
)

var (
	Champ = nook.Villager{
		Character:   champCharacter,
		Personality: nook.Jock,
		Phrase:      champPhrase}
)

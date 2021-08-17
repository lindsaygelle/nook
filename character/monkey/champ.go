package monkey

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
		Value:    "Thibaut"}

	champDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	champFrenchName = nook.Name{
		Language: language.French,
		Value:    "Thibaut"}

	champGermanName = nook.Name{
		Language: language.German,
		Value:    "Tschita"}

	champItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Champ"}

	champJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "さるお"}

	champKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "몽돌이"}

	champLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Mico"}

	champRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	champSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	champSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Mico"}

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
		Animal:   animal.Monkey,
		Birthday: champBirthday,
		Code:     champCode,
		Key:      character.Champ,
		Gender:   gender.Male,
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
		Value:    ""}

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
		Value:    ""}

	champSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	champSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "uh-ah-ah"}

	champTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	champPhrase = nook.Languages{
		language.AmericanEnglish:      champAmericanEnglishPhrase,
		language.CanadianFrench:       champCanadianFrenchPhrase,
		language.Dutch:                champDutchPhrase,
		language.French:               champFrenchPhrase,
		language.German:               champGermanPhrase,
		language.Italian:              champItalianPhrase,
		language.Japanese:             champJapanesePhrase,
		language.Korean:               champKoreanPhrase,
		language.LatinAmericanSpanish: champLatinAmericanSpanishPhrase,
		language.Russian:              champRussianPhrase,
		language.SimplifiedChinese:    champSimplifiedChinesePhrase,
		language.Spanish:              champSpanishPhrase,
		language.TraditionalChinese:   champTraditionalChinesePhrase}
)

var (
	Champ = nook.Villager{
		Character:   champCharacter,
		Personality: personality.Jock,
		Phrase:      champPhrase}
)

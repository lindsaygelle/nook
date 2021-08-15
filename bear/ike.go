package bear

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	ikeBirthday = nook.Birthday{
		Day:   16,
		Month: time.May}
)

var (
	ikeCode = nook.Code{
		Value: "bea11"}
)

var (
	ikeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Ike"}

	ikeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Isaacpot d'miel"}

	ikeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Ikeja denk"}

	ikeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Isaacpot d'miel"}

	ikeGermanName = nook.Name{
		Language: language.German,
		Value:    "Ikebrrruah"}

	ikeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ortensiosgrugno"}

	ikeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ダイクボウズ"}

	ikeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "대공터프"}

	ikeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Astillasgrurrr"}

	ikeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Айкгарр-сон"}

	ikeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "大功施主"}

	ikeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Astillastornillos"}

	ikeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "大功施主"}
)

var (
	ikeName = nook.Languages{
		language.AmericanEnglish:      ikeAmericanEnglishName,
		language.CanadianFrench:       ikeCanadianFrenchName,
		language.Dutch:                ikeDutchName,
		language.French:               ikeFrenchName,
		language.German:               ikeGermanName,
		language.Italian:              ikeItalianName,
		language.Japanese:             ikeJapaneseName,
		language.Korean:               ikeKoreanName,
		language.LatinAmericanSpanish: ikeLatinAmericanSpanishName,
		language.Russian:              ikeRussianName,
		language.SimplifiedChinese:    ikeSimplifiedChineseName,
		language.Spanish:              ikeSpanishName,
		language.TraditionalChinese:   ikeTraditionalChineseName}
)

var (
	ikeCharacter = nook.Character{
		Animal:   Bear,
		Birthday: ikeBirthday,
		Code:     ikeCode,
		Gender:   nook.Male,
		Name:     ikeName}
)

var (
	ikeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "roadie"}

	ikeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pot d'miel"}

	ikeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "ja denk"}

	ikeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pot d'miel"}

	ikeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "brrruah"}

	ikeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sgrugno"}

	ikeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ボウズ"}

	ikeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "터프"}

	ikeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "grurrr"}

	ikeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "гарр-сон"}

	ikeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "施主"}

	ikeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tornillos"}

	ikeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "施主"}
)

var (
	ikePhrase = nook.Languages{
		language.AmericanEnglish:      ikeAmericanEnglishName,
		language.CanadianFrench:       ikeCanadianFrenchName,
		language.Dutch:                ikeDutchName,
		language.French:               ikeFrenchName,
		language.German:               ikeGermanName,
		language.Italian:              ikeItalianName,
		language.Japanese:             ikeJapaneseName,
		language.Korean:               ikeKoreanName,
		language.LatinAmericanSpanish: ikeLatinAmericanSpanishName,
		language.Russian:              ikeRussianName,
		language.SimplifiedChinese:    ikeSimplifiedChineseName,
		language.Spanish:              ikeSpanishName,
		language.TraditionalChinese:   ikeTraditionalChineseName}
)

var (
	Ike = nook.Villager{
		Character:   ikeCharacter,
		Personality: nook.Cranky,
		Phrase:      ikePhrase}
)

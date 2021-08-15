package pig

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	sueeBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	sueeCode = nook.Code{
		Value: ""}
)

var (
	sueeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sue E"}

	sueeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	sueeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	sueeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Peguy"}

	sueeGermanName = nook.Name{
		Language: language.German,
		Value:    "Sabrina"}

	sueeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Piggy"}

	sueeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ブルジョア"}

	sueeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	sueeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	sueeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	sueeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "娇娇"}

	sueeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Porcia"}

	sueeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	sueeName = nook.Languages{
		language.AmericanEnglish:      sueeAmericanEnglishName,
		language.CanadianFrench:       sueeCanadianFrenchName,
		language.Dutch:                sueeDutchName,
		language.French:               sueeFrenchName,
		language.German:               sueeGermanName,
		language.Italian:              sueeItalianName,
		language.Japanese:             sueeJapaneseName,
		language.Korean:               sueeKoreanName,
		language.LatinAmericanSpanish: sueeLatinAmericanSpanishName,
		language.Russian:              sueeRussianName,
		language.SimplifiedChinese:    sueeSimplifiedChineseName,
		language.Spanish:              sueeSpanishName,
		language.TraditionalChinese:   sueeTraditionalChineseName}
)

var (
	sueeCharacter = nook.Character{
		Animal:   Pig,
		Birthday: sueeBirthday,
		Code:     sueeCode,
		Gender:   nook.Female,
		Name:     sueeName}
)

var (
	sueeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	sueeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	sueeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	sueeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "groingroin"}

	sueeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnaub"}

	sueeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "oinksnoink"}

	sueeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ざんす"}

	sueeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	sueeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	sueeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	sueeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "烦死"}

	sueeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "piojo"}

	sueeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	sueePhrase = nook.Languages{
		language.AmericanEnglish:      sueeAmericanEnglishName,
		language.CanadianFrench:       sueeCanadianFrenchName,
		language.Dutch:                sueeDutchName,
		language.French:               sueeFrenchName,
		language.German:               sueeGermanName,
		language.Italian:              sueeItalianName,
		language.Japanese:             sueeJapaneseName,
		language.Korean:               sueeKoreanName,
		language.LatinAmericanSpanish: sueeLatinAmericanSpanishName,
		language.Russian:              sueeRussianName,
		language.SimplifiedChinese:    sueeSimplifiedChineseName,
		language.Spanish:              sueeSpanishName,
		language.TraditionalChinese:   sueeTraditionalChineseName}
)

var (
	SueE = nook.Villager{
		Character:   sueeCharacter,
		Personality: nook.Snooty,
		Phrase:      sueePhrase}
)

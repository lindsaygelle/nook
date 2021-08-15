package bear

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	curtBirthday = nook.Birthday{
		Day:   1,
		Month: time.July}
)

var (
	curtCode = nook.Code{
		Value: "bea02"}
)

var (
	curtAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Curt"}

	curtCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Curt"}

	curtDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Curt"}

	curtFrenchName = nook.Name{
		Language: language.French,
		Value:    "Curt"}

	curtGermanName = nook.Name{
		Language: language.German,
		Value:    "Kurt"}

	curtItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Curt"}

	curtJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ガンテツ"}

	curtKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "뚝심"}

	curtLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Gorbaché"}

	curtRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Керт"}

	curtSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "铁熊"}

	curtSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Gorbaché"}

	curtTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鐵熊"}
)

var (
	curtName = nook.Languages{
		language.AmericanEnglish:      curtAmericanEnglishName,
		language.CanadianFrench:       curtCanadianFrenchName,
		language.Dutch:                curtDutchName,
		language.French:               curtFrenchName,
		language.German:               curtGermanName,
		language.Italian:              curtItalianName,
		language.Japanese:             curtJapaneseName,
		language.Korean:               curtKoreanName,
		language.LatinAmericanSpanish: curtLatinAmericanSpanishName,
		language.Russian:              curtRussianName,
		language.SimplifiedChinese:    curtSimplifiedChineseName,
		language.Spanish:              curtSpanishName,
		language.TraditionalChinese:   curtTraditionalChineseName}
)

var (
	curtCharacter = nook.Character{
		Animal:   Bear,
		Birthday: curtBirthday,
		Code:     curtCode,
		Gender:   nook.Male,
		Name:     curtName}
)

var (
	curtAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "fuzzball"}

	curtCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bouboule"}

	curtDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "brombeer"}

	curtFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bouboule"}

	curtGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "grumml"}

	curtItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "grrroan"}

	curtJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ウム"}

	curtKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "음"}

	curtLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "gromp"}

	curtRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "махрово"}

	curtSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嗯唔"}

	curtSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "oye"}

	curtTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嗯唔"}
)

var (
	curtPhrase = nook.Languages{
		language.AmericanEnglish:      curtAmericanEnglishName,
		language.CanadianFrench:       curtCanadianFrenchName,
		language.Dutch:                curtDutchName,
		language.French:               curtFrenchName,
		language.German:               curtGermanName,
		language.Italian:              curtItalianName,
		language.Japanese:             curtJapaneseName,
		language.Korean:               curtKoreanName,
		language.LatinAmericanSpanish: curtLatinAmericanSpanishName,
		language.Russian:              curtRussianName,
		language.SimplifiedChinese:    curtSimplifiedChineseName,
		language.Spanish:              curtSpanishName,
		language.TraditionalChinese:   curtTraditionalChineseName}
)

var (
	Curt = nook.Villager{
		Character:   curtCharacter,
		Personality: nook.Cranky,
		Phrase:      curtPhrase}
)

package deer

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	bamBirthday = nook.Birthday{
		Day:   7,
		Month: time.November}
)

var (
	bamCode = nook.Code{
		Value: "der01"}
)

var (
	bamAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bam"}

	bamCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Nacer"}

	bamDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bam"}

	bamFrenchName = nook.Name{
		Language: language.French,
		Value:    "Nacer"}

	bamGermanName = nook.Name{
		Language: language.German,
		Value:    "Benjamin"}

	bamItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cornelio"}

	bamJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "タケル"}

	bamKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "록키"}

	bamLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cornelio"}

	bamRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бам"}

	bamSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小健"}

	bamSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cornelio"}

	bamTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小健"}
)

var (
	bamName = nook.Languages{
		language.AmericanEnglish:      bamAmericanEnglishName,
		language.CanadianFrench:       bamCanadianFrenchName,
		language.Dutch:                bamDutchName,
		language.French:               bamFrenchName,
		language.German:               bamGermanName,
		language.Italian:              bamItalianName,
		language.Japanese:             bamJapaneseName,
		language.Korean:               bamKoreanName,
		language.LatinAmericanSpanish: bamLatinAmericanSpanishName,
		language.Russian:              bamRussianName,
		language.SimplifiedChinese:    bamSimplifiedChineseName,
		language.Spanish:              bamSpanishName,
		language.TraditionalChinese:   bamTraditionalChineseName}
)

var (
	bamCharacter = nook.Character{
		Animal:   Deer,
		Birthday: bamBirthday,
		Code:     bamCode,
		Gender:   nook.Male,
		Name:     bamName}
)

var (
	bamAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "boosh"}

	bamCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "têtedebois"}

	bamDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "gaffel"}

	bamFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "têtedebois"}

	bamGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "weihweih"}

	bamItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "hop hop"}

	bamJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ナラね"}

	bamKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "슉슉"}

	bamLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "caplúúú"}

	bamRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "скок-скок"}

	bamSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "梅花"}

	bamSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "galilla"}

	bamTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "梅花"}
)

var (
	bamPhrase = nook.Languages{
		language.AmericanEnglish:      bamAmericanEnglishName,
		language.CanadianFrench:       bamCanadianFrenchName,
		language.Dutch:                bamDutchName,
		language.French:               bamFrenchName,
		language.German:               bamGermanName,
		language.Italian:              bamItalianName,
		language.Japanese:             bamJapaneseName,
		language.Korean:               bamKoreanName,
		language.LatinAmericanSpanish: bamLatinAmericanSpanishName,
		language.Russian:              bamRussianName,
		language.SimplifiedChinese:    bamSimplifiedChineseName,
		language.Spanish:              bamSpanishName,
		language.TraditionalChinese:   bamTraditionalChineseName}
)

var (
	Bam = nook.Villager{
		Character:   bamCharacter,
		Personality: nook.Jock,
		Phrase:      bamPhrase}
)

package gorilla

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	booneBirthday = nook.Birthday{
		Day:   12,
		Month: time.September}
)

var (
	booneCode = nook.Code{
		Value: "gor02"}
)

var (
	booneAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Boone"}

	booneCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Babouinma noix"}

	booneDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Boonebavi-AAN"}

	booneFrenchName = nook.Name{
		Language: language.French,
		Value:    "Babouinma noix"}

	booneGermanName = nook.Name{
		Language: language.German,
		Value:    "Kongugh"}

	booneItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Babubaboom"}

	booneJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "まんたろうウルトラ"}

	booneKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "만복이야임마"}

	booneLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Babúangaua"}

	booneRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Буинбабум"}

	booneSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "万泰这家伙"}

	booneSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Babúangaua"}

	booneTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "萬泰超級"}
)

var (
	booneName = nook.Languages{
		language.AmericanEnglish:      booneAmericanEnglishName,
		language.CanadianFrench:       booneCanadianFrenchName,
		language.Dutch:                booneDutchName,
		language.French:               booneFrenchName,
		language.German:               booneGermanName,
		language.Italian:              booneItalianName,
		language.Japanese:             booneJapaneseName,
		language.Korean:               booneKoreanName,
		language.LatinAmericanSpanish: booneLatinAmericanSpanishName,
		language.Russian:              booneRussianName,
		language.SimplifiedChinese:    booneSimplifiedChineseName,
		language.Spanish:              booneSpanishName,
		language.TraditionalChinese:   booneTraditionalChineseName}
)

var (
	booneCharacter = nook.Character{
		Animal:   Gorilla,
		Birthday: booneBirthday,
		Code:     booneCode,
		Gender:   nook.Male,
		Name:     booneName}
)

var (
	booneAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "baboom"}

	booneCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ma noix"}

	booneDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "bavi-AAN"}

	booneFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ma noix"}

	booneGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "ugh"}

	booneItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "baboom"}

	booneJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ウルトラ"}

	booneKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "야임마"}

	booneLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "angaua"}

	booneRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бабум"}

	booneSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "这家伙"}

	booneSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "angaua"}

	booneTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "超級"}
)

var (
	boonePhrase = nook.Languages{
		language.AmericanEnglish:      booneAmericanEnglishName,
		language.CanadianFrench:       booneCanadianFrenchName,
		language.Dutch:                booneDutchName,
		language.French:               booneFrenchName,
		language.German:               booneGermanName,
		language.Italian:              booneItalianName,
		language.Japanese:             booneJapaneseName,
		language.Korean:               booneKoreanName,
		language.LatinAmericanSpanish: booneLatinAmericanSpanishName,
		language.Russian:              booneRussianName,
		language.SimplifiedChinese:    booneSimplifiedChineseName,
		language.Spanish:              booneSpanishName,
		language.TraditionalChinese:   booneTraditionalChineseName}
)

var (
	Boone = nook.Villager{
		Character:   booneCharacter,
		Personality: nook.Jock,
		Phrase:      boonePhrase}
)

package gorilla

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
		Value:    "Babouin"}

	booneDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Boone"}

	booneFrenchName = nook.Name{
		Language: language.French,
		Value:    "Babouin"}

	booneGermanName = nook.Name{
		Language: language.German,
		Value:    "Kong"}

	booneItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Babu"}

	booneJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "まんたろう"}

	booneKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "만복이"}

	booneLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Babú"}

	booneRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Буин"}

	booneSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "万泰"}

	booneSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Babú"}

	booneTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "萬泰"}
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
		Animal:   animal.Gorilla,
		Birthday: booneBirthday,
		Code:     booneCode,
		Key:      character.Boone,
		Gender:   gender.Male,
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
		Personality: personality.Jock,
		Phrase:      boonePhrase}
)

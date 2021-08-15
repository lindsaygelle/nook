package cat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	ankhaBirthday = nook.Birthday{
		Day:   22,
		Month: time.September}
)

var (
	ankhaCode = nook.Code{
		Value: "cat19"}
)

var (
	ankhaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Ankha"}

	ankhaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Nefertimi-i-i-ouh"}

	ankhaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Ankhasfinx"}

	ankhaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Nefertimi-i-i-ouh"}

	ankhaGermanName = nook.Name{
		Language: language.German,
		Value:    "Kleomi-miau"}

	ankhaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cleomiao miao"}

	ankhaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ナイルクフフ"}

	ankhaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "클레오파트라"}

	ankhaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Patrimarramius"}

	ankhaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Анками-мяу"}

	ankhaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "艳后尼罗"}

	ankhaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Patrimarramius"}

	ankhaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "艷后尼羅"}
)

var (
	ankhaName = nook.Languages{
		language.AmericanEnglish:      ankhaAmericanEnglishName,
		language.CanadianFrench:       ankhaCanadianFrenchName,
		language.Dutch:                ankhaDutchName,
		language.French:               ankhaFrenchName,
		language.German:               ankhaGermanName,
		language.Italian:              ankhaItalianName,
		language.Japanese:             ankhaJapaneseName,
		language.Korean:               ankhaKoreanName,
		language.LatinAmericanSpanish: ankhaLatinAmericanSpanishName,
		language.Russian:              ankhaRussianName,
		language.SimplifiedChinese:    ankhaSimplifiedChineseName,
		language.Spanish:              ankhaSpanishName,
		language.TraditionalChinese:   ankhaTraditionalChineseName}
)

var (
	ankhaCharacter = nook.Character{
		Animal:   Cat,
		Birthday: ankhaBirthday,
		Code:     ankhaCode,
		Gender:   nook.Female,
		Name:     ankhaName}
)

var (
	ankhaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "me meow"}

	ankhaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mi-i-i-ouh"}

	ankhaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "sfinx"}

	ankhaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mi-i-i-ouh"}

	ankhaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mi-miau"}

	ankhaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "miao miao"}

	ankhaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "クフフ"}

	ankhaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "파트라"}

	ankhaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "marramius"}

	ankhaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ми-мяу"}

	ankhaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "尼罗"}

	ankhaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "marramius"}

	ankhaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "尼羅"}
)

var (
	ankhaPhrase = nook.Languages{
		language.AmericanEnglish:      ankhaAmericanEnglishName,
		language.CanadianFrench:       ankhaCanadianFrenchName,
		language.Dutch:                ankhaDutchName,
		language.French:               ankhaFrenchName,
		language.German:               ankhaGermanName,
		language.Italian:              ankhaItalianName,
		language.Japanese:             ankhaJapaneseName,
		language.Korean:               ankhaKoreanName,
		language.LatinAmericanSpanish: ankhaLatinAmericanSpanishName,
		language.Russian:              ankhaRussianName,
		language.SimplifiedChinese:    ankhaSimplifiedChineseName,
		language.Spanish:              ankhaSpanishName,
		language.TraditionalChinese:   ankhaTraditionalChineseName}
)

var (
	Ankha = nook.Villager{
		Character:   ankhaCharacter,
		Personality: nook.Snooty,
		Phrase:      ankhaPhrase}
)

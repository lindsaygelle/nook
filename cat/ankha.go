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
		Value:    "Neferti"}

	ankhaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Ankha"}

	ankhaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Neferti"}

	ankhaGermanName = nook.Name{
		Language: language.German,
		Value:    "Kleo"}

	ankhaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cleo"}

	ankhaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ナイル"}

	ankhaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "클레오"}

	ankhaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Patri"}

	ankhaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Анка"}

	ankhaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "艳后"}

	ankhaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Patri"}

	ankhaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "艷后"}
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

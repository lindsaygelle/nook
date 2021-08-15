package deer

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	beauBirthday = nook.Birthday{
		Day:   5,
		Month: time.April}
)

var (
	beauCode = nook.Code{
		Value: "der07"}
)

var (
	beauAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Beau"}

	beauCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Stefaonfeufeuille"}

	beauDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Beauzoutsteen"}

	beauFrenchName = nook.Name{
		Language: language.French,
		Value:    "Stefaonfeufeuille"}

	beauGermanName = nook.Name{
		Language: language.German,
		Value:    "Martinknospel"}

	beauItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gigignam gnam"}

	beauJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ペーターおろおろ"}

	beauKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "피터우왕"}

	beauLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lopebabum"}

	beauRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бьюсоль-соль"}

	beauSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "彼得怎么办"}

	beauSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lopebabum"}

	beauTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "彼得怎麼辦"}
)

var (
	beauName = nook.Languages{
		language.AmericanEnglish:      beauAmericanEnglishName,
		language.CanadianFrench:       beauCanadianFrenchName,
		language.Dutch:                beauDutchName,
		language.French:               beauFrenchName,
		language.German:               beauGermanName,
		language.Italian:              beauItalianName,
		language.Japanese:             beauJapaneseName,
		language.Korean:               beauKoreanName,
		language.LatinAmericanSpanish: beauLatinAmericanSpanishName,
		language.Russian:              beauRussianName,
		language.SimplifiedChinese:    beauSimplifiedChineseName,
		language.Spanish:              beauSpanishName,
		language.TraditionalChinese:   beauTraditionalChineseName}
)

var (
	beauCharacter = nook.Character{
		Animal:   Deer,
		Birthday: beauBirthday,
		Code:     beauCode,
		Gender:   nook.Male,
		Name:     beauName}
)

var (
	beauAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "saltlick"}

	beauCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "feufeuille"}

	beauDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "zoutsteen"}

	beauFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "feufeuille"}

	beauGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "knospel"}

	beauItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "gnam gnam"}

	beauJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "おろおろ"}

	beauKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "우왕"}

	beauLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "babum"}

	beauRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "соль-соль"}

	beauSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "怎么办"}

	beauSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "babum"}

	beauTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "怎麼辦"}
)

var (
	beauPhrase = nook.Languages{
		language.AmericanEnglish:      beauAmericanEnglishName,
		language.CanadianFrench:       beauCanadianFrenchName,
		language.Dutch:                beauDutchName,
		language.French:               beauFrenchName,
		language.German:               beauGermanName,
		language.Italian:              beauItalianName,
		language.Japanese:             beauJapaneseName,
		language.Korean:               beauKoreanName,
		language.LatinAmericanSpanish: beauLatinAmericanSpanishName,
		language.Russian:              beauRussianName,
		language.SimplifiedChinese:    beauSimplifiedChineseName,
		language.Spanish:              beauSpanishName,
		language.TraditionalChinese:   beauTraditionalChineseName}
)

var (
	Beau = nook.Villager{
		Character:   beauCharacter,
		Personality: nook.Lazy,
		Phrase:      beauPhrase}
)

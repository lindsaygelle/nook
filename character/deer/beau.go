package deer

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
		Value:    "Stefaon"}

	beauDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Beau"}

	beauFrenchName = nook.Name{
		Language: language.French,
		Value:    "Stefaon"}

	beauGermanName = nook.Name{
		Language: language.German,
		Value:    "Martin"}

	beauItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gigi"}

	beauJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ペーター"}

	beauKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "피터"}

	beauLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lope"}

	beauRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бью"}

	beauSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "彼得"}

	beauSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lope"}

	beauTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "彼得"}
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
		Animal:   animal.Deer,
		Birthday: beauBirthday,
		Code:     beauCode,
		Key:      character.Beau,
		Gender:   gender.Male,
		Name:     beauName,
		Special:  false}
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
		language.AmericanEnglish:      beauAmericanEnglishPhrase,
		language.CanadianFrench:       beauCanadianFrenchPhrase,
		language.Dutch:                beauDutchPhrase,
		language.French:               beauFrenchPhrase,
		language.German:               beauGermanPhrase,
		language.Italian:              beauItalianPhrase,
		language.Japanese:             beauJapanesePhrase,
		language.Korean:               beauKoreanPhrase,
		language.LatinAmericanSpanish: beauLatinAmericanSpanishPhrase,
		language.Russian:              beauRussianPhrase,
		language.SimplifiedChinese:    beauSimplifiedChinesePhrase,
		language.Spanish:              beauSpanishPhrase,
		language.TraditionalChinese:   beauTraditionalChinesePhrase}
)

var (
	Beau = nook.Villager{
		Character:   beauCharacter,
		Personality: personality.Lazy,
		Phrase:      beauPhrase}
)

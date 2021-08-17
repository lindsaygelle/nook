package kangaroo

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
	kittBirthday = nook.Birthday{
		Day:   11,
		Month: time.October}
)

var (
	kittCode = nook.Code{
		Value: "kgr00"}
)

var (
	kittAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Kitt"}

	kittCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Poquette"}

	kittDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Kitt"}

	kittFrenchName = nook.Name{
		Language: language.French,
		Value:    "Poquette"}

	kittGermanName = nook.Name{
		Language: language.German,
		Value:    "Kerstin"}

	kittItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Kitt"}

	kittJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アップリケ"}

	kittKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "애플리케"}

	kittLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Antípoda"}

	kittRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Китт"}

	kittSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "缝缝"}

	kittSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Antípoda"}

	kittTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "縫縫"}
)

var (
	kittName = nook.Languages{
		language.AmericanEnglish:      kittAmericanEnglishName,
		language.CanadianFrench:       kittCanadianFrenchName,
		language.Dutch:                kittDutchName,
		language.French:               kittFrenchName,
		language.German:               kittGermanName,
		language.Italian:              kittItalianName,
		language.Japanese:             kittJapaneseName,
		language.Korean:               kittKoreanName,
		language.LatinAmericanSpanish: kittLatinAmericanSpanishName,
		language.Russian:              kittRussianName,
		language.SimplifiedChinese:    kittSimplifiedChineseName,
		language.Spanish:              kittSpanishName,
		language.TraditionalChinese:   kittTraditionalChineseName}
)

var (
	kittCharacter = nook.Character{
		Animal:   animal.Kangaroo,
		Birthday: kittBirthday,
		Code:     kittCode,
		Key:      character.Kitt,
		Gender:   gender.Female,
		Name:     kittName}
)

var (
	kittAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "child"}

	kittCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "gamin"}

	kittDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kind"}

	kittFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "gamin"}

	kittGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kindchen"}

	kittItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cucciolo"}

	kittJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ッポン"}

	kittKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "폴짝"}

	kittLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "chiquichí"}

	kittRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "детка"}

	kittSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "补补"}

	kittSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "mi alma"}

	kittTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "補補"}
)

var (
	kittPhrase = nook.Languages{
		language.AmericanEnglish:      kittAmericanEnglishPhrase,
		language.CanadianFrench:       kittCanadianFrenchPhrase,
		language.Dutch:                kittDutchPhrase,
		language.French:               kittFrenchPhrase,
		language.German:               kittGermanPhrase,
		language.Italian:              kittItalianPhrase,
		language.Japanese:             kittJapanesePhrase,
		language.Korean:               kittKoreanPhrase,
		language.LatinAmericanSpanish: kittLatinAmericanSpanishPhrase,
		language.Russian:              kittRussianPhrase,
		language.SimplifiedChinese:    kittSimplifiedChinesePhrase,
		language.Spanish:              kittSpanishPhrase,
		language.TraditionalChinese:   kittTraditionalChinesePhrase}
)

var (
	Kitt = nook.Villager{
		Character:   kittCharacter,
		Personality: personality.Normal,
		Phrase:      kittPhrase}
)

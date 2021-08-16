package deer

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	erikBirthday = nook.Birthday{
		Day:   27,
		Month: time.July}
)

var (
	erikCode = nook.Code{
		Value: "der09"}
)

var (
	erikAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Erik"}

	erikCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Abraham"}

	erikDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Erik"}

	erikFrenchName = nook.Name{
		Language: language.French,
		Value:    "Abraham"}

	erikGermanName = nook.Name{
		Language: language.German,
		Value:    "Erik"}

	erikItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cervasio"}

	erikJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チャック"}

	erikKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "자끄"}

	erikLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cervasio"}

	erikRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Эрик"}

	erikSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "查克"}

	erikSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cervasio"}

	erikTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "查克"}
)

var (
	erikName = nook.Languages{
		language.AmericanEnglish:      erikAmericanEnglishName,
		language.CanadianFrench:       erikCanadianFrenchName,
		language.Dutch:                erikDutchName,
		language.French:               erikFrenchName,
		language.German:               erikGermanName,
		language.Italian:              erikItalianName,
		language.Japanese:             erikJapaneseName,
		language.Korean:               erikKoreanName,
		language.LatinAmericanSpanish: erikLatinAmericanSpanishName,
		language.Russian:              erikRussianName,
		language.SimplifiedChinese:    erikSimplifiedChineseName,
		language.Spanish:              erikSpanishName,
		language.TraditionalChinese:   erikTraditionalChineseName}
)

var (
	erikCharacter = nook.Character{
		Animal:   animal.Deer,
		Birthday: erikBirthday,
		Code:     erikCode,
		Gender:   gender.Male,
		Name:     erikName}
)

var (
	erikAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "chow down"}

	erikCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "breuh"}

	erikDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "schorsie"}

	erikFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "breuh"}

	erikGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hirschl"}

	erikItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "relax"}

	erikJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "しかじか"}

	erikKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "뚜비뚜바"}

	erikLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ñume-ñume"}

	erikRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ем мох"}

	erikSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "鹿鹿"}

	erikSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ñume-ñume"}

	erikTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鹿鹿"}
)

var (
	erikPhrase = nook.Languages{
		language.AmericanEnglish:      erikAmericanEnglishName,
		language.CanadianFrench:       erikCanadianFrenchName,
		language.Dutch:                erikDutchName,
		language.French:               erikFrenchName,
		language.German:               erikGermanName,
		language.Italian:              erikItalianName,
		language.Japanese:             erikJapaneseName,
		language.Korean:               erikKoreanName,
		language.LatinAmericanSpanish: erikLatinAmericanSpanishName,
		language.Russian:              erikRussianName,
		language.SimplifiedChinese:    erikSimplifiedChineseName,
		language.Spanish:              erikSpanishName,
		language.TraditionalChinese:   erikTraditionalChineseName}
)

var (
	Erik = nook.Villager{
		Character:   erikCharacter,
		Personality: personality.Lazy,
		Phrase:      erikPhrase}
)

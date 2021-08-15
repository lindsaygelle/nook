package deer

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Abrahambreuh"}

	erikDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Erikschorsie"}

	erikFrenchName = nook.Name{
		Language: language.French,
		Value:    "Abrahambreuh"}

	erikGermanName = nook.Name{
		Language: language.German,
		Value:    "Erikhirschl"}

	erikItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cervasiorelax"}

	erikJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チャックしかじか"}

	erikKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "자끄뚜비뚜바"}

	erikLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cervasioñume-ñume"}

	erikRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Эрикем мох"}

	erikSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "查克鹿鹿"}

	erikSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cervasioñume-ñume"}

	erikTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "查克鹿鹿"}
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
		Animal:   Deer,
		Birthday: erikBirthday,
		Code:     erikCode,
		Gender:   nook.Male,
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
		Personality: nook.Lazy,
		Phrase:      erikPhrase}
)

package rabbit

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	ohareBirthday = nook.Birthday{
		Day:   24,
		Month: time.July}
)

var (
	ohareCode = nook.Code{
		Value: "rbt15"}
)

var (
	ohareAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "O'Hare"}

	ohareCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Civetamigo"}

	ohareDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "O'Hareamigo"}

	ohareFrenchName = nook.Name{
		Language: language.French,
		Value:    "Civetpécaïre"}

	ohareGermanName = nook.Name{
		Language: language.German,
		Value:    "Nicoamigo"}

	ohareItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Fritzaquesí"}

	ohareJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "サントスアミーゴ"}

	ohareKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "산토스아미고"}

	ohareLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ernestotralarí"}

	ohareRussianName = nook.Name{
		Language: language.Russian,
		Value:    "О'Хэарамиго"}

	ohareSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "桑托斯朋友"}

	ohareSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ernestotralarí"}

	ohareTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "桑托斯朋友"}
)

var (
	ohareName = nook.Languages{
		language.AmericanEnglish:      ohareAmericanEnglishName,
		language.CanadianFrench:       ohareCanadianFrenchName,
		language.Dutch:                ohareDutchName,
		language.French:               ohareFrenchName,
		language.German:               ohareGermanName,
		language.Italian:              ohareItalianName,
		language.Japanese:             ohareJapaneseName,
		language.Korean:               ohareKoreanName,
		language.LatinAmericanSpanish: ohareLatinAmericanSpanishName,
		language.Russian:              ohareRussianName,
		language.SimplifiedChinese:    ohareSimplifiedChineseName,
		language.Spanish:              ohareSpanishName,
		language.TraditionalChinese:   ohareTraditionalChineseName}
)

var (
	ohareCharacter = nook.Character{
		Animal:   Rabbit,
		Birthday: ohareBirthday,
		Code:     ohareCode,
		Gender:   nook.Male,
		Name:     ohareName}
)

var (
	ohareAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "amigo"}

	ohareCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "amigo"}

	ohareDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "amigo"}

	ohareFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pécaïre"}

	ohareGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "amigo"}

	ohareItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "aquesí"}

	ohareJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "アミーゴ"}

	ohareKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아미고"}

	ohareLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "tralarí"}

	ohareRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "амиго"}

	ohareSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "朋友"}

	ohareSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tralarí"}

	ohareTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "朋友"}
)

var (
	oharePhrase = nook.Languages{
		language.AmericanEnglish:      ohareAmericanEnglishName,
		language.CanadianFrench:       ohareCanadianFrenchName,
		language.Dutch:                ohareDutchName,
		language.French:               ohareFrenchName,
		language.German:               ohareGermanName,
		language.Italian:              ohareItalianName,
		language.Japanese:             ohareJapaneseName,
		language.Korean:               ohareKoreanName,
		language.LatinAmericanSpanish: ohareLatinAmericanSpanishName,
		language.Russian:              ohareRussianName,
		language.SimplifiedChinese:    ohareSimplifiedChineseName,
		language.Spanish:              ohareSpanishName,
		language.TraditionalChinese:   ohareTraditionalChineseName}
)

var (
	OHare = nook.Villager{
		Character:   ohareCharacter,
		Personality: nook.Smug,
		Phrase:      oharePhrase}
)

package wolf

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	kyleBirthday = nook.Birthday{
		Day:   6,
		Month: time.December}
)

var (
	kyleCode = nook.Code{
		Value: "wol10"}
)

var (
	kyleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Kyle"}

	kyleCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Garygarooouuu"}

	kyleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Kylealfa"}

	kyleFrenchName = nook.Name{
		Language: language.French,
		Value:    "Garygarooouuu"}

	kyleGermanName = nook.Name{
		Language: language.German,
		Value:    "Wolfgangahuuu"}

	kyleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ululosgnack"}

	kyleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "リカルドオゥイェ"}

	kyleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "리카르도오우예"}

	kyleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ataúlfoaujujú"}

	kyleRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кайлвожак"}

	kyleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "李可喔耶"}

	kyleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ataúlfoaujujú"}

	kyleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "李可喔耶"}
)

var (
	kyleName = nook.Languages{
		language.AmericanEnglish:      kyleAmericanEnglishName,
		language.CanadianFrench:       kyleCanadianFrenchName,
		language.Dutch:                kyleDutchName,
		language.French:               kyleFrenchName,
		language.German:               kyleGermanName,
		language.Italian:              kyleItalianName,
		language.Japanese:             kyleJapaneseName,
		language.Korean:               kyleKoreanName,
		language.LatinAmericanSpanish: kyleLatinAmericanSpanishName,
		language.Russian:              kyleRussianName,
		language.SimplifiedChinese:    kyleSimplifiedChineseName,
		language.Spanish:              kyleSpanishName,
		language.TraditionalChinese:   kyleTraditionalChineseName}
)

var (
	kyleCharacter = nook.Character{
		Animal:   Wolf,
		Birthday: kyleBirthday,
		Code:     kyleCode,
		Gender:   nook.Male,
		Name:     kyleName}
)

var (
	kyleAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "alpha"}

	kyleCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "garooouuu"}

	kyleDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "alfa"}

	kyleFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "garooouuu"}

	kyleGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "ahuuu"}

	kyleItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sgnack"}

	kyleJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "オゥイェ"}

	kyleKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "오우예"}

	kyleLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "aujujú"}

	kyleRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "вожак"}

	kyleSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喔耶"}

	kyleSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "aujujú"}

	kyleTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "喔耶"}
)

var (
	kylePhrase = nook.Languages{
		language.AmericanEnglish:      kyleAmericanEnglishName,
		language.CanadianFrench:       kyleCanadianFrenchName,
		language.Dutch:                kyleDutchName,
		language.French:               kyleFrenchName,
		language.German:               kyleGermanName,
		language.Italian:              kyleItalianName,
		language.Japanese:             kyleJapaneseName,
		language.Korean:               kyleKoreanName,
		language.LatinAmericanSpanish: kyleLatinAmericanSpanishName,
		language.Russian:              kyleRussianName,
		language.SimplifiedChinese:    kyleSimplifiedChineseName,
		language.Spanish:              kyleSpanishName,
		language.TraditionalChinese:   kyleTraditionalChineseName}
)

var (
	Kyle = nook.Villager{
		Character:   kyleCharacter,
		Personality: nook.Smug,
		Phrase:      kylePhrase}
)

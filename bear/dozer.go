package bear

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	dozerBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	dozerCode = nook.Code{
		Value: ""}
)

var (
	dozerAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Dozer"}

	dozerCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	dozerDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	dozerFrenchName = nook.Name{
		Language: language.French,
		Value:    "Doudou"}

	dozerGermanName = nook.Name{
		Language: language.German,
		Value:    "Nicki"}

	dozerItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ronfo"}

	dozerJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "スリープ"}

	dozerKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	dozerLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	dozerRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	dozerSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "忪忪"}

	dozerSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Toño"}

	dozerTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	dozerName = nook.Languages{
		language.AmericanEnglish:      dozerAmericanEnglishName,
		language.CanadianFrench:       dozerCanadianFrenchName,
		language.Dutch:                dozerDutchName,
		language.French:               dozerFrenchName,
		language.German:               dozerGermanName,
		language.Italian:              dozerItalianName,
		language.Japanese:             dozerJapaneseName,
		language.Korean:               dozerKoreanName,
		language.LatinAmericanSpanish: dozerLatinAmericanSpanishName,
		language.Russian:              dozerRussianName,
		language.SimplifiedChinese:    dozerSimplifiedChineseName,
		language.Spanish:              dozerSpanishName,
		language.TraditionalChinese:   dozerTraditionalChineseName}
)

var (
	dozerCharacter = nook.Character{
		Animal:   Bear,
		Birthday: dozerBirthday,
		Code:     dozerCode,
		Gender:   nook.Male,
		Name:     dozerName}
)

var (
	dozerAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	dozerCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	dozerDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	dozerFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "zzzzzz"}

	dozerGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnarch"}

	dozerItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "zzzzzz"}

	dozerJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でアリ"}

	dozerKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	dozerLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	dozerRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	dozerSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咕"}

	dozerSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "zzzzzz"}

	dozerTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	dozerPhrase = nook.Languages{
		language.AmericanEnglish:      dozerAmericanEnglishName,
		language.CanadianFrench:       dozerCanadianFrenchName,
		language.Dutch:                dozerDutchName,
		language.French:               dozerFrenchName,
		language.German:               dozerGermanName,
		language.Italian:              dozerItalianName,
		language.Japanese:             dozerJapaneseName,
		language.Korean:               dozerKoreanName,
		language.LatinAmericanSpanish: dozerLatinAmericanSpanishName,
		language.Russian:              dozerRussianName,
		language.SimplifiedChinese:    dozerSimplifiedChineseName,
		language.Spanish:              dozerSpanishName,
		language.TraditionalChinese:   dozerTraditionalChineseName}
)

var (
	Dozer = nook.Villager{
		Character:   dozerCharacter,
		Personality: nook.Lazy,
		Phrase:      dozerPhrase}
)

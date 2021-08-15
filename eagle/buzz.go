package eagle

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	buzzBirthday = nook.Birthday{
		Day:   7,
		Month: time.December}
)

var (
	buzzCode = nook.Code{
		Value: "pbr03"}
)

var (
	buzzAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Buzz"}

	buzzCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Phébusflap flap"}

	buzzDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Buzzkaptein"}

	buzzFrenchName = nook.Name{
		Language: language.French,
		Value:    "Phébusflap flap"}

	buzzGermanName = nook.Name{
		Language: language.German,
		Value:    "Balduinkäpten"}

	buzzItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Goliacapo"}

	buzzJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ひでよしッキーッ"}

	buzzKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "근엄짜란～"}

	buzzLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Nabarcapííí"}

	buzzRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Баззкапитан"}

	buzzSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "继光继继"}

	buzzSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Nabarcadete"}

	buzzTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "繼光繼繼"}
)

var (
	buzzName = nook.Languages{
		language.AmericanEnglish:      buzzAmericanEnglishName,
		language.CanadianFrench:       buzzCanadianFrenchName,
		language.Dutch:                buzzDutchName,
		language.French:               buzzFrenchName,
		language.German:               buzzGermanName,
		language.Italian:              buzzItalianName,
		language.Japanese:             buzzJapaneseName,
		language.Korean:               buzzKoreanName,
		language.LatinAmericanSpanish: buzzLatinAmericanSpanishName,
		language.Russian:              buzzRussianName,
		language.SimplifiedChinese:    buzzSimplifiedChineseName,
		language.Spanish:              buzzSpanishName,
		language.TraditionalChinese:   buzzTraditionalChineseName}
)

var (
	buzzCharacter = nook.Character{
		Animal:   Eagle,
		Birthday: buzzBirthday,
		Code:     buzzCode,
		Gender:   nook.Male,
		Name:     buzzName}
)

var (
	buzzAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "captain"}

	buzzCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "flap flap"}

	buzzDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kaptein"}

	buzzFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "flap flap"}

	buzzGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "käpten"}

	buzzItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "capo"}

	buzzJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ッキーッ"}

	buzzKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "짜란～"}

	buzzLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "capííí"}

	buzzRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "капитан"}

	buzzSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "继继"}

	buzzSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cadete"}

	buzzTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "繼繼"}
)

var (
	buzzPhrase = nook.Languages{
		language.AmericanEnglish:      buzzAmericanEnglishName,
		language.CanadianFrench:       buzzCanadianFrenchName,
		language.Dutch:                buzzDutchName,
		language.French:               buzzFrenchName,
		language.German:               buzzGermanName,
		language.Italian:              buzzItalianName,
		language.Japanese:             buzzJapaneseName,
		language.Korean:               buzzKoreanName,
		language.LatinAmericanSpanish: buzzLatinAmericanSpanishName,
		language.Russian:              buzzRussianName,
		language.SimplifiedChinese:    buzzSimplifiedChineseName,
		language.Spanish:              buzzSpanishName,
		language.TraditionalChinese:   buzzTraditionalChineseName}
)

var (
	Buzz = nook.Villager{
		Character:   buzzCharacter,
		Personality: nook.Cranky,
		Phrase:      buzzPhrase}
)

package cat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	raymondBirthday = nook.Birthday{
		Day:   1,
		Month: time.October}
)

var (
	raymondCode = nook.Code{
		Value: "cat23"}
)

var (
	raymondAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Raymond"}

	raymondCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Raymond"}

	raymondDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Raymond"}

	raymondFrenchName = nook.Name{
		Language: language.French,
		Value:    "Raymond"}

	raymondGermanName = nook.Name{
		Language: language.German,
		Value:    "Gunnar"}

	raymondItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Raimondo"}

	raymondJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジャック"}

	raymondKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "잭슨"}

	raymondLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Narciso"}

	raymondRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Реймонд"}

	raymondSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "杰克"}

	raymondSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Narciso"}

	raymondTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "傑克"}
)

var (
	raymondName = nook.Languages{
		language.AmericanEnglish:      raymondAmericanEnglishName,
		language.CanadianFrench:       raymondCanadianFrenchName,
		language.Dutch:                raymondDutchName,
		language.French:               raymondFrenchName,
		language.German:               raymondGermanName,
		language.Italian:              raymondItalianName,
		language.Japanese:             raymondJapaneseName,
		language.Korean:               raymondKoreanName,
		language.LatinAmericanSpanish: raymondLatinAmericanSpanishName,
		language.Russian:              raymondRussianName,
		language.SimplifiedChinese:    raymondSimplifiedChineseName,
		language.Spanish:              raymondSpanishName,
		language.TraditionalChinese:   raymondTraditionalChineseName}
)

var (
	raymondCharacter = nook.Character{
		Animal:   Cat,
		Birthday: raymondBirthday,
		Code:     raymondCode,
		Gender:   nook.Male,
		Name:     raymondName}
)

var (
	raymondAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "crisp"}

	raymondCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "panache"}

	raymondDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "netjes"}

	raymondFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "panache"}

	raymondGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "weeßte"}

	raymondItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "rrricooo"}

	raymondJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "キリッ"}

	raymondKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "우쭐"}

	raymondLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "atilda"}

	raymondRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "стильненько"}

	raymondSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "严肃"}

	raymondSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "atilda"}

	raymondTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嚴肅"}
)

var (
	raymondPhrase = nook.Languages{
		language.AmericanEnglish:      raymondAmericanEnglishName,
		language.CanadianFrench:       raymondCanadianFrenchName,
		language.Dutch:                raymondDutchName,
		language.French:               raymondFrenchName,
		language.German:               raymondGermanName,
		language.Italian:              raymondItalianName,
		language.Japanese:             raymondJapaneseName,
		language.Korean:               raymondKoreanName,
		language.LatinAmericanSpanish: raymondLatinAmericanSpanishName,
		language.Russian:              raymondRussianName,
		language.SimplifiedChinese:    raymondSimplifiedChineseName,
		language.Spanish:              raymondSpanishName,
		language.TraditionalChinese:   raymondTraditionalChineseName}
)

var (
	Raymond = nook.Villager{
		Character:   raymondCharacter,
		Personality: nook.Smug,
		Phrase:      raymondPhrase}
)

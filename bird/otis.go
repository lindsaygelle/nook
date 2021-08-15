package bird

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	otisBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	otisCode = nook.Code{
		Value: ""}
)

var (
	otisAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Otis"}

	otisCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	otisDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	otisFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ottisj'pense"}

	otisGermanName = nook.Name{
		Language: language.German,
		Value:    "Ottodenk ich"}

	otisItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pablotuitcì"}

	otisJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "たくあんなのじゃ"}

	otisKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	otisLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	otisRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	otisSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "长老神说"}

	otisSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Otodigo"}

	otisTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	otisName = nook.Languages{
		language.AmericanEnglish:      otisAmericanEnglishName,
		language.CanadianFrench:       otisCanadianFrenchName,
		language.Dutch:                otisDutchName,
		language.French:               otisFrenchName,
		language.German:               otisGermanName,
		language.Italian:              otisItalianName,
		language.Japanese:             otisJapaneseName,
		language.Korean:               otisKoreanName,
		language.LatinAmericanSpanish: otisLatinAmericanSpanishName,
		language.Russian:              otisRussianName,
		language.SimplifiedChinese:    otisSimplifiedChineseName,
		language.Spanish:              otisSpanishName,
		language.TraditionalChinese:   otisTraditionalChineseName}
)

var (
	otisCharacter = nook.Character{
		Animal:   Bird,
		Birthday: otisBirthday,
		Code:     otisCode,
		Gender:   nook.Male,
		Name:     otisName}
)

var (
	otisAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	otisCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	otisDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	otisFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "j'pense"}

	otisGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "denk ich"}

	otisItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "tuitcì"}

	otisJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なのじゃ"}

	otisKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	otisLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	otisRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	otisSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "神说"}

	otisSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "digo"}

	otisTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	otisPhrase = nook.Languages{
		language.AmericanEnglish:      otisAmericanEnglishName,
		language.CanadianFrench:       otisCanadianFrenchName,
		language.Dutch:                otisDutchName,
		language.French:               otisFrenchName,
		language.German:               otisGermanName,
		language.Italian:              otisItalianName,
		language.Japanese:             otisJapaneseName,
		language.Korean:               otisKoreanName,
		language.LatinAmericanSpanish: otisLatinAmericanSpanishName,
		language.Russian:              otisRussianName,
		language.SimplifiedChinese:    otisSimplifiedChineseName,
		language.Spanish:              otisSpanishName,
		language.TraditionalChinese:   otisTraditionalChineseName}
)

var (
	Otis = nook.Villager{
		Character:   otisCharacter,
		Personality: nook.Lazy,
		Phrase:      otisPhrase}
)

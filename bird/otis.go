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
		Value:    "N/A"}

	otisDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	otisFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ottis"}

	otisGermanName = nook.Name{
		Language: language.German,
		Value:    "Otto"}

	otisItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pablo"}

	otisJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "たくあん"}

	otisKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	otisLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	otisRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	otisSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "长老"}

	otisSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Oto"}

	otisTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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

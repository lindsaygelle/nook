package cat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	merryBirthday = nook.Birthday{
		Day:   29,
		Month: time.June}
)

var (
	merryCode = nook.Code{
		Value: "cat16"}
)

var (
	merryAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Merry"}

	merryCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Suzy"}

	merryDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Merry"}

	merryFrenchName = nook.Name{
		Language: language.French,
		Value:    "Suzy"}

	merryGermanName = nook.Name{
		Language: language.German,
		Value:    "Mischka"}

	merryItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Katy"}

	merryJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "さっち"}

	merryKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "유네찌"}

	merryLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Susi"}

	merryRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Мерри"}

	merrySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "莎莎"}

	merrySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Susi"}

	merryTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "莎莎"}
)

var (
	merryName = nook.Languages{
		language.AmericanEnglish:      merryAmericanEnglishName,
		language.CanadianFrench:       merryCanadianFrenchName,
		language.Dutch:                merryDutchName,
		language.French:               merryFrenchName,
		language.German:               merryGermanName,
		language.Italian:              merryItalianName,
		language.Japanese:             merryJapaneseName,
		language.Korean:               merryKoreanName,
		language.LatinAmericanSpanish: merryLatinAmericanSpanishName,
		language.Russian:              merryRussianName,
		language.SimplifiedChinese:    merrySimplifiedChineseName,
		language.Spanish:              merrySpanishName,
		language.TraditionalChinese:   merryTraditionalChineseName}
)

var (
	merryCharacter = nook.Character{
		Animal:   Cat,
		Birthday: merryBirthday,
		Code:     merryCode,
		Gender:   nook.Female,
		Name:     merryName}
)

var (
	merryAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "mweee"}

	merryCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "miou-mi"}

	merryDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "mi-jippie"}

	merryFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "miou-mi"}

	merryGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "prrrrr"}

	merryItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "pffffft"}

	merryJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "にゃん"}

	merryKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "냥냥"}

	merryLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "michifú"}

	merryRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "мяуи-и-и"}

	merrySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喵嗯"}

	merrySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "michifú"}

	merryTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "喵嗯"}
)

var (
	merryPhrase = nook.Languages{
		language.AmericanEnglish:      merryAmericanEnglishName,
		language.CanadianFrench:       merryCanadianFrenchName,
		language.Dutch:                merryDutchName,
		language.French:               merryFrenchName,
		language.German:               merryGermanName,
		language.Italian:              merryItalianName,
		language.Japanese:             merryJapaneseName,
		language.Korean:               merryKoreanName,
		language.LatinAmericanSpanish: merryLatinAmericanSpanishName,
		language.Russian:              merryRussianName,
		language.SimplifiedChinese:    merrySimplifiedChineseName,
		language.Spanish:              merrySpanishName,
		language.TraditionalChinese:   merryTraditionalChineseName}
)

var (
	Merry = nook.Villager{
		Character:   merryCharacter,
		Personality: nook.Peppy,
		Phrase:      merryPhrase}
)

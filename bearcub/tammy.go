package bearcub

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	tammyBirthday = nook.Birthday{
		Day:   23,
		Month: time.June}
)

var (
	tammyCode = nook.Code{
		Value: "cbr17"}
)

var (
	tammyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tammy"}

	tammyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Vanessagrifouille"}

	tammyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tammyweet je"}

	tammyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Vanessagrifouille"}

	tammyGermanName = nook.Name{
		Language: language.German,
		Value:    "Tatjanabossi"}

	tammyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Tamaraohè"}

	tammyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アネッサだってサ"}

	tammyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아네사차라리"}

	tammyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Aídatararí"}

	tammyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Таммипонимаешь"}

	tammySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "然姐竟然说"}

	tammySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Aídatararí"}

	tammyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "然姐竟然說"}
)

var (
	tammyName = nook.Languages{
		language.AmericanEnglish:      tammyAmericanEnglishName,
		language.CanadianFrench:       tammyCanadianFrenchName,
		language.Dutch:                tammyDutchName,
		language.French:               tammyFrenchName,
		language.German:               tammyGermanName,
		language.Italian:              tammyItalianName,
		language.Japanese:             tammyJapaneseName,
		language.Korean:               tammyKoreanName,
		language.LatinAmericanSpanish: tammyLatinAmericanSpanishName,
		language.Russian:              tammyRussianName,
		language.SimplifiedChinese:    tammySimplifiedChineseName,
		language.Spanish:              tammySpanishName,
		language.TraditionalChinese:   tammyTraditionalChineseName}
)

var (
	tammyCharacter = nook.Character{
		Animal:   Bearcub,
		Birthday: tammyBirthday,
		Code:     tammyCode,
		Gender:   nook.Female,
		Name:     tammyName}
)

var (
	tammyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ya heard"}

	tammyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "grifouille"}

	tammyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "weet je"}

	tammyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "grifouille"}

	tammyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bossi"}

	tammyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ohè"}

	tammyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だってサ"}

	tammyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "차라리"}

	tammyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "tararí"}

	tammyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "понимаешь"}

	tammySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "竟然说"}

	tammySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tararí"}

	tammyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "竟然說"}
)

var (
	tammyPhrase = nook.Languages{
		language.AmericanEnglish:      tammyAmericanEnglishName,
		language.CanadianFrench:       tammyCanadianFrenchName,
		language.Dutch:                tammyDutchName,
		language.French:               tammyFrenchName,
		language.German:               tammyGermanName,
		language.Italian:              tammyItalianName,
		language.Japanese:             tammyJapaneseName,
		language.Korean:               tammyKoreanName,
		language.LatinAmericanSpanish: tammyLatinAmericanSpanishName,
		language.Russian:              tammyRussianName,
		language.SimplifiedChinese:    tammySimplifiedChineseName,
		language.Spanish:              tammySpanishName,
		language.TraditionalChinese:   tammyTraditionalChineseName}
)

var (
	Tammy = nook.Villager{
		Character:   tammyCharacter,
		Personality: nook.BigSister,
		Phrase:      tammyPhrase}
)

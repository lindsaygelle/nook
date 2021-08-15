package mouse

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	flossieBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	flossieCode = nook.Code{
		Value: ""}
)

var (
	flossieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Flossie"}

	flossieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	flossieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	flossieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Chiquita"}

	flossieGermanName = nook.Name{
		Language: language.German,
		Value:    "Amalie"}

	flossieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Florinda"}

	flossieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ピスタチオ"}

	flossieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	flossieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	flossieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	flossieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	flossieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Oregilda"}

	flossieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	flossieName = nook.Languages{
		language.AmericanEnglish:      flossieAmericanEnglishName,
		language.CanadianFrench:       flossieCanadianFrenchName,
		language.Dutch:                flossieDutchName,
		language.French:               flossieFrenchName,
		language.German:               flossieGermanName,
		language.Italian:              flossieItalianName,
		language.Japanese:             flossieJapaneseName,
		language.Korean:               flossieKoreanName,
		language.LatinAmericanSpanish: flossieLatinAmericanSpanishName,
		language.Russian:              flossieRussianName,
		language.SimplifiedChinese:    flossieSimplifiedChineseName,
		language.Spanish:              flossieSpanishName,
		language.TraditionalChinese:   flossieTraditionalChineseName}
)

var (
	flossieCharacter = nook.Character{
		Animal:   Mouse,
		Birthday: flossieBirthday,
		Code:     flossieCode,
		Gender:   nook.Female,
		Name:     flossieName}
)

var (
	flossieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	flossieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	flossieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	flossieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "microbe"}

	flossieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "quietsch"}

	flossieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "fischietto"}

	flossieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ガハハ"}

	flossieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	flossieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	flossieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	flossieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	flossieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "o sea"}

	flossieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	flossiePhrase = nook.Languages{
		language.AmericanEnglish:      flossieAmericanEnglishName,
		language.CanadianFrench:       flossieCanadianFrenchName,
		language.Dutch:                flossieDutchName,
		language.French:               flossieFrenchName,
		language.German:               flossieGermanName,
		language.Italian:              flossieItalianName,
		language.Japanese:             flossieJapaneseName,
		language.Korean:               flossieKoreanName,
		language.LatinAmericanSpanish: flossieLatinAmericanSpanishName,
		language.Russian:              flossieRussianName,
		language.SimplifiedChinese:    flossieSimplifiedChineseName,
		language.Spanish:              flossieSpanishName,
		language.TraditionalChinese:   flossieTraditionalChineseName}
)

var (
	Flossie = nook.Villager{
		Character:   flossieCharacter,
		Personality: nook.Peppy,
		Phrase:      flossiePhrase}
)

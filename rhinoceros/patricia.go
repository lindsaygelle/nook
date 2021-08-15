package rhinoceros

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	patriciaBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	patriciaCode = nook.Code{
		Value: ""}
)

var (
	patriciaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Patricia"}

	patriciaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	patriciaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	patriciaFrenchName = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	patriciaGermanName = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	patriciaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	patriciaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "パトリシア"}

	patriciaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	patriciaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	patriciaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	patriciaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	patriciaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	patriciaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	patriciaName = nook.Languages{
		language.AmericanEnglish:      patriciaAmericanEnglishName,
		language.CanadianFrench:       patriciaCanadianFrenchName,
		language.Dutch:                patriciaDutchName,
		language.French:               patriciaFrenchName,
		language.German:               patriciaGermanName,
		language.Italian:              patriciaItalianName,
		language.Japanese:             patriciaJapaneseName,
		language.Korean:               patriciaKoreanName,
		language.LatinAmericanSpanish: patriciaLatinAmericanSpanishName,
		language.Russian:              patriciaRussianName,
		language.SimplifiedChinese:    patriciaSimplifiedChineseName,
		language.Spanish:              patriciaSpanishName,
		language.TraditionalChinese:   patriciaTraditionalChineseName}
)

var (
	patriciaCharacter = nook.Character{
		Animal:   Rhinoceros,
		Birthday: patriciaBirthday,
		Code:     patriciaCode,
		Gender:   nook.Female,
		Name:     patriciaName}
)

var (
	patriciaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	patriciaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	patriciaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	patriciaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	patriciaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	patriciaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	patriciaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "えへっ"}

	patriciaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	patriciaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	patriciaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	patriciaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	patriciaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	patriciaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	patriciaPhrase = nook.Languages{
		language.AmericanEnglish:      patriciaAmericanEnglishName,
		language.CanadianFrench:       patriciaCanadianFrenchName,
		language.Dutch:                patriciaDutchName,
		language.French:               patriciaFrenchName,
		language.German:               patriciaGermanName,
		language.Italian:              patriciaItalianName,
		language.Japanese:             patriciaJapaneseName,
		language.Korean:               patriciaKoreanName,
		language.LatinAmericanSpanish: patriciaLatinAmericanSpanishName,
		language.Russian:              patriciaRussianName,
		language.SimplifiedChinese:    patriciaSimplifiedChineseName,
		language.Spanish:              patriciaSpanishName,
		language.TraditionalChinese:   patriciaTraditionalChineseName}
)

var (
	Patricia = nook.Villager{
		Character:   patriciaCharacter,
		Personality: nook.Normal,
		Phrase:      patriciaPhrase}
)

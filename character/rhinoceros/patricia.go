package rhinoceros

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
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
		Value:    ""}

	patriciaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	patriciaFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	patriciaGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	patriciaItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	patriciaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "パトリシア"}

	patriciaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	patriciaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	patriciaRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	patriciaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	patriciaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	patriciaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
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
		Animal:   animal.Rhinoceros,
		Birthday: patriciaBirthday,
		Code:     patriciaCode,
		Key:      character.Patricia,
		Gender:   gender.Female,
		Name:     patriciaName,
		Special:  false}
)

var (
	patriciaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "えへっ"}

	patriciaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	patriciaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	patriciaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	patriciaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	patriciaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	patriciaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "えへっ"}

	patriciaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	patriciaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	patriciaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	patriciaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	patriciaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	patriciaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	patriciaPhrase = nook.Languages{
		language.AmericanEnglish:      patriciaAmericanEnglishPhrase,
		language.CanadianFrench:       patriciaCanadianFrenchPhrase,
		language.Dutch:                patriciaDutchPhrase,
		language.French:               patriciaFrenchPhrase,
		language.German:               patriciaGermanPhrase,
		language.Italian:              patriciaItalianPhrase,
		language.Japanese:             patriciaJapanesePhrase,
		language.Korean:               patriciaKoreanPhrase,
		language.LatinAmericanSpanish: patriciaLatinAmericanSpanishPhrase,
		language.Russian:              patriciaRussianPhrase,
		language.SimplifiedChinese:    patriciaSimplifiedChinesePhrase,
		language.Spanish:              patriciaSpanishPhrase,
		language.TraditionalChinese:   patriciaTraditionalChinesePhrase}
)

var (
	Patricia = nook.Villager{
		Character:   patriciaCharacter,
		Personality: personality.Normal,
		Phrase:      patriciaPhrase}
)

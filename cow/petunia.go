package cow

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	petuniaBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	petuniaCode = nook.Code{
		Value: ""}
)

var (
	petuniaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Petunia"}

	petuniaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	petuniaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	petuniaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pétunia"}

	petuniaGermanName = nook.Name{
		Language: language.German,
		Value:    "Petunia"}

	petuniaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Petunia"}

	petuniaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "しもふり"}

	petuniaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	petuniaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	petuniaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	petuniaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "彤彤"}

	petuniaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Petunia"}

	petuniaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	petuniaName = nook.Languages{
		language.AmericanEnglish:      petuniaAmericanEnglishName,
		language.CanadianFrench:       petuniaCanadianFrenchName,
		language.Dutch:                petuniaDutchName,
		language.French:               petuniaFrenchName,
		language.German:               petuniaGermanName,
		language.Italian:              petuniaItalianName,
		language.Japanese:             petuniaJapaneseName,
		language.Korean:               petuniaKoreanName,
		language.LatinAmericanSpanish: petuniaLatinAmericanSpanishName,
		language.Russian:              petuniaRussianName,
		language.SimplifiedChinese:    petuniaSimplifiedChineseName,
		language.Spanish:              petuniaSpanishName,
		language.TraditionalChinese:   petuniaTraditionalChineseName}
)

var (
	petuniaCharacter = nook.Character{
		Animal:   Cow,
		Birthday: petuniaBirthday,
		Code:     petuniaCode,
		Gender:   nook.Female,
		Name:     petuniaName}
)

var (
	petuniaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	petuniaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	petuniaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	petuniaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mouh mouh"}

	petuniaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "muh-la-la"}

	petuniaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "muu la la"}

	petuniaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ザマス"}

	petuniaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	petuniaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	petuniaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	petuniaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哞呜"}

	petuniaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "muu-la-la"}

	petuniaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	petuniaPhrase = nook.Languages{
		language.AmericanEnglish:      petuniaAmericanEnglishName,
		language.CanadianFrench:       petuniaCanadianFrenchName,
		language.Dutch:                petuniaDutchName,
		language.French:               petuniaFrenchName,
		language.German:               petuniaGermanName,
		language.Italian:              petuniaItalianName,
		language.Japanese:             petuniaJapaneseName,
		language.Korean:               petuniaKoreanName,
		language.LatinAmericanSpanish: petuniaLatinAmericanSpanishName,
		language.Russian:              petuniaRussianName,
		language.SimplifiedChinese:    petuniaSimplifiedChineseName,
		language.Spanish:              petuniaSpanishName,
		language.TraditionalChinese:   petuniaTraditionalChineseName}
)

var (
	Petunia = nook.Villager{
		Character:   petuniaCharacter,
		Personality: nook.Snooty,
		Phrase:      petuniaPhrase}
)

package cow

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
		Value:    ""}

	petuniaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

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
		Value:    ""}

	petuniaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	petuniaRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	petuniaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "彤彤"}

	petuniaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Petunia"}

	petuniaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
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
		Animal:   animal.Cow,
		Birthday: petuniaBirthday,
		Code:     petuniaCode,
		Key:      character.Petunia,
		Gender:   gender.Female,
		Name:     petuniaName,
		Special:  false}
)

var (
	petuniaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "moo la la"}

	petuniaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	petuniaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

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
		Value:    ""}

	petuniaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	petuniaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	petuniaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哞呜"}

	petuniaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "muu-la-la"}

	petuniaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	petuniaPhrase = nook.Languages{
		language.AmericanEnglish:      petuniaAmericanEnglishPhrase,
		language.CanadianFrench:       petuniaCanadianFrenchPhrase,
		language.Dutch:                petuniaDutchPhrase,
		language.French:               petuniaFrenchPhrase,
		language.German:               petuniaGermanPhrase,
		language.Italian:              petuniaItalianPhrase,
		language.Japanese:             petuniaJapanesePhrase,
		language.Korean:               petuniaKoreanPhrase,
		language.LatinAmericanSpanish: petuniaLatinAmericanSpanishPhrase,
		language.Russian:              petuniaRussianPhrase,
		language.SimplifiedChinese:    petuniaSimplifiedChinesePhrase,
		language.Spanish:              petuniaSpanishPhrase,
		language.TraditionalChinese:   petuniaTraditionalChinesePhrase}
)

var (
	Petunia = nook.Villager{
		Character:   petuniaCharacter,
		Personality: personality.Snooty,
		Phrase:      petuniaPhrase}
)

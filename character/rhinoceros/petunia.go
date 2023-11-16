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
	// petuniaBirthday represents Petunia's birthday (unknown).
	petuniaBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// petuniaCode represents Petunia's unique code (unknown).
	petuniaCode = nook.Code{
		Value: ""}
)

var (
	// petuniaAmericanEnglishName represents Petunia's name in American English.
	petuniaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Petunia"}

	// petuniaCanadianFrenchName represents Petunia's name in Canadian French.
	petuniaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// petuniaDutchName represents Petunia's name in Dutch.
	petuniaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// petuniaFrenchName represents Petunia's name in French.
	petuniaFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	// petuniaGermanName represents Petunia's name in German.
	petuniaGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	// petuniaItalianName represents Petunia's name in Italian.
	petuniaItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// petuniaJapaneseName represents Petunia's name in Japanese.
	petuniaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ペチュニア"}

	// petuniaKoreanName represents Petunia's name in Korean.
	petuniaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// petuniaLatinAmericanSpanishName represents Petunia's name in Latin American Spanish.
	petuniaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// petuniaRussianName represents Petunia's name in Russian.
	petuniaRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// petuniaSimplifiedChineseName represents Petunia's name in Simplified Chinese.
	petuniaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// petuniaSpanishName represents Petunia's name in Spanish.
	petuniaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// petuniaTraditionalChineseName represents Petunia's name in Traditional Chinese.
	petuniaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// petuniaName represents Petunia's name in different languages.
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
	// petuniaCharacter represents Petunia's character information.
	petuniaCharacter = nook.Character{
		Animal:   animal.Rhinoceros,
		Birthday: petuniaBirthday,
		Code:     petuniaCode,
		Key:      character.Petunia,
		Gender:   gender.Female,
		Name:     petuniaName,
		Special:  false}
)

var (
	// petuniaAmericanEnglishPhrase represents Petunia's phrase in American English.
	petuniaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "メルシィ"}

	// petuniaCanadianFrenchPhrase represents Petunia's phrase in Canadian French.
	petuniaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// petuniaDutchPhrase represents Petunia's phrase in Dutch.
	petuniaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// petuniaFrenchPhrase represents Petunia's phrase in French.
	petuniaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	// petuniaGermanPhrase represents Petunia's phrase in German.
	petuniaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	// petuniaItalianPhrase represents Petunia's phrase in Italian.
	petuniaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// petuniaJapanesePhrase represents Petunia's phrase in Japanese.
	petuniaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "メルシィ"}

	// petuniaKoreanPhrase represents Petunia's phrase in Korean.
	petuniaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// petuniaLatinAmericanSpanishPhrase represents Petunia's phrase in Latin American Spanish.
	petuniaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// petuniaRussianPhrase represents Petunia's phrase in Russian.
	petuniaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// petuniaSimplifiedChinesePhrase represents Petunia's phrase in Simplified Chinese.
	petuniaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// petuniaSpanishPhrase represents Petunia's phrase in Spanish.
	petuniaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// petuniaTraditionalChinesePhrase represents Petunia's phrase in Traditional Chinese.
	petuniaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// petuniaPhrase represents Petunia's phrase in different languages.
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
	// Petunia represents the character Petunia with her complete information.
	Petunia = nook.Villager{
		Character:   petuniaCharacter,
		Personality: personality.Snooty,
		Phrase:      petuniaPhrase}
)

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
	// patriciaBirthday represents Patricia's birthday.
	patriciaBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// patriciaCode represents Patricia's unique code.
	patriciaCode = nook.Code{
		Value: ""}
)

var (
	// patriciaAmericanEnglishName represents Patricia's name in American English.
	patriciaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Patricia"}

	// patriciaCanadianFrenchName represents Patricia's name in Canadian French.
	patriciaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// patriciaDutchName represents Patricia's name in Dutch.
	patriciaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// patriciaFrenchName represents Patricia's name in French.
	patriciaFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	// patriciaGermanName represents Patricia's name in German.
	patriciaGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	// patriciaItalianName represents Patricia's name in Italian.
	patriciaItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// patriciaJapaneseName represents Patricia's name in Japanese.
	patriciaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "パトリシア"}

	// patriciaKoreanName represents Patricia's name in Korean.
	patriciaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// patriciaLatinAmericanSpanishName represents Patricia's name in Latin American Spanish.
	patriciaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// patriciaRussianName represents Patricia's name in Russian.
	patriciaRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// patriciaSimplifiedChineseName represents Patricia's name in Simplified Chinese.
	patriciaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// patriciaSpanishName represents Patricia's name in Spanish.
	patriciaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// patriciaTraditionalChineseName represents Patricia's name in Traditional Chinese.
	patriciaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// patriciaName represents Patricia's name in different languages.
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
	// patriciaCharacter represents Patricia's character information.
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
	// patriciaAmericanEnglishPhrase represents Patricia's phrase in American English.
	patriciaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "えへっ"}

	// patriciaCanadianFrenchPhrase represents Patricia's phrase in Canadian French.
	patriciaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// patriciaDutchPhrase represents Patricia's phrase in Dutch.
	patriciaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// patriciaFrenchPhrase represents Patricia's phrase in French.
	patriciaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	// patriciaGermanPhrase represents Patricia's phrase in German.
	patriciaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	// patriciaItalianPhrase represents Patricia's phrase in Italian.
	patriciaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// patriciaJapanesePhrase represents Patricia's phrase in Japanese.
	patriciaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "えへっ"}

	// patriciaKoreanPhrase represents Patricia's phrase in Korean.
	patriciaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// patriciaLatinAmericanSpanishPhrase represents Patricia's phrase in Latin American Spanish.
	patriciaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// patriciaRussianPhrase represents Patricia's phrase in Russian.
	patriciaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// patriciaSimplifiedChinesePhrase represents Patricia's phrase in Simplified Chinese.
	patriciaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// patriciaSpanishPhrase represents Patricia's phrase in Spanish.
	patriciaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// patriciaTraditionalChinesePhrase represents Patricia's phrase in Traditional Chinese.
	patriciaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// patriciaPhrase represents Patricia's phrase in different languages.
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
	// Patricia represents the character Patricia with her complete information.
	Patricia = nook.Villager{
		Character:   patriciaCharacter,
		Personality: personality.Normal,
		Phrase:      patriciaPhrase}
)

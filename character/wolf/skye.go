package wolf

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
	// skyeBirthday represents Skye's birthday.
	skyeBirthday = nook.Birthday{
		Day:   24,
		Month: time.March}
)

var (
	// skyeCode represents Skye's unique code.
	skyeCode = nook.Code{
		Value: "wol09"}
)

var (
	// skyeAmericanEnglishName represents Skye's name in American English.
	skyeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Skye"}

	// skyeCanadianFrenchName represents Skye's name in Canadian French.
	skyeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Marilou"}

	// skyeDutchName represents Skye's name in Dutch.
	skyeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Skye"}

	// skyeFrenchName represents Skye's name in French.
	skyeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Marilou"}

	// skyeGermanName represents Skye's name in German.
	skyeGermanName = nook.Name{
		Language: language.German,
		Value:    "Sabine"}

	// skyeItalianName represents Skye's name in Italian.
	skyeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lupilla"}

	// skyeJapaneseName represents Skye's name in Japanese.
	skyeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "リリィ"}

	// skyeKoreanName represents Skye's name in Korean.
	skyeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "릴리"}

	// skyeLatinAmericanSpanishName represents Skye's name in Latin American Spanish.
	skyeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Alderia"}

	// skyeRussianName represents Skye's name in Russian.
	skyeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Скай"}

	// skyeSimplifiedChineseName represents Skye's name in Simplified Chinese.
	skyeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "百合"}

	// skyeSpanishName represents Skye's name in Spanish.
	skyeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Alderia"}

	// skyeTraditionalChineseName represents Skye's name in Traditional Chinese.
	skyeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "百合"}
)

var (
	// skyeName represents Skye's name in different languages.
	skyeName = nook.Languages{
		language.AmericanEnglish:      skyeAmericanEnglishName,
		language.CanadianFrench:       skyeCanadianFrenchName,
		language.Dutch:                skyeDutchName,
		language.French:               skyeFrenchName,
		language.German:               skyeGermanName,
		language.Italian:              skyeItalianName,
		language.Japanese:             skyeJapaneseName,
		language.Korean:               skyeKoreanName,
		language.LatinAmericanSpanish: skyeLatinAmericanSpanishName,
		language.Russian:              skyeRussianName,
		language.SimplifiedChinese:    skyeSimplifiedChineseName,
		language.Spanish:              skyeSpanishName,
		language.TraditionalChinese:   skyeTraditionalChineseName}
)

var (
	// skyeCharacter represents Skye's character information.
	skyeCharacter = nook.Character{
		Animal:   animal.Wolf,
		Birthday: skyeBirthday,
		Code:     skyeCode,
		Key:      character.Skye,
		Gender:   gender.Female,
		Name:     skyeName,
		Special:  false}
)

var (
	// skyeAmericanEnglishPhrase represents Skye's phrase in American English.
	skyeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "airmail"}

	// skyeCanadianFrenchPhrase represents Skye's phrase in Canadian French.
	skyeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mumuseau"}

	// skyeDutchPhrase represents Skye's phrase in Dutch.
	skyeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "luchtpost"}

	// skyeFrenchPhrase represents Skye's phrase in French.
	skyeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mumuseau"}

	// skyeGermanPhrase represents Skye's phrase in German.
	skyeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "ahiii"}

	// skyeItalianPhrase represents Skye's phrase in Italian.
	skyeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "uuuhlalà"}

	// skyeJapanesePhrase represents Skye's phrase in Japanese.
	skyeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "フワワ"}

	// skyeKoreanPhrase represents Skye's phrase in Korean.
	skyeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "후와와"}

	// skyeLatinAmericanSpanishPhrase represents Skye's phrase in Latin American Spanish.
	skyeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "auuu"}

	// skyeRussianPhrase represents Skye's phrase in Russian.
	skyeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "пилот"}

	// skyeSimplifiedChinesePhrase represents Skye's phrase in Simplified Chinese.
	skyeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "花花"}

	// skyeSpanishPhrase represents Skye's phrase in Spanish.
	skyeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "auuu"}

	// skyeTraditionalChinesePhrase represents Skye's phrase in Traditional Chinese.
	skyeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "花花"}
)

var (
	// skyePhrase represents Skye's phrase in different languages.
	skyePhrase = nook.Languages{
		language.AmericanEnglish:      skyeAmericanEnglishPhrase,
		language.CanadianFrench:       skyeCanadianFrenchPhrase,
		language.Dutch:                skyeDutchPhrase,
		language.French:               skyeFrenchPhrase,
		language.German:               skyeGermanPhrase,
		language.Italian:              skyeItalianPhrase,
		language.Japanese:             skyeJapanesePhrase,
		language.Korean:               skyeKoreanPhrase,
		language.LatinAmericanSpanish: skyeLatinAmericanSpanishPhrase,
		language.Russian:              skyeRussianPhrase,
		language.SimplifiedChinese:    skyeSimplifiedChinesePhrase,
		language.Spanish:              skyeSpanishPhrase,
		language.TraditionalChinese:   skyeTraditionalChinesePhrase}
)

var (
	// Skye represents the character Skye with her complete information.
	Skye = nook.Villager{
		Character:   skyeCharacter,
		Personality: personality.Normal,
		Phrase:      skyePhrase}
)

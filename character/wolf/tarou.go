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
	// tarouBirthday represents Tarou's birthday (January 0th).
	tarouBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// tarouCode represents Tarou's unique code (empty string).
	tarouCode = nook.Code{
		Value: ""}
)

var (
	// tarouAmericanEnglishName represents Tarou's name in American English.
	tarouAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tarou"}

	// tarouCanadianFrenchName represents Tarou's name in Canadian French.
	tarouCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// tarouDutchName represents Tarou's name in Dutch.
	tarouDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// tarouFrenchName represents Tarou's name in French.
	tarouFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	// tarouGermanName represents Tarou's name in German.
	tarouGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	// tarouItalianName represents Tarou's name in Italian.
	tarouItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// tarouJapaneseName represents Tarou's name in Japanese.
	tarouJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "タロウ"}

	// tarouKoreanName represents Tarou's name in Korean.
	tarouKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// tarouLatinAmericanSpanishName represents Tarou's name in Latin American Spanish.
	tarouLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// tarouRussianName represents Tarou's name in Russian.
	tarouRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// tarouSimplifiedChineseName represents Tarou's name in Simplified Chinese.
	tarouSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// tarouSpanishName represents Tarou's name in Spanish.
	tarouSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// tarouTraditionalChineseName represents Tarou's name in Traditional Chinese.
	tarouTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// tarouName represents Tarou's name in different languages.
	tarouName = nook.Languages{
		language.AmericanEnglish:      tarouAmericanEnglishName,
		language.CanadianFrench:       tarouCanadianFrenchName,
		language.Dutch:                tarouDutchName,
		language.French:               tarouFrenchName,
		language.German:               tarouGermanName,
		language.Italian:              tarouItalianName,
		language.Japanese:             tarouJapaneseName,
		language.Korean:               tarouKoreanName,
		language.LatinAmericanSpanish: tarouLatinAmericanSpanishName,
		language.Russian:              tarouRussianName,
		language.SimplifiedChinese:    tarouSimplifiedChineseName,
		language.Spanish:              tarouSpanishName,
		language.TraditionalChinese:   tarouTraditionalChineseName}
)

var (
	// tarouCharacter represents Tarou's character information.
	tarouCharacter = nook.Character{
		Animal:   animal.Wolf,
		Birthday: tarouBirthday,
		Code:     tarouCode,
		Key:      character.Tarou,
		Gender:   gender.Male,
		Name:     tarouName,
		Special:  false}
)

var (
	// tarouAmericanEnglishPhrase represents Tarou's phrase in American English.
	tarouAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ワオーン"}

	// tarouCanadianFrenchPhrase represents Tarou's phrase in Canadian French.
	tarouCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// tarouDutchPhrase represents Tarou's phrase in Dutch.
	tarouDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// tarouFrenchPhrase represents Tarou's phrase in French.
	tarouFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	// tarouGermanPhrase represents Tarou's phrase in German.
	tarouGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	// tarouItalianPhrase represents Tarou's phrase in Italian.
	tarouItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// tarouJapanesePhrase represents Tarou's phrase in Japanese.
	tarouJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ワオーン"}

	// tarouKoreanPhrase represents Tarou's phrase in Korean.
	tarouKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// tarouLatinAmericanSpanishPhrase represents Tarou's phrase in Latin American Spanish.
	tarouLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// tarouRussianPhrase represents Tarou's phrase in Russian.
	tarouRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// tarouSimplifiedChinesePhrase represents Tarou's phrase in Simplified Chinese.
	tarouSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// tarouSpanishPhrase represents Tarou's phrase in Spanish.
	tarouSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// tarouTraditionalChinesePhrase represents Tarou's phrase in Traditional Chinese.
	tarouTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// tarouPhrase represents Tarou's phrase in different languages.
	tarouPhrase = nook.Languages{
		language.AmericanEnglish:      tarouAmericanEnglishPhrase,
		language.CanadianFrench:       tarouCanadianFrenchPhrase,
		language.Dutch:                tarouDutchPhrase,
		language.French:               tarouFrenchPhrase,
		language.German:               tarouGermanPhrase,
		language.Italian:              tarouItalianPhrase,
		language.Japanese:             tarouJapanesePhrase,
		language.Korean:               tarouKoreanPhrase,
		language.LatinAmericanSpanish: tarouLatinAmericanSpanishPhrase,
		language.Russian:              tarouRussianPhrase,
		language.SimplifiedChinese:    tarouSimplifiedChinesePhrase,
		language.Spanish:              tarouSpanishPhrase,
		language.TraditionalChinese:   tarouTraditionalChinesePhrase}
)

var (
	// Tarou represents the character Tarou with his complete information.
	Tarou = nook.Villager{
		Character:   tarouCharacter,
		Personality: personality.Jock,
		Phrase:      tarouPhrase}
)

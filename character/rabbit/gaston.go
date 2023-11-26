package rabbit

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
	// gastonBirthday represents Gaston's birthday.
	gastonBirthday = nook.Birthday{
		Day:   28,
		Month: time.October}
)

var (
	// gastonCode represents Gaston's unique code.
	gastonCode = nook.Code{
		Value: "rbt04"}
)

var (
	// gastonAmericanEnglishName represents Gaston's name in American English.
	gastonAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Gaston"}

	// gastonCanadianFrenchName represents Gaston's name in Canadian French.
	gastonCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gaston"}

	// gastonDutchName represents Gaston's name in Dutch.
	gastonDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Gaston"}

	// gastonFrenchName represents Gaston's name in French.
	gastonFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gaston"}

	// gastonGermanName represents Gaston's name in German.
	gastonGermanName = nook.Name{
		Language: language.German,
		Value:    "Gaston"}

	// gastonItalianName represents Gaston's name in Italian.
	gastonItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gaston"}

	// gastonJapaneseName represents Gaston's name in Japanese.
	gastonJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "モサキチ"}

	// gastonKoreanName represents Gaston's name in Korean.
	gastonKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "대길"}

	// gastonLatinAmericanSpanishName represents Gaston's name in Latin American Spanish.
	gastonLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Gastón"}

	// gastonRussianName represents Gaston's name in Russian.
	gastonRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Гастон"}

	// gastonSimplifiedChineseName represents Gaston's name in Simplified Chinese.
	gastonSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "猛兔"}

	// gastonSpanishName represents Gaston's name in Spanish.
	gastonSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Gastón"}

	// gastonTraditionalChineseName represents Gaston's name in Traditional Chinese.
	gastonTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "猛兔"}
)

var (
	// gastonName represents Gaston's name in different languages.
	gastonName = nook.Languages{
		language.AmericanEnglish:      gastonAmericanEnglishName,
		language.CanadianFrench:       gastonCanadianFrenchName,
		language.Dutch:                gastonDutchName,
		language.French:               gastonFrenchName,
		language.German:               gastonGermanName,
		language.Italian:              gastonItalianName,
		language.Japanese:             gastonJapaneseName,
		language.Korean:               gastonKoreanName,
		language.LatinAmericanSpanish: gastonLatinAmericanSpanishName,
		language.Russian:              gastonRussianName,
		language.SimplifiedChinese:    gastonSimplifiedChineseName,
		language.Spanish:              gastonSpanishName,
		language.TraditionalChinese:   gastonTraditionalChineseName}
)

var (
	// gastonCharacter represents Gaston's character information.
	gastonCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: gastonBirthday,
		Code:     gastonCode,
		Key:      character.Gaston,
		Gender:   gender.Male,
		Name:     gastonName,
		Special:  false}
)

var (
	// gastonAmericanEnglishPhrase represents Gaston's phrase in American English.
	gastonAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "mon chou"}

	// gastonCanadianFrenchPhrase represents Gaston's phrase in Canadian French.
	gastonCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "cahuète"}

	// gastonDutchPhrase represents Gaston's phrase in Dutch.
	gastonDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "mon amour"}

	// gastonFrenchPhrase represents Gaston's phrase in French.
	gastonFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "cahuète"}

	// gastonGermanPhrase represents Gaston's phrase in German.
	gastonGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "herzblatt"}

	// gastonItalianPhrase represents Gaston's phrase in Italian.
	gastonItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cherie"}

	// gastonJapanesePhrase represents Gaston's phrase in Japanese.
	gastonJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ムホッ"}

	// gastonKoreanPhrase represents Gaston's phrase in Korean.
	gastonKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "쿨럭"}

	// gastonLatinAmericanSpanishPhrase represents Gaston's phrase in Latin American Spanish.
	gastonLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "mon chou"}

	// gastonRussianPhrase represents Gaston's phrase in Russian.
	gastonRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "мон зайшу"}

	// gastonSimplifiedChinesePhrase represents Gaston's phrase in Simplified Chinese.
	gastonSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "唔贺"}

	// gastonSpanishPhrase represents Gaston's phrase in Spanish.
	gastonSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "mon chou"}

	// gastonTraditionalChinesePhrase represents Gaston's phrase in Traditional Chinese.
	gastonTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "唔賀"}
)

var (
	// gastonPhrase represents Gaston's phrases in different languages.
	gastonPhrase = nook.Languages{
		language.AmericanEnglish:      gastonAmericanEnglishPhrase,
		language.CanadianFrench:       gastonCanadianFrenchPhrase,
		language.Dutch:                gastonDutchPhrase,
		language.French:               gastonFrenchPhrase,
		language.German:               gastonGermanPhrase,
		language.Italian:              gastonItalianPhrase,
		language.Japanese:             gastonJapanesePhrase,
		language.Korean:               gastonKoreanPhrase,
		language.LatinAmericanSpanish: gastonLatinAmericanSpanishPhrase,
		language.Russian:              gastonRussianPhrase,
		language.SimplifiedChinese:    gastonSimplifiedChinesePhrase,
		language.Spanish:              gastonSpanishPhrase,
		language.TraditionalChinese:   gastonTraditionalChinesePhrase}
)

var (
	// Gaston represents the character Gaston with his complete information.
	Gaston = nook.Villager{
		Character:   gastonCharacter,
		Personality: personality.Cranky,
		Phrase:      gastonPhrase}
)

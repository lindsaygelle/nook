package sheep

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
	// baabaraBirthday represents Baabara's birthday (March 28th).
	baabaraBirthday = nook.Birthday{
		Day:   28,
		Month: time.March}
)

var (
	// baabaraCode represents Baabara's unique code ("shp01").
	baabaraCode = nook.Code{
		Value: "shp01"}
)

var (
	// baabaraAmericanEnglishName represents Baabara's name in American English.
	baabaraAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Baabara"}

	// baabaraCanadianFrenchName represents Baabara's name in Canadian French.
	baabaraCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bêêtty"}

	// baabaraDutchName represents Baabara's name in Dutch.
	baabaraDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Baabara"}

	// baabaraFrenchName represents Baabara's name in French.
	baabaraFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bêêtty"}

	// baabaraGermanName represents Baabara's name in German.
	baabaraGermanName = nook.Name{
		Language: language.German,
		Value:    "Babsi"}

	// baabaraItalianName represents Baabara's name in Italian.
	baabaraItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Bea"}

	// baabaraJapaneseName represents Baabara's name in Japanese.
	baabaraJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "トロワ"}

	// baabaraKoreanName represents Baabara's name in Korean.
	baabaraKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "트로와"}

	// baabaraLatinAmericanSpanishName represents Baabara's name in Latin American Spanish.
	baabaraLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Beelén"}

	// baabaraRussianName represents Baabara's name in Russian.
	baabaraRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бабара"}

	// baabaraSimplifiedChineseName represents Baabara's name in Simplified Chinese.
	baabaraSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "华尔滋"}

	// baabaraSpanishName represents Baabara's name in Spanish.
	baabaraSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Beelén"}

	// baabaraTraditionalChineseName represents Baabara's name in Traditional Chinese.
	baabaraTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "華爾滋"}
)

var (
	// baabaraName represents Baabara's name in different languages.
	baabaraName = nook.Languages{
		language.AmericanEnglish:      baabaraAmericanEnglishName,
		language.CanadianFrench:       baabaraCanadianFrenchName,
		language.Dutch:                baabaraDutchName,
		language.French:               baabaraFrenchName,
		language.German:               baabaraGermanName,
		language.Italian:              baabaraItalianName,
		language.Japanese:             baabaraJapaneseName,
		language.Korean:               baabaraKoreanName,
		language.LatinAmericanSpanish: baabaraLatinAmericanSpanishName,
		language.Russian:              baabaraRussianName,
		language.SimplifiedChinese:    baabaraSimplifiedChineseName,
		language.Spanish:              baabaraSpanishName,
		language.TraditionalChinese:   baabaraTraditionalChineseName}
)

var (
	// baabaraCharacter represents Baabara's character information.
	baabaraCharacter = nook.Character{
		Animal:   animal.Sheep,
		Birthday: baabaraBirthday,
		Code:     baabaraCode,
		Key:      character.Baabara,
		Gender:   gender.Female,
		Name:     baabaraName,
		Special:  false}
)

var (
	// baabaraAmericanEnglishPhrase represents Baabara's phrase in American English.
	baabaraAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "daahling"}

	// baabaraCanadianFrenchPhrase represents Baabara's phrase in Canadian French.
	baabaraCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "chêêêri"}

	// baabaraDutchPhrase represents Baabara's phrase in Dutch.
	baabaraDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "schaaaapje"}

	// baabaraFrenchPhrase represents Baabara's phrase in French.
	baabaraFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "chêêêri"}

	// baabaraGermanPhrase represents Baabara's phrase in German.
	baabaraGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "däääää"}

	// baabaraItalianPhrase represents Baabara's phrase in Italian.
	baabaraItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "breeezza"}

	// baabaraJapanesePhrase represents Baabara's phrase in Japanese.
	baabaraJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "アンドゥ"}

	// baabaraKoreanPhrase represents Baabara's phrase in Korean.
	baabaraKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "앤쥬"}

	// baabaraLatinAmericanSpanishPhrase represents Baabara's phrase in Latin American Spanish.
	baabaraLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "peluuusas"}

	// baabaraRussianPhrase represents Baabara's phrase in Russian.
	baabaraRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "мила-ашка"}

	// baabaraSimplifiedChinesePhrase represents Baabara's phrase in Simplified Chinese.
	baabaraSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "一二三"}

	// baabaraSpanishPhrase represents Baabara's phrase in Spanish.
	baabaraSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "peeelusa"}

	// baabaraTraditionalChinesePhrase represents Baabara's phrase in Traditional Chinese.
	baabaraTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "一二三"}
)

var (
	// baabaraPhrase represents Baabara's phrases in different languages.
	baabaraPhrase = nook.Languages{
		language.AmericanEnglish:      baabaraAmericanEnglishPhrase,
		language.CanadianFrench:       baabaraCanadianFrenchPhrase,
		language.Dutch:                baabaraDutchPhrase,
		language.French:               baabaraFrenchPhrase,
		language.German:               baabaraGermanPhrase,
		language.Italian:              baabaraItalianPhrase,
		language.Japanese:             baabaraJapanesePhrase,
		language.Korean:               baabaraKoreanPhrase,
		language.LatinAmericanSpanish: baabaraLatinAmericanSpanishPhrase,
		language.Russian:              baabaraRussianPhrase,
		language.SimplifiedChinese:    baabaraSimplifiedChinesePhrase,
		language.Spanish:              baabaraSpanishPhrase,
		language.TraditionalChinese:   baabaraTraditionalChinesePhrase}
)

var (
	// Baabara represents the character Baabara with her complete information.
	Baabara = nook.Villager{
		Character:   baabaraCharacter,
		Personality: personality.Snooty,
		Phrase:      baabaraPhrase}
)

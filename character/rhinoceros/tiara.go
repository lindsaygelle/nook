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
	// tiaraBirthday represents Tiara's birthday.
	tiaraBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// tiaraCode represents Tiara's unique code.
	tiaraCode = nook.Code{
		Value: ""}
)

var (
	// tiaraAmericanEnglishName represents Tiara's name in American English.
	tiaraAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tiara"}

	// tiaraCanadianFrenchName represents Tiara's name in Canadian French.
	tiaraCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// tiaraDutchName represents Tiara's name in Dutch.
	tiaraDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// tiaraFrenchName represents Tiara's name in French.
	tiaraFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tiara"}

	// tiaraGermanName represents Tiara's name in German.
	tiaraGermanName = nook.Name{
		Language: language.German,
		Value:    "Roswitha"}

	// tiaraItalianName represents Tiara's name in Italian.
	tiaraItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cornelia"}

	// tiaraJapaneseName represents Tiara's name in Japanese.
	tiaraJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ワイハ"}

	// tiaraKoreanName represents Tiara's name in Korean.
	tiaraKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// tiaraLatinAmericanSpanishName represents Tiara's name in Latin American Spanish.
	tiaraLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// tiaraRussianName represents Tiara's name in Russian.
	tiaraRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// tiaraSimplifiedChineseName represents Tiara's name in Simplified Chinese.
	tiaraSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "锋锋"}

	// tiaraSpanishName represents Tiara's name in Spanish.
	tiaraSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tiara"}

	// tiaraTraditionalChineseName represents Tiara's name in Traditional Chinese.
	tiaraTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// tiaraName represents Tiara's name in different languages.
	tiaraName = nook.Languages{
		language.AmericanEnglish:      tiaraAmericanEnglishName,
		language.CanadianFrench:       tiaraCanadianFrenchName,
		language.Dutch:                tiaraDutchName,
		language.French:               tiaraFrenchName,
		language.German:               tiaraGermanName,
		language.Italian:              tiaraItalianName,
		language.Japanese:             tiaraJapaneseName,
		language.Korean:               tiaraKoreanName,
		language.LatinAmericanSpanish: tiaraLatinAmericanSpanishName,
		language.Russian:              tiaraRussianName,
		language.SimplifiedChinese:    tiaraSimplifiedChineseName,
		language.Spanish:              tiaraSpanishName,
		language.TraditionalChinese:   tiaraTraditionalChineseName}
)

var (
	// tiaraCharacter represents Tiara's character information.
	tiaraCharacter = nook.Character{
		Animal:   animal.Rhinoceros,
		Birthday: tiaraBirthday,
		Code:     tiaraCode,
		Key:      character.Tiara,
		Gender:   gender.Female,
		Name:     tiaraName,
		Special:  false}
)

var (
	// tiaraAmericanEnglishPhrase represents Tiara's phrase in American English.
	tiaraAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "lovey"}

	// tiaraCanadianFrenchPhrase represents Tiara's phrase in Canadian French.
	tiaraCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// tiaraDutchPhrase represents Tiara's phrase in Dutch.
	tiaraDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// tiaraFrenchPhrase represents Tiara's phrase in French.
	tiaraFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mon hippy"}

	// tiaraGermanPhrase represents Tiara's phrase in German.
	tiaraGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "herzchen"}

	// tiaraItalianPhrase represents Tiara's phrase in Italian.
	tiaraItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "dolcezza"}

	// tiaraJapanesePhrase represents Tiara's phrase in Japanese.
	tiaraJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "てゆーか"}

	// tiaraKoreanPhrase represents Tiara's phrase in Korean.
	tiaraKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// tiaraLatinAmericanSpanishPhrase represents Tiara's phrase in Latin American Spanish.
	tiaraLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// tiaraRussianPhrase represents Tiara's phrase in Russian.
	tiaraRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// tiaraSimplifiedChinesePhrase represents Tiara's phrase in Simplified Chinese.
	tiaraSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "呀嗨"}

	// tiaraSpanishPhrase represents Tiara's phrase in Spanish.
	tiaraSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "amorcito"}

	// tiaraTraditionalChinesePhrase represents Tiara's phrase in Traditional Chinese.
	tiaraTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// tiaraPhrase represents Tiara's phrase in different languages.
	tiaraPhrase = nook.Languages{
		language.AmericanEnglish:      tiaraAmericanEnglishPhrase,
		language.CanadianFrench:       tiaraCanadianFrenchPhrase,
		language.Dutch:                tiaraDutchPhrase,
		language.French:               tiaraFrenchPhrase,
		language.German:               tiaraGermanPhrase,
		language.Italian:              tiaraItalianPhrase,
		language.Japanese:             tiaraJapanesePhrase,
		language.Korean:               tiaraKoreanPhrase,
		language.LatinAmericanSpanish: tiaraLatinAmericanSpanishPhrase,
		language.Russian:              tiaraRussianPhrase,
		language.SimplifiedChinese:    tiaraSimplifiedChinesePhrase,
		language.Spanish:              tiaraSpanishPhrase,
		language.TraditionalChinese:   tiaraTraditionalChinesePhrase}
)

var (
	// Tiara represents the character Tiara with her complete information.
	Tiara = nook.Villager{
		Character:   tiaraCharacter,
		Personality: personality.Snooty,
		Phrase:      tiaraPhrase}
)

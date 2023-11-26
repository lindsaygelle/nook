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
	// loboBirthday represents Lobo's birthday.
	loboBirthday = nook.Birthday{
		Day:   5,
		Month: time.November}
)

var (
	// loboCode represents Lobo's unique code.
	loboCode = nook.Code{
		Value: "wol01"}
)

var (
	// loboAmericanEnglishName represents Lobo's name in American English.
	loboAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Lobo"}

	// loboCanadianFrenchName represents Lobo's name in Canadian French.
	loboCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Lobo"}

	// loboDutchName represents Lobo's name in Dutch.
	loboDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Lobo"}

	// loboFrenchName represents Lobo's name in French.
	loboFrenchName = nook.Name{
		Language: language.French,
		Value:    "Lobo"}

	// loboGermanName represents Lobo's name in German.
	loboGermanName = nook.Name{
		Language: language.German,
		Value:    "Lupo"}

	// loboItalianName represents Lobo's name in Italian.
	loboItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lupen"}

	// loboJapaneseName represents Lobo's name in Japanese.
	loboJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ブンジロウ"}

	// loboKoreanName represents Lobo's name in Korean.
	loboKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "늑태"}

	// loboLatinAmericanSpanishName represents Lobo's name in Latin American Spanish.
	loboLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lupo"}

	// loboRussianName represents Lobo's name in Russian.
	loboRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лобо"}

	// loboSimplifiedChineseName represents Lobo's name in Simplified Chinese.
	loboSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "潘奇隆"}

	// loboSpanishName represents Lobo's name in Spanish.
	loboSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lupo"}

	// loboTraditionalChineseName represents Lobo's name in Traditional Chinese.
	loboTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "潘奇隆"}
)

var (
	// loboName represents Lobo's name in different languages.
	loboName = nook.Languages{
		language.AmericanEnglish:      loboAmericanEnglishName,
		language.CanadianFrench:       loboCanadianFrenchName,
		language.Dutch:                loboDutchName,
		language.French:               loboFrenchName,
		language.German:               loboGermanName,
		language.Italian:              loboItalianName,
		language.Japanese:             loboJapaneseName,
		language.Korean:               loboKoreanName,
		language.LatinAmericanSpanish: loboLatinAmericanSpanishName,
		language.Russian:              loboRussianName,
		language.SimplifiedChinese:    loboSimplifiedChineseName,
		language.Spanish:              loboSpanishName,
		language.TraditionalChinese:   loboTraditionalChineseName}
)

var (
	// loboCharacter represents Lobo's character information.
	loboCharacter = nook.Character{
		Animal:   animal.Wolf,
		Birthday: loboBirthday,
		Code:     loboCode,
		Key:      character.Lobo,
		Gender:   gender.Male,
		Name:     loboName,
		Special:  false}
)

var (
	// loboAmericanEnglishPhrase represents Lobo's phrase in American English.
	loboAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ah-rooooo"}

	// loboCanadianFrenchPhrase represents Lobo's phrase in Canadian French.
	loboCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ah rooooo"}

	// loboDutchPhrase represents Lobo's phrase in Dutch.
	loboDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "aooooeee"}

	// loboFrenchPhrase represents Lobo's phrase in French.
	loboFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ah rooooo"}

	// loboGermanPhrase represents Lobo's phrase in German.
	loboGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "auuuuuu"}

	// loboItalianPhrase represents Lobo's phrase in Italian.
	loboItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "etchuuuu"}

	// loboJapanesePhrase represents Lobo's phrase in Japanese.
	loboJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だぜよ"}

	// loboKoreanPhrase represents Lobo's phrase in Korean.
	loboKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "봐봐"}

	// loboLatinAmericanSpanishPhrase represents Lobo's phrase in Latin American Spanish.
	loboLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "auuuh"}

	// loboRussianPhrase represents Lobo's phrase in Russian.
	loboRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ау-у-у"}

	// loboSimplifiedChinesePhrase represents Lobo's phrase in Simplified Chinese.
	loboSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "不然呢"}

	// loboSpanishPhrase represents Lobo's phrase in Spanish.
	loboSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "auuuh"}

	// loboTraditionalChinesePhrase represents Lobo's phrase in Traditional Chinese.
	loboTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "不然呢"}
)

var (
	// loboPhrase represents Lobo's phrase in different languages.
	loboPhrase = nook.Languages{
		language.AmericanEnglish:      loboAmericanEnglishPhrase,
		language.CanadianFrench:       loboCanadianFrenchPhrase,
		language.Dutch:                loboDutchPhrase,
		language.French:               loboFrenchPhrase,
		language.German:               loboGermanPhrase,
		language.Italian:              loboItalianPhrase,
		language.Japanese:             loboJapanesePhrase,
		language.Korean:               loboKoreanPhrase,
		language.LatinAmericanSpanish: loboLatinAmericanSpanishPhrase,
		language.Russian:              loboRussianPhrase,
		language.SimplifiedChinese:    loboSimplifiedChinesePhrase,
		language.Spanish:              loboSpanishPhrase,
		language.TraditionalChinese:   loboTraditionalChinesePhrase}
)

var (
	// Lobo represents the character Lobo with his complete information.
	Lobo = nook.Villager{
		Character:   loboCharacter,
		Personality: personality.Cranky,
		Phrase:      loboPhrase}
)

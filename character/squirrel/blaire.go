package squirrel

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
	// blaireBirthday represents Blaire's birthday (July 3rd).
	blaireBirthday = nook.Birthday{
		Day:   3,
		Month: time.July}
)

var (
	// blaireCode represents Blaire's unique code ("squ01").
	blaireCode = nook.Code{
		Value: "squ01"}
)

var (
	// blaireAmericanEnglishName represents Blaire's name in American English.
	blaireAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Blaire"}

	// blaireCanadianFrenchName represents Blaire's name in Canadian French.
	blaireCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Cachou"}

	// blaireDutchName represents Blaire's name in Dutch.
	blaireDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Blaire"}

	// blaireFrenchName represents Blaire's name in French.
	blaireFrenchName = nook.Name{
		Language: language.French,
		Value:    "Cachou"}

	// blaireGermanName represents Blaire's name in German.
	blaireGermanName = nook.Name{
		Language: language.German,
		Value:    "Klara"}

	// blaireItalianName represents Blaire's name in Italian.
	blaireItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ghianda"}

	// blaireJapaneseName represents Blaire's name in Japanese.
	blaireJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "シルエット"}

	// blaireKoreanName represents Blaire's name in Korean.
	blaireKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "실루엣"}

	// blaireLatinAmericanSpanishName represents Blaire's name in Latin American Spanish.
	blaireLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Azabache"}

	// blaireRussianName represents Blaire's name in Russian.
	blaireRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Блер"}

	// blaireSimplifiedChineseName represents Blaire's name in Simplified Chinese.
	blaireSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小影"}

	// blaireSpanishName represents Blaire's name in Spanish.
	blaireSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Azabache"}

	// blaireTraditionalChineseName represents Blaire's name in Traditional Chinese.
	blaireTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小影"}
)

var (
	// blaireName represents Blaire's name in different languages.
	blaireName = nook.Languages{
		language.AmericanEnglish:      blaireAmericanEnglishName,
		language.CanadianFrench:       blaireCanadianFrenchName,
		language.Dutch:                blaireDutchName,
		language.French:               blaireFrenchName,
		language.German:               blaireGermanName,
		language.Italian:              blaireItalianName,
		language.Japanese:             blaireJapaneseName,
		language.Korean:               blaireKoreanName,
		language.LatinAmericanSpanish: blaireLatinAmericanSpanishName,
		language.Russian:              blaireRussianName,
		language.SimplifiedChinese:    blaireSimplifiedChineseName,
		language.Spanish:              blaireSpanishName,
		language.TraditionalChinese:   blaireTraditionalChineseName}
)

var (
	// blaireCharacter represents Blaire's character information.
	blaireCharacter = nook.Character{
		Animal:   animal.Squirrel,
		Birthday: blaireBirthday,
		Code:     blaireCode,
		Key:      character.Blaire,
		Gender:   gender.Female,
		Name:     blaireName,
		Special:  false}
)

var (
	// blaireAmericanEnglishPhrase represents Blaire's phrase in American English.
	blaireAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "nutlet"}

	// blaireCanadianFrenchPhrase represents Blaire's phrase in Canadian French.
	blaireCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "coquille"}

	// blaireDutchPhrase represents Blaire's phrase in Dutch.
	blaireDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "notendop"}

	// blaireFrenchPhrase represents Blaire's phrase in French.
	blaireFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "coquille"}

	// blaireGermanPhrase represents Blaire's phrase in German.
	blaireGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "knack"}

	// blaireItalianPhrase represents Blaire's phrase in Italian.
	blaireItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "scrick"}

	// blaireJapanesePhrase represents Blaire's phrase in Japanese.
	blaireJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ふふん"}

	// blaireKoreanPhrase represents Blaire's phrase in Korean.
	blaireKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "헤헤헤"}

	// blaireLatinAmericanSpanishPhrase represents Blaire's phrase in Latin American Spanish.
	blaireLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ñicuá"}

	// blaireRussianPhrase represents Blaire's phrase in Russian.
	blaireRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "крошка"}

	// blaireSimplifiedChinesePhrase represents Blaire's phrase in Simplified Chinese.
	blaireSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哼哼"}

	// blaireSpanishPhrase represents Blaire's phrase in Spanish.
	blaireSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "bayita"}

	// blaireTraditionalChinesePhrase represents Blaire's phrase in Traditional Chinese.
	blaireTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哼哼"}
)

var (
	// blairePhrase represents Blaire's phrases in different languages.
	blairePhrase = nook.Languages{
		language.AmericanEnglish:      blaireAmericanEnglishPhrase,
		language.CanadianFrench:       blaireCanadianFrenchPhrase,
		language.Dutch:                blaireDutchPhrase,
		language.French:               blaireFrenchPhrase,
		language.German:               blaireGermanPhrase,
		language.Italian:              blaireItalianPhrase,
		language.Japanese:             blaireJapanesePhrase,
		language.Korean:               blaireKoreanPhrase,
		language.LatinAmericanSpanish: blaireLatinAmericanSpanishPhrase,
		language.Russian:              blaireRussianPhrase,
		language.SimplifiedChinese:    blaireSimplifiedChinesePhrase,
		language.Spanish:              blaireSpanishPhrase,
		language.TraditionalChinese:   blaireTraditionalChinesePhrase}
)

var (
	// Blaire represents the character Blaire with her complete information.
	Blaire = nook.Villager{
		Character:   blaireCharacter,
		Personality: personality.Snooty,
		Phrase:      blairePhrase}
)

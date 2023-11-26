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
	// miraBirthday represents Mira's birthday.
	miraBirthday = nook.Birthday{
		Day:   6,
		Month: time.July}
)

var (
	// miraCode represents Mira's unique code.
	miraCode = nook.Code{
		Value: "rbt19"}
)

var (
	// miraAmericanEnglishName represents Mira's name in American English.
	miraAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Mira"}

	// miraCanadianFrenchName represents Mira's name in Canadian French.
	miraCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Grisette"}

	// miraDutchName represents Mira's name in Dutch.
	miraDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Mira"}

	// miraFrenchName represents Mira's name in French.
	miraFrenchName = nook.Name{
		Language: language.French,
		Value:    "Grisette"}

	// miraGermanName represents Mira's name in German.
	miraGermanName = nook.Name{
		Language: language.German,
		Value:    "Mira"}

	// miraItalianName represents Mira's name in Italian.
	miraItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Martina"}

	// miraJapaneseName represents Mira's name in Japanese.
	miraJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ミラコ"}

	// miraKoreanName represents Mira's name in Korean.
	miraKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "미랑"}

	// miraLatinAmericanSpanishName represents Mira's name in Latin American Spanish.
	miraLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Amparo"}

	// miraRussianName represents Mira's name in Russian.
	miraRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Мира"}

	// miraSimplifiedChineseName represents Mira's name in Simplified Chinese.
	miraSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "蜜拉"}

	// miraSpanishName represents Mira's name in Spanish.
	miraSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Amparo"}

	// miraTraditionalChineseName represents Mira's name in Traditional Chinese.
	miraTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "蜜拉"}
)

var (
	// miraName represents Mira's name in different languages.
	miraName = nook.Languages{
		language.AmericanEnglish:      miraAmericanEnglishName,
		language.CanadianFrench:       miraCanadianFrenchName,
		language.Dutch:                miraDutchName,
		language.French:               miraFrenchName,
		language.German:               miraGermanName,
		language.Italian:              miraItalianName,
		language.Japanese:             miraJapaneseName,
		language.Korean:               miraKoreanName,
		language.LatinAmericanSpanish: miraLatinAmericanSpanishName,
		language.Russian:              miraRussianName,
		language.SimplifiedChinese:    miraSimplifiedChineseName,
		language.Spanish:              miraSpanishName,
		language.TraditionalChinese:   miraTraditionalChineseName}
)

var (
	// miraCharacter represents Mira's character information.
	miraCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: miraBirthday,
		Code:     miraCode,
		Key:      character.Mira,
		Gender:   gender.Female,
		Name:     miraName,
		Special:  false}
)

var (
	// miraAmericanEnglishPhrase represents Mira's phrase in American English.
	miraAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cottontail"}

	// miraCanadianFrenchPhrase represents Mira's phrase in Canadian French.
	miraCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "zip zip"}

	// miraDutchPhrase represents Mira's phrase in Dutch.
	miraDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "staartje"}

	// miraFrenchPhrase represents Mira's phrase in French.
	miraFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "auch auch"}

	// miraGermanPhrase represents Mira's phrase in German.
	miraGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "jahopp"}

	// miraItalianPhrase represents Mira's phrase in Italian.
	miraItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "baffoplà"}

	// miraJapanesePhrase represents Mira's phrase in Japanese.
	miraJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だムーン"}

	// miraKoreanPhrase represents Mira's phrase in Korean.
	miraKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "거봐라"}

	// miraLatinAmericanSpanishPhrase represents Mira's phrase in Latin American Spanish.
	miraLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "hep-hep"}

	// miraRussianPhrase represents Mira's phrase in Russian.
	miraRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "зайцехвост"}

	// miraSimplifiedChinesePhrase represents Mira's phrase in Simplified Chinese.
	miraSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "月球"}

	// miraSpanishPhrase represents Mira's phrase in Spanish.
	miraSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "muesli"}

	// miraTraditionalChinesePhrase represents Mira's phrase in Traditional Chinese.
	miraTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "月球"}
)

var (
	// miraPhrase represents Mira's phrases in different languages.
	miraPhrase = nook.Languages{
		language.AmericanEnglish:      miraAmericanEnglishPhrase,
		language.CanadianFrench:       miraCanadianFrenchPhrase,
		language.Dutch:                miraDutchPhrase,
		language.French:               miraFrenchPhrase,
		language.German:               miraGermanPhrase,
		language.Italian:              miraItalianPhrase,
		language.Japanese:             miraJapanesePhrase,
		language.Korean:               miraKoreanPhrase,
		language.LatinAmericanSpanish: miraLatinAmericanSpanishPhrase,
		language.Russian:              miraRussianPhrase,
		language.SimplifiedChinese:    miraSimplifiedChinesePhrase,
		language.Spanish:              miraSpanishPhrase,
		language.TraditionalChinese:   miraTraditionalChinesePhrase}
)

var (
	// Mira represents the character Mira with her complete information.
	Mira = nook.Villager{
		Character:   miraCharacter,
		Personality: personality.BigSister,
		Phrase:      miraPhrase}
)

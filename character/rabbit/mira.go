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
	miraBirthday = nook.Birthday{
		Day:   6,
		Month: time.July}
)

var (
	miraCode = nook.Code{
		Value: "rbt19"}
)

var (
	miraAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Mira"}

	miraCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Grisette"}

	miraDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Mira"}

	miraFrenchName = nook.Name{
		Language: language.French,
		Value:    "Grisette"}

	miraGermanName = nook.Name{
		Language: language.German,
		Value:    "Mira"}

	miraItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Martina"}

	miraJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ミラコ"}

	miraKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "미랑"}

	miraLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Amparo"}

	miraRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Мира"}

	miraSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "蜜拉"}

	miraSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Amparo"}

	miraTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "蜜拉"}
)

var (
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
	miraCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: miraBirthday,
		Code:     miraCode,
		Key:      character.Mira,
		Gender:   gender.Female,
		Name:     miraName}
)

var (
	miraAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cottontail"}

	miraCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "zip zip"}

	miraDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "staartje"}

	miraFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "auch auch"}

	miraGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "jahopp"}

	miraItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "baffoplà"}

	miraJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だムーン"}

	miraKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "거봐라"}

	miraLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "hep-hep"}

	miraRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "зайцехвост"}

	miraSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "月球"}

	miraSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "muesli"}

	miraTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "月球"}
)

var (
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
	Mira = nook.Villager{
		Character:   miraCharacter,
		Personality: personality.BigSister,
		Phrase:      miraPhrase}
)

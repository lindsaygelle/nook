package hamster

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
	// clayBirthday represents clay birthday.
	clayBirthday = nook.Birthday{
		Day:   19,
		Month: time.October}
)

var (
	// clayCode represents clay code.
	clayCode = nook.Code{
		Value: "ham05"}
)

var (
	// clayAmericanEnglishName represents clay american english name.
	clayAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Clay"}

	// clayCanadianFrenchName represents clay canadian french name.
	clayCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Guido"}

	// clayDutchName represents clay dutch name.
	clayDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Clay"}

	// clayFrenchName represents clay french name.
	clayFrenchName = nook.Name{
		Language: language.French,
		Value:    "Guido"}

	// clayGermanName represents clay german name.
	clayGermanName = nook.Name{
		Language: language.German,
		Value:    "Dietmar"}

	// clayItalianName represents clay italian name.
	clayItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Carmelo"}

	// clayJapaneseName represents clay japanese name.
	clayJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "どぐろう"}

	// clayKoreanName represents clay korean name.
	clayKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "햄둥"}

	// clayLatinAmericanSpanishName represents clay latin american spanish name.
	clayLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Boliche"}

	// clayRussianName represents clay russian name.
	clayRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Клэй"}

	// claySimplifiedChineseName represents clay simplified chinese name.
	claySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "子墨"}

	// claySpanishName represents clay spanish name.
	claySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Boliche"}

	// clayTraditionalChineseName represents clay traditional chinese name.
	clayTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "子墨"}
)

var (
	// clayName represents clay name.
	clayName = nook.Languages{
		language.AmericanEnglish:      clayAmericanEnglishName,
		language.CanadianFrench:       clayCanadianFrenchName,
		language.Dutch:                clayDutchName,
		language.French:               clayFrenchName,
		language.German:               clayGermanName,
		language.Italian:              clayItalianName,
		language.Japanese:             clayJapaneseName,
		language.Korean:               clayKoreanName,
		language.LatinAmericanSpanish: clayLatinAmericanSpanishName,
		language.Russian:              clayRussianName,
		language.SimplifiedChinese:    claySimplifiedChineseName,
		language.Spanish:              claySpanishName,
		language.TraditionalChinese:   clayTraditionalChineseName}
)

var (
	// clayCharacter represents clay character.
	clayCharacter = nook.Character{
		Animal:   animal.Hamster,
		Birthday: clayBirthday,
		Code:     clayCode,
		Key:      character.Clay,
		Gender:   gender.Male,
		Name:     clayName,
		Special:  false}
)

var (
	// clayAmericanEnglishPhrase represents clay american english phrase.
	clayAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "thump"}

	// clayCanadianFrenchPhrase represents clay canadian french phrase.
	clayCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pindépice"}

	// clayDutchPhrase represents clay dutch phrase.
	clayDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "boink"}

	// clayFrenchPhrase represents clay french phrase.
	clayFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pindépice"}

	// clayGermanPhrase represents clay german phrase.
	clayGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "backi"}

	// clayItalianPhrase represents clay italian phrase.
	clayItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "popomf"}

	// clayJapanesePhrase represents clay japanese phrase.
	clayJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "どきどき"}

	// clayKoreanPhrase represents clay korean phrase.
	clayKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "둥게둥게"}

	// clayLatinAmericanSpanishPhrase represents clay latin american spanish phrase.
	clayLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "tucu-tucu"}

	// clayRussianPhrase represents clay russian phrase.
	clayRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "стук-стук"}

	// claySimplifiedChinesePhrase represents clay simplified chinese phrase.
	claySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "忐忑不安"}

	// claySpanishPhrase represents clay spanish phrase.
	claySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tucu-tucu"}

	// clayTraditionalChinesePhrase represents clay traditional chinese phrase.
	clayTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "忐忑不安"}
)

var (
	// clayPhrase represents clay phrase.
	clayPhrase = nook.Languages{
		language.AmericanEnglish:      clayAmericanEnglishPhrase,
		language.CanadianFrench:       clayCanadianFrenchPhrase,
		language.Dutch:                clayDutchPhrase,
		language.French:               clayFrenchPhrase,
		language.German:               clayGermanPhrase,
		language.Italian:              clayItalianPhrase,
		language.Japanese:             clayJapanesePhrase,
		language.Korean:               clayKoreanPhrase,
		language.LatinAmericanSpanish: clayLatinAmericanSpanishPhrase,
		language.Russian:              clayRussianPhrase,
		language.SimplifiedChinese:    claySimplifiedChinesePhrase,
		language.Spanish:              claySpanishPhrase,
		language.TraditionalChinese:   clayTraditionalChinesePhrase}
)

var (
	// Clay represents clay.
	Clay = nook.Villager{
		Character:   clayCharacter,
		Personality: personality.Lazy,
		Phrase:      clayPhrase}
)

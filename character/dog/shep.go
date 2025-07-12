package dog

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
	// shepBirthday represents shep birthday.
	shepBirthday = nook.Birthday{
		Day:   24,
		Month: time.November}
)

var (
	// shepCode represents shep code.
	shepCode = nook.Code{
		Value: "dog18"}
)

var (
	// shepAmericanEnglishName represents shep american english name.
	shepAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Shep"}

	// shepCanadianFrenchName represents shep canadian french name.
	shepCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Mehdi"}

	// shepDutchName represents shep dutch name.
	shepDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Shep"}

	// shepFrenchName represents shep french name.
	shepFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mehdi"}

	// shepGermanName represents shep german name.
	shepGermanName = nook.Name{
		Language: language.German,
		Value:    "Thomas"}

	// shepItalianName represents shep italian name.
	shepItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Frangino"}

	// shepJapaneseName represents shep japanese name.
	shepJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ボブ"}

	// shepKoreanName represents shep korean name.
	shepKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "밥"}

	// shepLatinAmericanSpanishName represents shep latin american spanish name.
	shepLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Fleco"}

	// shepRussianName represents shep russian name.
	shepRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Шеп"}

	// shepSimplifiedChineseName represents shep simplified chinese name.
	shepSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "包伯"}

	// shepSpanishName represents shep spanish name.
	shepSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Fleco"}

	// shepTraditionalChineseName represents shep traditional chinese name.
	shepTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "包伯"}
)

var (
	// shepName represents shep name.
	shepName = nook.Languages{
		language.AmericanEnglish:      shepAmericanEnglishName,
		language.CanadianFrench:       shepCanadianFrenchName,
		language.Dutch:                shepDutchName,
		language.French:               shepFrenchName,
		language.German:               shepGermanName,
		language.Italian:              shepItalianName,
		language.Japanese:             shepJapaneseName,
		language.Korean:               shepKoreanName,
		language.LatinAmericanSpanish: shepLatinAmericanSpanishName,
		language.Russian:              shepRussianName,
		language.SimplifiedChinese:    shepSimplifiedChineseName,
		language.Spanish:              shepSpanishName,
		language.TraditionalChinese:   shepTraditionalChineseName}
)

var (
	// shepCharacter represents shep character.
	shepCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: shepBirthday,
		Code:     shepCode,
		Key:      character.Shep,
		Gender:   gender.Male,
		Name:     shepName,
		Special:  false}
)

var (
	// shepAmericanEnglishPhrase represents shep american english phrase.
	shepAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "baaa man"}

	// shepCanadianFrenchPhrase represents shep canadian french phrase.
	shepCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pedigree"}

	// shepDutchPhrase represents shep dutch phrase.
	shepDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "schapie"}

	// shepFrenchPhrase represents shep french phrase.
	shepFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pedigree"}

	// shepGermanPhrase represents shep german phrase.
	shepGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "heffheff"}

	// shepItalianPhrase represents shep italian phrase.
	shepItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "guau guau"}

	// shepJapanesePhrase represents shep japanese phrase.
	shepJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ワンダー"}

	// shepKoreanPhrase represents shep korean phrase.
	shepKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "안보여"}

	// shepLatinAmericanSpanishPhrase represents shep latin american spanish phrase.
	shepLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "grrruau"}

	// shepRussianPhrase represents shep russian phrase.
	shepRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бр-р-разер"}

	// shepSimplifiedChinesePhrase represents shep simplified chinese phrase.
	shepSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "汪想想"}

	// shepSpanishPhrase represents shep spanish phrase.
	shepSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tufos"}

	// shepTraditionalChinesePhrase represents shep traditional chinese phrase.
	shepTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "汪想想"}
)

var (
	// shepPhrase represents shep phrase.
	shepPhrase = nook.Languages{
		language.AmericanEnglish:      shepAmericanEnglishPhrase,
		language.CanadianFrench:       shepCanadianFrenchPhrase,
		language.Dutch:                shepDutchPhrase,
		language.French:               shepFrenchPhrase,
		language.German:               shepGermanPhrase,
		language.Italian:              shepItalianPhrase,
		language.Japanese:             shepJapanesePhrase,
		language.Korean:               shepKoreanPhrase,
		language.LatinAmericanSpanish: shepLatinAmericanSpanishPhrase,
		language.Russian:              shepRussianPhrase,
		language.SimplifiedChinese:    shepSimplifiedChinesePhrase,
		language.Spanish:              shepSpanishPhrase,
		language.TraditionalChinese:   shepTraditionalChinesePhrase}
)

var (
	// Shep represents shep.
	Shep = nook.Villager{
		Character:   shepCharacter,
		Personality: personality.Smug,
		Phrase:      shepPhrase}
)

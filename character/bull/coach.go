package bull

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
	// coachBirthday represents coach birthday.
	coachBirthday = nook.Birthday{
		Day:   29,
		Month: time.April}
)

var (
	// coachCode represents coach code.
	coachCode = nook.Code{
		Value: "bul07"}
)

var (
	// coachAmericanEnglishName represents coach american english name.
	coachAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Coach"}

	// coachCanadianFrenchName represents coach canadian french name.
	coachCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Arnold"}

	// coachDutchName represents coach dutch name.
	coachDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Coach"}

	// coachFrenchName represents coach french name.
	coachFrenchName = nook.Name{
		Language: language.French,
		Value:    "Arnold"}

	// coachGermanName represents coach german name.
	coachGermanName = nook.Name{
		Language: language.German,
		Value:    "Arnold"}

	// coachItalianName represents coach italian name.
	coachItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ercole"}

	// coachJapaneseName represents coach japanese name.
	coachJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "テッチャン"}

	// coachKoreanName represents coach korean name.
	coachKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "철소"}

	// coachLatinAmericanSpanishName represents coach latin american spanish name.
	coachLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cacho"}

	// coachRussianName represents coach russian name.
	coachRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Коуч"}

	// coachSimplifiedChineseName represents coach simplified chinese name.
	coachSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "大常"}

	// coachSpanishName represents coach spanish name.
	coachSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cacho"}

	// coachTraditionalChineseName represents coach traditional chinese name.
	coachTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "大常"}
)

var (
	// coachName represents coach name.
	coachName = nook.Languages{
		language.AmericanEnglish:      coachAmericanEnglishName,
		language.CanadianFrench:       coachCanadianFrenchName,
		language.Dutch:                coachDutchName,
		language.French:               coachFrenchName,
		language.German:               coachGermanName,
		language.Italian:              coachItalianName,
		language.Japanese:             coachJapaneseName,
		language.Korean:               coachKoreanName,
		language.LatinAmericanSpanish: coachLatinAmericanSpanishName,
		language.Russian:              coachRussianName,
		language.SimplifiedChinese:    coachSimplifiedChineseName,
		language.Spanish:              coachSpanishName,
		language.TraditionalChinese:   coachTraditionalChineseName}
)

var (
	// coachCharacter represents coach character.
	coachCharacter = nook.Character{
		Animal:   animal.Bull,
		Birthday: coachBirthday,
		Code:     coachCode,
		Key:      character.Coach,
		Gender:   gender.Male,
		Name:     coachName,
		Special:  false}
)

var (
	// coachAmericanEnglishPhrase represents coach american english phrase.
	coachAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "stubble"}

	// coachCanadianFrenchPhrase represents coach canadian french phrase.
	coachCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bétail"}

	// coachDutchPhrase represents coach dutch phrase.
	coachDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "stoppel"}

	// coachFrenchPhrase represents coach french phrase.
	coachFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mon cou"}

	// coachGermanPhrase represents coach german phrase.
	coachGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnoff"}

	// coachItalianPhrase represents coach italian phrase.
	coachItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "oooh issa"}

	// coachJapanesePhrase represents coach japanese phrase.
	coachJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ジョリッ"}

	// coachKoreanPhrase represents coach korean phrase.
	coachKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "땅땅"}

	// coachLatinAmericanSpanishPhrase represents coach latin american spanish phrase.
	coachLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "chacho"}

	// coachRussianPhrase represents coach russian phrase.
	coachRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "раз-два"}

	// coachSimplifiedChinesePhrase represents coach simplified chinese phrase.
	coachSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "胡渣"}

	// coachSpanishPhrase represents coach spanish phrase.
	coachSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chacho"}

	// coachTraditionalChinesePhrase represents coach traditional chinese phrase.
	coachTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鬍渣"}
)

var (
	// coachPhrase represents coach phrase.
	coachPhrase = nook.Languages{
		language.AmericanEnglish:      coachAmericanEnglishPhrase,
		language.CanadianFrench:       coachCanadianFrenchPhrase,
		language.Dutch:                coachDutchPhrase,
		language.French:               coachFrenchPhrase,
		language.German:               coachGermanPhrase,
		language.Italian:              coachItalianPhrase,
		language.Japanese:             coachJapanesePhrase,
		language.Korean:               coachKoreanPhrase,
		language.LatinAmericanSpanish: coachLatinAmericanSpanishPhrase,
		language.Russian:              coachRussianPhrase,
		language.SimplifiedChinese:    coachSimplifiedChinesePhrase,
		language.Spanish:              coachSpanishPhrase,
		language.TraditionalChinese:   coachTraditionalChinesePhrase}
)

var (
	// Coach represents coach.
	Coach = nook.Villager{
		Character:   coachCharacter,
		Personality: personality.Jock,
		Phrase:      coachPhrase}
)

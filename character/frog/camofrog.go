package frog

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
	// camofrogBirthday represents camofrog birthday.
	camofrogBirthday = nook.Birthday{
		Day:   5,
		Month: time.June}
)

var (
	// camofrogCode represents camofrog code.
	camofrogCode = nook.Code{
		Value: "flg03"}
)

var (
	// camofrogAmericanEnglishName represents camofrog american english name.
	camofrogAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Camofrog"}

	// camofrogCanadianFrenchName represents camofrog canadian french name.
	camofrogCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Milos"}

	// camofrogDutchName represents camofrog dutch name.
	camofrogDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Camofrog"}

	// camofrogFrenchName represents camofrog french name.
	camofrogFrenchName = nook.Name{
		Language: language.French,
		Value:    "Milos"}

	// camofrogGermanName represents camofrog german name.
	camofrogGermanName = nook.Name{
		Language: language.German,
		Value:    "Tarno"}

	// camofrogItalianName represents camofrog italian name.
	camofrogItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Crambo"}

	// camofrogJapaneseName represents camofrog japanese name.
	camofrogJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フルメタル"}

	// camofrogKoreanName represents camofrog korean name.
	camofrogKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "충성"}

	// camofrogLatinAmericanSpanishName represents camofrog latin american spanish name.
	camofrogLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Comando"}

	// camofrogRussianName represents camofrog russian name.
	camofrogRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Камофрог"}

	// camofrogSimplifiedChineseName represents camofrog simplified chinese name.
	camofrogSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "迷仔"}

	// camofrogSpanishName represents camofrog spanish name.
	camofrogSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Comando"}

	// camofrogTraditionalChineseName represents camofrog traditional chinese name.
	camofrogTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "迷仔"}
)

var (
	// camofrogName represents camofrog name.
	camofrogName = nook.Languages{
		language.AmericanEnglish:      camofrogAmericanEnglishName,
		language.CanadianFrench:       camofrogCanadianFrenchName,
		language.Dutch:                camofrogDutchName,
		language.French:               camofrogFrenchName,
		language.German:               camofrogGermanName,
		language.Italian:              camofrogItalianName,
		language.Japanese:             camofrogJapaneseName,
		language.Korean:               camofrogKoreanName,
		language.LatinAmericanSpanish: camofrogLatinAmericanSpanishName,
		language.Russian:              camofrogRussianName,
		language.SimplifiedChinese:    camofrogSimplifiedChineseName,
		language.Spanish:              camofrogSpanishName,
		language.TraditionalChinese:   camofrogTraditionalChineseName}
)

var (
	// camofrogCharacter represents camofrog character.
	camofrogCharacter = nook.Character{
		Animal:   animal.Frog,
		Birthday: camofrogBirthday,
		Code:     camofrogCode,
		Key:      character.Camofrog,
		Gender:   gender.Male,
		Name:     camofrogName,
		Special:  false}
)

var (
	// camofrogAmericanEnglishPhrase represents camofrog american english phrase.
	camofrogAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ten-hut"}

	// camofrogCanadianFrenchPhrase represents camofrog canadian french phrase.
	camofrogCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "kaki"}

	// camofrogDutchPhrase represents camofrog dutch phrase.
	camofrogDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "geef acht"}

	// camofrogFrenchPhrase represents camofrog french phrase.
	camofrogFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "taïaut"}

	// camofrogGermanPhrase represents camofrog german phrase.
	camofrogGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "quak-quak"}

	// camofrogItalianPhrase represents camofrog italian phrase.
	camofrogItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "attacroak"}

	// camofrogJapanesePhrase represents camofrog japanese phrase.
	camofrogJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "わい"}

	// camofrogKoreanPhrase represents camofrog korean phrase.
	camofrogKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "와이"}

	// camofrogLatinAmericanSpanishPhrase represents camofrog latin american spanish phrase.
	camofrogLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "croooa"}

	// camofrogRussianPhrase represents camofrog russian phrase.
	camofrogRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "шагом-квак"}

	// camofrogSimplifiedChinesePhrase represents camofrog simplified chinese phrase.
	camofrogSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喂"}

	// camofrogSpanishPhrase represents camofrog spanish phrase.
	camofrogSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "recluta"}

	// camofrogTraditionalChinesePhrase represents camofrog traditional chinese phrase.
	camofrogTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "喂"}
)

var (
	// camofrogPhrase represents camofrog phrase.
	camofrogPhrase = nook.Languages{
		language.AmericanEnglish:      camofrogAmericanEnglishPhrase,
		language.CanadianFrench:       camofrogCanadianFrenchPhrase,
		language.Dutch:                camofrogDutchPhrase,
		language.French:               camofrogFrenchPhrase,
		language.German:               camofrogGermanPhrase,
		language.Italian:              camofrogItalianPhrase,
		language.Japanese:             camofrogJapanesePhrase,
		language.Korean:               camofrogKoreanPhrase,
		language.LatinAmericanSpanish: camofrogLatinAmericanSpanishPhrase,
		language.Russian:              camofrogRussianPhrase,
		language.SimplifiedChinese:    camofrogSimplifiedChinesePhrase,
		language.Spanish:              camofrogSpanishPhrase,
		language.TraditionalChinese:   camofrogTraditionalChinesePhrase}
)

var (
	// Camofrog represents camofrog.
	Camofrog = nook.Villager{
		Character:   camofrogCharacter,
		Personality: personality.Cranky,
		Phrase:      camofrogPhrase}
)

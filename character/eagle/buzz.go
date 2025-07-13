package eagle

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
	// buzzBirthday represents buzz birthday.
	buzzBirthday = nook.Birthday{
		Day:   7,
		Month: time.December}
)

var (
	// buzzCode represents buzz code.
	buzzCode = nook.Code{
		Value: "pbr03"}
)

var (
	// buzzAmericanEnglishName represents buzz american english name.
	buzzAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Buzz"}

	// buzzCanadianFrenchName represents buzz canadian french name.
	buzzCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Phébus"}

	// buzzDutchName represents buzz dutch name.
	buzzDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Buzz"}

	// buzzFrenchName represents buzz french name.
	buzzFrenchName = nook.Name{
		Language: language.French,
		Value:    "Phébus"}

	// buzzGermanName represents buzz german name.
	buzzGermanName = nook.Name{
		Language: language.German,
		Value:    "Balduin"}

	// buzzItalianName represents buzz italian name.
	buzzItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Golia"}

	// buzzJapaneseName represents buzz japanese name.
	buzzJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ひでよし"}

	// buzzKoreanName represents buzz korean name.
	buzzKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "근엄"}

	// buzzLatinAmericanSpanishName represents buzz latin american spanish name.
	buzzLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Nabar"}

	// buzzRussianName represents buzz russian name.
	buzzRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Базз"}

	// buzzSimplifiedChineseName represents buzz simplified chinese name.
	buzzSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "继光"}

	// buzzSpanishName represents buzz spanish name.
	buzzSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Nabar"}

	// buzzTraditionalChineseName represents buzz traditional chinese name.
	buzzTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "繼光"}
)

var (
	// buzzName represents buzz name.
	buzzName = nook.Languages{
		language.AmericanEnglish:      buzzAmericanEnglishName,
		language.CanadianFrench:       buzzCanadianFrenchName,
		language.Dutch:                buzzDutchName,
		language.French:               buzzFrenchName,
		language.German:               buzzGermanName,
		language.Italian:              buzzItalianName,
		language.Japanese:             buzzJapaneseName,
		language.Korean:               buzzKoreanName,
		language.LatinAmericanSpanish: buzzLatinAmericanSpanishName,
		language.Russian:              buzzRussianName,
		language.SimplifiedChinese:    buzzSimplifiedChineseName,
		language.Spanish:              buzzSpanishName,
		language.TraditionalChinese:   buzzTraditionalChineseName}
)

var (
	// buzzCharacter represents buzz character.
	buzzCharacter = nook.Character{
		Animal:   animal.Eagle,
		Birthday: buzzBirthday,
		Code:     buzzCode,
		Key:      character.Buzz,
		Gender:   gender.Male,
		Name:     buzzName,
		Special:  false}
)

var (
	// buzzAmericanEnglishPhrase represents buzz american english phrase.
	buzzAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "captain"}

	// buzzCanadianFrenchPhrase represents buzz canadian french phrase.
	buzzCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "flap flap"}

	// buzzDutchPhrase represents buzz dutch phrase.
	buzzDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kaptein"}

	// buzzFrenchPhrase represents buzz french phrase.
	buzzFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "flap flap"}

	// buzzGermanPhrase represents buzz german phrase.
	buzzGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "käpten"}

	// buzzItalianPhrase represents buzz italian phrase.
	buzzItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "capo"}

	// buzzJapanesePhrase represents buzz japanese phrase.
	buzzJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ッキーッ"}

	// buzzKoreanPhrase represents buzz korean phrase.
	buzzKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "짜란～"}

	// buzzLatinAmericanSpanishPhrase represents buzz latin american spanish phrase.
	buzzLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "capííí"}

	// buzzRussianPhrase represents buzz russian phrase.
	buzzRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "капитан"}

	// buzzSimplifiedChinesePhrase represents buzz simplified chinese phrase.
	buzzSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "继继"}

	// buzzSpanishPhrase represents buzz spanish phrase.
	buzzSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cadete"}

	// buzzTraditionalChinesePhrase represents buzz traditional chinese phrase.
	buzzTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "繼繼"}
)

var (
	// buzzPhrase represents buzz phrase.
	buzzPhrase = nook.Languages{
		language.AmericanEnglish:      buzzAmericanEnglishPhrase,
		language.CanadianFrench:       buzzCanadianFrenchPhrase,
		language.Dutch:                buzzDutchPhrase,
		language.French:               buzzFrenchPhrase,
		language.German:               buzzGermanPhrase,
		language.Italian:              buzzItalianPhrase,
		language.Japanese:             buzzJapanesePhrase,
		language.Korean:               buzzKoreanPhrase,
		language.LatinAmericanSpanish: buzzLatinAmericanSpanishPhrase,
		language.Russian:              buzzRussianPhrase,
		language.SimplifiedChinese:    buzzSimplifiedChinesePhrase,
		language.Spanish:              buzzSpanishPhrase,
		language.TraditionalChinese:   buzzTraditionalChinesePhrase}
)

var (
	// Buzz represents buzz.
	Buzz = nook.Villager{
		Character:   buzzCharacter,
		Personality: personality.Cranky,
		Phrase:      buzzPhrase}
)

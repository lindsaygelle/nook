package hippo

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
	// harryBirthday represents harry birthday.
	harryBirthday = nook.Birthday{
		Day:   7,
		Month: time.January}
)

var (
	// harryCode represents harry code.
	harryCode = nook.Code{
		Value: "hip08"}
)

var (
	// harryAmericanEnglishName represents harry american english name.
	harryAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Harry"}

	// harryCanadianFrenchName represents harry canadian french name.
	harryCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bob"}

	// harryDutchName represents harry dutch name.
	harryDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Harry"}

	// harryFrenchName represents harry french name.
	harryFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bob"}

	// harryGermanName represents harry german name.
	harryGermanName = nook.Name{
		Language: language.German,
		Value:    "Jürgen"}

	// harryItalianName represents harry italian name.
	harryItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ercolino"}

	// harryJapaneseName represents harry japanese name.
	harryJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "オリバー"}

	// harryKoreanName represents harry korean name.
	harryKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "올리버"}

	// harryLatinAmericanSpanishName represents harry latin american spanish name.
	harryLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Harpo"}

	// harryRussianName represents harry russian name.
	harryRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Гарри"}

	// harrySimplifiedChineseName represents harry simplified chinese name.
	harrySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "欧世豪"}

	// harrySpanishName represents harry spanish name.
	harrySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Harpo"}

	// harryTraditionalChineseName represents harry traditional chinese name.
	harryTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "歐世豪"}
)

var (
	// harryName represents harry name.
	harryName = nook.Languages{
		language.AmericanEnglish:      harryAmericanEnglishName,
		language.CanadianFrench:       harryCanadianFrenchName,
		language.Dutch:                harryDutchName,
		language.French:               harryFrenchName,
		language.German:               harryGermanName,
		language.Italian:              harryItalianName,
		language.Japanese:             harryJapaneseName,
		language.Korean:               harryKoreanName,
		language.LatinAmericanSpanish: harryLatinAmericanSpanishName,
		language.Russian:              harryRussianName,
		language.SimplifiedChinese:    harrySimplifiedChineseName,
		language.Spanish:              harrySpanishName,
		language.TraditionalChinese:   harryTraditionalChineseName}
)

var (
	// harryCharacter represents harry character.
	harryCharacter = nook.Character{
		Animal:   animal.Hippo,
		Birthday: harryBirthday,
		Code:     harryCode,
		Key:      character.Harry,
		Gender:   gender.Male,
		Name:     harryName,
		Special:  false}
)

var (
	// harryAmericanEnglishPhrase represents harry american english phrase.
	harryAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "beach bum"}

	// harryCanadianFrenchPhrase represents harry canadian french phrase.
	harryCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mille"}

	// harryDutchPhrase represents harry dutch phrase.
	harryDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "strandgast"}

	// harryFrenchPhrase represents harry french phrase.
	harryFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mille"}

	// harryGermanPhrase represents harry german phrase.
	harryGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "höhöhö"}

	// harryItalianPhrase represents harry italian phrase.
	harryItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "uhuhuh"}

	// harryJapanesePhrase represents harry japanese phrase.
	harryJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でっせ"}

	// harryKoreanPhrase represents harry korean phrase.
	harryKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "냉큼오슈"}

	// harryLatinAmericanSpanishPhrase represents harry latin american spanish phrase.
	harryLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ho-hop"}

	// harryRussianPhrase represents harry russian phrase.
	harryRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "лежебока"}

	// harrySimplifiedChinesePhrase represents harry simplified chinese phrase.
	harrySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "世世"}

	// harrySpanishPhrase represents harry spanish phrase.
	harrySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "algas"}

	// harryTraditionalChinesePhrase represents harry traditional chinese phrase.
	harryTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "世世"}
)

var (
	// harryPhrase represents harry phrase.
	harryPhrase = nook.Languages{
		language.AmericanEnglish:      harryAmericanEnglishPhrase,
		language.CanadianFrench:       harryCanadianFrenchPhrase,
		language.Dutch:                harryDutchPhrase,
		language.French:               harryFrenchPhrase,
		language.German:               harryGermanPhrase,
		language.Italian:              harryItalianPhrase,
		language.Japanese:             harryJapanesePhrase,
		language.Korean:               harryKoreanPhrase,
		language.LatinAmericanSpanish: harryLatinAmericanSpanishPhrase,
		language.Russian:              harryRussianPhrase,
		language.SimplifiedChinese:    harrySimplifiedChinesePhrase,
		language.Spanish:              harrySpanishPhrase,
		language.TraditionalChinese:   harryTraditionalChinesePhrase}
)

var (
	// Harry represents harry.
	Harry = nook.Villager{
		Character:   harryCharacter,
		Personality: personality.Cranky,
		Phrase:      harryPhrase}
)

package mouse

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
	// mooseBirthday represents moose birthday.
	mooseBirthday = nook.Birthday{
		Day:   13,
		Month: time.September}
)

var (
	// mooseCode represents moose code.
	mooseCode = nook.Code{
		Value: "mus14"}
)

var (
	// mooseAmericanEnglishName represents moose american english name.
	mooseAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Moose"}

	// mooseCanadianFrenchName represents moose canadian french name.
	mooseCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Joachim"}

	// mooseDutchName represents moose dutch name.
	mooseDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Moose"}

	// mooseFrenchName represents moose french name.
	mooseFrenchName = nook.Name{
		Language: language.French,
		Value:    "Joachim"}

	// mooseGermanName represents moose german name.
	mooseGermanName = nook.Name{
		Language: language.German,
		Value:    "Mausbert"}

	// mooseItalianName represents moose italian name.
	mooseItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Aldo"}

	// mooseJapaneseName represents moose japanese name.
	mooseJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ピン"}

	// mooseKoreanName represents moose korean name.
	mooseKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "핑"}

	// mooseLatinAmericanSpanishName represents moose latin american spanish name.
	mooseLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Sato"}

	// mooseRussianName represents moose russian name.
	mooseRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Муз"}

	// mooseSimplifiedChineseName represents moose simplified chinese name.
	mooseSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿聘"}

	// mooseSpanishName represents moose spanish name.
	mooseSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Sato"}

	// mooseTraditionalChineseName represents moose traditional chinese name.
	mooseTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿聘"}
)

var (
	// mooseName represents moose name.
	mooseName = nook.Languages{
		language.AmericanEnglish:      mooseAmericanEnglishName,
		language.CanadianFrench:       mooseCanadianFrenchName,
		language.Dutch:                mooseDutchName,
		language.French:               mooseFrenchName,
		language.German:               mooseGermanName,
		language.Italian:              mooseItalianName,
		language.Japanese:             mooseJapaneseName,
		language.Korean:               mooseKoreanName,
		language.LatinAmericanSpanish: mooseLatinAmericanSpanishName,
		language.Russian:              mooseRussianName,
		language.SimplifiedChinese:    mooseSimplifiedChineseName,
		language.Spanish:              mooseSpanishName,
		language.TraditionalChinese:   mooseTraditionalChineseName}
)

var (
	// mooseCharacter represents moose character.
	mooseCharacter = nook.Character{
		Animal:   animal.Mouse,
		Birthday: mooseBirthday,
		Code:     mooseCode,
		Key:      character.Moose,
		Gender:   gender.Male,
		Name:     mooseName,
		Special:  false}
)

var (
	// mooseAmericanEnglishPhrase represents moose american english phrase.
	mooseAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "shorty"}

	// mooseCanadianFrenchPhrase represents moose canadian french phrase.
	mooseCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bééléé"}

	// mooseDutchPhrase represents moose dutch phrase.
	mooseDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kleintje"}

	// mooseFrenchPhrase represents moose french phrase.
	mooseFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bééléé"}

	// mooseGermanPhrase represents moose german phrase.
	mooseGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "käääse"}

	// mooseItalianPhrase represents moose italian phrase.
	mooseItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "quink"}

	// mooseJapanesePhrase represents moose japanese phrase.
	mooseJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "にゃろ"}

	// mooseKoreanPhrase represents moose korean phrase.
	mooseKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "짜샤"}

	// mooseLatinAmericanSpanishPhrase represents moose latin american spanish phrase.
	mooseLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "atiza"}

	// mooseRussianPhrase represents moose russian phrase.
	mooseRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "коротышка"}

	// mooseSimplifiedChinesePhrase represents moose simplified chinese phrase.
	mooseSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "可恶"}

	// mooseSpanishPhrase represents moose spanish phrase.
	mooseSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "atiza"}

	// mooseTraditionalChinesePhrase represents moose traditional chinese phrase.
	mooseTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "可惡"}
)

var (
	// moosePhrase represents moose phrase.
	moosePhrase = nook.Languages{
		language.AmericanEnglish:      mooseAmericanEnglishPhrase,
		language.CanadianFrench:       mooseCanadianFrenchPhrase,
		language.Dutch:                mooseDutchPhrase,
		language.French:               mooseFrenchPhrase,
		language.German:               mooseGermanPhrase,
		language.Italian:              mooseItalianPhrase,
		language.Japanese:             mooseJapanesePhrase,
		language.Korean:               mooseKoreanPhrase,
		language.LatinAmericanSpanish: mooseLatinAmericanSpanishPhrase,
		language.Russian:              mooseRussianPhrase,
		language.SimplifiedChinese:    mooseSimplifiedChinesePhrase,
		language.Spanish:              mooseSpanishPhrase,
		language.TraditionalChinese:   mooseTraditionalChinesePhrase}
)

var (
	// Moose represents moose.
	Moose = nook.Villager{
		Character:   mooseCharacter,
		Personality: personality.Jock,
		Phrase:      moosePhrase}
)

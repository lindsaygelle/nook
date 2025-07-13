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
	// chadderBirthday represents chadder birthday.
	chadderBirthday = nook.Birthday{
		Day:   15,
		Month: time.December}
)

var (
	// chadderCode represents chadder code.
	chadderCode = nook.Code{
		Value: "mus18"}
)

var (
	// chadderAmericanEnglishName represents chadder american english name.
	chadderAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Chadder"}

	// chadderCanadianFrenchName represents chadder canadian french name.
	chadderCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Mozzar"}

	// chadderDutchName represents chadder dutch name.
	chadderDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Chadder"}

	// chadderFrenchName represents chadder french name.
	chadderFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mozzar"}

	// chadderGermanName represents chadder german name.
	chadderGermanName = nook.Name{
		Language: language.German,
		Value:    "Charlie"}

	// chadderItalianName represents chadder italian name.
	chadderItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gruviero"}

	// chadderJapaneseName represents chadder japanese name.
	chadderJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チーズ"}

	// chadderKoreanName represents chadder korean name.
	chadderKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "치즈"}

	// chadderLatinAmericanSpanishName represents chadder latin american spanish name.
	chadderLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Roque"}

	// chadderRussianName represents chadder russian name.
	chadderRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Чаддер"}

	// chadderSimplifiedChineseName represents chadder simplified chinese name.
	chadderSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "起司"}

	// chadderSpanishName represents chadder spanish name.
	chadderSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Roque"}

	// chadderTraditionalChineseName represents chadder traditional chinese name.
	chadderTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "起司"}
)

var (
	// chadderName represents chadder name.
	chadderName = nook.Languages{
		language.AmericanEnglish:      chadderAmericanEnglishName,
		language.CanadianFrench:       chadderCanadianFrenchName,
		language.Dutch:                chadderDutchName,
		language.French:               chadderFrenchName,
		language.German:               chadderGermanName,
		language.Italian:              chadderItalianName,
		language.Japanese:             chadderJapaneseName,
		language.Korean:               chadderKoreanName,
		language.LatinAmericanSpanish: chadderLatinAmericanSpanishName,
		language.Russian:              chadderRussianName,
		language.SimplifiedChinese:    chadderSimplifiedChineseName,
		language.Spanish:              chadderSpanishName,
		language.TraditionalChinese:   chadderTraditionalChineseName}
)

var (
	// chadderCharacter represents chadder character.
	chadderCharacter = nook.Character{
		Animal:   animal.Mouse,
		Birthday: chadderBirthday,
		Code:     chadderCode,
		Key:      character.Chadder,
		Gender:   gender.Male,
		Name:     chadderName,
		Special:  false}
)

var (
	// chadderAmericanEnglishPhrase represents chadder american english phrase.
	chadderAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "fromage"}

	// chadderCanadianFrenchPhrase represents chadder canadian french phrase.
	chadderCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "cheese"}

	// chadderDutchPhrase represents chadder dutch phrase.
	chadderDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "cheese"}

	// chadderFrenchPhrase represents chadder french phrase.
	chadderFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "frometon"}

	// chadderGermanPhrase represents chadder german phrase.
	chadderGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hoho"}

	// chadderItalianPhrase represents chadder italian phrase.
	chadderItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "iiik"}

	// chadderJapanesePhrase represents chadder japanese phrase.
	chadderJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "まあね"}

	// chadderKoreanPhrase represents chadder korean phrase.
	chadderKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "꼬리꼬리"}

	// chadderLatinAmericanSpanishPhrase represents chadder latin american spanish phrase.
	chadderLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "fundifún"}

	// chadderRussianPhrase represents chadder russian phrase.
	chadderRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "сырок"}

	// chadderSimplifiedChinesePhrase represents chadder simplified chinese phrase.
	chadderSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "也是啦"}

	// chadderSpanishPhrase represents chadder spanish phrase.
	chadderSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "fundifún"}

	// chadderTraditionalChinesePhrase represents chadder traditional chinese phrase.
	chadderTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "也是啦"}
)

var (
	// chadderPhrase represents chadder phrase.
	chadderPhrase = nook.Languages{
		language.AmericanEnglish:      chadderAmericanEnglishPhrase,
		language.CanadianFrench:       chadderCanadianFrenchPhrase,
		language.Dutch:                chadderDutchPhrase,
		language.French:               chadderFrenchPhrase,
		language.German:               chadderGermanPhrase,
		language.Italian:              chadderItalianPhrase,
		language.Japanese:             chadderJapanesePhrase,
		language.Korean:               chadderKoreanPhrase,
		language.LatinAmericanSpanish: chadderLatinAmericanSpanishPhrase,
		language.Russian:              chadderRussianPhrase,
		language.SimplifiedChinese:    chadderSimplifiedChinesePhrase,
		language.Spanish:              chadderSpanishPhrase,
		language.TraditionalChinese:   chadderTraditionalChinesePhrase}
)

var (
	// Chadder represents chadder.
	Chadder = nook.Villager{
		Character:   chadderCharacter,
		Personality: personality.Smug,
		Phrase:      chadderPhrase}
)

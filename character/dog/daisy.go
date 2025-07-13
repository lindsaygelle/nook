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
	// daisyBirthday represents daisy birthday.
	daisyBirthday = nook.Birthday{
		Day:   16,
		Month: time.November}
)

var (
	// daisyCode represents daisy code.
	daisyCode = nook.Code{
		Value: "dog07"}
)

var (
	// daisyAmericanEnglishName represents daisy american english name.
	daisyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Daisy"}

	// daisyCanadianFrenchName represents daisy canadian french name.
	daisyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Naomie"}

	// daisyDutchName represents daisy dutch name.
	daisyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Daisy"}

	// daisyFrenchName represents daisy french name.
	daisyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Naomie"}

	// daisyGermanName represents daisy german name.
	daisyGermanName = nook.Name{
		Language: language.German,
		Value:    "Doris"}

	// daisyItalianName represents daisy italian name.
	daisyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Fiorella"}

	// daisyJapaneseName represents daisy japanese name.
	daisyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "バニラ"}

	// daisyKoreanName represents daisy korean name.
	daisyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "바닐라"}

	// daisyLatinAmericanSpanishName represents daisy latin american spanish name.
	daisyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Luisa"}

	// daisyRussianName represents daisy russian name.
	daisyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дейзи"}

	// daisySimplifiedChineseName represents daisy simplified chinese name.
	daisySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "香草"}

	// daisySpanishName represents daisy spanish name.
	daisySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Luisa"}

	// daisyTraditionalChineseName represents daisy traditional chinese name.
	daisyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "香草"}
)

var (
	// daisyName represents daisy name.
	daisyName = nook.Languages{
		language.AmericanEnglish:      daisyAmericanEnglishName,
		language.CanadianFrench:       daisyCanadianFrenchName,
		language.Dutch:                daisyDutchName,
		language.French:               daisyFrenchName,
		language.German:               daisyGermanName,
		language.Italian:              daisyItalianName,
		language.Japanese:             daisyJapaneseName,
		language.Korean:               daisyKoreanName,
		language.LatinAmericanSpanish: daisyLatinAmericanSpanishName,
		language.Russian:              daisyRussianName,
		language.SimplifiedChinese:    daisySimplifiedChineseName,
		language.Spanish:              daisySpanishName,
		language.TraditionalChinese:   daisyTraditionalChineseName}
)

var (
	// daisyCharacter represents daisy character.
	daisyCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: daisyBirthday,
		Code:     daisyCode,
		Key:      character.Daisy,
		Gender:   gender.Female,
		Name:     daisyName,
		Special:  false}
)

var (
	// daisyAmericanEnglishPhrase represents daisy american english phrase.
	daisyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bow-WOW"}

	// daisyCanadianFrenchPhrase represents daisy canadian french phrase.
	daisyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mon chou"}

	// daisyDutchPhrase represents daisy dutch phrase.
	daisyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hazewind"}

	// daisyFrenchPhrase represents daisy french phrase.
	daisyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mon chou"}

	// daisyGermanPhrase represents daisy german phrase.
	daisyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "wauwau"}

	// daisyItalianPhrase represents daisy italian phrase.
	daisyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bau WOW"}

	// daisyJapanesePhrase represents daisy japanese phrase.
	daisyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だよね"}

	// daisyKoreanPhrase represents daisy korean phrase.
	daisyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그렇죠"}

	// daisyLatinAmericanSpanishPhrase represents daisy latin american spanish phrase.
	daisyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "guaucito"}

	// daisyRussianPhrase represents daisy russian phrase.
	daisyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "тяв-ого"}

	// daisySimplifiedChinesePhrase represents daisy simplified chinese phrase.
	daisySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "对不对"}

	// daisySpanishPhrase represents daisy spanish phrase.
	daisySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "guaay"}

	// daisyTraditionalChinesePhrase represents daisy traditional chinese phrase.
	daisyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "對不對"}
)

var (
	// daisyPhrase represents daisy phrase.
	daisyPhrase = nook.Languages{
		language.AmericanEnglish:      daisyAmericanEnglishPhrase,
		language.CanadianFrench:       daisyCanadianFrenchPhrase,
		language.Dutch:                daisyDutchPhrase,
		language.French:               daisyFrenchPhrase,
		language.German:               daisyGermanPhrase,
		language.Italian:              daisyItalianPhrase,
		language.Japanese:             daisyJapanesePhrase,
		language.Korean:               daisyKoreanPhrase,
		language.LatinAmericanSpanish: daisyLatinAmericanSpanishPhrase,
		language.Russian:              daisyRussianPhrase,
		language.SimplifiedChinese:    daisySimplifiedChinesePhrase,
		language.Spanish:              daisySpanishPhrase,
		language.TraditionalChinese:   daisyTraditionalChinesePhrase}
)

var (
	// Daisy represents daisy.
	Daisy = nook.Villager{
		Character:   daisyCharacter,
		Personality: personality.Normal,
		Phrase:      daisyPhrase}
)

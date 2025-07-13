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
	// weldonBirthday represents weldon birthday.
	weldonBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// weldonCode represents weldon code.
	weldonCode = nook.Code{
		Value: ""}
)

var (
	// weldonAmericanEnglishName represents weldon american english name.
	weldonAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Weldon"}

	// weldonCanadianFrenchName represents weldon canadian french name.
	weldonCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// weldonDutchName represents weldon dutch name.
	weldonDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// weldonFrenchName represents weldon french name.
	weldonFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	// weldonGermanName represents weldon german name.
	weldonGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	// weldonItalianName represents weldon italian name.
	weldonItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// weldonJapaneseName represents weldon japanese name.
	weldonJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ウェルダン"}

	// weldonKoreanName represents weldon korean name.
	weldonKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// weldonLatinAmericanSpanishName represents weldon latin american spanish name.
	weldonLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// weldonRussianName represents weldon russian name.
	weldonRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// weldonSimplifiedChineseName represents weldon simplified chinese name.
	weldonSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// weldonSpanishName represents weldon spanish name.
	weldonSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// weldonTraditionalChineseName represents weldon traditional chinese name.
	weldonTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// weldonName represents weldon name.
	weldonName = nook.Languages{
		language.AmericanEnglish:      weldonAmericanEnglishName,
		language.CanadianFrench:       weldonCanadianFrenchName,
		language.Dutch:                weldonDutchName,
		language.French:               weldonFrenchName,
		language.German:               weldonGermanName,
		language.Italian:              weldonItalianName,
		language.Japanese:             weldonJapaneseName,
		language.Korean:               weldonKoreanName,
		language.LatinAmericanSpanish: weldonLatinAmericanSpanishName,
		language.Russian:              weldonRussianName,
		language.SimplifiedChinese:    weldonSimplifiedChineseName,
		language.Spanish:              weldonSpanishName,
		language.TraditionalChinese:   weldonTraditionalChineseName}
)

var (
	// weldonCharacter represents weldon character.
	weldonCharacter = nook.Character{
		Animal:   animal.Bull,
		Birthday: weldonBirthday,
		Code:     weldonCode,
		Key:      character.Weldon,
		Gender:   gender.Male,
		Name:     weldonName,
		Special:  false}
)

var (
	// weldonAmericanEnglishPhrase represents weldon american english phrase.
	weldonAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ムーチョ"}

	// weldonCanadianFrenchPhrase represents weldon canadian french phrase.
	weldonCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// weldonDutchPhrase represents weldon dutch phrase.
	weldonDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// weldonFrenchPhrase represents weldon french phrase.
	weldonFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	// weldonGermanPhrase represents weldon german phrase.
	weldonGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	// weldonItalianPhrase represents weldon italian phrase.
	weldonItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// weldonJapanesePhrase represents weldon japanese phrase.
	weldonJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ムーチョ"}

	// weldonKoreanPhrase represents weldon korean phrase.
	weldonKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// weldonLatinAmericanSpanishPhrase represents weldon latin american spanish phrase.
	weldonLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// weldonRussianPhrase represents weldon russian phrase.
	weldonRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// weldonSimplifiedChinesePhrase represents weldon simplified chinese phrase.
	weldonSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// weldonSpanishPhrase represents weldon spanish phrase.
	weldonSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// weldonTraditionalChinesePhrase represents weldon traditional chinese phrase.
	weldonTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// weldonPhrase represents weldon phrase.
	weldonPhrase = nook.Languages{
		language.AmericanEnglish:      weldonAmericanEnglishPhrase,
		language.CanadianFrench:       weldonCanadianFrenchPhrase,
		language.Dutch:                weldonDutchPhrase,
		language.French:               weldonFrenchPhrase,
		language.German:               weldonGermanPhrase,
		language.Italian:              weldonItalianPhrase,
		language.Japanese:             weldonJapanesePhrase,
		language.Korean:               weldonKoreanPhrase,
		language.LatinAmericanSpanish: weldonLatinAmericanSpanishPhrase,
		language.Russian:              weldonRussianPhrase,
		language.SimplifiedChinese:    weldonSimplifiedChinesePhrase,
		language.Spanish:              weldonSpanishPhrase,
		language.TraditionalChinese:   weldonTraditionalChinesePhrase}
)

var (
	// Weldon represents weldon.
	Weldon = nook.Villager{
		Character:   weldonCharacter,
		Personality: personality.Cranky,
		Phrase:      weldonPhrase}
)

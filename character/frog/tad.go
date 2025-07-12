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
	// tadBirthday represents tad birthday.
	tadBirthday = nook.Birthday{
		Day:   3,
		Month: time.August}
)

var (
	// tadCode represents tad code.
	tadCode = nook.Code{
		Value: "flg09"}
)

var (
	// tadAmericanEnglishName represents tad american english name.
	tadAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tad"}

	// tadCanadianFrenchName represents tad canadian french name.
	tadCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rénato"}

	// tadDutchName represents tad dutch name.
	tadDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tad"}

	// tadFrenchName represents tad french name.
	tadFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rénato"}

	// tadGermanName represents tad german name.
	tadGermanName = nook.Name{
		Language: language.German,
		Value:    "Paul"}

	// tadItalianName represents tad italian name.
	tadItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Achille"}

	// tadJapaneseName represents tad japanese name.
	tadJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "タンボ"}

	// tadKoreanName represents tad korean name.
	tadKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "텀보"}

	// tadLatinAmericanSpanishName represents tad latin american spanish name.
	tadLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Saltonio"}

	// tadRussianName represents tad russian name.
	tadRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Тэд"}

	// tadSimplifiedChineseName represents tad simplified chinese name.
	tadSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿田"}

	// tadSpanishName represents tad spanish name.
	tadSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Saltonio"}

	// tadTraditionalChineseName represents tad traditional chinese name.
	tadTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿田"}
)

var (
	// tadName represents tad name.
	tadName = nook.Languages{
		language.AmericanEnglish:      tadAmericanEnglishName,
		language.CanadianFrench:       tadCanadianFrenchName,
		language.Dutch:                tadDutchName,
		language.French:               tadFrenchName,
		language.German:               tadGermanName,
		language.Italian:              tadItalianName,
		language.Japanese:             tadJapaneseName,
		language.Korean:               tadKoreanName,
		language.LatinAmericanSpanish: tadLatinAmericanSpanishName,
		language.Russian:              tadRussianName,
		language.SimplifiedChinese:    tadSimplifiedChineseName,
		language.Spanish:              tadSpanishName,
		language.TraditionalChinese:   tadTraditionalChineseName}
)

var (
	// tadCharacter represents tad character.
	tadCharacter = nook.Character{
		Animal:   animal.Frog,
		Birthday: tadBirthday,
		Code:     tadCode,
		Key:      character.Tad,
		Gender:   gender.Male,
		Name:     tadName,
		Special:  false}
)

var (
	// tadAmericanEnglishPhrase represents tad american english phrase.
	tadAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sluuuurp"}

	// tadCanadianFrenchPhrase represents tad canadian french phrase.
	tadCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "nénuf"}

	// tadDutchPhrase represents tad dutch phrase.
	tadDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "sluuuurp"}

	// tadFrenchPhrase represents tad french phrase.
	tadFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "nénuf"}

	// tadGermanPhrase represents tad german phrase.
	tadGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schlürpf"}

	// tadItalianPhrase represents tad italian phrase.
	tadItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sluuuurp"}

	// tadJapanesePhrase represents tad japanese phrase.
	tadJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だよん"}

	// tadKoreanPhrase represents tad korean phrase.
	tadKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "흐압"}

	// tadLatinAmericanSpanishPhrase represents tad latin american spanish phrase.
	tadLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "zump"}

	// tadRussianPhrase represents tad russian phrase.
	tadRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "жабррр"}

	// tadSimplifiedChinesePhrase represents tad simplified chinese phrase.
	tadSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "蝌蚪"}

	// tadSpanishPhrase represents tad spanish phrase.
	tadSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "zump"}

	// tadTraditionalChinesePhrase represents tad traditional chinese phrase.
	tadTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "蝌蚪"}
)

var (
	// tadPhrase represents tad phrase.
	tadPhrase = nook.Languages{
		language.AmericanEnglish:      tadAmericanEnglishPhrase,
		language.CanadianFrench:       tadCanadianFrenchPhrase,
		language.Dutch:                tadDutchPhrase,
		language.French:               tadFrenchPhrase,
		language.German:               tadGermanPhrase,
		language.Italian:              tadItalianPhrase,
		language.Japanese:             tadJapanesePhrase,
		language.Korean:               tadKoreanPhrase,
		language.LatinAmericanSpanish: tadLatinAmericanSpanishPhrase,
		language.Russian:              tadRussianPhrase,
		language.SimplifiedChinese:    tadSimplifiedChinesePhrase,
		language.Spanish:              tadSpanishPhrase,
		language.TraditionalChinese:   tadTraditionalChinesePhrase}
)

var (
	// Tad represents tad.
	Tad = nook.Villager{
		Character:   tadCharacter,
		Personality: personality.Jock,
		Phrase:      tadPhrase}
)

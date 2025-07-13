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
	// chicoBirthday represents chico birthday.
	chicoBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// chicoCode represents chico code.
	chicoCode = nook.Code{
		Value: ""}
)

var (
	// chicoAmericanEnglishName represents chico american english name.
	chicoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Chico"}

	// chicoCanadianFrenchName represents chico canadian french name.
	chicoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// chicoDutchName represents chico dutch name.
	chicoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// chicoFrenchName represents chico french name.
	chicoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rapido"}

	// chicoGermanName represents chico german name.
	chicoGermanName = nook.Name{
		Language: language.German,
		Value:    "Pablo"}

	// chicoItalianName represents chico italian name.
	chicoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ciro"}

	// chicoJapaneseName represents chico japanese name.
	chicoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チューボー"}

	// chicoKoreanName represents chico korean name.
	chicoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// chicoLatinAmericanSpanishName represents chico latin american spanish name.
	chicoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// chicoRussianName represents chico russian name.
	chicoRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// chicoSimplifiedChineseName represents chico simplified chinese name.
	chicoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "吱吱"}

	// chicoSpanishName represents chico spanish name.
	chicoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Chico"}

	// chicoTraditionalChineseName represents chico traditional chinese name.
	chicoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// chicoName represents chico name.
	chicoName = nook.Languages{
		language.AmericanEnglish:      chicoAmericanEnglishName,
		language.CanadianFrench:       chicoCanadianFrenchName,
		language.Dutch:                chicoDutchName,
		language.French:               chicoFrenchName,
		language.German:               chicoGermanName,
		language.Italian:              chicoItalianName,
		language.Japanese:             chicoJapaneseName,
		language.Korean:               chicoKoreanName,
		language.LatinAmericanSpanish: chicoLatinAmericanSpanishName,
		language.Russian:              chicoRussianName,
		language.SimplifiedChinese:    chicoSimplifiedChineseName,
		language.Spanish:              chicoSpanishName,
		language.TraditionalChinese:   chicoTraditionalChineseName}
)

var (
	// chicoCharacter represents chico character.
	chicoCharacter = nook.Character{
		Animal:   animal.Mouse,
		Birthday: chicoBirthday,
		Code:     chicoCode,
		Key:      character.Chico,
		Gender:   gender.Male,
		Name:     chicoName,
		Special:  false}
)

var (
	// chicoAmericanEnglishPhrase represents chico american english phrase.
	chicoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cheeeese"}

	// chicoCanadianFrenchPhrase represents chico canadian french phrase.
	chicoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// chicoDutchPhrase represents chico dutch phrase.
	chicoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// chicoFrenchPhrase represents chico french phrase.
	chicoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ma croûte"}

	// chicoGermanPhrase represents chico german phrase.
	chicoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "griiins"}

	// chicoItalianPhrase represents chico italian phrase.
	chicoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "caciotta"}

	// chicoJapanesePhrase represents chico japanese phrase.
	chicoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ナリ"}

	// chicoKoreanPhrase represents chico korean phrase.
	chicoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// chicoLatinAmericanSpanishPhrase represents chico latin american spanish phrase.
	chicoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// chicoRussianPhrase represents chico russian phrase.
	chicoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// chicoSimplifiedChinesePhrase represents chico simplified chinese phrase.
	chicoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哗哗"}

	// chicoSpanishPhrase represents chico spanish phrase.
	chicoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "quesito"}

	// chicoTraditionalChinesePhrase represents chico traditional chinese phrase.
	chicoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// chicoPhrase represents chico phrase.
	chicoPhrase = nook.Languages{
		language.AmericanEnglish:      chicoAmericanEnglishPhrase,
		language.CanadianFrench:       chicoCanadianFrenchPhrase,
		language.Dutch:                chicoDutchPhrase,
		language.French:               chicoFrenchPhrase,
		language.German:               chicoGermanPhrase,
		language.Italian:              chicoItalianPhrase,
		language.Japanese:             chicoJapanesePhrase,
		language.Korean:               chicoKoreanPhrase,
		language.LatinAmericanSpanish: chicoLatinAmericanSpanishPhrase,
		language.Russian:              chicoRussianPhrase,
		language.SimplifiedChinese:    chicoSimplifiedChinesePhrase,
		language.Spanish:              chicoSpanishPhrase,
		language.TraditionalChinese:   chicoTraditionalChinesePhrase}
)

var (
	// Chico represents chico.
	Chico = nook.Villager{
		Character:   chicoCharacter,
		Personality: personality.Lazy,
		Phrase:      chicoPhrase}
)

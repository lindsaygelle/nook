package kangaroo

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
	// koharuBirthday represents koharu birthday.
	koharuBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// koharuCode represents koharu code.
	koharuCode = nook.Code{
		Value: ""}
)

var (
	// koharuAmericanEnglishName represents koharu american english name.
	koharuAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Koharu"}

	// koharuCanadianFrenchName represents koharu canadian french name.
	koharuCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// koharuDutchName represents koharu dutch name.
	koharuDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// koharuFrenchName represents koharu french name.
	koharuFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	// koharuGermanName represents koharu german name.
	koharuGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	// koharuItalianName represents koharu italian name.
	koharuItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// koharuJapaneseName represents koharu japanese name.
	koharuJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "こはる"}

	// koharuKoreanName represents koharu korean name.
	koharuKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// koharuLatinAmericanSpanishName represents koharu latin american spanish name.
	koharuLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// koharuRussianName represents koharu russian name.
	koharuRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// koharuSimplifiedChineseName represents koharu simplified chinese name.
	koharuSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// koharuSpanishName represents koharu spanish name.
	koharuSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// koharuTraditionalChineseName represents koharu traditional chinese name.
	koharuTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// koharuName represents koharu name.
	koharuName = nook.Languages{
		language.AmericanEnglish:      koharuAmericanEnglishName,
		language.CanadianFrench:       koharuCanadianFrenchName,
		language.Dutch:                koharuDutchName,
		language.French:               koharuFrenchName,
		language.German:               koharuGermanName,
		language.Italian:              koharuItalianName,
		language.Japanese:             koharuJapaneseName,
		language.Korean:               koharuKoreanName,
		language.LatinAmericanSpanish: koharuLatinAmericanSpanishName,
		language.Russian:              koharuRussianName,
		language.SimplifiedChinese:    koharuSimplifiedChineseName,
		language.Spanish:              koharuSpanishName,
		language.TraditionalChinese:   koharuTraditionalChineseName}
)

var (
	// koharuCharacter represents koharu character.
	koharuCharacter = nook.Character{
		Animal:   animal.Kangaroo,
		Birthday: koharuBirthday,
		Code:     koharuCode,
		Key:      character.Koharu,
		Gender:   gender.Female,
		Name:     koharuName,
		Special:  false}
)

var (
	// koharuAmericanEnglishPhrase represents koharu american english phrase.
	koharuAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ちょいと"}

	// koharuCanadianFrenchPhrase represents koharu canadian french phrase.
	koharuCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// koharuDutchPhrase represents koharu dutch phrase.
	koharuDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// koharuFrenchPhrase represents koharu french phrase.
	koharuFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	// koharuGermanPhrase represents koharu german phrase.
	koharuGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	// koharuItalianPhrase represents koharu italian phrase.
	koharuItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// koharuJapanesePhrase represents koharu japanese phrase.
	koharuJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ちょいと"}

	// koharuKoreanPhrase represents koharu korean phrase.
	koharuKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// koharuLatinAmericanSpanishPhrase represents koharu latin american spanish phrase.
	koharuLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// koharuRussianPhrase represents koharu russian phrase.
	koharuRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// koharuSimplifiedChinesePhrase represents koharu simplified chinese phrase.
	koharuSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// koharuSpanishPhrase represents koharu spanish phrase.
	koharuSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// koharuTraditionalChinesePhrase represents koharu traditional chinese phrase.
	koharuTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// koharuPhrase represents koharu phrase.
	koharuPhrase = nook.Languages{
		language.AmericanEnglish:      koharuAmericanEnglishPhrase,
		language.CanadianFrench:       koharuCanadianFrenchPhrase,
		language.Dutch:                koharuDutchPhrase,
		language.French:               koharuFrenchPhrase,
		language.German:               koharuGermanPhrase,
		language.Italian:              koharuItalianPhrase,
		language.Japanese:             koharuJapanesePhrase,
		language.Korean:               koharuKoreanPhrase,
		language.LatinAmericanSpanish: koharuLatinAmericanSpanishPhrase,
		language.Russian:              koharuRussianPhrase,
		language.SimplifiedChinese:    koharuSimplifiedChinesePhrase,
		language.Spanish:              koharuSpanishPhrase,
		language.TraditionalChinese:   koharuTraditionalChinesePhrase}
)

var (
	// Koharu represents koharu.
	Koharu = nook.Villager{
		Character:   koharuCharacter,
		Personality: personality.Peppy,
		Phrase:      koharuPhrase}
)

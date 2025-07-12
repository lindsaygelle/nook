package penguin

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
	// nobuoBirthday represents nobuo birthday.
	nobuoBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// nobuoCode represents nobuo code.
	nobuoCode = nook.Code{
		Value: ""}
)

var (
	// nobuoAmericanEnglishName represents nobuo american english name.
	nobuoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Nobuo"}

	// nobuoCanadianFrenchName represents nobuo canadian french name.
	nobuoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// nobuoDutchName represents nobuo dutch name.
	nobuoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// nobuoFrenchName represents nobuo french name.
	nobuoFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	// nobuoGermanName represents nobuo german name.
	nobuoGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	// nobuoItalianName represents nobuo italian name.
	nobuoItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// nobuoJapaneseName represents nobuo japanese name.
	nobuoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "のぶお"}

	// nobuoKoreanName represents nobuo korean name.
	nobuoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// nobuoLatinAmericanSpanishName represents nobuo latin american spanish name.
	nobuoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// nobuoRussianName represents nobuo russian name.
	nobuoRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// nobuoSimplifiedChineseName represents nobuo simplified chinese name.
	nobuoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// nobuoSpanishName represents nobuo spanish name.
	nobuoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// nobuoTraditionalChineseName represents nobuo traditional chinese name.
	nobuoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// nobuoName represents nobuo name.
	nobuoName = nook.Languages{
		language.AmericanEnglish:      nobuoAmericanEnglishName,
		language.CanadianFrench:       nobuoCanadianFrenchName,
		language.Dutch:                nobuoDutchName,
		language.French:               nobuoFrenchName,
		language.German:               nobuoGermanName,
		language.Italian:              nobuoItalianName,
		language.Japanese:             nobuoJapaneseName,
		language.Korean:               nobuoKoreanName,
		language.LatinAmericanSpanish: nobuoLatinAmericanSpanishName,
		language.Russian:              nobuoRussianName,
		language.SimplifiedChinese:    nobuoSimplifiedChineseName,
		language.Spanish:              nobuoSpanishName,
		language.TraditionalChinese:   nobuoTraditionalChineseName}
)

var (
	// nobuoCharacter represents nobuo character.
	nobuoCharacter = nook.Character{
		Animal:   animal.Penguin,
		Birthday: nobuoBirthday,
		Code:     nobuoCode,
		Key:      character.Nobuo,
		Gender:   gender.Male,
		Name:     nobuoName,
		Special:  false}
)

var (
	// nobuoAmericanEnglishPhrase represents nobuo american english phrase.
	nobuoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ブツブツ"}

	// nobuoCanadianFrenchPhrase represents nobuo canadian french phrase.
	nobuoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// nobuoDutchPhrase represents nobuo dutch phrase.
	nobuoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// nobuoFrenchPhrase represents nobuo french phrase.
	nobuoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	// nobuoGermanPhrase represents nobuo german phrase.
	nobuoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	// nobuoItalianPhrase represents nobuo italian phrase.
	nobuoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// nobuoJapanesePhrase represents nobuo japanese phrase.
	nobuoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ブツブツ"}

	// nobuoKoreanPhrase represents nobuo korean phrase.
	nobuoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// nobuoLatinAmericanSpanishPhrase represents nobuo latin american spanish phrase.
	nobuoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// nobuoRussianPhrase represents nobuo russian phrase.
	nobuoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// nobuoSimplifiedChinesePhrase represents nobuo simplified chinese phrase.
	nobuoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// nobuoSpanishPhrase represents nobuo spanish phrase.
	nobuoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// nobuoTraditionalChinesePhrase represents nobuo traditional chinese phrase.
	nobuoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// nobuoPhrase represents nobuo phrase.
	nobuoPhrase = nook.Languages{
		language.AmericanEnglish:      nobuoAmericanEnglishPhrase,
		language.CanadianFrench:       nobuoCanadianFrenchPhrase,
		language.Dutch:                nobuoDutchPhrase,
		language.French:               nobuoFrenchPhrase,
		language.German:               nobuoGermanPhrase,
		language.Italian:              nobuoItalianPhrase,
		language.Japanese:             nobuoJapanesePhrase,
		language.Korean:               nobuoKoreanPhrase,
		language.LatinAmericanSpanish: nobuoLatinAmericanSpanishPhrase,
		language.Russian:              nobuoRussianPhrase,
		language.SimplifiedChinese:    nobuoSimplifiedChinesePhrase,
		language.Spanish:              nobuoSpanishPhrase,
		language.TraditionalChinese:   nobuoTraditionalChinesePhrase}
)

var (
	// Nobuo represents nobuo.
	Nobuo = nook.Villager{
		Character:   nobuoCharacter,
		Personality: personality.Lazy,
		Phrase:      nobuoPhrase}
)

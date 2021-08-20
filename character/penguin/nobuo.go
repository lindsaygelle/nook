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
	nobuoBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	nobuoCode = nook.Code{
		Value: ""}
)

var (
	nobuoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Nobuo"}

	nobuoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	nobuoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	nobuoFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	nobuoGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	nobuoItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	nobuoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "のぶお"}

	nobuoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	nobuoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	nobuoRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	nobuoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	nobuoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	nobuoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
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
	nobuoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ブツブツ"}

	nobuoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	nobuoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	nobuoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	nobuoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	nobuoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	nobuoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ブツブツ"}

	nobuoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	nobuoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	nobuoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	nobuoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	nobuoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	nobuoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
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
	Nobuo = nook.Villager{
		Character:   nobuoCharacter,
		Personality: personality.Lazy,
		Phrase:      nobuoPhrase}
)

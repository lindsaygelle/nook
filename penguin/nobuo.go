package penguin

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "N/A"}

	nobuoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	nobuoFrenchName = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	nobuoGermanName = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	nobuoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	nobuoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "のぶお"}

	nobuoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	nobuoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	nobuoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	nobuoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	nobuoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	nobuoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
		Animal:   Penguin,
		Birthday: nobuoBirthday,
		Code:     nobuoCode,
		Gender:   nook.Male,
		Name:     nobuoName}
)

var (
	nobuoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	nobuoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	nobuoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	nobuoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	nobuoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	nobuoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	nobuoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ブツブツ"}

	nobuoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	nobuoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	nobuoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	nobuoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	nobuoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	nobuoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	nobuoPhrase = nook.Languages{
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
	Nobuo = nook.Villager{
		Character:   nobuoCharacter,
		Personality: nook.Lazy,
		Phrase:      nobuoPhrase}
)

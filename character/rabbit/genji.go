package rabbit

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
	genjiBirthday = nook.Birthday{
		Day:   21,
		Month: time.January}
)

var (
	genjiCode = nook.Code{
		Value: "rbt08"}
)

var (
	genjiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Genji"}

	genjiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Kali"}

	genjiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Genji"}

	genjiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Kali"}

	genjiGermanName = nook.Name{
		Language: language.German,
		Value:    "Aki"}

	genjiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Genji"}

	genjiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ゲンジ"}

	genjiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "토시"}

	genjiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Sumo"}

	genjiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Гэндзи"}

	genjiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "儒林"}

	genjiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Sumo"}

	genjiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "儒林"}
)

var (
	genjiName = nook.Languages{
		language.AmericanEnglish:      genjiAmericanEnglishName,
		language.CanadianFrench:       genjiCanadianFrenchName,
		language.Dutch:                genjiDutchName,
		language.French:               genjiFrenchName,
		language.German:               genjiGermanName,
		language.Italian:              genjiItalianName,
		language.Japanese:             genjiJapaneseName,
		language.Korean:               genjiKoreanName,
		language.LatinAmericanSpanish: genjiLatinAmericanSpanishName,
		language.Russian:              genjiRussianName,
		language.SimplifiedChinese:    genjiSimplifiedChineseName,
		language.Spanish:              genjiSpanishName,
		language.TraditionalChinese:   genjiTraditionalChineseName}
)

var (
	genjiCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: genjiBirthday,
		Code:     genjiCode,
		Key:      character.Genji,
		Gender:   gender.Male,
		Name:     genjiName}
)

var (
	genjiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "otaku"}

	genjiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mochi"}

	genjiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "mochi"}

	genjiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mochi"}

	genjiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "goseimas"}

	genjiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "mochi"}

	genjiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "まろ"}

	genjiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "땡글"}

	genjiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "mochi"}

	genjiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "зайпончик"}

	genjiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "臣"}

	genjiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "mochi"}

	genjiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "臣"}
)

var (
	genjiPhrase = nook.Languages{
		language.AmericanEnglish:      genjiAmericanEnglishPhrase,
		language.CanadianFrench:       genjiCanadianFrenchPhrase,
		language.Dutch:                genjiDutchPhrase,
		language.French:               genjiFrenchPhrase,
		language.German:               genjiGermanPhrase,
		language.Italian:              genjiItalianPhrase,
		language.Japanese:             genjiJapanesePhrase,
		language.Korean:               genjiKoreanPhrase,
		language.LatinAmericanSpanish: genjiLatinAmericanSpanishPhrase,
		language.Russian:              genjiRussianPhrase,
		language.SimplifiedChinese:    genjiSimplifiedChinesePhrase,
		language.Spanish:              genjiSpanishPhrase,
		language.TraditionalChinese:   genjiTraditionalChinesePhrase}
)

var (
	Genji = nook.Villager{
		Character:   genjiCharacter,
		Personality: personality.Jock,
		Phrase:      genjiPhrase}
)

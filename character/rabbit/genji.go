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
	// genjiBirthday represents Genji's birthday.
	genjiBirthday = nook.Birthday{
		Day:   21,
		Month: time.January}
)

var (
	// genjiCode represents Genji's unique code ("rbt08").
	genjiCode = nook.Code{
		Value: "rbt08"}
)

var (
	// genjiAmericanEnglishName represents Genji's name in American English.
	genjiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Genji"}

	// genjiCanadianFrenchName represents Genji's name in Canadian French.
	genjiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Kali"}

	// genjiDutchName represents Genji's name in Dutch.
	genjiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Genji"}

	// genjiFrenchName represents Genji's name in French.
	genjiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Kali"}

	// genjiGermanName represents Genji's name in German.
	genjiGermanName = nook.Name{
		Language: language.German,
		Value:    "Aki"}

	// genjiItalianName represents Genji's name in Italian.
	genjiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Genji"}

	// genjiJapaneseName represents Genji's name in Japanese.
	genjiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ゲンジ"}

	// genjiKoreanName represents Genji's name in Korean.
	genjiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "토시"}

	// genjiLatinAmericanSpanishName represents Genji's name in Latin American Spanish.
	genjiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Sumo"}

	// genjiRussianName represents Genji's name in Russian.
	genjiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Гэндзи"}

	// genjiSimplifiedChineseName represents Genji's name in Simplified Chinese.
	genjiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "儒林"}

	// genjiSpanishName represents Genji's name in Spanish.
	genjiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Sumo"}

	// genjiTraditionalChineseName represents Genji's name in Traditional Chinese.
	genjiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "儒林"}
)

var (
	// genjiName represents Genji's name in different languages.
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
	// genjiCharacter represents Genji's character information.
	genjiCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: genjiBirthday,
		Code:     genjiCode,
		Key:      character.Genji,
		Gender:   gender.Male,
		Name:     genjiName,
		Special:  false}
)

var (
	// genjiAmericanEnglishPhrase represents Genji's phrase in American English.
	genjiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "otaku"}

	// genjiCanadianFrenchPhrase represents Genji's phrase in Canadian French.
	genjiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mochi"}

	// genjiDutchPhrase represents Genji's phrase in Dutch.
	genjiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "mochi"}

	// genjiFrenchPhrase represents Genji's phrase in French.
	genjiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mochi"}

	// genjiGermanPhrase represents Genji's phrase in German.
	genjiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "goseimas"}

	// genjiItalianPhrase represents Genji's phrase in Italian.
	genjiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "mochi"}

	// genjiJapanesePhrase represents Genji's phrase in Japanese.
	genjiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "まろ"}

	// genjiKoreanPhrase represents Genji's phrase in Korean.
	genjiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "땡글"}

	// genjiLatinAmericanSpanishPhrase represents Genji's phrase in Latin American Spanish.
	genjiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "mochi"}

	// genjiRussianPhrase represents Genji's phrase in Russian.
	genjiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "зайпончик"}

	// genjiSimplifiedChinesePhrase represents Genji's phrase in Simplified Chinese.
	genjiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "臣"}

	// genjiSpanishPhrase represents Genji's phrase in Spanish.
	genjiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "mochi"}

	// genjiTraditionalChinesePhrase represents Genji's phrase in Traditional Chinese.
	genjiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "臣"}
)

var (
	// genjiPhrase represents Genji's phrases in different languages.
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
	// Genji represents the character Genji with his complete information.
	Genji = nook.Villager{
		Character:   genjiCharacter,
		Personality: personality.Jock,
		Phrase:      genjiPhrase}
)

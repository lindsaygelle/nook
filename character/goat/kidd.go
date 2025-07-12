package goat

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
	// kiddBirthday represents kidd birthday.
	kiddBirthday = nook.Birthday{
		Day:   28,
		Month: time.June}
)

var (
	// kiddCode represents kidd code.
	kiddCode = nook.Code{
		Value: "goa07"}
)

var (
	// kiddAmericanEnglishName represents kidd american english name.
	kiddAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Kidd"}

	// kiddCanadianFrenchName represents kidd canadian french name.
	kiddCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Moktar"}

	// kiddDutchName represents kidd dutch name.
	kiddDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Kidd"}

	// kiddFrenchName represents kidd french name.
	kiddFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mokhtar"}

	// kiddGermanName represents kidd german name.
	kiddGermanName = nook.Name{
		Language: language.German,
		Value:    "Bocki"}

	// kiddItalianName represents kidd italian name.
	kiddItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Vittorio"}

	// kiddJapaneseName represents kidd japanese name.
	kiddJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "やさお"}

	// kiddKoreanName represents kidd korean name.
	kiddKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "염두리"}

	// kiddLatinAmericanSpanishName represents kidd latin american spanish name.
	kiddLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cabrálex"}

	// kiddRussianName represents kidd russian name.
	kiddRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кидд"}

	// kiddSimplifiedChineseName represents kidd simplified chinese name.
	kiddSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "文青"}

	// kiddSpanishName represents kidd spanish name.
	kiddSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cabrálex"}

	// kiddTraditionalChineseName represents kidd traditional chinese name.
	kiddTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "文青"}
)

var (
	// kiddName represents kidd name.
	kiddName = nook.Languages{
		language.AmericanEnglish:      kiddAmericanEnglishName,
		language.CanadianFrench:       kiddCanadianFrenchName,
		language.Dutch:                kiddDutchName,
		language.French:               kiddFrenchName,
		language.German:               kiddGermanName,
		language.Italian:              kiddItalianName,
		language.Japanese:             kiddJapaneseName,
		language.Korean:               kiddKoreanName,
		language.LatinAmericanSpanish: kiddLatinAmericanSpanishName,
		language.Russian:              kiddRussianName,
		language.SimplifiedChinese:    kiddSimplifiedChineseName,
		language.Spanish:              kiddSpanishName,
		language.TraditionalChinese:   kiddTraditionalChineseName}
)

var (
	// kiddCharacter represents kidd character.
	kiddCharacter = nook.Character{
		Animal:   animal.Goat,
		Birthday: kiddBirthday,
		Code:     kiddCode,
		Key:      character.Kidd,
		Gender:   gender.Male,
		Name:     kiddName,
		Special:  false}
)

var (
	// kiddAmericanEnglishPhrase represents kidd american english phrase.
	kiddAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "wut"}

	// kiddCanadianFrenchPhrase represents kidd canadian french phrase.
	kiddCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mouhai"}

	// kiddDutchPhrase represents kidd dutch phrase.
	kiddDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "watte"}

	// kiddFrenchPhrase represents kidd french phrase.
	kiddFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mouhai"}

	// kiddGermanPhrase represents kidd german phrase.
	kiddGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mehehe"}

	// kiddItalianPhrase represents kidd italian phrase.
	kiddItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "balido"}

	// kiddJapanesePhrase represents kidd japanese phrase.
	kiddJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "はぁ～"}

	// kiddKoreanPhrase represents kidd korean phrase.
	kiddKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "웁스"}

	// kiddLatinAmericanSpanishPhrase represents kidd latin american spanish phrase.
	kiddLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ein"}

	// kiddRussianPhrase represents kidd russian phrase.
	kiddRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "так"}

	// kiddSimplifiedChinesePhrase represents kidd simplified chinese phrase.
	kiddSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啊"}

	// kiddSpanishPhrase represents kidd spanish phrase.
	kiddSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ein"}

	// kiddTraditionalChinesePhrase represents kidd traditional chinese phrase.
	kiddTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "啊"}
)

var (
	// kiddPhrase represents kidd phrase.
	kiddPhrase = nook.Languages{
		language.AmericanEnglish:      kiddAmericanEnglishPhrase,
		language.CanadianFrench:       kiddCanadianFrenchPhrase,
		language.Dutch:                kiddDutchPhrase,
		language.French:               kiddFrenchPhrase,
		language.German:               kiddGermanPhrase,
		language.Italian:              kiddItalianPhrase,
		language.Japanese:             kiddJapanesePhrase,
		language.Korean:               kiddKoreanPhrase,
		language.LatinAmericanSpanish: kiddLatinAmericanSpanishPhrase,
		language.Russian:              kiddRussianPhrase,
		language.SimplifiedChinese:    kiddSimplifiedChinesePhrase,
		language.Spanish:              kiddSpanishPhrase,
		language.TraditionalChinese:   kiddTraditionalChinesePhrase}
)

var (
	// Kidd represents kidd.
	Kidd = nook.Villager{
		Character:   kiddCharacter,
		Personality: personality.Smug,
		Phrase:      kiddPhrase}
)

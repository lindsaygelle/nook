package cat

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
	// kattBirthday represents katt birthday.
	kattBirthday = nook.Birthday{
		Day:   27,
		Month: time.April}
)

var (
	// kattCode represents katt code.
	kattCode = nook.Code{
		Value: "cat21"}
)

var (
	// kattAmericanEnglishName represents katt american english name.
	kattAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Katt"}

	// kattCanadianFrenchName represents katt canadian french name.
	kattCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Kat"}

	// kattDutchName represents katt dutch name.
	kattDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Katt"}

	// kattFrenchName represents katt french name.
	kattFrenchName = nook.Name{
		Language: language.French,
		Value:    "Kat"}

	// kattGermanName represents katt german name.
	kattGermanName = nook.Name{
		Language: language.German,
		Value:    "Janine"}

	// kattItalianName represents katt italian name.
	kattItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ofelia"}

	// kattJapaneseName represents katt japanese name.
	kattJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ちょい"}

	// kattKoreanName represents katt korean name.
	kattKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "쵸이"}

	// kattLatinAmericanSpanishName represents katt latin american spanish name.
	kattLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Melina"}

	// kattRussianName represents katt russian name.
	kattRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кэт"}

	// kattSimplifiedChineseName represents katt simplified chinese name.
	kattSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "巧巧"}

	// kattSpanishName represents katt spanish name.
	kattSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Melina"}

	// kattTraditionalChineseName represents katt traditional chinese name.
	kattTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "巧巧"}
)

var (
	// kattName represents katt name.
	kattName = nook.Languages{
		language.AmericanEnglish:      kattAmericanEnglishName,
		language.CanadianFrench:       kattCanadianFrenchName,
		language.Dutch:                kattDutchName,
		language.French:               kattFrenchName,
		language.German:               kattGermanName,
		language.Italian:              kattItalianName,
		language.Japanese:             kattJapaneseName,
		language.Korean:               kattKoreanName,
		language.LatinAmericanSpanish: kattLatinAmericanSpanishName,
		language.Russian:              kattRussianName,
		language.SimplifiedChinese:    kattSimplifiedChineseName,
		language.Spanish:              kattSpanishName,
		language.TraditionalChinese:   kattTraditionalChineseName}
)

var (
	// kattCharacter represents katt character.
	kattCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: kattBirthday,
		Code:     kattCode,
		Key:      character.Katt,
		Gender:   gender.Female,
		Name:     kattName,
		Special:  false}
)

var (
	// kattAmericanEnglishPhrase represents katt american english phrase.
	kattAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "purrty"}

	// kattCanadianFrenchPhrase represents katt canadian french phrase.
	kattCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "miaouille"}

	// kattDutchPhrase represents katt dutch phrase.
	kattDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "prrrachtig"}

	// kattFrenchPhrase represents katt french phrase.
	kattFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "miaouille"}

	// kattGermanPhrase represents katt german phrase.
	kattGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "zischel"}

	// kattItalianPhrase represents katt italian phrase.
	kattItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "fffz"}

	// kattJapanesePhrase represents katt japanese phrase.
	kattJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "じゃね"}

	// kattKoreanPhrase represents katt korean phrase.
	kattKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "휘리릭"}

	// kattLatinAmericanSpanishPhrase represents katt latin american spanish phrase.
	kattLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "purrrs"}

	// kattRussianPhrase represents katt russian phrase.
	kattRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хор-р-рошо"}

	// kattSimplifiedChinesePhrase represents katt simplified chinese phrase.
	kattSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "掰掰"}

	// kattSpanishPhrase represents katt spanish phrase.
	kattSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "salmonete"}

	// kattTraditionalChinesePhrase represents katt traditional chinese phrase.
	kattTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "掰掰"}
)

var (
	// kattPhrase represents katt phrase.
	kattPhrase = nook.Languages{
		language.AmericanEnglish:      kattAmericanEnglishPhrase,
		language.CanadianFrench:       kattCanadianFrenchPhrase,
		language.Dutch:                kattDutchPhrase,
		language.French:               kattFrenchPhrase,
		language.German:               kattGermanPhrase,
		language.Italian:              kattItalianPhrase,
		language.Japanese:             kattJapanesePhrase,
		language.Korean:               kattKoreanPhrase,
		language.LatinAmericanSpanish: kattLatinAmericanSpanishPhrase,
		language.Russian:              kattRussianPhrase,
		language.SimplifiedChinese:    kattSimplifiedChinesePhrase,
		language.Spanish:              kattSpanishPhrase,
		language.TraditionalChinese:   kattTraditionalChinesePhrase}
)

var (
	// Katt represents katt.
	Katt = nook.Villager{
		Character:   kattCharacter,
		Personality: personality.BigSister,
		Phrase:      kattPhrase}
)

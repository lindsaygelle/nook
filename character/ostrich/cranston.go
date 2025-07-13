package ostrich

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
	// cranstonBirthday represents cranston birthday.
	cranstonBirthday = nook.Birthday{
		Day:   23,
		Month: time.September}
)

var (
	// cranstonCode represents cranston code.
	cranstonCode = nook.Code{
		Value: "ost06"}
)

var (
	// cranstonAmericanEnglishName represents cranston american english name.
	cranstonAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cranston"}

	// cranstonCanadianFrenchName represents cranston canadian french name.
	cranstonCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gabin"}

	// cranstonDutchName represents cranston dutch name.
	cranstonDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cranston"}

	// cranstonFrenchName represents cranston french name.
	cranstonFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gabin"}

	// cranstonGermanName represents cranston german name.
	cranstonGermanName = nook.Name{
		Language: language.German,
		Value:    "Guido"}

	// cranstonItalianName represents cranston italian name.
	cranstonItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Carmine"}

	// cranstonJapaneseName represents cranston japanese name.
	cranstonJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "トキオ"}

	// cranstonKoreanName represents cranston korean name.
	cranstonKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "타키"}

	// cranstonLatinAmericanSpanishName represents cranston latin american spanish name.
	cranstonLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Carmelo"}

	// cranstonRussianName represents cranston russian name.
	cranstonRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Крэнстон"}

	// cranstonSimplifiedChineseName represents cranston simplified chinese name.
	cranstonSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "朱禄"}

	// cranstonSpanishName represents cranston spanish name.
	cranstonSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Carmelo"}

	// cranstonTraditionalChineseName represents cranston traditional chinese name.
	cranstonTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "朱祿"}
)

var (
	// cranstonName represents cranston name.
	cranstonName = nook.Languages{
		language.AmericanEnglish:      cranstonAmericanEnglishName,
		language.CanadianFrench:       cranstonCanadianFrenchName,
		language.Dutch:                cranstonDutchName,
		language.French:               cranstonFrenchName,
		language.German:               cranstonGermanName,
		language.Italian:              cranstonItalianName,
		language.Japanese:             cranstonJapaneseName,
		language.Korean:               cranstonKoreanName,
		language.LatinAmericanSpanish: cranstonLatinAmericanSpanishName,
		language.Russian:              cranstonRussianName,
		language.SimplifiedChinese:    cranstonSimplifiedChineseName,
		language.Spanish:              cranstonSpanishName,
		language.TraditionalChinese:   cranstonTraditionalChineseName}
)

var (
	// cranstonCharacter represents cranston character.
	cranstonCharacter = nook.Character{
		Animal:   animal.Ostrich,
		Birthday: cranstonBirthday,
		Code:     cranstonCode,
		Key:      character.Cranston,
		Gender:   gender.Male,
		Name:     cranstonName,
		Special:  false}
)

var (
	// cranstonAmericanEnglishPhrase represents cranston american english phrase.
	cranstonAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sweatband"}

	// cranstonCanadianFrenchPhrase represents cranston canadian french phrase.
	cranstonCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "omelette"}

	// cranstonDutchPhrase represents cranston dutch phrase.
	cranstonDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "zweetband"}

	// cranstonFrenchPhrase represents cranston french phrase.
	cranstonFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "omelette"}

	// cranstonGermanPhrase represents cranston german phrase.
	cranstonGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "happahappa"}

	// cranstonItalianPhrase represents cranston italian phrase.
	cranstonItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "boa"}

	// cranstonJapanesePhrase represents cranston japanese phrase.
	cranstonJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "およよ"}

	// cranstonKoreanPhrase represents cranston korean phrase.
	cranstonKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "오요용"}

	// cranstonLatinAmericanSpanishPhrase represents cranston latin american spanish phrase.
	cranstonLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "gruqui"}

	// cranstonRussianPhrase represents cranston russian phrase.
	cranstonRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "не потей"}

	// cranstonSimplifiedChinesePhrase represents cranston simplified chinese phrase.
	cranstonSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哎呀呀"}

	// cranstonSpanishPhrase represents cranston spanish phrase.
	cranstonSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "gruqui"}

	// cranstonTraditionalChinesePhrase represents cranston traditional chinese phrase.
	cranstonTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哎呀呀"}
)

var (
	// cranstonPhrase represents cranston phrase.
	cranstonPhrase = nook.Languages{
		language.AmericanEnglish:      cranstonAmericanEnglishPhrase,
		language.CanadianFrench:       cranstonCanadianFrenchPhrase,
		language.Dutch:                cranstonDutchPhrase,
		language.French:               cranstonFrenchPhrase,
		language.German:               cranstonGermanPhrase,
		language.Italian:              cranstonItalianPhrase,
		language.Japanese:             cranstonJapanesePhrase,
		language.Korean:               cranstonKoreanPhrase,
		language.LatinAmericanSpanish: cranstonLatinAmericanSpanishPhrase,
		language.Russian:              cranstonRussianPhrase,
		language.SimplifiedChinese:    cranstonSimplifiedChinesePhrase,
		language.Spanish:              cranstonSpanishPhrase,
		language.TraditionalChinese:   cranstonTraditionalChinesePhrase}
)

var (
	// Cranston represents cranston.
	Cranston = nook.Villager{
		Character:   cranstonCharacter,
		Personality: personality.Lazy,
		Phrase:      cranstonPhrase}
)

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
	// mathildaBirthday represents mathilda birthday.
	mathildaBirthday = nook.Birthday{
		Day:   12,
		Month: time.November}
)

var (
	// mathildaCode represents mathilda code.
	mathildaCode = nook.Code{
		Value: "kgr01"}
)

var (
	// mathildaAmericanEnglishName represents mathilda american english name.
	mathildaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Mathilda"}

	// mathildaCanadianFrenchName represents mathilda canadian french name.
	mathildaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Mathilde"}

	// mathildaDutchName represents mathilda dutch name.
	mathildaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Mathilda"}

	// mathildaFrenchName represents mathilda french name.
	mathildaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mathilde"}

	// mathildaGermanName represents mathilda german name.
	mathildaGermanName = nook.Name{
		Language: language.German,
		Value:    "Marga"}

	// mathildaItalianName represents mathilda italian name.
	mathildaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Matilda"}

	// mathildaJapaneseName represents mathilda japanese name.
	mathildaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アザラク"}

	// mathildaKoreanName represents mathilda korean name.
	mathildaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "안젤라"}

	// mathildaLatinAmericanSpanishName represents mathilda latin american spanish name.
	mathildaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pugilda"}

	// mathildaRussianName represents mathilda russian name.
	mathildaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Матильда"}

	// mathildaSimplifiedChineseName represents mathilda simplified chinese name.
	mathildaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "亚莎"}

	// mathildaSpanishName represents mathilda spanish name.
	mathildaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pugilda"}

	// mathildaTraditionalChineseName represents mathilda traditional chinese name.
	mathildaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "亞莎"}
)

var (
	// mathildaName represents mathilda name.
	mathildaName = nook.Languages{
		language.AmericanEnglish:      mathildaAmericanEnglishName,
		language.CanadianFrench:       mathildaCanadianFrenchName,
		language.Dutch:                mathildaDutchName,
		language.French:               mathildaFrenchName,
		language.German:               mathildaGermanName,
		language.Italian:              mathildaItalianName,
		language.Japanese:             mathildaJapaneseName,
		language.Korean:               mathildaKoreanName,
		language.LatinAmericanSpanish: mathildaLatinAmericanSpanishName,
		language.Russian:              mathildaRussianName,
		language.SimplifiedChinese:    mathildaSimplifiedChineseName,
		language.Spanish:              mathildaSpanishName,
		language.TraditionalChinese:   mathildaTraditionalChineseName}
)

var (
	// mathildaCharacter represents mathilda character.
	mathildaCharacter = nook.Character{
		Animal:   animal.Kangaroo,
		Birthday: mathildaBirthday,
		Code:     mathildaCode,
		Key:      character.Mathilda,
		Gender:   gender.Female,
		Name:     mathildaName,
		Special:  false}
)

var (
	// mathildaAmericanEnglishPhrase represents mathilda american english phrase.
	mathildaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "wee baby"}

	// mathildaCanadianFrenchPhrase represents mathilda canadian french phrase.
	mathildaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "riquiqui"}

	// mathildaDutchPhrase represents mathilda dutch phrase.
	mathildaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "ukkepuk"}

	// mathildaFrenchPhrase represents mathilda french phrase.
	mathildaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "riquiqui"}

	// mathildaGermanPhrase represents mathilda german phrase.
	mathildaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "knirps"}

	// mathildaItalianPhrase represents mathilda italian phrase.
	mathildaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "oplà"}

	// mathildaJapanesePhrase represents mathilda japanese phrase.
	mathildaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ッハ"}

	// mathildaKoreanPhrase represents mathilda korean phrase.
	mathildaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "응"}

	// mathildaLatinAmericanSpanishPhrase represents mathilda latin american spanish phrase.
	mathildaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ala-rorro"}

	// mathildaRussianPhrase represents mathilda russian phrase.
	mathildaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "малютка"}

	// mathildaSimplifiedChinesePhrase represents mathilda simplified chinese phrase.
	mathildaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嘿"}

	// mathildaSpanishPhrase represents mathilda spanish phrase.
	mathildaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "peso mosca"}

	// mathildaTraditionalChinesePhrase represents mathilda traditional chinese phrase.
	mathildaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘿"}
)

var (
	// mathildaPhrase represents mathilda phrase.
	mathildaPhrase = nook.Languages{
		language.AmericanEnglish:      mathildaAmericanEnglishPhrase,
		language.CanadianFrench:       mathildaCanadianFrenchPhrase,
		language.Dutch:                mathildaDutchPhrase,
		language.French:               mathildaFrenchPhrase,
		language.German:               mathildaGermanPhrase,
		language.Italian:              mathildaItalianPhrase,
		language.Japanese:             mathildaJapanesePhrase,
		language.Korean:               mathildaKoreanPhrase,
		language.LatinAmericanSpanish: mathildaLatinAmericanSpanishPhrase,
		language.Russian:              mathildaRussianPhrase,
		language.SimplifiedChinese:    mathildaSimplifiedChinesePhrase,
		language.Spanish:              mathildaSpanishPhrase,
		language.TraditionalChinese:   mathildaTraditionalChinesePhrase}
)

var (
	// Mathilda represents mathilda.
	Mathilda = nook.Villager{
		Character:   mathildaCharacter,
		Personality: personality.Snooty,
		Phrase:      mathildaPhrase}
)

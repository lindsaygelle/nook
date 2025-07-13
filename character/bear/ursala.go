package bear

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

// ursalaBirthday represents Ursala's birthday.
var (
	// ursalaBirthday represents ursala birthday.
	ursalaBirthday = nook.Birthday{
		Day:   16,
		Month: time.January}
)

// ursalaCode represents Ursala's unique code.
var (
	// ursalaCode represents ursala code.
	ursalaCode = nook.Code{
		Value: "bea08"}
)

// Different names for Ursala in various languages.
var (
	// ursalaAmericanEnglishName represents ursala american english name.
	ursalaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Ursala"}

	// ursalaCanadianFrenchName represents ursala canadian french name.
	ursalaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Oursula"}

	// ursalaDutchName represents ursala dutch name.
	ursalaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Ursala"}

	// ursalaFrenchName represents ursala french name.
	ursalaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Oursula"}

	// ursalaGermanName represents ursala german name.
	ursalaGermanName = nook.Name{
		Language: language.German,
		Value:    "Ursula"}

	// ursalaItalianName represents ursala italian name.
	ursalaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ursula"}

	// ursalaJapaneseName represents ursala japanese name.
	ursalaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ネーヤ"}

	// ursalaKoreanName represents ursala korean name.
	ursalaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "네이아"}

	// ursalaLatinAmericanSpanishName represents ursala latin american spanish name.
	ursalaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Úrsula"}

	// ursalaRussianName represents ursala russian name.
	ursalaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Урсала"}

	// ursalaSimplifiedChineseName represents ursala simplified chinese name.
	ursalaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "妮雅"}

	// ursalaSpanishName represents ursala spanish name.
	ursalaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Úrsula"}

	// ursalaTraditionalChineseName represents ursala traditional chinese name.
	ursalaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "妮雅"}
)

// ursalaName represents Ursala's name in different languages.
var (
	// ursalaName represents ursala name.
	ursalaName = nook.Languages{
		language.AmericanEnglish:      ursalaAmericanEnglishName,
		language.CanadianFrench:       ursalaCanadianFrenchName,
		language.Dutch:                ursalaDutchName,
		language.French:               ursalaFrenchName,
		language.German:               ursalaGermanName,
		language.Italian:              ursalaItalianName,
		language.Japanese:             ursalaJapaneseName,
		language.Korean:               ursalaKoreanName,
		language.LatinAmericanSpanish: ursalaLatinAmericanSpanishName,
		language.Russian:              ursalaRussianName,
		language.SimplifiedChinese:    ursalaSimplifiedChineseName,
		language.Spanish:              ursalaSpanishName,
		language.TraditionalChinese:   ursalaTraditionalChineseName}
)

// ursalaCharacter represents Ursala's character information.
var (
	// ursalaCharacter represents ursala character.
	ursalaCharacter = nook.Character{
		Animal:   animal.Bear,
		Birthday: ursalaBirthday,
		Code:     ursalaCode,
		Key:      character.Ursala,
		Gender:   gender.Female,
		Name:     ursalaName,
		Special:  false}
)

// Different phrases spoken by Ursala in various languages.
var (
	// ursalaAmericanEnglishPhrase represents ursala american english phrase.
	ursalaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "grooomph"}

	// ursalaCanadianFrenchPhrase represents ursala canadian french phrase.
	ursalaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "groumpf"}

	// ursalaDutchPhrase represents ursala dutch phrase.
	ursalaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "oemf"}

	// ursalaFrenchPhrase represents ursala french phrase.
	ursalaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "groumpf"}

	// ursalaGermanPhrase represents ursala german phrase.
	ursalaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "brumbrum"}

	// ursalaItalianPhrase represents ursala italian phrase.
	ursalaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "gruuunf"}

	// ursalaJapanesePhrase represents ursala japanese phrase.
	ursalaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "やーねぇ"}

	// ursalaKoreanPhrase represents ursala korean phrase.
	ursalaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그라지"}

	// ursalaLatinAmericanSpanishPhrase represents ursala latin american spanish phrase.
	ursalaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "grumf"}

	// ursalaRussianPhrase represents ursala russian phrase.
	ursalaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "гррум"}

	// ursalaSimplifiedChinesePhrase represents ursala simplified chinese phrase.
	ursalaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "呀呐"}

	// ursalaSpanishPhrase represents ursala spanish phrase.
	ursalaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "grumf"}

	// ursalaTraditionalChinesePhrase represents ursala traditional chinese phrase.
	ursalaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "呀吶"}
)

// ursalaPhrase represents Ursala's phrases in different languages.
var (
	// ursalaPhrase represents ursala phrase.
	ursalaPhrase = nook.Languages{
		language.AmericanEnglish:      ursalaAmericanEnglishPhrase,
		language.CanadianFrench:       ursalaCanadianFrenchPhrase,
		language.Dutch:                ursalaDutchPhrase,
		language.French:               ursalaFrenchPhrase,
		language.German:               ursalaGermanPhrase,
		language.Italian:              ursalaItalianPhrase,
		language.Japanese:             ursalaJapanesePhrase,
		language.Korean:               ursalaKoreanPhrase,
		language.LatinAmericanSpanish: ursalaLatinAmericanSpanishPhrase,
		language.Russian:              ursalaRussianPhrase,
		language.SimplifiedChinese:    ursalaSimplifiedChinesePhrase,
		language.Spanish:              ursalaSpanishPhrase,
		language.TraditionalChinese:   ursalaTraditionalChinesePhrase}
)

// Ursala represents the character Ursala with her complete information.
var (
	// Ursala represents ursala.
	Ursala = nook.Villager{
		Character:   ursalaCharacter,
		Personality: personality.BigSister,
		Phrase:      ursalaPhrase}
)

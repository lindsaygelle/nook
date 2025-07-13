package mouse

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
	// broccoloBirthday represents broccolo birthday.
	broccoloBirthday = nook.Birthday{
		Day:   30,
		Month: time.June}
)

var (
	// broccoloCode represents broccolo code.
	broccoloCode = nook.Code{
		Value: "mus12"}
)

var (
	// broccoloAmericanEnglishName represents broccolo american english name.
	broccoloAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Broccolo"}

	// broccoloCanadianFrenchName represents broccolo canadian french name.
	broccoloCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Steven"}

	// broccoloDutchName represents broccolo dutch name.
	broccoloDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Broccolo"}

	// broccoloFrenchName represents broccolo french name.
	broccoloFrenchName = nook.Name{
		Language: language.French,
		Value:    "Steven"}

	// broccoloGermanName represents broccolo german name.
	broccoloGermanName = nook.Name{
		Language: language.German,
		Value:    "Fritzi"}

	// broccoloItalianName represents broccolo italian name.
	broccoloItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Franco"}

	// broccoloJapaneseName represents broccolo japanese name.
	broccoloJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ブロッコリー"}

	// broccoloKoreanName represents broccolo korean name.
	broccoloKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "브로콜리"}

	// broccoloLatinAmericanSpanishName represents broccolo latin american spanish name.
	broccoloLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Brócoli"}

	// broccoloRussianName represents broccolo russian name.
	broccoloRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Брокколо"}

	// broccoloSimplifiedChineseName represents broccolo simplified chinese name.
	broccoloSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "茜蓝"}

	// broccoloSpanishName represents broccolo spanish name.
	broccoloSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Brócoli"}

	// broccoloTraditionalChineseName represents broccolo traditional chinese name.
	broccoloTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "茜藍"}
)

var (
	// broccoloName represents broccolo name.
	broccoloName = nook.Languages{
		language.AmericanEnglish:      broccoloAmericanEnglishName,
		language.CanadianFrench:       broccoloCanadianFrenchName,
		language.Dutch:                broccoloDutchName,
		language.French:               broccoloFrenchName,
		language.German:               broccoloGermanName,
		language.Italian:              broccoloItalianName,
		language.Japanese:             broccoloJapaneseName,
		language.Korean:               broccoloKoreanName,
		language.LatinAmericanSpanish: broccoloLatinAmericanSpanishName,
		language.Russian:              broccoloRussianName,
		language.SimplifiedChinese:    broccoloSimplifiedChineseName,
		language.Spanish:              broccoloSpanishName,
		language.TraditionalChinese:   broccoloTraditionalChineseName}
)

var (
	// broccoloCharacter represents broccolo character.
	broccoloCharacter = nook.Character{
		Animal:   animal.Mouse,
		Birthday: broccoloBirthday,
		Code:     broccoloCode,
		Key:      character.Broccolo,
		Gender:   gender.Male,
		Name:     broccoloName,
		Special:  false}
)

var (
	// broccoloAmericanEnglishPhrase represents broccolo american english phrase.
	broccoloAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "eat it"}

	// broccoloCanadianFrenchPhrase represents broccolo canadian french phrase.
	broccoloCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "flim-flam"}

	// broccoloDutchPhrase represents broccolo dutch phrase.
	broccoloDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "smakelijk"}

	// broccoloFrenchPhrase represents broccolo french phrase.
	broccoloFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "flim-flam"}

	// broccoloGermanPhrase represents broccolo german phrase.
	broccoloGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "knabba"}

	// broccoloItalianPhrase represents broccolo italian phrase.
	broccoloItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "squikko"}

	// broccoloJapanesePhrase represents broccolo japanese phrase.
	broccoloJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ピコ"}

	// broccoloKoreanPhrase represents broccolo korean phrase.
	broccoloKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "앗차"}

	// broccoloLatinAmericanSpanishPhrase represents broccolo latin american spanish phrase.
	broccoloLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "flinflán"}

	// broccoloRussianPhrase represents broccolo russian phrase.
	broccoloRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ням-ням"}

	// broccoloSimplifiedChinesePhrase represents broccolo simplified chinese phrase.
	broccoloSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "微微"}

	// broccoloSpanishPhrase represents broccolo spanish phrase.
	broccoloSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "dientes"}

	// broccoloTraditionalChinesePhrase represents broccolo traditional chinese phrase.
	broccoloTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "微微"}
)

var (
	// broccoloPhrase represents broccolo phrase.
	broccoloPhrase = nook.Languages{
		language.AmericanEnglish:      broccoloAmericanEnglishPhrase,
		language.CanadianFrench:       broccoloCanadianFrenchPhrase,
		language.Dutch:                broccoloDutchPhrase,
		language.French:               broccoloFrenchPhrase,
		language.German:               broccoloGermanPhrase,
		language.Italian:              broccoloItalianPhrase,
		language.Japanese:             broccoloJapanesePhrase,
		language.Korean:               broccoloKoreanPhrase,
		language.LatinAmericanSpanish: broccoloLatinAmericanSpanishPhrase,
		language.Russian:              broccoloRussianPhrase,
		language.SimplifiedChinese:    broccoloSimplifiedChinesePhrase,
		language.Spanish:              broccoloSpanishPhrase,
		language.TraditionalChinese:   broccoloTraditionalChinesePhrase}
)

var (
	// Broccolo represents broccolo.
	Broccolo = nook.Villager{
		Character:   broccoloCharacter,
		Personality: personality.Lazy,
		Phrase:      broccoloPhrase}
)

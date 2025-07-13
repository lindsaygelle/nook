package octopus

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
	// inkwellBirthday represents inkwell birthday.
	inkwellBirthday = nook.Birthday{
		Day:   2,
		Month: time.June}
)

var (
	// inkwellCode represents inkwell code.
	inkwellCode = nook.Code{
		Value: "ocp03"}
)

var (
	// inkwellAmericanEnglishName represents inkwell american english name.
	inkwellAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Inkwell"}

	// inkwellCanadianFrenchName represents inkwell canadian french name.
	inkwellCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Pigmento"}

	// inkwellDutchName represents inkwell dutch name.
	inkwellDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// inkwellFrenchName represents inkwell french name.
	inkwellFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pigmento"}

	// inkwellGermanName represents inkwell german name.
	inkwellGermanName = nook.Name{
		Language: language.German,
		Value:    "Klecks"}

	// inkwellItalianName represents inkwell italian name.
	inkwellItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Polpotto"}

	// inkwellJapaneseName represents inkwell japanese name.
	inkwellJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "スミダクン"}

	// inkwellKoreanName represents inkwell korean name.
	inkwellKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "멍무리"}

	// inkwellLatinAmericanSpanishName represents inkwell latin american spanish name.
	inkwellLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Tintelio"}

	// inkwellRussianName represents inkwell russian name.
	inkwellRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// inkwellSimplifiedChineseName represents inkwell simplified chinese name.
	inkwellSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// inkwellSpanishName represents inkwell spanish name.
	inkwellSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tintelio"}

	// inkwellTraditionalChineseName represents inkwell traditional chinese name.
	inkwellTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// inkwellName represents inkwell name.
	inkwellName = nook.Languages{
		language.AmericanEnglish:      inkwellAmericanEnglishName,
		language.CanadianFrench:       inkwellCanadianFrenchName,
		language.Dutch:                inkwellDutchName,
		language.French:               inkwellFrenchName,
		language.German:               inkwellGermanName,
		language.Italian:              inkwellItalianName,
		language.Japanese:             inkwellJapaneseName,
		language.Korean:               inkwellKoreanName,
		language.LatinAmericanSpanish: inkwellLatinAmericanSpanishName,
		language.Russian:              inkwellRussianName,
		language.SimplifiedChinese:    inkwellSimplifiedChineseName,
		language.Spanish:              inkwellSpanishName,
		language.TraditionalChinese:   inkwellTraditionalChineseName}
)

var (
	// inkwellCharacter represents inkwell character.
	inkwellCharacter = nook.Character{
		Animal:   animal.Octopus,
		Birthday: inkwellBirthday,
		Code:     inkwellCode,
		Key:      character.Inkwell,
		Gender:   gender.Male,
		Name:     inkwellName,
		Special:  false}
)

var (
	// inkwellAmericanEnglishPhrase represents inkwell american english phrase.
	inkwellAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "splat"}

	// inkwellCanadianFrenchPhrase represents inkwell canadian french phrase.
	inkwellCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "splat"}

	// inkwellDutchPhrase represents inkwell dutch phrase.
	inkwellDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// inkwellFrenchPhrase represents inkwell french phrase.
	inkwellFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "splat"}

	// inkwellGermanPhrase represents inkwell german phrase.
	inkwellGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "splat"}

	// inkwellItalianPhrase represents inkwell italian phrase.
	inkwellItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "splat"}

	// inkwellJapanesePhrase represents inkwell japanese phrase.
	inkwellJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ガチで"}

	// inkwellKoreanPhrase represents inkwell korean phrase.
	inkwellKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "진심"}

	// inkwellLatinAmericanSpanishPhrase represents inkwell latin american spanish phrase.
	inkwellLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "tinta va"}

	// inkwellRussianPhrase represents inkwell russian phrase.
	inkwellRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// inkwellSimplifiedChinesePhrase represents inkwell simplified chinese phrase.
	inkwellSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// inkwellSpanishPhrase represents inkwell spanish phrase.
	inkwellSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tinta va"}

	// inkwellTraditionalChinesePhrase represents inkwell traditional chinese phrase.
	inkwellTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// inkwellPhrase represents inkwell phrase.
	inkwellPhrase = nook.Languages{
		language.AmericanEnglish:      inkwellAmericanEnglishPhrase,
		language.CanadianFrench:       inkwellCanadianFrenchPhrase,
		language.Dutch:                inkwellDutchPhrase,
		language.French:               inkwellFrenchPhrase,
		language.German:               inkwellGermanPhrase,
		language.Italian:              inkwellItalianPhrase,
		language.Japanese:             inkwellJapanesePhrase,
		language.Korean:               inkwellKoreanPhrase,
		language.LatinAmericanSpanish: inkwellLatinAmericanSpanishPhrase,
		language.Russian:              inkwellRussianPhrase,
		language.SimplifiedChinese:    inkwellSimplifiedChinesePhrase,
		language.Spanish:              inkwellSpanishPhrase,
		language.TraditionalChinese:   inkwellTraditionalChinesePhrase}
)

var (
	// Inkwell represents inkwell.
	Inkwell = nook.Villager{
		Character:   inkwellCharacter,
		Personality: personality.Jock,
		Phrase:      inkwellPhrase}
)

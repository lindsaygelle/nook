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
	inkwellBirthday = nook.Birthday{
		Day:   2,
		Month: time.June}
)

var (
	inkwellCode = nook.Code{
		Value: "ocp03"}
)

var (
	inkwellAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Inkwell"}

	inkwellCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Pigmento"}

	inkwellDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	inkwellFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pigmento"}

	inkwellGermanName = nook.Name{
		Language: language.German,
		Value:    "Klecks"}

	inkwellItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Polpotto"}

	inkwellJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "スミダクン"}

	inkwellKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "멍무리"}

	inkwellLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Tintelio"}

	inkwellRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	inkwellSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	inkwellSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tintelio"}

	inkwellTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
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
	inkwellCharacter = nook.Character{
		Animal:   animal.Octopus,
		Birthday: inkwellBirthday,
		Code:     inkwellCode,
		Key:      character.Inkwell,
		Gender:   gender.Male,
		Name:     inkwellName}
)

var (
	inkwellAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "splat"}

	inkwellCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "splat"}

	inkwellDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	inkwellFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "splat"}

	inkwellGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "splat"}

	inkwellItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "splat"}

	inkwellJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ガチで"}

	inkwellKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "진심"}

	inkwellLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "tinta va"}

	inkwellRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	inkwellSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	inkwellSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tinta va"}

	inkwellTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
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
	Inkwell = nook.Villager{
		Character:   inkwellCharacter,
		Personality: personality.Jock,
		Phrase:      inkwellPhrase}
)

package chicken

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
	leighBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	leighCode = nook.Code{
		Value: ""}
)

var (
	leighAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Leigh"}

	leighCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	leighDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	leighFrenchName = nook.Name{
		Language: language.French,
		Value:    "Huguette"}

	leighGermanName = nook.Name{
		Language: language.German,
		Value:    "Loretta"}

	leighItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Piumetta"}

	leighJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ディアナ"}

	leighKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	leighLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	leighRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	leighSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "画眉"}

	leighSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tabita"}

	leighTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	leighName = nook.Languages{
		language.AmericanEnglish:      leighAmericanEnglishName,
		language.CanadianFrench:       leighCanadianFrenchName,
		language.Dutch:                leighDutchName,
		language.French:               leighFrenchName,
		language.German:               leighGermanName,
		language.Italian:              leighItalianName,
		language.Japanese:             leighJapaneseName,
		language.Korean:               leighKoreanName,
		language.LatinAmericanSpanish: leighLatinAmericanSpanishName,
		language.Russian:              leighRussianName,
		language.SimplifiedChinese:    leighSimplifiedChineseName,
		language.Spanish:              leighSpanishName,
		language.TraditionalChinese:   leighTraditionalChineseName}
)

var (
	leighCharacter = nook.Character{
		Animal:   animal.Chicken,
		Birthday: leighBirthday,
		Code:     leighCode,
		Key:      character.Leigh,
		Gender:   gender.Female,
		Name:     leighName,
		Special:  false}
)

var (
	leighAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cutie"}

	leighCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	leighDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	leighFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "poussinet"}

	leighGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "goldi"}

	leighItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cocò"}

	leighJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だもん"}

	leighKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	leighLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	leighRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	leighSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啄啄"}

	leighSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chuli"}

	leighTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	leighPhrase = nook.Languages{
		language.AmericanEnglish:      leighAmericanEnglishPhrase,
		language.CanadianFrench:       leighCanadianFrenchPhrase,
		language.Dutch:                leighDutchPhrase,
		language.French:               leighFrenchPhrase,
		language.German:               leighGermanPhrase,
		language.Italian:              leighItalianPhrase,
		language.Japanese:             leighJapanesePhrase,
		language.Korean:               leighKoreanPhrase,
		language.LatinAmericanSpanish: leighLatinAmericanSpanishPhrase,
		language.Russian:              leighRussianPhrase,
		language.SimplifiedChinese:    leighSimplifiedChinesePhrase,
		language.Spanish:              leighSpanishPhrase,
		language.TraditionalChinese:   leighTraditionalChinesePhrase}
)

var (
	Leigh = nook.Villager{
		Character:   leighCharacter,
		Personality: personality.Peppy,
		Phrase:      leighPhrase}
)

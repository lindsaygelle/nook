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
	// leighBirthday represents leigh birthday.
	leighBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// leighCode represents leigh code.
	leighCode = nook.Code{
		Value: ""}
)

var (
	// leighAmericanEnglishName represents leigh american english name.
	leighAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Leigh"}

	// leighCanadianFrenchName represents leigh canadian french name.
	leighCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// leighDutchName represents leigh dutch name.
	leighDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// leighFrenchName represents leigh french name.
	leighFrenchName = nook.Name{
		Language: language.French,
		Value:    "Huguette"}

	// leighGermanName represents leigh german name.
	leighGermanName = nook.Name{
		Language: language.German,
		Value:    "Loretta"}

	// leighItalianName represents leigh italian name.
	leighItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Piumetta"}

	// leighJapaneseName represents leigh japanese name.
	leighJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ディアナ"}

	// leighKoreanName represents leigh korean name.
	leighKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// leighLatinAmericanSpanishName represents leigh latin american spanish name.
	leighLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// leighRussianName represents leigh russian name.
	leighRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// leighSimplifiedChineseName represents leigh simplified chinese name.
	leighSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "画眉"}

	// leighSpanishName represents leigh spanish name.
	leighSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tabita"}

	// leighTraditionalChineseName represents leigh traditional chinese name.
	leighTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// leighName represents leigh name.
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
	// leighCharacter represents leigh character.
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
	// leighAmericanEnglishPhrase represents leigh american english phrase.
	leighAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cutie"}

	// leighCanadianFrenchPhrase represents leigh canadian french phrase.
	leighCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// leighDutchPhrase represents leigh dutch phrase.
	leighDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// leighFrenchPhrase represents leigh french phrase.
	leighFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "poussinet"}

	// leighGermanPhrase represents leigh german phrase.
	leighGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "goldi"}

	// leighItalianPhrase represents leigh italian phrase.
	leighItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cocò"}

	// leighJapanesePhrase represents leigh japanese phrase.
	leighJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だもん"}

	// leighKoreanPhrase represents leigh korean phrase.
	leighKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// leighLatinAmericanSpanishPhrase represents leigh latin american spanish phrase.
	leighLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// leighRussianPhrase represents leigh russian phrase.
	leighRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// leighSimplifiedChinesePhrase represents leigh simplified chinese phrase.
	leighSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啄啄"}

	// leighSpanishPhrase represents leigh spanish phrase.
	leighSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chuli"}

	// leighTraditionalChinesePhrase represents leigh traditional chinese phrase.
	leighTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// leighPhrase represents leigh phrase.
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
	// Leigh represents leigh.
	Leigh = nook.Villager{
		Character:   leighCharacter,
		Personality: personality.Peppy,
		Phrase:      leighPhrase}
)

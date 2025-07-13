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
	// rooneyBirthday represents rooney birthday.
	rooneyBirthday = nook.Birthday{
		Day:   1,
		Month: time.December}
)

var (
	// rooneyCode represents rooney code.
	rooneyCode = nook.Code{
		Value: "kgr09"}
)

var (
	// rooneyAmericanEnglishName represents rooney american english name.
	rooneyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rooney"}

	// rooneyCanadianFrenchName represents rooney canadian french name.
	rooneyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Mike"}

	// rooneyDutchName represents rooney dutch name.
	rooneyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rooney"}

	// rooneyFrenchName represents rooney french name.
	rooneyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mike"}

	// rooneyGermanName represents rooney german name.
	rooneyGermanName = nook.Name{
		Language: language.German,
		Value:    "Robert"}

	// rooneyItalianName represents rooney italian name.
	rooneyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Balzak"}

	// rooneyJapaneseName represents rooney japanese name.
	rooneyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マイク"}

	// rooneyKoreanName represents rooney korean name.
	rooneyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마이크"}

	// rooneyLatinAmericanSpanishName represents rooney latin american spanish name.
	rooneyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cerillo"}

	// rooneyRussianName represents rooney russian name.
	rooneyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Руни"}

	// rooneySimplifiedChineseName represents rooney simplified chinese name.
	rooneySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "麦克"}

	// rooneySpanishName represents rooney spanish name.
	rooneySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cerillo"}

	// rooneyTraditionalChineseName represents rooney traditional chinese name.
	rooneyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "麥克"}
)

var (
	// rooneyName represents rooney name.
	rooneyName = nook.Languages{
		language.AmericanEnglish:      rooneyAmericanEnglishName,
		language.CanadianFrench:       rooneyCanadianFrenchName,
		language.Dutch:                rooneyDutchName,
		language.French:               rooneyFrenchName,
		language.German:               rooneyGermanName,
		language.Italian:              rooneyItalianName,
		language.Japanese:             rooneyJapaneseName,
		language.Korean:               rooneyKoreanName,
		language.LatinAmericanSpanish: rooneyLatinAmericanSpanishName,
		language.Russian:              rooneyRussianName,
		language.SimplifiedChinese:    rooneySimplifiedChineseName,
		language.Spanish:              rooneySpanishName,
		language.TraditionalChinese:   rooneyTraditionalChineseName}
)

var (
	// rooneyCharacter represents rooney character.
	rooneyCharacter = nook.Character{
		Animal:   animal.Kangaroo,
		Birthday: rooneyBirthday,
		Code:     rooneyCode,
		Key:      character.Rooney,
		Gender:   gender.Male,
		Name:     rooneyName,
		Special:  false}
)

var (
	// rooneyAmericanEnglishPhrase represents rooney american english phrase.
	rooneyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "punches"}

	// rooneyCanadianFrenchPhrase represents rooney canadian french phrase.
	rooneyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "savate"}

	// rooneyDutchPhrase represents rooney dutch phrase.
	rooneyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "boks"}

	// rooneyFrenchPhrase represents rooney french phrase.
	rooneyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "savate"}

	// rooneyGermanPhrase represents rooney german phrase.
	rooneyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bonk"}

	// rooneyItalianPhrase represents rooney italian phrase.
	rooneyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "balz balz"}

	// rooneyJapanesePhrase represents rooney japanese phrase.
	rooneyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "やるぜ"}

	// rooneyKoreanPhrase represents rooney korean phrase.
	rooneyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아자"}

	// rooneyLatinAmericanSpanishPhrase represents rooney latin american spanish phrase.
	rooneyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "chimpón"}

	// rooneyRussianPhrase represents rooney russian phrase.
	rooneyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хук"}

	// rooneySimplifiedChinesePhrase represents rooney simplified chinese phrase.
	rooneySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "做就对了"}

	// rooneySpanishPhrase represents rooney spanish phrase.
	rooneySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chimpón"}

	// rooneyTraditionalChinesePhrase represents rooney traditional chinese phrase.
	rooneyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "做就對了"}
)

var (
	// rooneyPhrase represents rooney phrase.
	rooneyPhrase = nook.Languages{
		language.AmericanEnglish:      rooneyAmericanEnglishPhrase,
		language.CanadianFrench:       rooneyCanadianFrenchPhrase,
		language.Dutch:                rooneyDutchPhrase,
		language.French:               rooneyFrenchPhrase,
		language.German:               rooneyGermanPhrase,
		language.Italian:              rooneyItalianPhrase,
		language.Japanese:             rooneyJapanesePhrase,
		language.Korean:               rooneyKoreanPhrase,
		language.LatinAmericanSpanish: rooneyLatinAmericanSpanishPhrase,
		language.Russian:              rooneyRussianPhrase,
		language.SimplifiedChinese:    rooneySimplifiedChinesePhrase,
		language.Spanish:              rooneySpanishPhrase,
		language.TraditionalChinese:   rooneyTraditionalChinesePhrase}
)

var (
	// Rooney represents rooney.
	Rooney = nook.Villager{
		Character:   rooneyCharacter,
		Personality: personality.Cranky,
		Phrase:      rooneyPhrase}
)

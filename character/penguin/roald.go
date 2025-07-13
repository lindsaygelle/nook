package penguin

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
	// roaldBirthday represents roald birthday.
	roaldBirthday = nook.Birthday{
		Day:   5,
		Month: time.January}
)

var (
	// roaldCode represents roald code.
	roaldCode = nook.Code{
		Value: "pgn01"}
)

var (
	// roaldAmericanEnglishName represents roald american english name.
	roaldAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Roald"}

	// roaldCanadianFrenchName represents roald canadian french name.
	roaldCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Reynald"}

	// roaldDutchName represents roald dutch name.
	roaldDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Roald"}

	// roaldFrenchName represents roald french name.
	roaldFrenchName = nook.Name{
		Language: language.French,
		Value:    "Reynald"}

	// roaldGermanName represents roald german name.
	roaldGermanName = nook.Name{
		Language: language.German,
		Value:    "Roland"}

	// roaldItalianName represents roald italian name.
	roaldItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Angelino"}

	// roaldJapaneseName represents roald japanese name.
	roaldJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ペンタ"}

	// roaldKoreanName represents roald korean name.
	roaldKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "펭수"}

	// roaldLatinAmericanSpanishName represents roald latin american spanish name.
	roaldLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bobi"}

	// roaldRussianName represents roald russian name.
	roaldRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Роальд"}

	// roaldSimplifiedChineseName represents roald simplified chinese name.
	roaldSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "企鹅达"}

	// roaldSpanishName represents roald spanish name.
	roaldSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bobi"}

	// roaldTraditionalChineseName represents roald traditional chinese name.
	roaldTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "企鵝達"}
)

var (
	// roaldName represents roald name.
	roaldName = nook.Languages{
		language.AmericanEnglish:      roaldAmericanEnglishName,
		language.CanadianFrench:       roaldCanadianFrenchName,
		language.Dutch:                roaldDutchName,
		language.French:               roaldFrenchName,
		language.German:               roaldGermanName,
		language.Italian:              roaldItalianName,
		language.Japanese:             roaldJapaneseName,
		language.Korean:               roaldKoreanName,
		language.LatinAmericanSpanish: roaldLatinAmericanSpanishName,
		language.Russian:              roaldRussianName,
		language.SimplifiedChinese:    roaldSimplifiedChineseName,
		language.Spanish:              roaldSpanishName,
		language.TraditionalChinese:   roaldTraditionalChineseName}
)

var (
	// roaldCharacter represents roald character.
	roaldCharacter = nook.Character{
		Animal:   animal.Penguin,
		Birthday: roaldBirthday,
		Code:     roaldCode,
		Key:      character.Roald,
		Gender:   gender.Male,
		Name:     roaldName,
		Special:  false}
)

var (
	// roaldAmericanEnglishPhrase represents roald american english phrase.
	roaldAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "b-b-buddy"}

	// roaldCanadianFrenchPhrase represents roald canadian french phrase.
	roaldCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pépépère"}

	// roaldDutchPhrase represents roald dutch phrase.
	roaldDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "m-m-maatje"}

	// roaldFrenchPhrase represents roald french phrase.
	roaldFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pépépère"}

	// roaldGermanPhrase represents roald german phrase.
	roaldGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "k-k-kumpel"}

	// roaldItalianPhrase represents roald italian phrase.
	roaldItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "br-br-brr"}

	// roaldJapanesePhrase represents roald japanese phrase.
	roaldJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だペン"}

	// roaldKoreanPhrase represents roald korean phrase.
	roaldKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "팽팽"}

	// roaldLatinAmericanSpanishPhrase represents roald latin american spanish phrase.
	roaldLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "titirito"}

	// roaldRussianPhrase represents roald russian phrase.
	roaldRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "д-дружище"}

	// roaldSimplifiedChinesePhrase represents roald simplified chinese phrase.
	roaldSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "企鹅"}

	// roaldSpanishPhrase represents roald spanish phrase.
	roaldSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "titirito"}

	// roaldTraditionalChinesePhrase represents roald traditional chinese phrase.
	roaldTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "企鵝"}
)

var (
	// roaldPhrase represents roald phrase.
	roaldPhrase = nook.Languages{
		language.AmericanEnglish:      roaldAmericanEnglishPhrase,
		language.CanadianFrench:       roaldCanadianFrenchPhrase,
		language.Dutch:                roaldDutchPhrase,
		language.French:               roaldFrenchPhrase,
		language.German:               roaldGermanPhrase,
		language.Italian:              roaldItalianPhrase,
		language.Japanese:             roaldJapanesePhrase,
		language.Korean:               roaldKoreanPhrase,
		language.LatinAmericanSpanish: roaldLatinAmericanSpanishPhrase,
		language.Russian:              roaldRussianPhrase,
		language.SimplifiedChinese:    roaldSimplifiedChinesePhrase,
		language.Spanish:              roaldSpanishPhrase,
		language.TraditionalChinese:   roaldTraditionalChinesePhrase}
)

var (
	// Roald represents roald.
	Roald = nook.Villager{
		Character:   roaldCharacter,
		Personality: personality.Jock,
		Phrase:      roaldPhrase}
)

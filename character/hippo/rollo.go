package hippo

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
	// rolloBirthday represents rollo birthday.
	rolloBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// rolloCode represents rollo code.
	rolloCode = nook.Code{
		Value: ""}
)

var (
	// rolloAmericanEnglishName represents rollo american english name.
	rolloAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rollo"}

	// rolloCanadianFrenchName represents rollo canadian french name.
	rolloCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// rolloDutchName represents rollo dutch name.
	rolloDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// rolloFrenchName represents rollo french name.
	rolloFrenchName = nook.Name{
		Language: language.French,
		Value:    "Hippo"}

	// rolloGermanName represents rollo german name.
	rolloGermanName = nook.Name{
		Language: language.German,
		Value:    "Winfried"}

	// rolloItalianName represents rollo italian name.
	rolloItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rollo"}

	// rolloJapaneseName represents rollo japanese name.
	rolloJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ヒポクラテス"}

	// rolloKoreanName represents rollo korean name.
	rolloKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// rolloLatinAmericanSpanishName represents rollo latin american spanish name.
	rolloLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// rolloRussianName represents rollo russian name.
	rolloRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// rolloSimplifiedChineseName represents rollo simplified chinese name.
	rolloSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "罗素"}

	// rolloSpanishName represents rollo spanish name.
	rolloSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Hipo"}

	// rolloTraditionalChineseName represents rollo traditional chinese name.
	rolloTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// rolloName represents rollo name.
	rolloName = nook.Languages{
		language.AmericanEnglish:      rolloAmericanEnglishName,
		language.CanadianFrench:       rolloCanadianFrenchName,
		language.Dutch:                rolloDutchName,
		language.French:               rolloFrenchName,
		language.German:               rolloGermanName,
		language.Italian:              rolloItalianName,
		language.Japanese:             rolloJapaneseName,
		language.Korean:               rolloKoreanName,
		language.LatinAmericanSpanish: rolloLatinAmericanSpanishName,
		language.Russian:              rolloRussianName,
		language.SimplifiedChinese:    rolloSimplifiedChineseName,
		language.Spanish:              rolloSpanishName,
		language.TraditionalChinese:   rolloTraditionalChineseName}
)

var (
	// rolloCharacter represents rollo character.
	rolloCharacter = nook.Character{
		Animal:   animal.Hippo,
		Birthday: rolloBirthday,
		Code:     rolloCode,
		Key:      character.Rollo,
		Gender:   gender.Male,
		Name:     rolloName,
		Special:  false}
)

var (
	// rolloAmericanEnglishPhrase represents rollo american english phrase.
	rolloAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "beaulch"}

	// rolloCanadianFrenchPhrase represents rollo canadian french phrase.
	rolloCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// rolloDutchPhrase represents rollo dutch phrase.
	rolloDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// rolloFrenchPhrase represents rollo french phrase.
	rolloFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "que diable"}

	// rolloGermanPhrase represents rollo german phrase.
	rolloGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "rülps"}

	// rolloItalianPhrase represents rollo italian phrase.
	rolloItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "trippo"}

	// rolloJapanesePhrase represents rollo japanese phrase.
	rolloJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なのな"}

	// rolloKoreanPhrase represents rollo korean phrase.
	rolloKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// rolloLatinAmericanSpanishPhrase represents rollo latin american spanish phrase.
	rolloLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// rolloRussianPhrase represents rollo russian phrase.
	rolloRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// rolloSimplifiedChinesePhrase represents rollo simplified chinese phrase.
	rolloSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哈"}

	// rolloSpanishPhrase represents rollo spanish phrase.
	rolloSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "hip-hip"}

	// rolloTraditionalChinesePhrase represents rollo traditional chinese phrase.
	rolloTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// rolloPhrase represents rollo phrase.
	rolloPhrase = nook.Languages{
		language.AmericanEnglish:      rolloAmericanEnglishPhrase,
		language.CanadianFrench:       rolloCanadianFrenchPhrase,
		language.Dutch:                rolloDutchPhrase,
		language.French:               rolloFrenchPhrase,
		language.German:               rolloGermanPhrase,
		language.Italian:              rolloItalianPhrase,
		language.Japanese:             rolloJapanesePhrase,
		language.Korean:               rolloKoreanPhrase,
		language.LatinAmericanSpanish: rolloLatinAmericanSpanishPhrase,
		language.Russian:              rolloRussianPhrase,
		language.SimplifiedChinese:    rolloSimplifiedChinesePhrase,
		language.Spanish:              rolloSpanishPhrase,
		language.TraditionalChinese:   rolloTraditionalChinesePhrase}
)

var (
	// Rollo represents rollo.
	Rollo = nook.Villager{
		Character:   rolloCharacter,
		Personality: personality.Lazy,
		Phrase:      rolloPhrase}
)

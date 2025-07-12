package bull

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
	// oxfordBirthday represents oxford birthday.
	oxfordBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// oxfordCode represents oxford code.
	oxfordCode = nook.Code{
		Value: ""}
)

var (
	// oxfordAmericanEnglishName represents oxford american english name.
	oxfordAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Oxford"}

	// oxfordCanadianFrenchName represents oxford canadian french name.
	oxfordCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// oxfordDutchName represents oxford dutch name.
	oxfordDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// oxfordFrenchName represents oxford french name.
	oxfordFrenchName = nook.Name{
		Language: language.French,
		Value:    "Sorbonne"}

	// oxfordGermanName represents oxford german name.
	oxfordGermanName = nook.Name{
		Language: language.German,
		Value:    "Rolf"}

	// oxfordItalianName represents oxford italian name.
	oxfordItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Angus"}

	// oxfordJapaneseName represents oxford japanese name.
	oxfordJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "イワオ"}

	// oxfordKoreanName represents oxford korean name.
	oxfordKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// oxfordLatinAmericanSpanishName represents oxford latin american spanish name.
	oxfordLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// oxfordRussianName represents oxford russian name.
	oxfordRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// oxfordSimplifiedChineseName represents oxford simplified chinese name.
	oxfordSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "岩岩"}

	// oxfordSpanishName represents oxford spanish name.
	oxfordSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Astaúlfo"}

	// oxfordTraditionalChineseName represents oxford traditional chinese name.
	oxfordTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// oxfordName represents oxford name.
	oxfordName = nook.Languages{
		language.AmericanEnglish:      oxfordAmericanEnglishName,
		language.CanadianFrench:       oxfordCanadianFrenchName,
		language.Dutch:                oxfordDutchName,
		language.French:               oxfordFrenchName,
		language.German:               oxfordGermanName,
		language.Italian:              oxfordItalianName,
		language.Japanese:             oxfordJapaneseName,
		language.Korean:               oxfordKoreanName,
		language.LatinAmericanSpanish: oxfordLatinAmericanSpanishName,
		language.Russian:              oxfordRussianName,
		language.SimplifiedChinese:    oxfordSimplifiedChineseName,
		language.Spanish:              oxfordSpanishName,
		language.TraditionalChinese:   oxfordTraditionalChineseName}
)

var (
	// oxfordCharacter represents oxford character.
	oxfordCharacter = nook.Character{
		Animal:   animal.Bull,
		Birthday: oxfordBirthday,
		Code:     oxfordCode,
		Key:      character.Oxford,
		Gender:   gender.Male,
		Name:     oxfordName,
		Special:  false}
)

var (
	// oxfordAmericanEnglishPhrase represents oxford american english phrase.
	oxfordAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bully, eh"}

	// oxfordCanadianFrenchPhrase represents oxford canadian french phrase.
	oxfordCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// oxfordDutchPhrase represents oxford dutch phrase.
	oxfordDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// oxfordFrenchPhrase represents oxford french phrase.
	oxfordFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "la vache"}

	// oxfordGermanPhrase represents oxford german phrase.
	oxfordGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bulli, hä"}

	// oxfordItalianPhrase represents oxford italian phrase.
	oxfordItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "muuh"}

	// oxfordJapanesePhrase represents oxford japanese phrase.
	oxfordJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でんがな"}

	// oxfordKoreanPhrase represents oxford korean phrase.
	oxfordKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// oxfordLatinAmericanSpanishPhrase represents oxford latin american spanish phrase.
	oxfordLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// oxfordRussianPhrase represents oxford russian phrase.
	oxfordRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// oxfordSimplifiedChinesePhrase represents oxford simplified chinese phrase.
	oxfordSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "可不"}

	// oxfordSpanishPhrase represents oxford spanish phrase.
	oxfordSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "torito"}

	// oxfordTraditionalChinesePhrase represents oxford traditional chinese phrase.
	oxfordTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// oxfordPhrase represents oxford phrase.
	oxfordPhrase = nook.Languages{
		language.AmericanEnglish:      oxfordAmericanEnglishPhrase,
		language.CanadianFrench:       oxfordCanadianFrenchPhrase,
		language.Dutch:                oxfordDutchPhrase,
		language.French:               oxfordFrenchPhrase,
		language.German:               oxfordGermanPhrase,
		language.Italian:              oxfordItalianPhrase,
		language.Japanese:             oxfordJapanesePhrase,
		language.Korean:               oxfordKoreanPhrase,
		language.LatinAmericanSpanish: oxfordLatinAmericanSpanishPhrase,
		language.Russian:              oxfordRussianPhrase,
		language.SimplifiedChinese:    oxfordSimplifiedChinesePhrase,
		language.Spanish:              oxfordSpanishPhrase,
		language.TraditionalChinese:   oxfordTraditionalChinesePhrase}
)

var (
	// Oxford represents oxford.
	Oxford = nook.Villager{
		Character:   oxfordCharacter,
		Personality: personality.Cranky,
		Phrase:      oxfordPhrase}
)

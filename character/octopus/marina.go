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
	// marinaBirthday represents marina birthday.
	marinaBirthday = nook.Birthday{
		Day:   26,
		Month: time.June}
)

var (
	// marinaCode represents marina code.
	marinaCode = nook.Code{
		Value: "ocp01"}
)

var (
	// marinaAmericanEnglishName represents marina american english name.
	marinaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Marina"}

	// marinaCanadianFrenchName represents marina canadian french name.
	marinaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Marina"}

	// marinaDutchName represents marina dutch name.
	marinaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Marina"}

	// marinaFrenchName represents marina french name.
	marinaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Marina"}

	// marinaGermanName represents marina german name.
	marinaGermanName = nook.Name{
		Language: language.German,
		Value:    "Marianne"}

	// marinaItalianName represents marina italian name.
	marinaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Marina"}

	// marinaJapaneseName represents marina japanese name.
	marinaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "タコリーナ"}

	// marinaKoreanName represents marina korean name.
	marinaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "문리나"}

	// marinaLatinAmericanSpanishName represents marina latin american spanish name.
	marinaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Marina"}

	// marinaRussianName represents marina russian name.
	marinaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Марина"}

	// marinaSimplifiedChineseName represents marina simplified chinese name.
	marinaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "章立娜"}

	// marinaSpanishName represents marina spanish name.
	marinaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Marina"}

	// marinaTraditionalChineseName represents marina traditional chinese name.
	marinaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "章立娜"}
)

var (
	// marinaName represents marina name.
	marinaName = nook.Languages{
		language.AmericanEnglish:      marinaAmericanEnglishName,
		language.CanadianFrench:       marinaCanadianFrenchName,
		language.Dutch:                marinaDutchName,
		language.French:               marinaFrenchName,
		language.German:               marinaGermanName,
		language.Italian:              marinaItalianName,
		language.Japanese:             marinaJapaneseName,
		language.Korean:               marinaKoreanName,
		language.LatinAmericanSpanish: marinaLatinAmericanSpanishName,
		language.Russian:              marinaRussianName,
		language.SimplifiedChinese:    marinaSimplifiedChineseName,
		language.Spanish:              marinaSpanishName,
		language.TraditionalChinese:   marinaTraditionalChineseName}
)

var (
	// marinaCharacter represents marina character.
	marinaCharacter = nook.Character{
		Animal:   animal.Octopus,
		Birthday: marinaBirthday,
		Code:     marinaCode,
		Key:      character.Marina,
		Gender:   gender.Female,
		Name:     marinaName,
		Special:  false}
)

var (
	// marinaAmericanEnglishPhrase represents marina american english phrase.
	marinaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "blurp"}

	// marinaCanadianFrenchPhrase represents marina canadian french phrase.
	marinaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "beurp"}

	// marinaDutchPhrase represents marina dutch phrase.
	marinaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "blorp"}

	// marinaFrenchPhrase represents marina french phrase.
	marinaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "beurp"}

	// marinaGermanPhrase represents marina german phrase.
	marinaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "blubblub"}

	// marinaItalianPhrase represents marina italian phrase.
	marinaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "blurp"}

	// marinaJapanesePhrase represents marina japanese phrase.
	marinaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "きゃ"}

	// marinaKoreanPhrase represents marina korean phrase.
	marinaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "캬캬"}

	// marinaLatinAmericanSpanishPhrase represents marina latin american spanish phrase.
	marinaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "bliup"}

	// marinaRussianPhrase represents marina russian phrase.
	marinaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хлюп"}

	// marinaSimplifiedChinesePhrase represents marina simplified chinese phrase.
	marinaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咔"}

	// marinaSpanishPhrase represents marina spanish phrase.
	marinaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "dospatas"}

	// marinaTraditionalChinesePhrase represents marina traditional chinese phrase.
	marinaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咔"}
)

var (
	// marinaPhrase represents marina phrase.
	marinaPhrase = nook.Languages{
		language.AmericanEnglish:      marinaAmericanEnglishPhrase,
		language.CanadianFrench:       marinaCanadianFrenchPhrase,
		language.Dutch:                marinaDutchPhrase,
		language.French:               marinaFrenchPhrase,
		language.German:               marinaGermanPhrase,
		language.Italian:              marinaItalianPhrase,
		language.Japanese:             marinaJapanesePhrase,
		language.Korean:               marinaKoreanPhrase,
		language.LatinAmericanSpanish: marinaLatinAmericanSpanishPhrase,
		language.Russian:              marinaRussianPhrase,
		language.SimplifiedChinese:    marinaSimplifiedChinesePhrase,
		language.Spanish:              marinaSpanishPhrase,
		language.TraditionalChinese:   marinaTraditionalChinesePhrase}
)

var (
	// Marina represents marina.
	Marina = nook.Villager{
		Character:   marinaCharacter,
		Personality: personality.Normal,
		Phrase:      marinaPhrase}
)

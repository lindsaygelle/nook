package anteater

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

// snootyBirthday represents Snooty's birthday.
var (
	// snootyBirthday represents snooty birthday.
	snootyBirthday = nook.Birthday{
		Day:   24,
		Month: time.October}
)

// snootyCode represents Snooty's unique code.
var (
	// snootyCode represents snooty code.
	snootyCode = nook.Code{
		Value: "ant06"}
)

// Different names for Snooty in various languages.
var (
	// snootyAmericanEnglishName represents snooty american english name.
	snootyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Snooty"}

	// snootyCanadianFrenchName represents snooty canadian french name.
	snootyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Tarina"}

	// snootyDutchName represents snooty dutch name.
	snootyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Snooty"}

	// snootyFrenchName represents snooty french name.
	snootyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tarina"}

	// snootyGermanName represents snooty german name.
	snootyGermanName = nook.Name{
		Language: language.German,
		Value:    "Hilde"}

	// snootyItalianName represents snooty italian name.
	snootyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lucrezia"}

	// snootyJapaneseName represents snooty japanese name.
	snootyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "こまち"}

	// snootyKoreanName represents snooty korean name.
	snootyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "스누티"}

	// snootyLatinAmericanSpanishName represents snooty latin american spanish name.
	snootyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rosanari"}

	// snootyRussianName represents snooty russian name.
	snootyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Снути"}

	// snootySimplifiedChineseName represents snooty simplified chinese name.
	snootySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "贵妃"}

	// snootySpanishName represents snooty spanish name.
	snootySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rosanari"}

	// snootyTraditionalChineseName represents snooty traditional chinese name.
	snootyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "貴妃"}
)

// snootyName represents Snooty's name in different languages.
var (
	// snootyName represents snooty name.
	snootyName = nook.Languages{
		language.AmericanEnglish:      snootyAmericanEnglishName,
		language.CanadianFrench:       snootyCanadianFrenchName,
		language.Dutch:                snootyDutchName,
		language.French:               snootyFrenchName,
		language.German:               snootyGermanName,
		language.Italian:              snootyItalianName,
		language.Japanese:             snootyJapaneseName,
		language.Korean:               snootyKoreanName,
		language.LatinAmericanSpanish: snootyLatinAmericanSpanishName,
		language.Russian:              snootyRussianName,
		language.SimplifiedChinese:    snootySimplifiedChineseName,
		language.Spanish:              snootySpanishName,
		language.TraditionalChinese:   snootyTraditionalChineseName}
)

// snootyCharacter represents Snooty's character information.
var (
	// snootyCharacter represents snooty character.
	snootyCharacter = nook.Character{
		Animal:   animal.Anteater,
		Birthday: snootyBirthday,
		Code:     snootyCode,
		Key:      character.Snooty,
		Gender:   gender.Female,
		Name:     snootyName,
		Special:  false}
)

// Different phrases spoken by Snooty in various languages.
var (
	// snootyAmericanEnglishPhrase represents snooty american english phrase.
	snootyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sniffff"}

	// snootyCanadianFrenchPhrase represents snooty canadian french phrase.
	snootyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "snif snif"}

	// snootyDutchPhrase represents snooty dutch phrase.
	snootyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snuifff"}

	// snootyFrenchPhrase represents snooty french phrase.
	snootyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "snif snif"}

	// snootyGermanPhrase represents snooty german phrase.
	snootyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnuff"}

	// snootyItalianPhrase represents snooty italian phrase.
	snootyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sniffff"}

	// snootyJapanesePhrase represents snooty japanese phrase.
	snootyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "たもれ"}

	// snootyKoreanPhrase represents snooty korean phrase.
	snootyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "추릅"}

	// snootyLatinAmericanSpanishPhrase represents snooty latin american spanish phrase.
	snootyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "sniff"}

	// snootyRussianPhrase represents snooty russian phrase.
	snootyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "шмыг"}

	// snootySimplifiedChinesePhrase represents snooty simplified chinese phrase.
	snootySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "麻烦了"}

	// snootySpanishPhrase represents snooty spanish phrase.
	snootySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "sniff"}

	// snootyTraditionalChinesePhrase represents snooty traditional chinese phrase.
	snootyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "麻煩了"}
)

// snootyPhrase represents Snooty's phrases in different languages.
var (
	// snootyPhrase represents snooty phrase.
	snootyPhrase = nook.Languages{
		language.AmericanEnglish:      snootyAmericanEnglishPhrase,
		language.CanadianFrench:       snootyCanadianFrenchPhrase,
		language.Dutch:                snootyDutchPhrase,
		language.French:               snootyFrenchPhrase,
		language.German:               snootyGermanPhrase,
		language.Italian:              snootyItalianPhrase,
		language.Japanese:             snootyJapanesePhrase,
		language.Korean:               snootyKoreanPhrase,
		language.LatinAmericanSpanish: snootyLatinAmericanSpanishPhrase,
		language.Russian:              snootyRussianPhrase,
		language.SimplifiedChinese:    snootySimplifiedChinesePhrase,
		language.Spanish:              snootySpanishPhrase,
		language.TraditionalChinese:   snootyTraditionalChinesePhrase}
)

// Snooty represents the character Snooty with her complete information.
var (
	// Snooty represents snooty.
	Snooty = nook.Villager{
		Character:   snootyCharacter,
		Personality: personality.Snooty,
		Phrase:      snootyPhrase}
)

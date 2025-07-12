package monkey

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
	// eliseBirthday represents elise birthday.
	eliseBirthday = nook.Birthday{
		Day:   21,
		Month: time.March}
)

var (
	// eliseCode represents elise code.
	eliseCode = nook.Code{
		Value: "mnk05"}
)

var (
	// eliseAmericanEnglishName represents elise american english name.
	eliseAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Elise"}

	// eliseCanadianFrenchName represents elise canadian french name.
	eliseCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Élise"}

	// eliseDutchName represents elise dutch name.
	eliseDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Elise"}

	// eliseFrenchName represents elise french name.
	eliseFrenchName = nook.Name{
		Language: language.French,
		Value:    "Élise"}

	// eliseGermanName represents elise german name.
	eliseGermanName = nook.Name{
		Language: language.German,
		Value:    "Steffi"}

	// eliseItalianName represents elise italian name.
	eliseItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Elisa"}

	// eliseJapaneseName represents elise japanese name.
	eliseJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "モンこ"}

	// eliseKoreanName represents elise korean name.
	eliseKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "몽자"}

	// eliseLatinAmericanSpanishName represents elise latin american spanish name.
	eliseLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Mayra"}

	// eliseRussianName represents elise russian name.
	eliseRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Элиза"}

	// eliseSimplifiedChineseName represents elise simplified chinese name.
	eliseSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "孟琪"}

	// eliseSpanishName represents elise spanish name.
	eliseSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Mayra"}

	// eliseTraditionalChineseName represents elise traditional chinese name.
	eliseTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "孟琪"}
)

var (
	// eliseName represents elise name.
	eliseName = nook.Languages{
		language.AmericanEnglish:      eliseAmericanEnglishName,
		language.CanadianFrench:       eliseCanadianFrenchName,
		language.Dutch:                eliseDutchName,
		language.French:               eliseFrenchName,
		language.German:               eliseGermanName,
		language.Italian:              eliseItalianName,
		language.Japanese:             eliseJapaneseName,
		language.Korean:               eliseKoreanName,
		language.LatinAmericanSpanish: eliseLatinAmericanSpanishName,
		language.Russian:              eliseRussianName,
		language.SimplifiedChinese:    eliseSimplifiedChineseName,
		language.Spanish:              eliseSpanishName,
		language.TraditionalChinese:   eliseTraditionalChineseName}
)

var (
	// eliseCharacter represents elise character.
	eliseCharacter = nook.Character{
		Animal:   animal.Monkey,
		Birthday: eliseBirthday,
		Code:     eliseCode,
		Key:      character.Elise,
		Gender:   gender.Female,
		Name:     eliseName,
		Special:  false}
)

var (
	// eliseAmericanEnglishPhrase represents elise american english phrase.
	eliseAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "puh-lease"}

	// eliseCanadianFrenchPhrase represents elise canadian french phrase.
	eliseCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "steup"}

	// eliseDutchPhrase represents elise dutch phrase.
	eliseDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kom op zeg"}

	// eliseFrenchPhrase represents elise french phrase.
	eliseFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "steup"}

	// eliseGermanPhrase represents elise german phrase.
	eliseGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "strebstreb"}

	// eliseItalianPhrase represents elise italian phrase.
	eliseItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "per favore"}

	// eliseJapanesePhrase represents elise japanese phrase.
	eliseJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だモン"}

	// eliseKoreanPhrase represents elise korean phrase.
	eliseKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "몽몽"}

	// eliseLatinAmericanSpanishPhrase represents elise latin american spanish phrase.
	eliseLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "uh-lalá"}

	// eliseRussianPhrase represents elise russian phrase.
	eliseRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "само собой"}

	// eliseSimplifiedChinesePhrase represents elise simplified chinese phrase.
	eliseSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "孟孟"}

	// eliseSpanishPhrase represents elise spanish phrase.
	eliseSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ricura"}

	// eliseTraditionalChinesePhrase represents elise traditional chinese phrase.
	eliseTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "孟孟"}
)

var (
	// elisePhrase represents elise phrase.
	elisePhrase = nook.Languages{
		language.AmericanEnglish:      eliseAmericanEnglishPhrase,
		language.CanadianFrench:       eliseCanadianFrenchPhrase,
		language.Dutch:                eliseDutchPhrase,
		language.French:               eliseFrenchPhrase,
		language.German:               eliseGermanPhrase,
		language.Italian:              eliseItalianPhrase,
		language.Japanese:             eliseJapanesePhrase,
		language.Korean:               eliseKoreanPhrase,
		language.LatinAmericanSpanish: eliseLatinAmericanSpanishPhrase,
		language.Russian:              eliseRussianPhrase,
		language.SimplifiedChinese:    eliseSimplifiedChinesePhrase,
		language.Spanish:              eliseSpanishPhrase,
		language.TraditionalChinese:   eliseTraditionalChinesePhrase}
)

var (
	// Elise represents elise.
	Elise = nook.Villager{
		Character:   eliseCharacter,
		Personality: personality.Snooty,
		Phrase:      elisePhrase}
)

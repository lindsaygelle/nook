package horse

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
	// elmerBirthday represents elmer birthday.
	elmerBirthday = nook.Birthday{
		Day:   5,
		Month: time.October}
)

var (
	// elmerCode represents elmer code.
	elmerCode = nook.Code{
		Value: "hrs03"}
)

var (
	// elmerAmericanEnglishName represents elmer american english name.
	elmerAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Elmer"}

	// elmerCanadianFrenchName represents elmer canadian french name.
	elmerCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Martin"}

	// elmerDutchName represents elmer dutch name.
	elmerDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Elmer"}

	// elmerFrenchName represents elmer french name.
	elmerFrenchName = nook.Name{
		Language: language.French,
		Value:    "Martin"}

	// elmerGermanName represents elmer german name.
	elmerGermanName = nook.Name{
		Language: language.German,
		Value:    "Emil"}

	// elmerItalianName represents elmer italian name.
	elmerItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Oreste"}

	// elmerJapaneseName represents elmer japanese name.
	elmerJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "サブレ"}

	// elmerKoreanName represents elmer korean name.
	elmerKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "샤브렌"}

	// elmerLatinAmericanSpanishName represents elmer latin american spanish name.
	elmerLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Jacinto"}

	// elmerRussianName represents elmer russian name.
	elmerRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Элмер"}

	// elmerSimplifiedChineseName represents elmer simplified chinese name.
	elmerSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "酥饼"}

	// elmerSpanishName represents elmer spanish name.
	elmerSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Jacinto"}

	// elmerTraditionalChineseName represents elmer traditional chinese name.
	elmerTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "酥餅"}
)

var (
	// elmerName represents elmer name.
	elmerName = nook.Languages{
		language.AmericanEnglish:      elmerAmericanEnglishName,
		language.CanadianFrench:       elmerCanadianFrenchName,
		language.Dutch:                elmerDutchName,
		language.French:               elmerFrenchName,
		language.German:               elmerGermanName,
		language.Italian:              elmerItalianName,
		language.Japanese:             elmerJapaneseName,
		language.Korean:               elmerKoreanName,
		language.LatinAmericanSpanish: elmerLatinAmericanSpanishName,
		language.Russian:              elmerRussianName,
		language.SimplifiedChinese:    elmerSimplifiedChineseName,
		language.Spanish:              elmerSpanishName,
		language.TraditionalChinese:   elmerTraditionalChineseName}
)

var (
	// elmerCharacter represents elmer character.
	elmerCharacter = nook.Character{
		Animal:   animal.Horse,
		Birthday: elmerBirthday,
		Code:     elmerCode,
		Key:      character.Elmer,
		Gender:   gender.Male,
		Name:     elmerName,
		Special:  false}
)

var (
	// elmerAmericanEnglishPhrase represents elmer american english phrase.
	elmerAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "tenderfoot"}

	// elmerCanadianFrenchPhrase represents elmer canadian french phrase.
	elmerCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mouuuuiii"}

	// elmerDutchPhrase represents elmer dutch phrase.
	elmerDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "jonkie"}

	// elmerFrenchPhrase represents elmer french phrase.
	elmerFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mouuuuiii"}

	// elmerGermanPhrase represents elmer german phrase.
	elmerGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "anfänger"}

	// elmerItalianPhrase represents elmer italian phrase.
	elmerItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "pinopino"}

	// elmerJapanesePhrase represents elmer japanese phrase.
	elmerJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だヒヒン"}

	// elmerKoreanPhrase represents elmer korean phrase.
	elmerKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "히히힝"}

	// elmerLatinAmericanSpanishPhrase represents elmer latin american spanish phrase.
	elmerLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "hop-uah"}

	// elmerRussianPhrase represents elmer russian phrase.
	elmerRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "жеребенок"}

	// elmerSimplifiedChinesePhrase represents elmer simplified chinese phrase.
	elmerSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "酥酥"}

	// elmerSpanishPhrase represents elmer spanish phrase.
	elmerSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pezuñitas"}

	// elmerTraditionalChinesePhrase represents elmer traditional chinese phrase.
	elmerTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "酥酥"}
)

var (
	// elmerPhrase represents elmer phrase.
	elmerPhrase = nook.Languages{
		language.AmericanEnglish:      elmerAmericanEnglishPhrase,
		language.CanadianFrench:       elmerCanadianFrenchPhrase,
		language.Dutch:                elmerDutchPhrase,
		language.French:               elmerFrenchPhrase,
		language.German:               elmerGermanPhrase,
		language.Italian:              elmerItalianPhrase,
		language.Japanese:             elmerJapanesePhrase,
		language.Korean:               elmerKoreanPhrase,
		language.LatinAmericanSpanish: elmerLatinAmericanSpanishPhrase,
		language.Russian:              elmerRussianPhrase,
		language.SimplifiedChinese:    elmerSimplifiedChinesePhrase,
		language.Spanish:              elmerSpanishPhrase,
		language.TraditionalChinese:   elmerTraditionalChinesePhrase}
)

var (
	// Elmer represents elmer.
	Elmer = nook.Villager{
		Character:   elmerCharacter,
		Personality: personality.Lazy,
		Phrase:      elmerPhrase}
)

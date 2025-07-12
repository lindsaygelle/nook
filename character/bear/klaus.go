package bear

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

// klausBirthday represents Klaus's birthday.
var (
	// klausBirthday represents klaus birthday.
	klausBirthday = nook.Birthday{
		Day:   31,
		Month: time.March}
)

// klausCode represents Klaus's unique code.
var (
	// klausCode represents klaus code.
	klausCode = nook.Code{
		Value: "bea14"}
)

// Different names for Klaus in various languages.
var (
	// klausAmericanEnglishName represents klaus american english name.
	klausAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Klaus"}

	// klausCanadianFrenchName represents klaus canadian french name.
	klausCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Klaus"}

	// klausDutchName represents klaus dutch name.
	klausDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Klaus"}

	// klausFrenchName represents klaus french name.
	klausFrenchName = nook.Name{
		Language: language.French,
		Value:    "Klaus"}

	// klausGermanName represents klaus german name.
	klausGermanName = nook.Name{
		Language: language.German,
		Value:    "Grischa"}

	// klausItalianName represents klaus italian name.
	klausItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Maciste"}

	// klausJapaneseName represents klaus japanese name.
	klausJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クマロス"}

	// klausKoreanName represents klaus korean name.
	klausKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "곰도로스"}

	// klausLatinAmericanSpanishName represents klaus latin american spanish name.
	klausLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Gruñerto"}

	// klausRussianName represents klaus russian name.
	klausRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Клаус"}

	// klausSimplifiedChineseName represents klaus simplified chinese name.
	klausSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "熊战士"}

	// klausSpanishName represents klaus spanish name.
	klausSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Gruñerto"}

	// klausTraditionalChineseName represents klaus traditional chinese name.
	klausTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "熊战士"}
)

// klausName represents Klaus's name in different languages.
var (
	// klausName represents klaus name.
	klausName = nook.Languages{
		language.AmericanEnglish:      klausAmericanEnglishName,
		language.CanadianFrench:       klausCanadianFrenchName,
		language.Dutch:                klausDutchName,
		language.French:               klausFrenchName,
		language.German:               klausGermanName,
		language.Italian:              klausItalianName,
		language.Japanese:             klausJapaneseName,
		language.Korean:               klausKoreanName,
		language.LatinAmericanSpanish: klausLatinAmericanSpanishName,
		language.Russian:              klausRussianName,
		language.SimplifiedChinese:    klausSimplifiedChineseName,
		language.Spanish:              klausSpanishName,
		language.TraditionalChinese:   klausTraditionalChineseName}
)

// klausCharacter represents Klaus's character information.
var (
	// klausCharacter represents klaus character.
	klausCharacter = nook.Character{
		Animal:   animal.Bear,
		Birthday: klausBirthday,
		Code:     klausCode,
		Key:      character.Klaus,
		Gender:   gender.Male,
		Name:     klausName,
		Special:  false}
)

// Different phrases spoken by Klaus in various languages.
var (
	// klausAmericanEnglishPhrase represents klaus american english phrase.
	klausAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "strudel"}

	// klausCanadianFrenchPhrase represents klaus canadian french phrase.
	klausCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "achtung"}

	// klausDutchPhrase represents klaus dutch phrase.
	klausDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "strudel"}

	// klausFrenchPhrase represents klaus french phrase.
	klausFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "achtung"}

	// klausGermanPhrase represents klaus german phrase.
	klausGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "gruaaa"}

	// klausItalianPhrase represents klaus italian phrase.
	klausItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bailar"}

	// klausJapanesePhrase represents klaus japanese phrase.
	klausJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "オパー"}

	// klausKoreanPhrase represents klaus korean phrase.
	klausKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "라저"}

	// klausLatinAmericanSpanishPhrase represents klaus latin american spanish phrase.
	klausLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "achurf"}

	// klausRussianPhrase represents klaus russian phrase.
	klausRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "штрудель"}

	// klausSimplifiedChinesePhrase represents klaus simplified chinese phrase.
	klausSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Over"}

	// klausSpanishPhrase represents klaus spanish phrase.
	klausSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ricitos"}

	// klausTraditionalChinesePhrase represents klaus traditional chinese phrase.
	klausTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Over"}
)

// klausPhrase represents Klaus's phrases in different languages.
var (
	// klausPhrase represents klaus phrase.
	klausPhrase = nook.Languages{
		language.AmericanEnglish:      klausAmericanEnglishPhrase,
		language.CanadianFrench:       klausCanadianFrenchPhrase,
		language.Dutch:                klausDutchPhrase,
		language.French:               klausFrenchPhrase,
		language.German:               klausGermanPhrase,
		language.Italian:              klausItalianPhrase,
		language.Japanese:             klausJapanesePhrase,
		language.Korean:               klausKoreanPhrase,
		language.LatinAmericanSpanish: klausLatinAmericanSpanishPhrase,
		language.Russian:              klausRussianPhrase,
		language.SimplifiedChinese:    klausSimplifiedChinesePhrase,
		language.Spanish:              klausSpanishPhrase,
		language.TraditionalChinese:   klausTraditionalChinesePhrase}
)

// Klaus represents the character Klaus with his complete information.
var (
	// Klaus represents klaus.
	Klaus = nook.Villager{
		Character:   klausCharacter,
		Personality: personality.Smug,
		Phrase:      klausPhrase}
)

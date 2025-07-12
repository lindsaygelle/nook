package duck

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
	// drakeBirthday represents drake birthday.
	drakeBirthday = nook.Birthday{
		Day:   25,
		Month: time.June}
)

var (
	// drakeCode represents drake code.
	drakeCode = nook.Code{
		Value: "duk09"}
)

var (
	// drakeAmericanEnglishName represents drake american english name.
	drakeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Drake"}

	// drakeCanadianFrenchName represents drake canadian french name.
	drakeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Colvert"}

	// drakeDutchName represents drake dutch name.
	drakeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Drake"}

	// drakeFrenchName represents drake french name.
	drakeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Colvert"}

	// drakeGermanName represents drake german name.
	drakeGermanName = nook.Name{
		Language: language.German,
		Value:    "Gustav"}

	// drakeItalianName represents drake italian name.
	drakeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Drago"}

	// drakeJapaneseName represents drake japanese name.
	drakeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フォアグラ"}

	// drakeKoreanName represents drake korean name.
	drakeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "푸아그라"}

	// drakeLatinAmericanSpanishName represents drake latin american spanish name.
	drakeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Draco"}

	// drakeRussianName represents drake russian name.
	drakeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дрейк"}

	// drakeSimplifiedChineseName represents drake simplified chinese name.
	drakeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "肥肝"}

	// drakeSpanishName represents drake spanish name.
	drakeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Draco"}

	// drakeTraditionalChineseName represents drake traditional chinese name.
	drakeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "肥肝"}
)

var (
	// drakeName represents drake name.
	drakeName = nook.Languages{
		language.AmericanEnglish:      drakeAmericanEnglishName,
		language.CanadianFrench:       drakeCanadianFrenchName,
		language.Dutch:                drakeDutchName,
		language.French:               drakeFrenchName,
		language.German:               drakeGermanName,
		language.Italian:              drakeItalianName,
		language.Japanese:             drakeJapaneseName,
		language.Korean:               drakeKoreanName,
		language.LatinAmericanSpanish: drakeLatinAmericanSpanishName,
		language.Russian:              drakeRussianName,
		language.SimplifiedChinese:    drakeSimplifiedChineseName,
		language.Spanish:              drakeSpanishName,
		language.TraditionalChinese:   drakeTraditionalChineseName}
)

var (
	// drakeCharacter represents drake character.
	drakeCharacter = nook.Character{
		Animal:   animal.Duck,
		Birthday: drakeBirthday,
		Code:     drakeCode,
		Key:      character.Drake,
		Gender:   gender.Male,
		Name:     drakeName,
		Special:  false}
)

var (
	// drakeAmericanEnglishPhrase represents drake american english phrase.
	drakeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "quacko"}

	// drakeCanadianFrenchPhrase represents drake canadian french phrase.
	drakeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ben couac"}

	// drakeDutchPhrase represents drake dutch phrase.
	drakeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snabbel"}

	// drakeFrenchPhrase represents drake french phrase.
	drakeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "canard"}

	// drakeGermanPhrase represents drake german phrase.
	drakeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnabbl"}

	// drakeItalianPhrase represents drake italian phrase.
	drakeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "quaqua"}

	// drakeJapanesePhrase represents drake japanese phrase.
	drakeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "かもね"}

	// drakeKoreanPhrase represents drake korean phrase.
	drakeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "어쩌면"}

	// drakeLatinAmericanSpanishPhrase represents drake latin american spanish phrase.
	drakeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cui-cuac"}

	// drakeRussianPhrase represents drake russian phrase.
	drakeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кра-а"}

	// drakeSimplifiedChinesePhrase represents drake simplified chinese phrase.
	drakeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "或许鸭"}

	// drakeSpanishPhrase represents drake spanish phrase.
	drakeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "polluelo"}

	// drakeTraditionalChinesePhrase represents drake traditional chinese phrase.
	drakeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "或許鴨"}
)

var (
	// drakePhrase represents drake phrase.
	drakePhrase = nook.Languages{
		language.AmericanEnglish:      drakeAmericanEnglishPhrase,
		language.CanadianFrench:       drakeCanadianFrenchPhrase,
		language.Dutch:                drakeDutchPhrase,
		language.French:               drakeFrenchPhrase,
		language.German:               drakeGermanPhrase,
		language.Italian:              drakeItalianPhrase,
		language.Japanese:             drakeJapanesePhrase,
		language.Korean:               drakeKoreanPhrase,
		language.LatinAmericanSpanish: drakeLatinAmericanSpanishPhrase,
		language.Russian:              drakeRussianPhrase,
		language.SimplifiedChinese:    drakeSimplifiedChinesePhrase,
		language.Spanish:              drakeSpanishPhrase,
		language.TraditionalChinese:   drakeTraditionalChinesePhrase}
)

var (
	// Drake represents drake.
	Drake = nook.Villager{
		Character:   drakeCharacter,
		Personality: personality.Lazy,
		Phrase:      drakePhrase}
)

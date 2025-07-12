package hamster

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
	// flurryBirthday represents flurry birthday.
	flurryBirthday = nook.Birthday{
		Day:   30,
		Month: time.January}
)

var (
	// flurryCode represents flurry code.
	flurryCode = nook.Code{
		Value: "ham06"}
)

var (
	// flurryAmericanEnglishName represents flurry american english name.
	flurryAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Flurry"}

	// flurryCanadianFrenchName represents flurry canadian french name.
	flurryCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Emma"}

	// flurryDutchName represents flurry dutch name.
	flurryDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Flurry"}

	// flurryFrenchName represents flurry french name.
	flurryFrenchName = nook.Name{
		Language: language.French,
		Value:    "Emma"}

	// flurryGermanName represents flurry german name.
	flurryGermanName = nook.Name{
		Language: language.German,
		Value:    "Emilie"}

	// flurryItalianName represents flurry italian name.
	flurryItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Meringa"}

	// flurryJapaneseName represents flurry japanese name.
	flurryJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ゆきみ"}

	// flurryKoreanName represents flurry korean name.
	flurryKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "뽀야미"}

	// flurryLatinAmericanSpanishName represents flurry latin american spanish name.
	flurryLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lluvia"}

	// flurryRussianName represents flurry russian name.
	flurryRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Фларри"}

	// flurrySimplifiedChineseName represents flurry simplified chinese name.
	flurrySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "雪美"}

	// flurrySpanishName represents flurry spanish name.
	flurrySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lluvia"}

	// flurryTraditionalChineseName represents flurry traditional chinese name.
	flurryTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "雪美"}
)

var (
	// flurryName represents flurry name.
	flurryName = nook.Languages{
		language.AmericanEnglish:      flurryAmericanEnglishName,
		language.CanadianFrench:       flurryCanadianFrenchName,
		language.Dutch:                flurryDutchName,
		language.French:               flurryFrenchName,
		language.German:               flurryGermanName,
		language.Italian:              flurryItalianName,
		language.Japanese:             flurryJapaneseName,
		language.Korean:               flurryKoreanName,
		language.LatinAmericanSpanish: flurryLatinAmericanSpanishName,
		language.Russian:              flurryRussianName,
		language.SimplifiedChinese:    flurrySimplifiedChineseName,
		language.Spanish:              flurrySpanishName,
		language.TraditionalChinese:   flurryTraditionalChineseName}
)

var (
	// flurryCharacter represents flurry character.
	flurryCharacter = nook.Character{
		Animal:   animal.Hamster,
		Birthday: flurryBirthday,
		Code:     flurryCode,
		Key:      character.Flurry,
		Gender:   gender.Female,
		Name:     flurryName,
		Special:  false}
)

var (
	// flurryAmericanEnglishPhrase represents flurry american english phrase.
	flurryAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "powderpuff"}

	// flurryCanadianFrenchPhrase represents flurry canadian french phrase.
	flurryCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "abajoujou"}

	// flurryDutchPhrase represents flurry dutch phrase.
	flurryDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "donsbal"}

	// flurryFrenchPhrase represents flurry french phrase.
	flurryFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "abajoujou"}

	// flurryGermanPhrase represents flurry german phrase.
	flurryGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "millimilli"}

	// flurryItalianPhrase represents flurry italian phrase.
	flurryItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "fondue"}

	// flurryJapanesePhrase represents flurry japanese phrase.
	flurryJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なのです"}

	// flurryKoreanPhrase represents flurry korean phrase.
	flurryKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "뽀드득"}

	// flurryLatinAmericanSpanishPhrase represents flurry latin american spanish phrase.
	flurryLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "nomnom"}

	// flurryRussianPhrase represents flurry russian phrase.
	flurryRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "пушистик"}

	// flurrySimplifiedChinesePhrase represents flurry simplified chinese phrase.
	flurrySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "对啊"}

	// flurrySpanishPhrase represents flurry spanish phrase.
	flurrySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "nomnom"}

	// flurryTraditionalChinesePhrase represents flurry traditional chinese phrase.
	flurryTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "對啊"}
)

var (
	// flurryPhrase represents flurry phrase.
	flurryPhrase = nook.Languages{
		language.AmericanEnglish:      flurryAmericanEnglishPhrase,
		language.CanadianFrench:       flurryCanadianFrenchPhrase,
		language.Dutch:                flurryDutchPhrase,
		language.French:               flurryFrenchPhrase,
		language.German:               flurryGermanPhrase,
		language.Italian:              flurryItalianPhrase,
		language.Japanese:             flurryJapanesePhrase,
		language.Korean:               flurryKoreanPhrase,
		language.LatinAmericanSpanish: flurryLatinAmericanSpanishPhrase,
		language.Russian:              flurryRussianPhrase,
		language.SimplifiedChinese:    flurrySimplifiedChinesePhrase,
		language.Spanish:              flurrySpanishPhrase,
		language.TraditionalChinese:   flurryTraditionalChinesePhrase}
)

var (
	// Flurry represents flurry.
	Flurry = nook.Villager{
		Character:   flurryCharacter,
		Personality: personality.Normal,
		Phrase:      flurryPhrase}
)

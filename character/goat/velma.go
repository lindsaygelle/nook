package goat

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
	// velmaBirthday represents velma birthday.
	velmaBirthday = nook.Birthday{
		Day:   14,
		Month: time.January}
)

var (
	// velmaCode represents velma code.
	velmaCode = nook.Code{
		Value: "goa06"}
)

var (
	// velmaAmericanEnglishName represents velma american english name.
	velmaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Velma"}

	// velmaCanadianFrenchName represents velma canadian french name.
	velmaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Véra"}

	// velmaDutchName represents velma dutch name.
	velmaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Velma"}

	// velmaFrenchName represents velma french name.
	velmaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Véra"}

	// velmaGermanName represents velma german name.
	velmaGermanName = nook.Name{
		Language: language.German,
		Value:    "Wilma"}

	// velmaItalianName represents velma italian name.
	velmaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Belarda"}

	// velmaJapaneseName represents velma japanese name.
	velmaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ピティエ"}

	// velmaKoreanName represents velma korean name.
	velmaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "벨마"}

	// velmaLatinAmericanSpanishName represents velma latin american spanish name.
	velmaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Erika"}

	// velmaRussianName represents velma russian name.
	velmaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Велма"}

	// velmaSimplifiedChineseName represents velma simplified chinese name.
	velmaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "班长"}

	// velmaSpanishName represents velma spanish name.
	velmaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Erika"}

	// velmaTraditionalChineseName represents velma traditional chinese name.
	velmaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "班長"}
)

var (
	// velmaName represents velma name.
	velmaName = nook.Languages{
		language.AmericanEnglish:      velmaAmericanEnglishName,
		language.CanadianFrench:       velmaCanadianFrenchName,
		language.Dutch:                velmaDutchName,
		language.French:               velmaFrenchName,
		language.German:               velmaGermanName,
		language.Italian:              velmaItalianName,
		language.Japanese:             velmaJapaneseName,
		language.Korean:               velmaKoreanName,
		language.LatinAmericanSpanish: velmaLatinAmericanSpanishName,
		language.Russian:              velmaRussianName,
		language.SimplifiedChinese:    velmaSimplifiedChineseName,
		language.Spanish:              velmaSpanishName,
		language.TraditionalChinese:   velmaTraditionalChineseName}
)

var (
	// velmaCharacter represents velma character.
	velmaCharacter = nook.Character{
		Animal:   animal.Goat,
		Birthday: velmaBirthday,
		Code:     velmaCode,
		Key:      character.Velma,
		Gender:   gender.Female,
		Name:     velmaName,
		Special:  false}
)

var (
	// velmaAmericanEnglishPhrase represents velma american english phrase.
	velmaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "blih"}

	// velmaCanadianFrenchPhrase represents velma canadian french phrase.
	velmaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tu vois"}

	// velmaDutchPhrase represents velma dutch phrase.
	velmaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "meh"}

	// velmaFrenchPhrase represents velma french phrase.
	velmaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tu vois"}

	// velmaGermanPhrase represents velma german phrase.
	velmaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mecker"}

	// velmaItalianPhrase represents velma italian phrase.
	velmaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "blip"}

	// velmaJapanesePhrase represents velma japanese phrase.
	velmaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ザーマス"}

	// velmaKoreanPhrase represents velma korean phrase.
	velmaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "똑바로해"}

	// velmaLatinAmericanSpanishPhrase represents velma latin american spanish phrase.
	velmaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "verááás-tú"}

	// velmaRussianPhrase represents velma russian phrase.
	velmaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ме-е-е"}

	// velmaSimplifiedChinesePhrase represents velma simplified chinese phrase.
	velmaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "起立"}

	// velmaSpanishPhrase represents velma spanish phrase.
	velmaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "veráaas-tú"}

	// velmaTraditionalChinesePhrase represents velma traditional chinese phrase.
	velmaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "起立"}
)

var (
	// velmaPhrase represents velma phrase.
	velmaPhrase = nook.Languages{
		language.AmericanEnglish:      velmaAmericanEnglishPhrase,
		language.CanadianFrench:       velmaCanadianFrenchPhrase,
		language.Dutch:                velmaDutchPhrase,
		language.French:               velmaFrenchPhrase,
		language.German:               velmaGermanPhrase,
		language.Italian:              velmaItalianPhrase,
		language.Japanese:             velmaJapanesePhrase,
		language.Korean:               velmaKoreanPhrase,
		language.LatinAmericanSpanish: velmaLatinAmericanSpanishPhrase,
		language.Russian:              velmaRussianPhrase,
		language.SimplifiedChinese:    velmaSimplifiedChinesePhrase,
		language.Spanish:              velmaSpanishPhrase,
		language.TraditionalChinese:   velmaTraditionalChinesePhrase}
)

var (
	// Velma represents velma.
	Velma = nook.Villager{
		Character:   velmaCharacter,
		Personality: personality.Snooty,
		Phrase:      velmaPhrase}
)

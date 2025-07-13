package koala

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
	// ozzieBirthday represents ozzie birthday.
	ozzieBirthday = nook.Birthday{
		Day:   7,
		Month: time.May}
)

var (
	// ozzieCode represents ozzie code.
	ozzieCode = nook.Code{
		Value: "kal05"}
)

var (
	// ozzieAmericanEnglishName represents ozzie american english name.
	ozzieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Ozzie"}

	// ozzieCanadianFrenchName represents ozzie canadian french name.
	ozzieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Koko"}

	// ozzieDutchName represents ozzie dutch name.
	ozzieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Ozzie"}

	// ozzieFrenchName represents ozzie french name.
	ozzieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Koko"}

	// ozzieGermanName represents ozzie german name.
	ozzieGermanName = nook.Name{
		Language: language.German,
		Value:    "Oskar"}

	// ozzieItalianName represents ozzie italian name.
	ozzieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ozzy"}

	// ozzieJapaneseName represents ozzie japanese name.
	ozzieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ドングリ"}

	// ozzieKoreanName represents ozzie korean name.
	ozzieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "동동이"}

	// ozzieLatinAmericanSpanishName represents ozzie latin american spanish name.
	ozzieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Koloa"}

	// ozzieRussianName represents ozzie russian name.
	ozzieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Оззи"}

	// ozzieSimplifiedChineseName represents ozzie simplified chinese name.
	ozzieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿栗"}

	// ozzieSpanishName represents ozzie spanish name.
	ozzieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Koloa"}

	// ozzieTraditionalChineseName represents ozzie traditional chinese name.
	ozzieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿栗"}
)

var (
	// ozzieName represents ozzie name.
	ozzieName = nook.Languages{
		language.AmericanEnglish:      ozzieAmericanEnglishName,
		language.CanadianFrench:       ozzieCanadianFrenchName,
		language.Dutch:                ozzieDutchName,
		language.French:               ozzieFrenchName,
		language.German:               ozzieGermanName,
		language.Italian:              ozzieItalianName,
		language.Japanese:             ozzieJapaneseName,
		language.Korean:               ozzieKoreanName,
		language.LatinAmericanSpanish: ozzieLatinAmericanSpanishName,
		language.Russian:              ozzieRussianName,
		language.SimplifiedChinese:    ozzieSimplifiedChineseName,
		language.Spanish:              ozzieSpanishName,
		language.TraditionalChinese:   ozzieTraditionalChineseName}
)

var (
	// ozzieCharacter represents ozzie character.
	ozzieCharacter = nook.Character{
		Animal:   animal.Koala,
		Birthday: ozzieBirthday,
		Code:     ozzieCode,
		Key:      character.Ozzie,
		Gender:   gender.Male,
		Name:     ozzieName,
		Special:  false}
)

var (
	// ozzieAmericanEnglishPhrase represents ozzie american english phrase.
	ozzieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ol' bear"}

	// ozzieCanadianFrenchPhrase represents ozzie canadian french phrase.
	ozzieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "koa là là"}

	// ozzieDutchPhrase represents ozzie dutch phrase.
	ozzieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "ouwe beer"}

	// ozzieFrenchPhrase represents ozzie french phrase.
	ozzieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "koa là là"}

	// ozzieGermanPhrase represents ozzie german phrase.
	ozzieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "altes Haus"}

	// ozzieItalianPhrase represents ozzie italian phrase.
	ozzieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ulabadula"}

	// ozzieJapanesePhrase represents ozzie japanese phrase.
	ozzieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ククッ"}

	// ozzieKoreanPhrase represents ozzie korean phrase.
	ozzieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "큐큐"}

	// ozzieLatinAmericanSpanishPhrase represents ozzie latin american spanish phrase.
	ozzieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "koahhh"}

	// ozzieRussianPhrase represents ozzie russian phrase.
	ozzieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "медведище"}

	// ozzieSimplifiedChinesePhrase represents ozzie simplified chinese phrase.
	ozzieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "颗颗"}

	// ozzieSpanishPhrase represents ozzie spanish phrase.
	ozzieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "osete"}

	// ozzieTraditionalChinesePhrase represents ozzie traditional chinese phrase.
	ozzieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "顆顆"}
)

var (
	// ozziePhrase represents ozzie phrase.
	ozziePhrase = nook.Languages{
		language.AmericanEnglish:      ozzieAmericanEnglishPhrase,
		language.CanadianFrench:       ozzieCanadianFrenchPhrase,
		language.Dutch:                ozzieDutchPhrase,
		language.French:               ozzieFrenchPhrase,
		language.German:               ozzieGermanPhrase,
		language.Italian:              ozzieItalianPhrase,
		language.Japanese:             ozzieJapanesePhrase,
		language.Korean:               ozzieKoreanPhrase,
		language.LatinAmericanSpanish: ozzieLatinAmericanSpanishPhrase,
		language.Russian:              ozzieRussianPhrase,
		language.SimplifiedChinese:    ozzieSimplifiedChinesePhrase,
		language.Spanish:              ozzieSpanishPhrase,
		language.TraditionalChinese:   ozzieTraditionalChinesePhrase}
)

var (
	// Ozzie represents ozzie.
	Ozzie = nook.Villager{
		Character:   ozzieCharacter,
		Personality: personality.Lazy,
		Phrase:      ozziePhrase}
)

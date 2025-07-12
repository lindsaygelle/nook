package bearcub

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
	// pokoBirthday represents poko birthday.
	pokoBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// pokoCode represents poko code.
	pokoCode = nook.Code{
		Value: ""}
)

var (
	// pokoAmericanEnglishName represents poko american english name.
	pokoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Poko"}

	// pokoCanadianFrenchName represents poko canadian french name.
	pokoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// pokoDutchName represents poko dutch name.
	pokoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// pokoFrenchName represents poko french name.
	pokoFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	// pokoGermanName represents poko german name.
	pokoGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	// pokoItalianName represents poko italian name.
	pokoItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// pokoJapaneseName represents poko japanese name.
	pokoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ポコ"}

	// pokoKoreanName represents poko korean name.
	pokoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// pokoLatinAmericanSpanishName represents poko latin american spanish name.
	pokoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// pokoRussianName represents poko russian name.
	pokoRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// pokoSimplifiedChineseName represents poko simplified chinese name.
	pokoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// pokoSpanishName represents poko spanish name.
	pokoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// pokoTraditionalChineseName represents poko traditional chinese name.
	pokoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// pokoName represents poko name.
	pokoName = nook.Languages{
		language.AmericanEnglish:      pokoAmericanEnglishName,
		language.CanadianFrench:       pokoCanadianFrenchName,
		language.Dutch:                pokoDutchName,
		language.French:               pokoFrenchName,
		language.German:               pokoGermanName,
		language.Italian:              pokoItalianName,
		language.Japanese:             pokoJapaneseName,
		language.Korean:               pokoKoreanName,
		language.LatinAmericanSpanish: pokoLatinAmericanSpanishName,
		language.Russian:              pokoRussianName,
		language.SimplifiedChinese:    pokoSimplifiedChineseName,
		language.Spanish:              pokoSpanishName,
		language.TraditionalChinese:   pokoTraditionalChineseName}
)

var (
	// pokoCharacter represents poko character.
	pokoCharacter = nook.Character{
		Animal:   animal.Bearcub,
		Birthday: pokoBirthday,
		Code:     pokoCode,
		Key:      character.Poko,
		Gender:   gender.Male,
		Name:     pokoName,
		Special:  false}
)

var (
	// pokoAmericanEnglishPhrase represents poko american english phrase.
	pokoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "へへ"}

	// pokoCanadianFrenchPhrase represents poko canadian french phrase.
	pokoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// pokoDutchPhrase represents poko dutch phrase.
	pokoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// pokoFrenchPhrase represents poko french phrase.
	pokoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	// pokoGermanPhrase represents poko german phrase.
	pokoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	// pokoItalianPhrase represents poko italian phrase.
	pokoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// pokoJapanesePhrase represents poko japanese phrase.
	pokoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "へへ"}

	// pokoKoreanPhrase represents poko korean phrase.
	pokoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// pokoLatinAmericanSpanishPhrase represents poko latin american spanish phrase.
	pokoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// pokoRussianPhrase represents poko russian phrase.
	pokoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// pokoSimplifiedChinesePhrase represents poko simplified chinese phrase.
	pokoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// pokoSpanishPhrase represents poko spanish phrase.
	pokoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// pokoTraditionalChinesePhrase represents poko traditional chinese phrase.
	pokoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// pokoPhrase represents poko phrase.
	pokoPhrase = nook.Languages{
		language.AmericanEnglish:      pokoAmericanEnglishPhrase,
		language.CanadianFrench:       pokoCanadianFrenchPhrase,
		language.Dutch:                pokoDutchPhrase,
		language.French:               pokoFrenchPhrase,
		language.German:               pokoGermanPhrase,
		language.Italian:              pokoItalianPhrase,
		language.Japanese:             pokoJapanesePhrase,
		language.Korean:               pokoKoreanPhrase,
		language.LatinAmericanSpanish: pokoLatinAmericanSpanishPhrase,
		language.Russian:              pokoRussianPhrase,
		language.SimplifiedChinese:    pokoSimplifiedChinesePhrase,
		language.Spanish:              pokoSpanishPhrase,
		language.TraditionalChinese:   pokoTraditionalChinesePhrase}
)

var (
	// Poko represents poko.
	Poko = nook.Villager{
		Character:   pokoCharacter,
		Personality: personality.Jock,
		Phrase:      pokoPhrase}
)

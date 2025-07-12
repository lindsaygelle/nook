package frog

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
	// driftBirthday represents drift birthday.
	driftBirthday = nook.Birthday{
		Day:   9,
		Month: time.October}
)

var (
	// driftCode represents drift code.
	driftCode = nook.Code{
		Value: "flg04"}
)

var (
	// driftAmericanEnglishName represents drift american english name.
	driftAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Drift"}

	// driftCanadianFrenchName represents drift canadian french name.
	driftCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gordon"}

	// driftDutchName represents drift dutch name.
	driftDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Drift"}

	// driftFrenchName represents drift french name.
	driftFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gordon"}

	// driftGermanName represents drift german name.
	driftGermanName = nook.Name{
		Language: language.German,
		Value:    "Caspar"}

	// driftItalianName represents drift italian name.
	driftItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Prospero"}

	// driftJapaneseName represents drift japanese name.
	driftJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ドク"}

	// driftKoreanName represents drift korean name.
	driftKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "덕"}

	// driftLatinAmericanSpanishName represents drift latin american spanish name.
	driftLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Surfín"}

	// driftRussianName represents drift russian name.
	driftRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дрифт"}

	// driftSimplifiedChineseName represents drift simplified chinese name.
	driftSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "毒仔"}

	// driftSpanishName represents drift spanish name.
	driftSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Surfín"}

	// driftTraditionalChineseName represents drift traditional chinese name.
	driftTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "毒仔"}
)

var (
	// driftName represents drift name.
	driftName = nook.Languages{
		language.AmericanEnglish:      driftAmericanEnglishName,
		language.CanadianFrench:       driftCanadianFrenchName,
		language.Dutch:                driftDutchName,
		language.French:               driftFrenchName,
		language.German:               driftGermanName,
		language.Italian:              driftItalianName,
		language.Japanese:             driftJapaneseName,
		language.Korean:               driftKoreanName,
		language.LatinAmericanSpanish: driftLatinAmericanSpanishName,
		language.Russian:              driftRussianName,
		language.SimplifiedChinese:    driftSimplifiedChineseName,
		language.Spanish:              driftSpanishName,
		language.TraditionalChinese:   driftTraditionalChineseName}
)

var (
	// driftCharacter represents drift character.
	driftCharacter = nook.Character{
		Animal:   animal.Frog,
		Birthday: driftBirthday,
		Code:     driftCode,
		Key:      character.Drift,
		Gender:   gender.Male,
		Name:     driftName,
		Special:  false}
)

var (
	// driftAmericanEnglishPhrase represents drift american english phrase.
	driftAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "brah"}

	// driftCanadianFrenchPhrase represents drift canadian french phrase.
	driftCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "braaaa"}

	// driftDutchPhrase represents drift dutch phrase.
	driftDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "gozer"}

	// driftFrenchPhrase represents drift french phrase.
	driftFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "braaaa"}

	// driftGermanPhrase represents drift german phrase.
	driftGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "quak"}

	// driftItalianPhrase represents drift italian phrase.
	driftItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "girin"}

	// driftJapanesePhrase represents drift japanese phrase.
	driftJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ゲロゲロ"}

	// driftKoreanPhrase represents drift korean phrase.
	driftKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "개굴개굴"}

	// driftLatinAmericanSpanishPhrase represents drift latin american spanish phrase.
	driftLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "dribit"}

	// driftRussianPhrase represents drift russian phrase.
	driftRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бро"}

	// driftSimplifiedChinesePhrase represents drift simplified chinese phrase.
	driftSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "聒聒"}

	// driftSpanishPhrase represents drift spanish phrase.
	driftSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "colega"}

	// driftTraditionalChinesePhrase represents drift traditional chinese phrase.
	driftTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘓嘓"}
)

var (
	// driftPhrase represents drift phrase.
	driftPhrase = nook.Languages{
		language.AmericanEnglish:      driftAmericanEnglishPhrase,
		language.CanadianFrench:       driftCanadianFrenchPhrase,
		language.Dutch:                driftDutchPhrase,
		language.French:               driftFrenchPhrase,
		language.German:               driftGermanPhrase,
		language.Italian:              driftItalianPhrase,
		language.Japanese:             driftJapanesePhrase,
		language.Korean:               driftKoreanPhrase,
		language.LatinAmericanSpanish: driftLatinAmericanSpanishPhrase,
		language.Russian:              driftRussianPhrase,
		language.SimplifiedChinese:    driftSimplifiedChinesePhrase,
		language.Spanish:              driftSpanishPhrase,
		language.TraditionalChinese:   driftTraditionalChinesePhrase}
)

var (
	// Drift represents drift.
	Drift = nook.Villager{
		Character:   driftCharacter,
		Personality: personality.Jock,
		Phrase:      driftPhrase}
)

package ostrich

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
	// nindoriBirthday represents nindori birthday.
	nindoriBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// nindoriCode represents nindori code.
	nindoriCode = nook.Code{
		Value: ""}
)

var (
	// nindoriAmericanEnglishName represents nindori american english name.
	nindoriAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Nindori"}

	// nindoriCanadianFrenchName represents nindori canadian french name.
	nindoriCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// nindoriDutchName represents nindori dutch name.
	nindoriDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// nindoriFrenchName represents nindori french name.
	nindoriFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	// nindoriGermanName represents nindori german name.
	nindoriGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	// nindoriItalianName represents nindori italian name.
	nindoriItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// nindoriJapaneseName represents nindori japanese name.
	nindoriJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ニンドリ"}

	// nindoriKoreanName represents nindori korean name.
	nindoriKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// nindoriLatinAmericanSpanishName represents nindori latin american spanish name.
	nindoriLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// nindoriRussianName represents nindori russian name.
	nindoriRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// nindoriSimplifiedChineseName represents nindori simplified chinese name.
	nindoriSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// nindoriSpanishName represents nindori spanish name.
	nindoriSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// nindoriTraditionalChineseName represents nindori traditional chinese name.
	nindoriTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// nindoriName represents nindori name.
	nindoriName = nook.Languages{
		language.AmericanEnglish:      nindoriAmericanEnglishName,
		language.CanadianFrench:       nindoriCanadianFrenchName,
		language.Dutch:                nindoriDutchName,
		language.French:               nindoriFrenchName,
		language.German:               nindoriGermanName,
		language.Italian:              nindoriItalianName,
		language.Japanese:             nindoriJapaneseName,
		language.Korean:               nindoriKoreanName,
		language.LatinAmericanSpanish: nindoriLatinAmericanSpanishName,
		language.Russian:              nindoriRussianName,
		language.SimplifiedChinese:    nindoriSimplifiedChineseName,
		language.Spanish:              nindoriSpanishName,
		language.TraditionalChinese:   nindoriTraditionalChineseName}
)

var (
	// nindoriCharacter represents nindori character.
	nindoriCharacter = nook.Character{
		Animal:   animal.Ostrich,
		Birthday: nindoriBirthday,
		Code:     nindoriCode,
		Key:      character.Nindori,
		Gender:   gender.Male,
		Name:     nindoriName,
		Special:  false}
)

var (
	// nindoriAmericanEnglishPhrase represents nindori american english phrase.
	nindoriAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ドクソウ"}

	// nindoriCanadianFrenchPhrase represents nindori canadian french phrase.
	nindoriCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// nindoriDutchPhrase represents nindori dutch phrase.
	nindoriDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// nindoriFrenchPhrase represents nindori french phrase.
	nindoriFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	// nindoriGermanPhrase represents nindori german phrase.
	nindoriGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	// nindoriItalianPhrase represents nindori italian phrase.
	nindoriItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// nindoriJapanesePhrase represents nindori japanese phrase.
	nindoriJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ドクソウ"}

	// nindoriKoreanPhrase represents nindori korean phrase.
	nindoriKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// nindoriLatinAmericanSpanishPhrase represents nindori latin american spanish phrase.
	nindoriLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// nindoriRussianPhrase represents nindori russian phrase.
	nindoriRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// nindoriSimplifiedChinesePhrase represents nindori simplified chinese phrase.
	nindoriSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// nindoriSpanishPhrase represents nindori spanish phrase.
	nindoriSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// nindoriTraditionalChinesePhrase represents nindori traditional chinese phrase.
	nindoriTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// nindoriPhrase represents nindori phrase.
	nindoriPhrase = nook.Languages{
		language.AmericanEnglish:      nindoriAmericanEnglishPhrase,
		language.CanadianFrench:       nindoriCanadianFrenchPhrase,
		language.Dutch:                nindoriDutchPhrase,
		language.French:               nindoriFrenchPhrase,
		language.German:               nindoriGermanPhrase,
		language.Italian:              nindoriItalianPhrase,
		language.Japanese:             nindoriJapanesePhrase,
		language.Korean:               nindoriKoreanPhrase,
		language.LatinAmericanSpanish: nindoriLatinAmericanSpanishPhrase,
		language.Russian:              nindoriRussianPhrase,
		language.SimplifiedChinese:    nindoriSimplifiedChinesePhrase,
		language.Spanish:              nindoriSpanishPhrase,
		language.TraditionalChinese:   nindoriTraditionalChinesePhrase}
)

var (
	// Nindori represents nindori.
	Nindori = nook.Villager{
		Character:   nindoriCharacter,
		Personality: personality.Jock,
		Phrase:      nindoriPhrase}
)

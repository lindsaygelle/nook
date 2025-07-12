package chicken

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
	// bettyBirthday represents betty birthday.
	bettyBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// bettyCode represents betty code.
	bettyCode = nook.Code{
		Value: ""}
)

var (
	// bettyAmericanEnglishName represents betty american english name.
	bettyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Betty"}

	// bettyCanadianFrenchName represents betty canadian french name.
	bettyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// bettyDutchName represents betty dutch name.
	bettyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// bettyFrenchName represents betty french name.
	bettyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Betty"}

	// bettyGermanName represents betty german name.
	bettyGermanName = nook.Name{
		Language: language.German,
		Value:    "Margot"}

	// bettyItalianName represents betty italian name.
	bettyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Crestina"}

	// bettyJapaneseName represents betty japanese name.
	bettyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ホイップ"}

	// bettyKoreanName represents betty korean name.
	bettyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// bettyLatinAmericanSpanishName represents betty latin american spanish name.
	bettyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// bettyRussianName represents betty russian name.
	bettyRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// bettySimplifiedChineseName represents betty simplified chinese name.
	bettySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嗒嗒"}

	// bettySpanishName represents betty spanish name.
	bettySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Clueca"}

	// bettyTraditionalChineseName represents betty traditional chinese name.
	bettyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// bettyName represents betty name.
	bettyName = nook.Languages{
		language.AmericanEnglish:      bettyAmericanEnglishName,
		language.CanadianFrench:       bettyCanadianFrenchName,
		language.Dutch:                bettyDutchName,
		language.French:               bettyFrenchName,
		language.German:               bettyGermanName,
		language.Italian:              bettyItalianName,
		language.Japanese:             bettyJapaneseName,
		language.Korean:               bettyKoreanName,
		language.LatinAmericanSpanish: bettyLatinAmericanSpanishName,
		language.Russian:              bettyRussianName,
		language.SimplifiedChinese:    bettySimplifiedChineseName,
		language.Spanish:              bettySpanishName,
		language.TraditionalChinese:   bettyTraditionalChineseName}
)

var (
	// bettyCharacter represents betty character.
	bettyCharacter = nook.Character{
		Animal:   animal.Chicken,
		Birthday: bettyBirthday,
		Code:     bettyCode,
		Key:      character.Betty,
		Gender:   gender.Female,
		Name:     bettyName,
		Special:  false}
)

var (
	// bettyAmericanEnglishPhrase represents betty american english phrase.
	bettyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cluckling"}

	// bettyCanadianFrenchPhrase represents betty canadian french phrase.
	bettyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// bettyDutchPhrase represents betty dutch phrase.
	bettyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// bettyFrenchPhrase represents betty french phrase.
	bettyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "côt côt"}

	// bettyGermanPhrase represents betty german phrase.
	bettyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "klackl"}

	// bettyItalianPhrase represents betty italian phrase.
	bettyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "coccodè"}

	// bettyJapanesePhrase represents betty japanese phrase.
	bettyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だがね"}

	// bettyKoreanPhrase represents betty korean phrase.
	bettyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// bettyLatinAmericanSpanishPhrase represents betty latin american spanish phrase.
	bettyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// bettyRussianPhrase represents betty russian phrase.
	bettyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// bettySimplifiedChinesePhrase represents betty simplified chinese phrase.
	bettySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "ok"}

	// bettySpanishPhrase represents betty spanish phrase.
	bettySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "clo-clo"}

	// bettyTraditionalChinesePhrase represents betty traditional chinese phrase.
	bettyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// bettyPhrase represents betty phrase.
	bettyPhrase = nook.Languages{
		language.AmericanEnglish:      bettyAmericanEnglishPhrase,
		language.CanadianFrench:       bettyCanadianFrenchPhrase,
		language.Dutch:                bettyDutchPhrase,
		language.French:               bettyFrenchPhrase,
		language.German:               bettyGermanPhrase,
		language.Italian:              bettyItalianPhrase,
		language.Japanese:             bettyJapanesePhrase,
		language.Korean:               bettyKoreanPhrase,
		language.LatinAmericanSpanish: bettyLatinAmericanSpanishPhrase,
		language.Russian:              bettyRussianPhrase,
		language.SimplifiedChinese:    bettySimplifiedChinesePhrase,
		language.Spanish:              bettySpanishPhrase,
		language.TraditionalChinese:   bettyTraditionalChinesePhrase}
)

var (
	// Betty represents betty.
	Betty = nook.Villager{
		Character:   bettyCharacter,
		Personality: personality.Normal,
		Phrase:      bettyPhrase}
)

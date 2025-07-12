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

// meganBirthday represents Megan's birthday.
var (
	// meganBirthday represents megan birthday.
	meganBirthday = nook.Birthday{
		Day:   13,
		Month: time.March}
)

// meganCode represents Megan's unique code.
var (
	// meganCode represents megan code.
	meganCode = nook.Code{
		Value: "bea15"}
)

// Different names for Megan in various languages.
var (
	// meganAmericanEnglishName represents megan american english name.
	meganAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Megan"}

	// meganCanadianFrenchName represents megan canadian french name.
	meganCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Candy"}

	// meganDutchName represents megan dutch name.
	meganDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Megan"}

	// meganFrenchName represents megan french name.
	meganFrenchName = nook.Name{
		Language: language.French,
		Value:    "Candy"}

	// meganGermanName represents megan german name.
	meganGermanName = nook.Name{
		Language: language.German,
		Value:    "Dagmar"}

	// meganItalianName represents megan italian name.
	meganItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Dolcinia"}

	// meganJapaneseName represents megan japanese name.
	meganJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "キャンディ"}

	// meganKoreanName represents megan korean name.
	meganKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "캔디"}

	// meganLatinAmericanSpanishName represents megan latin american spanish name.
	meganLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Carmela"}

	// meganRussianName represents megan russian name.
	meganRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Меган"}

	// meganSimplifiedChineseName represents megan simplified chinese name.
	meganSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "糖果"}

	// meganSpanishName represents megan spanish name.
	meganSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Carmela"}

	// meganTraditionalChineseName represents megan traditional chinese name.
	meganTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "糖果"}
)

// meganName represents Megan's name in different languages.
var (
	// meganName represents megan name.
	meganName = nook.Languages{
		language.AmericanEnglish:      meganAmericanEnglishName,
		language.CanadianFrench:       meganCanadianFrenchName,
		language.Dutch:                meganDutchName,
		language.French:               meganFrenchName,
		language.German:               meganGermanName,
		language.Italian:              meganItalianName,
		language.Japanese:             meganJapaneseName,
		language.Korean:               meganKoreanName,
		language.LatinAmericanSpanish: meganLatinAmericanSpanishName,
		language.Russian:              meganRussianName,
		language.SimplifiedChinese:    meganSimplifiedChineseName,
		language.Spanish:              meganSpanishName,
		language.TraditionalChinese:   meganTraditionalChineseName}
)

// meganCharacter represents Megan's character information.
var (
	// meganCharacter represents megan character.
	meganCharacter = nook.Character{
		Animal:   animal.Bear,
		Birthday: meganBirthday,
		Code:     meganCode,
		Key:      character.Megan,
		Gender:   gender.Female,
		Name:     meganName,
		Special:  false}
)

// Different phrases spoken by Megan in various languages.
var (
	// meganAmericanEnglishPhrase represents megan american english phrase.
	meganAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sundae"}

	// meganCanadianFrenchPhrase represents megan canadian french phrase.
	meganCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "berlingot"}

	// meganDutchPhrase represents megan dutch phrase.
	meganDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "lolly"}

	// meganFrenchPhrase represents megan french phrase.
	meganFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "berlingot"}

	// meganGermanPhrase represents megan german phrase.
	meganGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "lollipop"}

	// meganItalianPhrase represents megan italian phrase.
	meganItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sucré"}

	// meganJapanesePhrase represents megan japanese phrase.
	meganJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ぺろ"}

	// meganKoreanPhrase represents megan korean phrase.
	meganKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "낼름"}

	// meganLatinAmericanSpanishPhrase represents megan latin american spanish phrase.
	meganLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "tutifruti"}

	// meganRussianPhrase represents megan russian phrase.
	meganRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "сладенько"}

	// meganSimplifiedChinesePhrase represents megan simplified chinese phrase.
	meganSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "舔舔"}

	// meganSpanishPhrase represents megan spanish phrase.
	meganSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tutifruti"}

	// meganTraditionalChinesePhrase represents megan traditional chinese phrase.
	meganTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "舔舔"}
)

// meganPhrase represents Megan's phrases in different languages.
var (
	// meganPhrase represents megan phrase.
	meganPhrase = nook.Languages{
		language.AmericanEnglish:      meganAmericanEnglishPhrase,
		language.CanadianFrench:       meganCanadianFrenchPhrase,
		language.Dutch:                meganDutchPhrase,
		language.French:               meganFrenchPhrase,
		language.German:               meganGermanPhrase,
		language.Italian:              meganItalianPhrase,
		language.Japanese:             meganJapanesePhrase,
		language.Korean:               meganKoreanPhrase,
		language.LatinAmericanSpanish: meganLatinAmericanSpanishPhrase,
		language.Russian:              meganRussianPhrase,
		language.SimplifiedChinese:    meganSimplifiedChinesePhrase,
		language.Spanish:              meganSpanishPhrase,
		language.TraditionalChinese:   meganTraditionalChinesePhrase}
)

// Megan represents the character Megan with her complete information.
var (
	// Megan represents megan.
	Megan = nook.Villager{
		Character:   meganCharacter,
		Personality: personality.Normal,
		Phrase:      meganPhrase}
)

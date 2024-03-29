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
	meganBirthday = nook.Birthday{
		Day:   13,
		Month: time.March}
)

// meganCode represents Megan's unique code.
var (
	meganCode = nook.Code{
		Value: "bea15"}
)

// Different names for Megan in various languages.
var (
	meganAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Megan"}

	meganCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Candy"}

	meganDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Megan"}

	meganFrenchName = nook.Name{
		Language: language.French,
		Value:    "Candy"}

	meganGermanName = nook.Name{
		Language: language.German,
		Value:    "Dagmar"}

	meganItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Dolcinia"}

	meganJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "キャンディ"}

	meganKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "캔디"}

	meganLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Carmela"}

	meganRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Меган"}

	meganSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "糖果"}

	meganSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Carmela"}

	meganTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "糖果"}
)

// meganName represents Megan's name in different languages.
var (
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
	meganAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sundae"}

	meganCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "berlingot"}

	meganDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "lolly"}

	meganFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "berlingot"}

	meganGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "lollipop"}

	meganItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sucré"}

	meganJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ぺろ"}

	meganKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "낼름"}

	meganLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "tutifruti"}

	meganRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "сладенько"}

	meganSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "舔舔"}

	meganSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tutifruti"}

	meganTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "舔舔"}
)

// meganPhrase represents Megan's phrases in different languages.
var (
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
	Megan = nook.Villager{
		Character:   meganCharacter,
		Personality: personality.Normal,
		Phrase:      meganPhrase}
)

package alligator

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

// bootsBirthday represents Boots' birthday.
var (
	// bootsBirthday represents boots birthday.
	bootsBirthday = nook.Birthday{
		Day:   7,
		Month: time.August}
)

// bootsCode represents Boots' unique code.
var (
	// bootsCode represents boots code.
	bootsCode = nook.Code{
		Value: "crd02"}
)

// Different names for Boots in various languages.
var (
	// bootsAmericanEnglishName represents boots american english name.
	bootsAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Boots"}

	// bootsCanadianFrenchName represents boots canadian french name.
	bootsCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Croko"}

	// bootsDutchName represents boots dutch name.
	bootsDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Boots"}

	// bootsFrenchName represents boots french name.
	bootsFrenchName = nook.Name{
		Language: language.French,
		Value:    "Croko"}

	// bootsGermanName represents boots german name.
	bootsGermanName = nook.Name{
		Language: language.German,
		Value:    "Tilmann"}

	// bootsItalianName represents boots italian name.
	bootsItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Crocco"}

	// bootsJapaneseName represents boots japanese name.
	bootsJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ホウサク"}

	// bootsKoreanName represents boots korean name.
	bootsKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "풍작"}

	// bootsLatinAmericanSpanishName represents boots latin american spanish name.
	bootsLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Braulio"}

	// bootsRussianName represents boots russian name.
	bootsRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бутс"}

	// bootsSimplifiedChineseName represents boots simplified chinese name.
	bootsSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "丰年"}

	// bootsSpanishName represents boots spanish name.
	bootsSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Braulio"}

	// bootsTraditionalChineseName represents boots traditional chinese name.
	bootsTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "豐年"}
)

// bootsName represents Boots' name in different languages.
var (
	// bootsName represents boots name.
	bootsName = nook.Languages{
		language.AmericanEnglish:      bootsAmericanEnglishName,
		language.CanadianFrench:       bootsCanadianFrenchName,
		language.Dutch:                bootsDutchName,
		language.French:               bootsFrenchName,
		language.German:               bootsGermanName,
		language.Italian:              bootsItalianName,
		language.Japanese:             bootsJapaneseName,
		language.Korean:               bootsKoreanName,
		language.LatinAmericanSpanish: bootsLatinAmericanSpanishName,
		language.Russian:              bootsRussianName,
		language.SimplifiedChinese:    bootsSimplifiedChineseName,
		language.Spanish:              bootsSpanishName,
		language.TraditionalChinese:   bootsTraditionalChineseName}
)

// bootsCharacter represents Boots' character information.
var (
	// bootsCharacter represents boots character.
	bootsCharacter = nook.Character{
		Animal:   animal.Alligator,
		Birthday: bootsBirthday,
		Code:     bootsCode,
		Key:      character.Boots,
		Gender:   gender.Male,
		Name:     bootsName,
		Special:  false}
)

// Different phrases spoken by Boots in various languages.
var (
	// bootsAmericanEnglishPhrase represents boots american english phrase.
	bootsAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "munchie"}

	// bootsCanadianFrenchPhrase represents boots canadian french phrase.
	bootsCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "miam miam"}

	// bootsDutchPhrase represents boots dutch phrase.
	bootsDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "smikkel"}

	// bootsFrenchPhrase represents boots french phrase.
	bootsFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mariolle"}

	// bootsGermanPhrase represents boots german phrase.
	bootsGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schmatzi"}

	// bootsItalianPhrase represents boots italian phrase.
	bootsItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "gnammi"}

	// bootsJapanesePhrase represents boots japanese phrase.
	bootsJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だぴょん"}

	// bootsKoreanPhrase represents boots korean phrase.
	bootsKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "만세"}

	// bootsLatinAmericanSpanishPhrase represents boots latin american spanish phrase.
	bootsLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "crococroco"}

	// bootsRussianPhrase represents boots russian phrase.
	bootsRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кусь"}

	// bootsSimplifiedChinesePhrase represents boots simplified chinese phrase.
	bootsSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "跳"}

	// bootsSpanishPhrase represents boots spanish phrase.
	bootsSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "crococroco"}

	// bootsTraditionalChinesePhrase represents boots traditional chinese phrase.
	bootsTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "跳"}
)

// bootsPhrase represents Boots' phrases in different languages.
var (
	// bootsPhrase represents boots phrase.
	bootsPhrase = nook.Languages{
		language.AmericanEnglish:      bootsAmericanEnglishPhrase,
		language.CanadianFrench:       bootsCanadianFrenchPhrase,
		language.Dutch:                bootsDutchPhrase,
		language.French:               bootsFrenchPhrase,
		language.German:               bootsGermanPhrase,
		language.Italian:              bootsItalianPhrase,
		language.Japanese:             bootsJapanesePhrase,
		language.Korean:               bootsKoreanPhrase,
		language.LatinAmericanSpanish: bootsLatinAmericanSpanishPhrase,
		language.Russian:              bootsRussianPhrase,
		language.SimplifiedChinese:    bootsSimplifiedChinesePhrase,
		language.Spanish:              bootsSpanishPhrase,
		language.TraditionalChinese:   bootsTraditionalChinesePhrase}
)

// Boots represents the character Boots with his complete information.
var (
	// Boots represents boots.
	Boots = nook.Villager{
		Character:   bootsCharacter,
		Personality: personality.Jock,
		Phrase:      bootsPhrase}
)

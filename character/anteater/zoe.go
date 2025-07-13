package anteater

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

// zoeBirthday represents Zoe's birthday.
var (
	// zoeBirthday represents zoe birthday.
	zoeBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

// zoeCode represents Zoe's unique code.
var (
	// zoeCode represents zoe code.
	zoeCode = nook.Code{
		Value: ""}
)

// Different names for Zoe in various languages.
var (
	// zoeAmericanEnglishName represents zoe american english name.
	zoeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Zoe"}

	// zoeCanadianFrenchName represents zoe canadian french name.
	zoeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// zoeDutchName represents zoe dutch name.
	zoeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// zoeFrenchName represents zoe french name.
	zoeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Zoé"}

	// zoeGermanName represents zoe german name.
	zoeGermanName = nook.Name{
		Language: language.German,
		Value:    "Zoe"}

	// zoeItalianName represents zoe italian name.
	zoeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Zoe"}

	// zoeJapaneseName represents zoe japanese name.
	zoeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ビーフン"}

	// zoeKoreanName represents zoe korean name.
	zoeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// zoeLatinAmericanSpanishName represents zoe latin american spanish name.
	zoeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// zoeRussianName represents zoe russian name.
	zoeRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// zoeSimplifiedChineseName represents zoe simplified chinese name.
	zoeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "蔓蔓"}

	// zoeSpanishName represents zoe spanish name.
	zoeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Naricia"}

	// zoeTraditionalChineseName represents zoe traditional chinese name.
	zoeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

// zoeName represents Zoe's name in different languages.
var (
	// zoeName represents zoe name.
	zoeName = nook.Languages{
		language.AmericanEnglish:      zoeAmericanEnglishName,
		language.CanadianFrench:       zoeCanadianFrenchName,
		language.Dutch:                zoeDutchName,
		language.French:               zoeFrenchName,
		language.German:               zoeGermanName,
		language.Italian:              zoeItalianName,
		language.Japanese:             zoeJapaneseName,
		language.Korean:               zoeKoreanName,
		language.LatinAmericanSpanish: zoeLatinAmericanSpanishName,
		language.Russian:              zoeRussianName,
		language.SimplifiedChinese:    zoeSimplifiedChineseName,
		language.Spanish:              zoeSpanishName,
		language.TraditionalChinese:   zoeTraditionalChineseName}
)

// zoeCharacter represents Zoe's character information.
var (
	// zoeCharacter represents zoe character.
	zoeCharacter = nook.Character{
		Animal:   animal.Anteater,
		Birthday: zoeBirthday,
		Code:     zoeCode,
		Key:      character.Zoe,
		Gender:   gender.Female,
		Name:     zoeName,
		Special:  false}
)

// Different phrases spoken by Zoe in various languages.
var (
	// zoeAmericanEnglishPhrase represents zoe american english phrase.
	zoeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "whiiifff"}

	// zoeCanadianFrenchPhrase represents zoe canadian french phrase.
	zoeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// zoeDutchPhrase represents zoe dutch phrase.
	zoeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// zoeFrenchPhrase represents zoe french phrase.
	zoeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "fouf fouf"}

	// zoeGermanPhrase represents zoe german phrase.
	zoeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnieef"}

	// zoeItalianPhrase represents zoe italian phrase.
	zoeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "oiiinfff"}

	// zoeJapanesePhrase represents zoe japanese phrase.
	zoeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だモン"}

	// zoeKoreanPhrase represents zoe korean phrase.
	zoeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// zoeLatinAmericanSpanishPhrase represents zoe latin american spanish phrase.
	zoeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// zoeRussianPhrase represents zoe russian phrase.
	zoeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// zoeSimplifiedChinesePhrase represents zoe simplified chinese phrase.
	zoeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "没错"}

	// zoeSpanishPhrase represents zoe spanish phrase.
	zoeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "fufuf"}

	// zoeTraditionalChinesePhrase represents zoe traditional chinese phrase.
	zoeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

// zoePhrase represents Zoe's phrases in different languages.
var (
	// zoePhrase represents zoe phrase.
	zoePhrase = nook.Languages{
		language.AmericanEnglish:      zoeAmericanEnglishPhrase,
		language.CanadianFrench:       zoeCanadianFrenchPhrase,
		language.Dutch:                zoeDutchPhrase,
		language.French:               zoeFrenchPhrase,
		language.German:               zoeGermanPhrase,
		language.Italian:              zoeItalianPhrase,
		language.Japanese:             zoeJapanesePhrase,
		language.Korean:               zoeKoreanPhrase,
		language.LatinAmericanSpanish: zoeLatinAmericanSpanishPhrase,
		language.Russian:              zoeRussianPhrase,
		language.SimplifiedChinese:    zoeSimplifiedChinesePhrase,
		language.Spanish:              zoeSpanishPhrase,
		language.TraditionalChinese:   zoeTraditionalChinesePhrase}
)

// Zoe represents the character Zoe with her complete information.
var (
	// Zoe represents zoe.
	Zoe = nook.Villager{
		Character:   zoeCharacter,
		Personality: personality.Normal,
		Phrase:      zoePhrase}
)

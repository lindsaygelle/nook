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

// ikeBirthday represents Ike's birthday.
var (
	// ikeBirthday represents ike birthday.
	ikeBirthday = nook.Birthday{
		Day:   16,
		Month: time.May}
)

// ikeCode represents Ike's unique code.
var (
	// ikeCode represents ike code.
	ikeCode = nook.Code{
		Value: "bea11"}
)

// Different names for Ike in various languages.
var (
	// ikeAmericanEnglishName represents ike american english name.
	ikeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Ike"}

	// ikeCanadianFrenchName represents ike canadian french name.
	ikeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Isaac"}

	// ikeDutchName represents ike dutch name.
	ikeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Ike"}

	// ikeFrenchName represents ike french name.
	ikeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Isaac"}

	// ikeGermanName represents ike german name.
	ikeGermanName = nook.Name{
		Language: language.German,
		Value:    "Ike"}

	// ikeItalianName represents ike italian name.
	ikeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ortensio"}

	// ikeJapaneseName represents ike japanese name.
	ikeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ダイク"}

	// ikeKoreanName represents ike korean name.
	ikeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "대공"}

	// ikeLatinAmericanSpanishName represents ike latin american spanish name.
	ikeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Astillas"}

	// ikeRussianName represents ike russian name.
	ikeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Айк"}

	// ikeSimplifiedChineseName represents ike simplified chinese name.
	ikeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "大功"}

	// ikeSpanishName represents ike spanish name.
	ikeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Astillas"}

	// ikeTraditionalChineseName represents ike traditional chinese name.
	ikeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "大功"}
)

// ikeName represents Ike's name in different languages.
var (
	// ikeName represents ike name.
	ikeName = nook.Languages{
		language.AmericanEnglish:      ikeAmericanEnglishName,
		language.CanadianFrench:       ikeCanadianFrenchName,
		language.Dutch:                ikeDutchName,
		language.French:               ikeFrenchName,
		language.German:               ikeGermanName,
		language.Italian:              ikeItalianName,
		language.Japanese:             ikeJapaneseName,
		language.Korean:               ikeKoreanName,
		language.LatinAmericanSpanish: ikeLatinAmericanSpanishName,
		language.Russian:              ikeRussianName,
		language.SimplifiedChinese:    ikeSimplifiedChineseName,
		language.Spanish:              ikeSpanishName,
		language.TraditionalChinese:   ikeTraditionalChineseName}
)

// ikeCharacter represents Ike's character information.
var (
	// ikeCharacter represents ike character.
	ikeCharacter = nook.Character{
		Animal:   animal.Bear,
		Birthday: ikeBirthday,
		Code:     ikeCode,
		Key:      character.Ike,
		Gender:   gender.Male,
		Name:     ikeName,
		Special:  false}
)

// Different phrases spoken by Ike in various languages.
var (
	// ikeAmericanEnglishPhrase represents ike american english phrase.
	ikeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "roadie"}

	// ikeCanadianFrenchPhrase represents ike canadian french phrase.
	ikeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pot d'miel"}

	// ikeDutchPhrase represents ike dutch phrase.
	ikeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "ja denk"}

	// ikeFrenchPhrase represents ike french phrase.
	ikeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pot d'miel"}

	// ikeGermanPhrase represents ike german phrase.
	ikeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "brrruah"}

	// ikeItalianPhrase represents ike italian phrase.
	ikeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sgrugno"}

	// ikeJapanesePhrase represents ike japanese phrase.
	ikeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ボウズ"}

	// ikeKoreanPhrase represents ike korean phrase.
	ikeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "터프"}

	// ikeLatinAmericanSpanishPhrase represents ike latin american spanish phrase.
	ikeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "grurrr"}

	// ikeRussianPhrase represents ike russian phrase.
	ikeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "гарр-сон"}

	// ikeSimplifiedChinesePhrase represents ike simplified chinese phrase.
	ikeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "施主"}

	// ikeSpanishPhrase represents ike spanish phrase.
	ikeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tornillos"}

	// ikeTraditionalChinesePhrase represents ike traditional chinese phrase.
	ikeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "施主"}
)

// ikePhrase represents Ike's phrases in different languages.
var (
	// ikePhrase represents ike phrase.
	ikePhrase = nook.Languages{
		language.AmericanEnglish:      ikeAmericanEnglishPhrase,
		language.CanadianFrench:       ikeCanadianFrenchPhrase,
		language.Dutch:                ikeDutchPhrase,
		language.French:               ikeFrenchPhrase,
		language.German:               ikeGermanPhrase,
		language.Italian:              ikeItalianPhrase,
		language.Japanese:             ikeJapanesePhrase,
		language.Korean:               ikeKoreanPhrase,
		language.LatinAmericanSpanish: ikeLatinAmericanSpanishPhrase,
		language.Russian:              ikeRussianPhrase,
		language.SimplifiedChinese:    ikeSimplifiedChinesePhrase,
		language.Spanish:              ikeSpanishPhrase,
		language.TraditionalChinese:   ikeTraditionalChinesePhrase}
)

// Ike represents the character Ike with his complete information.
var (
	// Ike represents ike.
	Ike = nook.Villager{
		Character:   ikeCharacter,
		Personality: personality.Cranky,
		Phrase:      ikePhrase}
)

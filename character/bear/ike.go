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
	ikeBirthday = nook.Birthday{
		Day:   16,
		Month: time.May}
)

// ikeCode represents Ike's unique code.
var (
	ikeCode = nook.Code{
		Value: "bea11"}
)

// Different names for Ike in various languages.
var (
	ikeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Ike"}

	ikeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Isaac"}

	ikeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Ike"}

	ikeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Isaac"}

	ikeGermanName = nook.Name{
		Language: language.German,
		Value:    "Ike"}

	ikeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ortensio"}

	ikeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ダイク"}

	ikeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "대공"}

	ikeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Astillas"}

	ikeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Айк"}

	ikeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "大功"}

	ikeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Astillas"}

	ikeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "大功"}
)

// ikeName represents Ike's name in different languages.
var (
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
	ikeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "roadie"}

	ikeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pot d'miel"}

	ikeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "ja denk"}

	ikeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pot d'miel"}

	ikeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "brrruah"}

	ikeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sgrugno"}

	ikeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ボウズ"}

	ikeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "터프"}

	ikeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "grurrr"}

	ikeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "гарр-сон"}

	ikeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "施主"}

	ikeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tornillos"}

	ikeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "施主"}
)

// ikePhrase represents Ike's phrases in different languages.
var (
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
	Ike = nook.Villager{
		Character:   ikeCharacter,
		Personality: personality.Cranky,
		Phrase:      ikePhrase}
)

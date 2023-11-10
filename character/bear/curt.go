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

// curtBirthday represents Curt's birthday (July 1st).
var (
	curtBirthday = nook.Birthday{
		Day:   1,
		Month: time.July}
)

// curtCode represents Curt's unique code (bea02).
var (
	curtCode = nook.Code{
		Value: "bea02"}
)

// Different names for Curt in various languages.
var (
	curtAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Curt"}

	curtCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Curt"}

	curtDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Curt"}

	curtFrenchName = nook.Name{
		Language: language.French,
		Value:    "Curt"}

	curtGermanName = nook.Name{
		Language: language.German,
		Value:    "Kurt"}

	curtItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Curt"}

	curtJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ガンテツ"}

	curtKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "뚝심"}

	curtLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Gorbaché"}

	curtRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Керт"}

	curtSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "铁熊"}

	curtSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Gorbaché"}

	curtTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鐵熊"}
)

// curtName represents Curt's name in different languages.
var (
	curtName = nook.Languages{
		language.AmericanEnglish:      curtAmericanEnglishName,
		language.CanadianFrench:       curtCanadianFrenchName,
		language.Dutch:                curtDutchName,
		language.French:               curtFrenchName,
		language.German:               curtGermanName,
		language.Italian:              curtItalianName,
		language.Japanese:             curtJapaneseName,
		language.Korean:               curtKoreanName,
		language.LatinAmericanSpanish: curtLatinAmericanSpanishName,
		language.Russian:              curtRussianName,
		language.SimplifiedChinese:    curtSimplifiedChineseName,
		language.Spanish:              curtSpanishName,
		language.TraditionalChinese:   curtTraditionalChineseName}
)

// curtCharacter represents Curt's character information.
var (
	curtCharacter = nook.Character{
		Animal:   animal.Bear,
		Birthday: curtBirthday,
		Code:     curtCode,
		Key:      character.Curt,
		Gender:   gender.Male,
		Name:     curtName,
		Special:  false}
)

// Different phrases spoken by Curt in various languages.
var (
	curtAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "fuzzball"}

	curtCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bouboule"}

	curtDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "brombeer"}

	curtFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bouboule"}

	curtGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "grumml"}

	curtItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "grrroan"}

	curtJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ウム"}

	curtKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "음"}

	curtLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "gromp"}

	curtRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "махрово"}

	curtSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嗯唔"}

	curtSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "oye"}

	curtTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嗯唔"}
)

// curtPhrase represents Curt's phrases in different languages.
var (
	curtPhrase = nook.Languages{
		language.AmericanEnglish:      curtAmericanEnglishPhrase,
		language.CanadianFrench:       curtCanadianFrenchPhrase,
		language.Dutch:                curtDutchPhrase,
		language.French:               curtFrenchPhrase,
		language.German:               curtGermanPhrase,
		language.Italian:              curtItalianPhrase,
		language.Japanese:             curtJapanesePhrase,
		language.Korean:               curtKoreanPhrase,
		language.LatinAmericanSpanish: curtLatinAmericanSpanishPhrase,
		language.Russian:              curtRussianPhrase,
		language.SimplifiedChinese:    curtSimplifiedChinesePhrase,
		language.Spanish:              curtSpanishPhrase,
		language.TraditionalChinese:   curtTraditionalChinesePhrase}
)

// Curt represents the character Curt with his complete information.
var (
	Curt = nook.Villager{
		Character:   curtCharacter,
		Personality: personality.Cranky,
		Phrase:      curtPhrase}
)

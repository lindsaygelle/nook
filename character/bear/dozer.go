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

// dozerBirthday represents Dozer's birthday.
var (
	dozerBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

// dozerCode represents Dozer's unique code.
var (
	dozerCode = nook.Code{
		Value: ""}
)

// Different names for Dozer in various languages.
var (
	dozerAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Dozer"}

	dozerCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	dozerDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	dozerFrenchName = nook.Name{
		Language: language.French,
		Value:    "Doudou"}

	dozerGermanName = nook.Name{
		Language: language.German,
		Value:    "Nicki"}

	dozerItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ronfo"}

	dozerJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "スリープ"}

	dozerKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	dozerLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	dozerRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	dozerSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "忪忪"}

	dozerSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Toño"}

	dozerTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

// dozerName represents Dozer's name in different languages.
var (
	dozerName = nook.Languages{
		language.AmericanEnglish:      dozerAmericanEnglishName,
		language.CanadianFrench:       dozerCanadianFrenchName,
		language.Dutch:                dozerDutchName,
		language.French:               dozerFrenchName,
		language.German:               dozerGermanName,
		language.Italian:              dozerItalianName,
		language.Japanese:             dozerJapaneseName,
		language.Korean:               dozerKoreanName,
		language.LatinAmericanSpanish: dozerLatinAmericanSpanishName,
		language.Russian:              dozerRussianName,
		language.SimplifiedChinese:    dozerSimplifiedChineseName,
		language.Spanish:              dozerSpanishName,
		language.TraditionalChinese:   dozerTraditionalChineseName}
)

// dozerCharacter represents Dozer's character information.
var (
	dozerCharacter = nook.Character{
		Animal:   animal.Bear,
		Birthday: dozerBirthday,
		Code:     dozerCode,
		Key:      character.Dozer,
		Gender:   gender.Male,
		Name:     dozerName,
		Special:  false}
)

// Different phrases spoken by Dozer in various languages.
var (
	dozerAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "zzzzzz"}

	dozerCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	dozerDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	dozerFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "zzzzzz"}

	dozerGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnarch"}

	dozerItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "zzzzzz"}

	dozerJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でアリ"}

	dozerKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	dozerLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	dozerRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	dozerSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咕"}

	dozerSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "zzzzzz"}

	dozerTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

// dozerPhrase represents Dozer's phrases in different languages.
var (
	dozerPhrase = nook.Languages{
		language.AmericanEnglish:      dozerAmericanEnglishPhrase,
		language.CanadianFrench:       dozerCanadianFrenchPhrase,
		language.Dutch:                dozerDutchPhrase,
		language.French:               dozerFrenchPhrase,
		language.German:               dozerGermanPhrase,
		language.Italian:              dozerItalianPhrase,
		language.Japanese:             dozerJapanesePhrase,
		language.Korean:               dozerKoreanPhrase,
		language.LatinAmericanSpanish: dozerLatinAmericanSpanishPhrase,
		language.Russian:              dozerRussianPhrase,
		language.SimplifiedChinese:    dozerSimplifiedChinesePhrase,
		language.Spanish:              dozerSpanishPhrase,
		language.TraditionalChinese:   dozerTraditionalChinesePhrase}
)

// Dozer represents the character Dozer with his complete information.
var (
	Dozer = nook.Villager{
		Character:   dozerCharacter,
		Personality: personality.Lazy,
		Phrase:      dozerPhrase}
)

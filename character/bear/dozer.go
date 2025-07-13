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
	// dozerBirthday represents dozer birthday.
	dozerBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

// dozerCode represents Dozer's unique code.
var (
	// dozerCode represents dozer code.
	dozerCode = nook.Code{
		Value: ""}
)

// Different names for Dozer in various languages.
var (
	// dozerAmericanEnglishName represents dozer american english name.
	dozerAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Dozer"}

	// dozerCanadianFrenchName represents dozer canadian french name.
	dozerCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// dozerDutchName represents dozer dutch name.
	dozerDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// dozerFrenchName represents dozer french name.
	dozerFrenchName = nook.Name{
		Language: language.French,
		Value:    "Doudou"}

	// dozerGermanName represents dozer german name.
	dozerGermanName = nook.Name{
		Language: language.German,
		Value:    "Nicki"}

	// dozerItalianName represents dozer italian name.
	dozerItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ronfo"}

	// dozerJapaneseName represents dozer japanese name.
	dozerJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "スリープ"}

	// dozerKoreanName represents dozer korean name.
	dozerKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// dozerLatinAmericanSpanishName represents dozer latin american spanish name.
	dozerLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// dozerRussianName represents dozer russian name.
	dozerRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// dozerSimplifiedChineseName represents dozer simplified chinese name.
	dozerSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "忪忪"}

	// dozerSpanishName represents dozer spanish name.
	dozerSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Toño"}

	// dozerTraditionalChineseName represents dozer traditional chinese name.
	dozerTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

// dozerName represents Dozer's name in different languages.
var (
	// dozerName represents dozer name.
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
	// dozerCharacter represents dozer character.
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
	// dozerAmericanEnglishPhrase represents dozer american english phrase.
	dozerAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "zzzzzz"}

	// dozerCanadianFrenchPhrase represents dozer canadian french phrase.
	dozerCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// dozerDutchPhrase represents dozer dutch phrase.
	dozerDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// dozerFrenchPhrase represents dozer french phrase.
	dozerFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "zzzzzz"}

	// dozerGermanPhrase represents dozer german phrase.
	dozerGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnarch"}

	// dozerItalianPhrase represents dozer italian phrase.
	dozerItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "zzzzzz"}

	// dozerJapanesePhrase represents dozer japanese phrase.
	dozerJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でアリ"}

	// dozerKoreanPhrase represents dozer korean phrase.
	dozerKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// dozerLatinAmericanSpanishPhrase represents dozer latin american spanish phrase.
	dozerLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// dozerRussianPhrase represents dozer russian phrase.
	dozerRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// dozerSimplifiedChinesePhrase represents dozer simplified chinese phrase.
	dozerSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咕"}

	// dozerSpanishPhrase represents dozer spanish phrase.
	dozerSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "zzzzzz"}

	// dozerTraditionalChinesePhrase represents dozer traditional chinese phrase.
	dozerTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

// dozerPhrase represents Dozer's phrases in different languages.
var (
	// dozerPhrase represents dozer phrase.
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
	// Dozer represents dozer.
	Dozer = nook.Villager{
		Character:   dozerCharacter,
		Personality: personality.Lazy,
		Phrase:      dozerPhrase}
)

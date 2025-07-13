package dog

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
	// bowBirthday represents bow birthday.
	bowBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// bowCode represents bow code.
	bowCode = nook.Code{
		Value: ""}
)

var (
	// bowAmericanEnglishName represents bow american english name.
	bowAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bow"}

	// bowCanadianFrenchName represents bow canadian french name.
	bowCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// bowDutchName represents bow dutch name.
	bowDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// bowFrenchName represents bow french name.
	bowFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	// bowGermanName represents bow german name.
	bowGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	// bowItalianName represents bow italian name.
	bowItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// bowJapaneseName represents bow japanese name.
	bowJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "バウ"}

	// bowKoreanName represents bow korean name.
	bowKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// bowLatinAmericanSpanishName represents bow latin american spanish name.
	bowLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// bowRussianName represents bow russian name.
	bowRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// bowSimplifiedChineseName represents bow simplified chinese name.
	bowSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// bowSpanishName represents bow spanish name.
	bowSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// bowTraditionalChineseName represents bow traditional chinese name.
	bowTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// bowName represents bow name.
	bowName = nook.Languages{
		language.AmericanEnglish:      bowAmericanEnglishName,
		language.CanadianFrench:       bowCanadianFrenchName,
		language.Dutch:                bowDutchName,
		language.French:               bowFrenchName,
		language.German:               bowGermanName,
		language.Italian:              bowItalianName,
		language.Japanese:             bowJapaneseName,
		language.Korean:               bowKoreanName,
		language.LatinAmericanSpanish: bowLatinAmericanSpanishName,
		language.Russian:              bowRussianName,
		language.SimplifiedChinese:    bowSimplifiedChineseName,
		language.Spanish:              bowSpanishName,
		language.TraditionalChinese:   bowTraditionalChineseName}
)

var (
	// bowCharacter represents bow character.
	bowCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: bowBirthday,
		Code:     bowCode,
		Key:      character.Bow,
		Gender:   gender.Male,
		Name:     bowName,
		Special:  false}
)

var (
	// bowAmericanEnglishPhrase represents bow american english phrase.
	bowAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "バウ"}

	// bowCanadianFrenchPhrase represents bow canadian french phrase.
	bowCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// bowDutchPhrase represents bow dutch phrase.
	bowDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// bowFrenchPhrase represents bow french phrase.
	bowFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	// bowGermanPhrase represents bow german phrase.
	bowGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	// bowItalianPhrase represents bow italian phrase.
	bowItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// bowJapanesePhrase represents bow japanese phrase.
	bowJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "バウ"}

	// bowKoreanPhrase represents bow korean phrase.
	bowKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// bowLatinAmericanSpanishPhrase represents bow latin american spanish phrase.
	bowLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// bowRussianPhrase represents bow russian phrase.
	bowRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// bowSimplifiedChinesePhrase represents bow simplified chinese phrase.
	bowSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// bowSpanishPhrase represents bow spanish phrase.
	bowSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// bowTraditionalChinesePhrase represents bow traditional chinese phrase.
	bowTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// bowPhrase represents bow phrase.
	bowPhrase = nook.Languages{
		language.AmericanEnglish:      bowAmericanEnglishPhrase,
		language.CanadianFrench:       bowCanadianFrenchPhrase,
		language.Dutch:                bowDutchPhrase,
		language.French:               bowFrenchPhrase,
		language.German:               bowGermanPhrase,
		language.Italian:              bowItalianPhrase,
		language.Japanese:             bowJapanesePhrase,
		language.Korean:               bowKoreanPhrase,
		language.LatinAmericanSpanish: bowLatinAmericanSpanishPhrase,
		language.Russian:              bowRussianPhrase,
		language.SimplifiedChinese:    bowSimplifiedChinesePhrase,
		language.Spanish:              bowSpanishPhrase,
		language.TraditionalChinese:   bowTraditionalChinesePhrase}
)

var (
	// Bow represents bow.
	Bow = nook.Villager{
		Character:   bowCharacter,
		Personality: personality.Lazy,
		Phrase:      bowPhrase}
)

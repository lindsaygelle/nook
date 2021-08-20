package pig

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
	piglegBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	piglegCode = nook.Code{
		Value: ""}
)

var (
	piglegAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pigleg"}

	piglegCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Xavier"}

	piglegDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	piglegFrenchName = nook.Name{
		Language: language.French,
		Value:    "Xavier"}

	piglegGermanName = nook.Name{
		Language: language.German,
		Value:    "Jonas"}

	piglegItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Drake"}

	piglegJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "バチコーメ"}

	piglegKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	piglegLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Jamoncio"}

	piglegRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	piglegSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	piglegSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Jamoncio"}

	piglegTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	piglegName = nook.Languages{
		language.AmericanEnglish:      piglegAmericanEnglishName,
		language.CanadianFrench:       piglegCanadianFrenchName,
		language.Dutch:                piglegDutchName,
		language.French:               piglegFrenchName,
		language.German:               piglegGermanName,
		language.Italian:              piglegItalianName,
		language.Japanese:             piglegJapaneseName,
		language.Korean:               piglegKoreanName,
		language.LatinAmericanSpanish: piglegLatinAmericanSpanishName,
		language.Russian:              piglegRussianName,
		language.SimplifiedChinese:    piglegSimplifiedChineseName,
		language.Spanish:              piglegSpanishName,
		language.TraditionalChinese:   piglegTraditionalChineseName}
)

var (
	piglegCharacter = nook.Character{
		Animal:   animal.Pig,
		Birthday: piglegBirthday,
		Code:     piglegCode,
		Key:      character.Pigleg,
		Gender:   gender.Male,
		Name:     piglegName,
		Special:  false}
)

var (
	piglegAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "arrrn"}

	piglegCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Unknown"}

	piglegDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	piglegFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "Unknown"}

	piglegGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "Unknown"}

	piglegItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "grugno"}

	piglegJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ヨーソロ"}

	piglegKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	piglegLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Unknown"}

	piglegRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	piglegSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	piglegSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "Unknown"}

	piglegTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	piglegPhrase = nook.Languages{
		language.AmericanEnglish:      piglegAmericanEnglishPhrase,
		language.CanadianFrench:       piglegCanadianFrenchPhrase,
		language.Dutch:                piglegDutchPhrase,
		language.French:               piglegFrenchPhrase,
		language.German:               piglegGermanPhrase,
		language.Italian:              piglegItalianPhrase,
		language.Japanese:             piglegJapanesePhrase,
		language.Korean:               piglegKoreanPhrase,
		language.LatinAmericanSpanish: piglegLatinAmericanSpanishPhrase,
		language.Russian:              piglegRussianPhrase,
		language.SimplifiedChinese:    piglegSimplifiedChinesePhrase,
		language.Spanish:              piglegSpanishPhrase,
		language.TraditionalChinese:   piglegTraditionalChinesePhrase}
)

var (
	Pigleg = nook.Villager{
		Character:   piglegCharacter,
		Personality: personality.Jock,
		Phrase:      piglegPhrase}
)

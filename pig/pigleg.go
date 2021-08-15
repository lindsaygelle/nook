package pig

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "XavierUnknown"}

	piglegDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	piglegFrenchName = nook.Name{
		Language: language.French,
		Value:    "XavierUnknown"}

	piglegGermanName = nook.Name{
		Language: language.German,
		Value:    "JonasUnknown"}

	piglegItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Drakegrugno"}

	piglegJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "バチコーメヨーソロ"}

	piglegKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	piglegLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "JamoncioUnknown"}

	piglegRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	piglegSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	piglegSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "JamoncioUnknown"}

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
		Animal:   Pig,
		Birthday: piglegBirthday,
		Code:     piglegCode,
		Gender:   nook.Male,
		Name:     piglegName}
)

var (
	piglegAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	piglegCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Unknown"}

	piglegDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

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
		Value:    "N/A"}

	piglegLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Unknown"}

	piglegRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	piglegSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	piglegSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "Unknown"}

	piglegTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	piglegPhrase = nook.Languages{
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
	Pigleg = nook.Villager{
		Character:   piglegCharacter,
		Personality: nook.Jock,
		Phrase:      piglegPhrase}
)

package bull

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	oxfordBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	oxfordCode = nook.Code{
		Value: ""}
)

var (
	oxfordAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Oxford"}

	oxfordCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	oxfordDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	oxfordFrenchName = nook.Name{
		Language: language.French,
		Value:    "Sorbonnela vache"}

	oxfordGermanName = nook.Name{
		Language: language.German,
		Value:    "Rolfbulli, hä"}

	oxfordItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Angusmuuh"}

	oxfordJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "イワオでんがな"}

	oxfordKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	oxfordLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	oxfordRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	oxfordSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "岩岩可不"}

	oxfordSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Astaúlfotorito"}

	oxfordTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	oxfordName = nook.Languages{
		language.AmericanEnglish:      oxfordAmericanEnglishName,
		language.CanadianFrench:       oxfordCanadianFrenchName,
		language.Dutch:                oxfordDutchName,
		language.French:               oxfordFrenchName,
		language.German:               oxfordGermanName,
		language.Italian:              oxfordItalianName,
		language.Japanese:             oxfordJapaneseName,
		language.Korean:               oxfordKoreanName,
		language.LatinAmericanSpanish: oxfordLatinAmericanSpanishName,
		language.Russian:              oxfordRussianName,
		language.SimplifiedChinese:    oxfordSimplifiedChineseName,
		language.Spanish:              oxfordSpanishName,
		language.TraditionalChinese:   oxfordTraditionalChineseName}
)

var (
	oxfordCharacter = nook.Character{
		Animal:   Bull,
		Birthday: oxfordBirthday,
		Code:     oxfordCode,
		Gender:   nook.Male,
		Name:     oxfordName}
)

var (
	oxfordAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	oxfordCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	oxfordDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	oxfordFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "la vache"}

	oxfordGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bulli, hä"}

	oxfordItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "muuh"}

	oxfordJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でんがな"}

	oxfordKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	oxfordLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	oxfordRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	oxfordSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "可不"}

	oxfordSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "torito"}

	oxfordTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	oxfordPhrase = nook.Languages{
		language.AmericanEnglish:      oxfordAmericanEnglishName,
		language.CanadianFrench:       oxfordCanadianFrenchName,
		language.Dutch:                oxfordDutchName,
		language.French:               oxfordFrenchName,
		language.German:               oxfordGermanName,
		language.Italian:              oxfordItalianName,
		language.Japanese:             oxfordJapaneseName,
		language.Korean:               oxfordKoreanName,
		language.LatinAmericanSpanish: oxfordLatinAmericanSpanishName,
		language.Russian:              oxfordRussianName,
		language.SimplifiedChinese:    oxfordSimplifiedChineseName,
		language.Spanish:              oxfordSpanishName,
		language.TraditionalChinese:   oxfordTraditionalChineseName}
)

var (
	Oxford = nook.Villager{
		Character:   oxfordCharacter,
		Personality: nook.Cranky,
		Phrase:      oxfordPhrase}
)

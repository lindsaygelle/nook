package bull

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
		Value:    "Sorbonne"}

	oxfordGermanName = nook.Name{
		Language: language.German,
		Value:    "Rolf"}

	oxfordItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Angus"}

	oxfordJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "イワオ"}

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
		Value:    "岩岩"}

	oxfordSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Astaúlfo"}

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
		Animal:   animal.Bull,
		Birthday: oxfordBirthday,
		Code:     oxfordCode,
		Key:      character.Oxford,
		Gender:   gender.Male,
		Name:     oxfordName,
		Special:  false}
)

var (
	oxfordAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bully, eh"}

	oxfordCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	oxfordDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

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
		Value:    ""}

	oxfordLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	oxfordRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	oxfordSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "可不"}

	oxfordSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "torito"}

	oxfordTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	oxfordPhrase = nook.Languages{
		language.AmericanEnglish:      oxfordAmericanEnglishPhrase,
		language.CanadianFrench:       oxfordCanadianFrenchPhrase,
		language.Dutch:                oxfordDutchPhrase,
		language.French:               oxfordFrenchPhrase,
		language.German:               oxfordGermanPhrase,
		language.Italian:              oxfordItalianPhrase,
		language.Japanese:             oxfordJapanesePhrase,
		language.Korean:               oxfordKoreanPhrase,
		language.LatinAmericanSpanish: oxfordLatinAmericanSpanishPhrase,
		language.Russian:              oxfordRussianPhrase,
		language.SimplifiedChinese:    oxfordSimplifiedChinesePhrase,
		language.Spanish:              oxfordSpanishPhrase,
		language.TraditionalChinese:   oxfordTraditionalChinesePhrase}
)

var (
	Oxford = nook.Villager{
		Character:   oxfordCharacter,
		Personality: personality.Cranky,
		Phrase:      oxfordPhrase}
)

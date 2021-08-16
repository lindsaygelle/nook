package bird

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	twirpBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	twirpCode = nook.Code{
		Value: ""}
)

var (
	twirpAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Twirp"}

	twirpCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	twirpDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	twirpFrenchName = nook.Name{
		Language: language.French,
		Value:    "Neuneu"}

	twirpGermanName = nook.Name{
		Language: language.German,
		Value:    "Meik"}

	twirpItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Furbecco"}

	twirpJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "でチュン"}

	twirpKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	twirpLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	twirpRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	twirpSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "子谈"}

	twirpSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tomás"}

	twirpTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	twirpName = nook.Languages{
		language.AmericanEnglish:      twirpAmericanEnglishName,
		language.CanadianFrench:       twirpCanadianFrenchName,
		language.Dutch:                twirpDutchName,
		language.French:               twirpFrenchName,
		language.German:               twirpGermanName,
		language.Italian:              twirpItalianName,
		language.Japanese:             twirpJapaneseName,
		language.Korean:               twirpKoreanName,
		language.LatinAmericanSpanish: twirpLatinAmericanSpanishName,
		language.Russian:              twirpRussianName,
		language.SimplifiedChinese:    twirpSimplifiedChineseName,
		language.Spanish:              twirpSpanishName,
		language.TraditionalChinese:   twirpTraditionalChineseName}
)

var (
	twirpCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: twirpBirthday,
		Code:     twirpCode,
		Gender:   gender.Male,
		Name:     twirpName}
)

var (
	twirpAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "CHIRP"}

	twirpCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	twirpDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	twirpFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "oufffff"}

	twirpGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "tschirp"}

	twirpItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "CIRP"}

	twirpJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でチュン"}

	twirpKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	twirpLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	twirpRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	twirpSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咻"}

	twirpSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "PÍO-PÍO"}

	twirpTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	twirpPhrase = nook.Languages{
		language.AmericanEnglish:      twirpAmericanEnglishName,
		language.CanadianFrench:       twirpCanadianFrenchName,
		language.Dutch:                twirpDutchName,
		language.French:               twirpFrenchName,
		language.German:               twirpGermanName,
		language.Italian:              twirpItalianName,
		language.Japanese:             twirpJapaneseName,
		language.Korean:               twirpKoreanName,
		language.LatinAmericanSpanish: twirpLatinAmericanSpanishName,
		language.Russian:              twirpRussianName,
		language.SimplifiedChinese:    twirpSimplifiedChineseName,
		language.Spanish:              twirpSpanishName,
		language.TraditionalChinese:   twirpTraditionalChineseName}
)

var (
	Twirp = nook.Villager{
		Character:   twirpCharacter,
		Personality: personality.Cranky,
		Phrase:      twirpPhrase}
)

package penguin

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
	roaldBirthday = nook.Birthday{
		Day:   5,
		Month: time.January}
)

var (
	roaldCode = nook.Code{
		Value: "pgn01"}
)

var (
	roaldAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Roald"}

	roaldCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Reynald"}

	roaldDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Roald"}

	roaldFrenchName = nook.Name{
		Language: language.French,
		Value:    "Reynald"}

	roaldGermanName = nook.Name{
		Language: language.German,
		Value:    "Roland"}

	roaldItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Angelino"}

	roaldJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ペンタ"}

	roaldKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "펭수"}

	roaldLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bobi"}

	roaldRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Роальд"}

	roaldSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "企鹅达"}

	roaldSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bobi"}

	roaldTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "企鵝達"}
)

var (
	roaldName = nook.Languages{
		language.AmericanEnglish:      roaldAmericanEnglishName,
		language.CanadianFrench:       roaldCanadianFrenchName,
		language.Dutch:                roaldDutchName,
		language.French:               roaldFrenchName,
		language.German:               roaldGermanName,
		language.Italian:              roaldItalianName,
		language.Japanese:             roaldJapaneseName,
		language.Korean:               roaldKoreanName,
		language.LatinAmericanSpanish: roaldLatinAmericanSpanishName,
		language.Russian:              roaldRussianName,
		language.SimplifiedChinese:    roaldSimplifiedChineseName,
		language.Spanish:              roaldSpanishName,
		language.TraditionalChinese:   roaldTraditionalChineseName}
)

var (
	roaldCharacter = nook.Character{
		Animal:   animal.Penguin,
		Birthday: roaldBirthday,
		Code:     roaldCode,
		Key:      character.Roald,
		Gender:   gender.Male,
		Name:     roaldName}
)

var (
	roaldAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "b-b-buddy"}

	roaldCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pépépère"}

	roaldDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "m-m-maatje"}

	roaldFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pépépère"}

	roaldGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "k-k-kumpel"}

	roaldItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "br-br-brr"}

	roaldJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だペン"}

	roaldKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "팽팽"}

	roaldLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "titirito"}

	roaldRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "д-дружище"}

	roaldSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "企鹅"}

	roaldSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "titirito"}

	roaldTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "企鵝"}
)

var (
	roaldPhrase = nook.Languages{
		language.AmericanEnglish:      roaldAmericanEnglishName,
		language.CanadianFrench:       roaldCanadianFrenchName,
		language.Dutch:                roaldDutchName,
		language.French:               roaldFrenchName,
		language.German:               roaldGermanName,
		language.Italian:              roaldItalianName,
		language.Japanese:             roaldJapaneseName,
		language.Korean:               roaldKoreanName,
		language.LatinAmericanSpanish: roaldLatinAmericanSpanishName,
		language.Russian:              roaldRussianName,
		language.SimplifiedChinese:    roaldSimplifiedChineseName,
		language.Spanish:              roaldSpanishName,
		language.TraditionalChinese:   roaldTraditionalChineseName}
)

var (
	Roald = nook.Villager{
		Character:   roaldCharacter,
		Personality: personality.Jock,
		Phrase:      roaldPhrase}
)

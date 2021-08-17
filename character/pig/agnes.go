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
	agnesBirthday = nook.Birthday{
		Day:   21,
		Month: time.April}
)

var (
	agnesCode = nook.Code{
		Value: "pig17"}
)

var (
	agnesAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Agnes"}

	agnesCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Pansy"}

	agnesDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Agnes"}

	agnesFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pansy"}

	agnesGermanName = nook.Name{
		Language: language.German,
		Value:    "Nora"}

	agnesItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Maia"}

	agnesJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アグネス"}

	agnesKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아그네스"}

	agnesLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Nerea"}

	agnesRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Агнес"}

	agnesSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "古乃欣"}

	agnesSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Negrea"}

	agnesTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "古乃欣"}
)

var (
	agnesName = nook.Languages{
		language.AmericanEnglish:      agnesAmericanEnglishName,
		language.CanadianFrench:       agnesCanadianFrenchName,
		language.Dutch:                agnesDutchName,
		language.French:               agnesFrenchName,
		language.German:               agnesGermanName,
		language.Italian:              agnesItalianName,
		language.Japanese:             agnesJapaneseName,
		language.Korean:               agnesKoreanName,
		language.LatinAmericanSpanish: agnesLatinAmericanSpanishName,
		language.Russian:              agnesRussianName,
		language.SimplifiedChinese:    agnesSimplifiedChineseName,
		language.Spanish:              agnesSpanishName,
		language.TraditionalChinese:   agnesTraditionalChineseName}
)

var (
	agnesCharacter = nook.Character{
		Animal:   animal.Pig,
		Birthday: agnesBirthday,
		Code:     agnesCode,
		Key:      character.Agnes,
		Gender:   gender.Female,
		Name:     agnesName}
)

var (
	agnesAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "snuffle"}

	agnesCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "sauciflard"}

	agnesDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snuffelaar"}

	agnesFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "sauciflard"}

	agnesGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnorrz"}

	agnesItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "splosh"}

	agnesJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ブフフ"}

	agnesKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "꿀꾸루"}

	agnesLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "dindoink"}

	agnesRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хрю-хрю"}

	agnesSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "噗呼呼"}

	agnesSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "dindoink"}

	agnesTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "噗呼呼"}
)

var (
	agnesPhrase = nook.Languages{
		language.AmericanEnglish:      agnesAmericanEnglishPhrase,
		language.CanadianFrench:       agnesCanadianFrenchPhrase,
		language.Dutch:                agnesDutchPhrase,
		language.French:               agnesFrenchPhrase,
		language.German:               agnesGermanPhrase,
		language.Italian:              agnesItalianPhrase,
		language.Japanese:             agnesJapanesePhrase,
		language.Korean:               agnesKoreanPhrase,
		language.LatinAmericanSpanish: agnesLatinAmericanSpanishPhrase,
		language.Russian:              agnesRussianPhrase,
		language.SimplifiedChinese:    agnesSimplifiedChinesePhrase,
		language.Spanish:              agnesSpanishPhrase,
		language.TraditionalChinese:   agnesTraditionalChinesePhrase}
)

var (
	Agnes = nook.Villager{
		Character:   agnesCharacter,
		Personality: personality.BigSister,
		Phrase:      agnesPhrase}
)

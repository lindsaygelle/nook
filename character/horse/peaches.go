package horse

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
	peachesBirthday = nook.Birthday{
		Day:   28,
		Month: time.November}
)

var (
	peachesCode = nook.Code{
		Value: "hrs08"}
)

var (
	peachesAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Peaches"}

	peachesCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Prune"}

	peachesDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Peaches"}

	peachesFrenchName = nook.Name{
		Language: language.French,
		Value:    "Prune"}

	peachesGermanName = nook.Name{
		Language: language.German,
		Value:    "Claire"}

	peachesItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ronzina"}

	peachesJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ドサコ"}

	peachesKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "말자"}

	peachesLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Perla"}

	peachesRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пичес"}

	peachesSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小薰"}

	peachesSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Perla"}

	peachesTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小薰"}
)

var (
	peachesName = nook.Languages{
		language.AmericanEnglish:      peachesAmericanEnglishName,
		language.CanadianFrench:       peachesCanadianFrenchName,
		language.Dutch:                peachesDutchName,
		language.French:               peachesFrenchName,
		language.German:               peachesGermanName,
		language.Italian:              peachesItalianName,
		language.Japanese:             peachesJapaneseName,
		language.Korean:               peachesKoreanName,
		language.LatinAmericanSpanish: peachesLatinAmericanSpanishName,
		language.Russian:              peachesRussianName,
		language.SimplifiedChinese:    peachesSimplifiedChineseName,
		language.Spanish:              peachesSpanishName,
		language.TraditionalChinese:   peachesTraditionalChineseName}
)

var (
	peachesCharacter = nook.Character{
		Animal:   animal.Horse,
		Birthday: peachesBirthday,
		Code:     peachesCode,
		Key:      character.Peaches,
		Gender:   gender.Female,
		Name:     peachesName,
		Special:  false}
)

var (
	peachesAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "neighbor"}

	peachesCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mustang"}

	peachesDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hinnik"}

	peachesFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mustang"}

	peachesGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "nachbar"}

	peachesItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "oh mamma"}

	peachesJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だポン"}

	peachesKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "몰라요"}

	peachesLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "niiiih"}

	peachesRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "соседушка"}

	peachesSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "蹦"}

	peachesSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "niiiih"}

	peachesTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "蹦"}
)

var (
	peachesPhrase = nook.Languages{
		language.AmericanEnglish:      peachesAmericanEnglishPhrase,
		language.CanadianFrench:       peachesCanadianFrenchPhrase,
		language.Dutch:                peachesDutchPhrase,
		language.French:               peachesFrenchPhrase,
		language.German:               peachesGermanPhrase,
		language.Italian:              peachesItalianPhrase,
		language.Japanese:             peachesJapanesePhrase,
		language.Korean:               peachesKoreanPhrase,
		language.LatinAmericanSpanish: peachesLatinAmericanSpanishPhrase,
		language.Russian:              peachesRussianPhrase,
		language.SimplifiedChinese:    peachesSimplifiedChinesePhrase,
		language.Spanish:              peachesSpanishPhrase,
		language.TraditionalChinese:   peachesTraditionalChinesePhrase}
)

var (
	Peaches = nook.Villager{
		Character:   peachesCharacter,
		Personality: personality.Normal,
		Phrase:      peachesPhrase}
)

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
	// peachesBirthday represents peaches birthday.
	peachesBirthday = nook.Birthday{
		Day:   28,
		Month: time.November}
)

var (
	// peachesCode represents peaches code.
	peachesCode = nook.Code{
		Value: "hrs08"}
)

var (
	// peachesAmericanEnglishName represents peaches american english name.
	peachesAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Peaches"}

	// peachesCanadianFrenchName represents peaches canadian french name.
	peachesCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Prune"}

	// peachesDutchName represents peaches dutch name.
	peachesDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Peaches"}

	// peachesFrenchName represents peaches french name.
	peachesFrenchName = nook.Name{
		Language: language.French,
		Value:    "Prune"}

	// peachesGermanName represents peaches german name.
	peachesGermanName = nook.Name{
		Language: language.German,
		Value:    "Claire"}

	// peachesItalianName represents peaches italian name.
	peachesItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ronzina"}

	// peachesJapaneseName represents peaches japanese name.
	peachesJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ドサコ"}

	// peachesKoreanName represents peaches korean name.
	peachesKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "말자"}

	// peachesLatinAmericanSpanishName represents peaches latin american spanish name.
	peachesLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Perla"}

	// peachesRussianName represents peaches russian name.
	peachesRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пичес"}

	// peachesSimplifiedChineseName represents peaches simplified chinese name.
	peachesSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小薰"}

	// peachesSpanishName represents peaches spanish name.
	peachesSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Perla"}

	// peachesTraditionalChineseName represents peaches traditional chinese name.
	peachesTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小薰"}
)

var (
	// peachesName represents peaches name.
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
	// peachesCharacter represents peaches character.
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
	// peachesAmericanEnglishPhrase represents peaches american english phrase.
	peachesAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "neighbor"}

	// peachesCanadianFrenchPhrase represents peaches canadian french phrase.
	peachesCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mustang"}

	// peachesDutchPhrase represents peaches dutch phrase.
	peachesDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hinnik"}

	// peachesFrenchPhrase represents peaches french phrase.
	peachesFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mustang"}

	// peachesGermanPhrase represents peaches german phrase.
	peachesGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "nachbar"}

	// peachesItalianPhrase represents peaches italian phrase.
	peachesItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "oh mamma"}

	// peachesJapanesePhrase represents peaches japanese phrase.
	peachesJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だポン"}

	// peachesKoreanPhrase represents peaches korean phrase.
	peachesKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "몰라요"}

	// peachesLatinAmericanSpanishPhrase represents peaches latin american spanish phrase.
	peachesLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "niiiih"}

	// peachesRussianPhrase represents peaches russian phrase.
	peachesRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "соседушка"}

	// peachesSimplifiedChinesePhrase represents peaches simplified chinese phrase.
	peachesSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "蹦"}

	// peachesSpanishPhrase represents peaches spanish phrase.
	peachesSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "niiiih"}

	// peachesTraditionalChinesePhrase represents peaches traditional chinese phrase.
	peachesTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "蹦"}
)

var (
	// peachesPhrase represents peaches phrase.
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
	// Peaches represents peaches.
	Peaches = nook.Villager{
		Character:   peachesCharacter,
		Personality: personality.Normal,
		Phrase:      peachesPhrase}
)

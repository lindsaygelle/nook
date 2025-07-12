package cat

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
	// kidcatBirthday represents kidcat birthday.
	kidcatBirthday = nook.Birthday{
		Day:   1,
		Month: time.August}
)

var (
	// kidcatCode represents kidcat code.
	kidcatCode = nook.Code{
		Value: "cat10"}
)

var (
	// kidcatAmericanEnglishName represents kidcat american english name.
	kidcatAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Kid Cat"}

	// kidcatCanadianFrenchName represents kidcat canadian french name.
	kidcatCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Câlin"}

	// kidcatDutchName represents kidcat dutch name.
	kidcatDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Kid Cat"}

	// kidcatFrenchName represents kidcat french name.
	kidcatFrenchName = nook.Name{
		Language: language.French,
		Value:    "Câlin"}

	// kidcatGermanName represents kidcat german name.
	kidcatGermanName = nook.Name{
		Language: language.German,
		Value:    "Pit"}

	// kidcatItalianName represents kidcat italian name.
	kidcatItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Zampa"}

	// kidcatJapaneseName represents kidcat japanese name.
	kidcatJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "１ごう"}

	// kidcatKoreanName represents kidcat korean name.
	kidcatKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "1호"}

	// kidcatLatinAmericanSpanishName represents kidcat latin american spanish name.
	kidcatLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Gatomán"}

	// kidcatRussianName represents kidcat russian name.
	kidcatRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кид Кэт"}

	// kidcatSimplifiedChineseName represents kidcat simplified chinese name.
	kidcatSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿一"}

	// kidcatSpanishName represents kidcat spanish name.
	kidcatSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Gatomán"}

	// kidcatTraditionalChineseName represents kidcat traditional chinese name.
	kidcatTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿一"}
)

var (
	// kidcatName represents kidcat name.
	kidcatName = nook.Languages{
		language.AmericanEnglish:      kidcatAmericanEnglishName,
		language.CanadianFrench:       kidcatCanadianFrenchName,
		language.Dutch:                kidcatDutchName,
		language.French:               kidcatFrenchName,
		language.German:               kidcatGermanName,
		language.Italian:              kidcatItalianName,
		language.Japanese:             kidcatJapaneseName,
		language.Korean:               kidcatKoreanName,
		language.LatinAmericanSpanish: kidcatLatinAmericanSpanishName,
		language.Russian:              kidcatRussianName,
		language.SimplifiedChinese:    kidcatSimplifiedChineseName,
		language.Spanish:              kidcatSpanishName,
		language.TraditionalChinese:   kidcatTraditionalChineseName}
)

var (
	// kidcatCharacter represents kidcat character.
	kidcatCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: kidcatBirthday,
		Code:     kidcatCode,
		Key:      character.KidCat,
		Gender:   gender.Male,
		Name:     kidcatName,
		Special:  false}
)

var (
	// kidcatAmericanEnglishPhrase represents kidcat american english phrase.
	kidcatAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "psst"}

	// kidcatCanadianFrenchPhrase represents kidcat canadian french phrase.
	kidcatCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "psst"}

	// kidcatDutchPhrase represents kidcat dutch phrase.
	kidcatDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "psst"}

	// kidcatFrenchPhrase represents kidcat french phrase.
	kidcatFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "psst"}

	// kidcatGermanPhrase represents kidcat german phrase.
	kidcatGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "maunz"}

	// kidcatItalianPhrase represents kidcat italian phrase.
	kidcatItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "psst"}

	// kidcatJapanesePhrase represents kidcat japanese phrase.
	kidcatJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "とぉっ"}

	// kidcatKoreanPhrase represents kidcat korean phrase.
	kidcatKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "얍"}

	// kidcatLatinAmericanSpanishPhrase represents kidcat latin american spanish phrase.
	kidcatLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "psst"}

	// kidcatRussianPhrase represents kidcat russian phrase.
	kidcatRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "шш"}

	// kidcatSimplifiedChinesePhrase represents kidcat simplified chinese phrase.
	kidcatSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喝"}

	// kidcatSpanishPhrase represents kidcat spanish phrase.
	kidcatSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "soldado"}

	// kidcatTraditionalChinesePhrase represents kidcat traditional chinese phrase.
	kidcatTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "喝"}
)

var (
	// kidcatPhrase represents kidcat phrase.
	kidcatPhrase = nook.Languages{
		language.AmericanEnglish:      kidcatAmericanEnglishPhrase,
		language.CanadianFrench:       kidcatCanadianFrenchPhrase,
		language.Dutch:                kidcatDutchPhrase,
		language.French:               kidcatFrenchPhrase,
		language.German:               kidcatGermanPhrase,
		language.Italian:              kidcatItalianPhrase,
		language.Japanese:             kidcatJapanesePhrase,
		language.Korean:               kidcatKoreanPhrase,
		language.LatinAmericanSpanish: kidcatLatinAmericanSpanishPhrase,
		language.Russian:              kidcatRussianPhrase,
		language.SimplifiedChinese:    kidcatSimplifiedChinesePhrase,
		language.Spanish:              kidcatSpanishPhrase,
		language.TraditionalChinese:   kidcatTraditionalChinesePhrase}
)

var (
	// KidCat represents kid cat.
	KidCat = nook.Villager{
		Character:   kidcatCharacter,
		Personality: personality.Jock,
		Phrase:      kidcatPhrase}
)

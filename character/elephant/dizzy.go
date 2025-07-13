package elephant

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
	// dizzyBirthday represents dizzy birthday.
	dizzyBirthday = nook.Birthday{
		Day:   14,
		Month: time.July}
)

var (
	// dizzyCode represents dizzy code.
	dizzyCode = nook.Code{
		Value: "elp01"}
)

var (
	// dizzyAmericanEnglishName represents dizzy american english name.
	dizzyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Dizzy"}

	// dizzyCanadianFrenchName represents dizzy canadian french name.
	dizzyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Pachy"}

	// dizzyDutchName represents dizzy dutch name.
	dizzyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Dizzy"}

	// dizzyFrenchName represents dizzy french name.
	dizzyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pachy"}

	// dizzyGermanName represents dizzy german name.
	dizzyGermanName = nook.Name{
		Language: language.German,
		Value:    "Bastian"}

	// dizzyItalianName represents dizzy italian name.
	dizzyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Annibale"}

	// dizzyJapaneseName represents dizzy japanese name.
	dizzyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ヒュージ"}

	// dizzyKoreanName represents dizzy korean name.
	dizzyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "휴지"}

	// dizzyLatinAmericanSpanishName represents dizzy latin american spanish name.
	dizzyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Quique"}

	// dizzyRussianName represents dizzy russian name.
	dizzyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Диззи"}

	// dizzySimplifiedChineseName represents dizzy simplified chinese name.
	dizzySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "巨巨"}

	// dizzySpanishName represents dizzy spanish name.
	dizzySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Quique"}

	// dizzyTraditionalChineseName represents dizzy traditional chinese name.
	dizzyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "巨巨"}
)

var (
	// dizzyName represents dizzy name.
	dizzyName = nook.Languages{
		language.AmericanEnglish:      dizzyAmericanEnglishName,
		language.CanadianFrench:       dizzyCanadianFrenchName,
		language.Dutch:                dizzyDutchName,
		language.French:               dizzyFrenchName,
		language.German:               dizzyGermanName,
		language.Italian:              dizzyItalianName,
		language.Japanese:             dizzyJapaneseName,
		language.Korean:               dizzyKoreanName,
		language.LatinAmericanSpanish: dizzyLatinAmericanSpanishName,
		language.Russian:              dizzyRussianName,
		language.SimplifiedChinese:    dizzySimplifiedChineseName,
		language.Spanish:              dizzySpanishName,
		language.TraditionalChinese:   dizzyTraditionalChineseName}
)

var (
	// dizzyCharacter represents dizzy character.
	dizzyCharacter = nook.Character{
		Animal:   animal.Elephant,
		Birthday: dizzyBirthday,
		Code:     dizzyCode,
		Key:      character.Dizzy,
		Gender:   gender.Male,
		Name:     dizzyName,
		Special:  false}
)

var (
	// dizzyAmericanEnglishPhrase represents dizzy american english phrase.
	dizzyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "woo-oo"}

	// dizzyCanadianFrenchPhrase represents dizzy canadian french phrase.
	dizzyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "woutchou"}

	// dizzyDutchPhrase represents dizzy dutch phrase.
	dizzyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "trompet"}

	// dizzyFrenchPhrase represents dizzy french phrase.
	dizzyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "woutchou"}

	// dizzyGermanPhrase represents dizzy german phrase.
	dizzyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "taraaaa"}

	// dizzyItalianPhrase represents dizzy italian phrase.
	dizzyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ullà"}

	// dizzyJapanesePhrase represents dizzy japanese phrase.
	dizzyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だゾウ"}

	// dizzyKoreanPhrase represents dizzy korean phrase.
	dizzyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "휴휴"}

	// dizzyLatinAmericanSpanishPhrase represents dizzy latin american spanish phrase.
	dizzyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "wuuulá"}

	// dizzyRussianPhrase represents dizzy russian phrase.
	dizzyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "тру-ру"}

	// dizzySimplifiedChinesePhrase represents dizzy simplified chinese phrase.
	dizzySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "象象"}

	// dizzySpanishPhrase represents dizzy spanish phrase.
	dizzySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "naricita"}

	// dizzyTraditionalChinesePhrase represents dizzy traditional chinese phrase.
	dizzyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "象象"}
)

var (
	// dizzyPhrase represents dizzy phrase.
	dizzyPhrase = nook.Languages{
		language.AmericanEnglish:      dizzyAmericanEnglishPhrase,
		language.CanadianFrench:       dizzyCanadianFrenchPhrase,
		language.Dutch:                dizzyDutchPhrase,
		language.French:               dizzyFrenchPhrase,
		language.German:               dizzyGermanPhrase,
		language.Italian:              dizzyItalianPhrase,
		language.Japanese:             dizzyJapanesePhrase,
		language.Korean:               dizzyKoreanPhrase,
		language.LatinAmericanSpanish: dizzyLatinAmericanSpanishPhrase,
		language.Russian:              dizzyRussianPhrase,
		language.SimplifiedChinese:    dizzySimplifiedChinesePhrase,
		language.Spanish:              dizzySpanishPhrase,
		language.TraditionalChinese:   dizzyTraditionalChinesePhrase}
)

var (
	// Dizzy represents dizzy.
	Dizzy = nook.Villager{
		Character:   dizzyCharacter,
		Personality: personality.Lazy,
		Phrase:      dizzyPhrase}
)

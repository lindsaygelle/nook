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
	// frigaBirthday represents friga birthday.
	frigaBirthday = nook.Birthday{
		Day:   16,
		Month: time.October}
)

var (
	// frigaCode represents friga code.
	frigaCode = nook.Code{
		Value: "pgn04"}
)

var (
	// frigaAmericanEnglishName represents friga american english name.
	frigaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Friga"}

	// frigaCanadianFrenchName represents friga canadian french name.
	frigaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Friga"}

	// frigaDutchName represents friga dutch name.
	frigaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Friga"}

	// frigaFrenchName represents friga french name.
	frigaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Friga"}

	// frigaGermanName represents friga german name.
	frigaGermanName = nook.Name{
		Language: language.German,
		Value:    "Frieda"}

	// frigaItalianName represents friga italian name.
	frigaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Frida"}

	// frigaJapaneseName represents friga japanese name.
	frigaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "サブリナ"}

	// frigaKoreanName represents friga korean name.
	frigaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "사브리나"}

	// frigaLatinAmericanSpanishName represents friga latin american spanish name.
	frigaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Frida"}

	// frigaRussianName represents friga russian name.
	frigaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Фрига"}

	// frigaSimplifiedChineseName represents friga simplified chinese name.
	frigaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "谢宾娜"}

	// frigaSpanishName represents friga spanish name.
	frigaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Frida"}

	// frigaTraditionalChineseName represents friga traditional chinese name.
	frigaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "謝賓娜"}
)

var (
	// frigaName represents friga name.
	frigaName = nook.Languages{
		language.AmericanEnglish:      frigaAmericanEnglishName,
		language.CanadianFrench:       frigaCanadianFrenchName,
		language.Dutch:                frigaDutchName,
		language.French:               frigaFrenchName,
		language.German:               frigaGermanName,
		language.Italian:              frigaItalianName,
		language.Japanese:             frigaJapaneseName,
		language.Korean:               frigaKoreanName,
		language.LatinAmericanSpanish: frigaLatinAmericanSpanishName,
		language.Russian:              frigaRussianName,
		language.SimplifiedChinese:    frigaSimplifiedChineseName,
		language.Spanish:              frigaSpanishName,
		language.TraditionalChinese:   frigaTraditionalChineseName}
)

var (
	// frigaCharacter represents friga character.
	frigaCharacter = nook.Character{
		Animal:   animal.Penguin,
		Birthday: frigaBirthday,
		Code:     frigaCode,
		Key:      character.Friga,
		Gender:   gender.Female,
		Name:     frigaName,
		Special:  false}
)

var (
	// frigaAmericanEnglishPhrase represents friga american english phrase.
	frigaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "brrrmph"}

	// frigaCanadianFrenchPhrase represents friga canadian french phrase.
	frigaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bourfff"}

	// frigaDutchPhrase represents friga dutch phrase.
	frigaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "brrrhmpf"}

	// frigaFrenchPhrase represents friga french phrase.
	frigaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bourfff"}

	// frigaGermanPhrase represents friga german phrase.
	frigaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "grrmpf"}

	// frigaItalianPhrase represents friga italian phrase.
	frigaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "brrrumf"}

	// frigaJapanesePhrase represents friga japanese phrase.
	frigaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ツルルン"}

	// frigaKoreanPhrase represents friga korean phrase.
	frigaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "으쓱"}

	// frigaLatinAmericanSpanishPhrase represents friga latin american spanish phrase.
	frigaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "urfff"}

	// frigaRussianPhrase represents friga russian phrase.
	frigaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "брр-хм"}

	// frigaSimplifiedChinesePhrase represents friga simplified chinese phrase.
	frigaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "滑溜溜"}

	// frigaSpanishPhrase represents friga spanish phrase.
	frigaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "heladito"}

	// frigaTraditionalChinesePhrase represents friga traditional chinese phrase.
	frigaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "滑溜溜"}
)

var (
	// frigaPhrase represents friga phrase.
	frigaPhrase = nook.Languages{
		language.AmericanEnglish:      frigaAmericanEnglishPhrase,
		language.CanadianFrench:       frigaCanadianFrenchPhrase,
		language.Dutch:                frigaDutchPhrase,
		language.French:               frigaFrenchPhrase,
		language.German:               frigaGermanPhrase,
		language.Italian:              frigaItalianPhrase,
		language.Japanese:             frigaJapanesePhrase,
		language.Korean:               frigaKoreanPhrase,
		language.LatinAmericanSpanish: frigaLatinAmericanSpanishPhrase,
		language.Russian:              frigaRussianPhrase,
		language.SimplifiedChinese:    frigaSimplifiedChinesePhrase,
		language.Spanish:              frigaSpanishPhrase,
		language.TraditionalChinese:   frigaTraditionalChinesePhrase}
)

var (
	// Friga represents friga.
	Friga = nook.Villager{
		Character:   frigaCharacter,
		Personality: personality.Snooty,
		Phrase:      frigaPhrase}
)

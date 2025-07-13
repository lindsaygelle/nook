package bear

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

// paulaBirthday represents Paula's birthday.
var (
	// paulaBirthday represents paula birthday.
	paulaBirthday = nook.Birthday{
		Day:   22,
		Month: time.March}
)

// paulaCode represents Paula's unique code.
var (
	// paulaCode represents paula code.
	paulaCode = nook.Code{
		Value: "bea10"}
)

// Different names for Paula in various languages.
var (
	// paulaAmericanEnglishName represents paula american english name.
	paulaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Paula"}

	// paulaCanadianFrenchName represents paula canadian french name.
	paulaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Wendy"}

	// paulaDutchName represents paula dutch name.
	paulaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Paula"}

	// paulaFrenchName represents paula french name.
	paulaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Wendy"}

	// paulaGermanName represents paula german name.
	paulaGermanName = nook.Name{
		Language: language.German,
		Value:    "Paula"}

	// paulaItalianName represents paula italian name.
	paulaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Bruna"}

	// paulaJapaneseName represents paula japanese name.
	paulaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "レイチェル"}

	// paulaKoreanName represents paula korean name.
	paulaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "레이첼"}

	// paulaLatinAmericanSpanishName represents paula latin american spanish name.
	paulaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lizi"}

	// paulaRussianName represents paula russian name.
	paulaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Паула"}

	// paulaSimplifiedChineseName represents paula simplified chinese name.
	paulaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "瑞秋"}

	// paulaSpanishName represents paula spanish name.
	paulaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lizi"}

	// paulaTraditionalChineseName represents paula traditional chinese name.
	paulaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "瑞秋"}
)

// paulaName represents Paula's name in different languages.
var (
	// paulaName represents paula name.
	paulaName = nook.Languages{
		language.AmericanEnglish:      paulaAmericanEnglishName,
		language.CanadianFrench:       paulaCanadianFrenchName,
		language.Dutch:                paulaDutchName,
		language.French:               paulaFrenchName,
		language.German:               paulaGermanName,
		language.Italian:              paulaItalianName,
		language.Japanese:             paulaJapaneseName,
		language.Korean:               paulaKoreanName,
		language.LatinAmericanSpanish: paulaLatinAmericanSpanishName,
		language.Russian:              paulaRussianName,
		language.SimplifiedChinese:    paulaSimplifiedChineseName,
		language.Spanish:              paulaSpanishName,
		language.TraditionalChinese:   paulaTraditionalChineseName}
)

// paulaCharacter represents Paula's character information.
var (
	// paulaCharacter represents paula character.
	paulaCharacter = nook.Character{
		Animal:   animal.Bear,
		Birthday: paulaBirthday,
		Code:     paulaCode,
		Key:      character.Paula,
		Gender:   gender.Female,
		Name:     paulaName,
		Special:  false}
)

// Different phrases spoken by Paula in various languages.
var (
	// paulaAmericanEnglishPhrase represents paula american english phrase.
	paulaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "yodelay"}

	// paulaCanadianFrenchPhrase represents paula canadian french phrase.
	paulaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tourlou"}

	// paulaDutchPhrase represents paula dutch phrase.
	paulaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "jodelie"}

	// paulaFrenchPhrase represents paula french phrase.
	paulaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "léchouille"}

	// paulaGermanPhrase represents paula german phrase.
	paulaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bralala"}

	// paulaItalianPhrase represents paula italian phrase.
	paulaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bah"}

	// paulaJapanesePhrase represents paula japanese phrase.
	paulaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ヤッホー"}

	// paulaKoreanPhrase represents paula korean phrase.
	paulaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "야호야호"}

	// paulaLatinAmericanSpanishPhrase represents paula latin american spanish phrase.
	paulaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "borobó"}

	// paulaRussianPhrase represents paula russian phrase.
	paulaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "йодль"}

	// paulaSimplifiedChinesePhrase represents paula simplified chinese phrase.
	paulaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "耶呼"}

	// paulaSpanishPhrase represents paula spanish phrase.
	paulaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "borobó"}

	// paulaTraditionalChinesePhrase represents paula traditional chinese phrase.
	paulaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "耶呼"}
)

// paulaPhrase represents Paula's phrases in different languages.
var (
	// paulaPhrase represents paula phrase.
	paulaPhrase = nook.Languages{
		language.AmericanEnglish:      paulaAmericanEnglishPhrase,
		language.CanadianFrench:       paulaCanadianFrenchPhrase,
		language.Dutch:                paulaDutchPhrase,
		language.French:               paulaFrenchPhrase,
		language.German:               paulaGermanPhrase,
		language.Italian:              paulaItalianPhrase,
		language.Japanese:             paulaJapanesePhrase,
		language.Korean:               paulaKoreanPhrase,
		language.LatinAmericanSpanish: paulaLatinAmericanSpanishPhrase,
		language.Russian:              paulaRussianPhrase,
		language.SimplifiedChinese:    paulaSimplifiedChinesePhrase,
		language.Spanish:              paulaSpanishPhrase,
		language.TraditionalChinese:   paulaTraditionalChinesePhrase}
)

// Paula represents the character Paula with her complete information.
var (
	// Paula represents paula.
	Paula = nook.Villager{
		Character:   paulaCharacter,
		Personality: personality.BigSister,
		Phrase:      paulaPhrase}
)

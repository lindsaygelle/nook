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
	paulaBirthday = nook.Birthday{
		Day:   22,
		Month: time.March}
)

// paulaCode represents Paula's unique code.
var (
	paulaCode = nook.Code{
		Value: "bea10"}
)

// Different names for Paula in various languages.
var (
	paulaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Paula"}

	paulaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Wendy"}

	paulaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Paula"}

	paulaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Wendy"}

	paulaGermanName = nook.Name{
		Language: language.German,
		Value:    "Paula"}

	paulaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Bruna"}

	paulaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "レイチェル"}

	paulaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "레이첼"}

	paulaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lizi"}

	paulaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Паула"}

	paulaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "瑞秋"}

	paulaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lizi"}

	paulaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "瑞秋"}
)

// paulaName represents Paula's name in different languages.
var (
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
	paulaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "yodelay"}

	paulaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tourlou"}

	paulaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "jodelie"}

	paulaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "léchouille"}

	paulaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bralala"}

	paulaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bah"}

	paulaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ヤッホー"}

	paulaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "야호야호"}

	paulaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "borobó"}

	paulaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "йодль"}

	paulaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "耶呼"}

	paulaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "borobó"}

	paulaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "耶呼"}
)

// paulaPhrase represents Paula's phrases in different languages.
var (
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
	Paula = nook.Villager{
		Character:   paulaCharacter,
		Personality: personality.BigSister,
		Phrase:      paulaPhrase}
)

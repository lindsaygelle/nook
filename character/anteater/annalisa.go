package anteater

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

// annalisaBirthday represents Annalisa's birthday.
var (
	// annalisaBirthday represents annalisa birthday.
	annalisaBirthday = nook.Birthday{
		Day:   6,
		Month: time.February}
)

// annalisaCode represents Annalisa's unique code.
var (
	// annalisaCode represents annalisa code.
	annalisaCode = nook.Code{
		Value: "ant08"}
)

// Different names for Annalisa in various languages.
var (
	// annalisaAmericanEnglishName represents annalisa american english name.
	annalisaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Annalisa"}

	// annalisaCanadianFrenchName represents annalisa canadian french name.
	annalisaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Roberta"}

	// annalisaDutchName represents annalisa dutch name.
	annalisaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Annalisa"}

	// annalisaFrenchName represents annalisa french name.
	annalisaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Roberta"}

	// annalisaGermanName represents annalisa german name.
	annalisaGermanName = nook.Name{
		Language: language.German,
		Value:    "Annalena"}

	// annalisaItalianName represents annalisa italian name.
	annalisaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Anita"}

	// annalisaJapaneseName represents annalisa japanese name.
	annalisaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "みやび"}

	// annalisaKoreanName represents annalisa korean name.
	annalisaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "설백"}

	// annalisaLatinAmericanSpanishName represents annalisa latin american spanish name.
	annalisaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Alba"}

	// annalisaRussianName represents annalisa russian name.
	annalisaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Аннализа"}

	// annalisaSimplifiedChineseName represents annalisa simplified chinese name.
	annalisaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小雅"}

	// annalisaSpanishName represents annalisa spanish name.
	annalisaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Alba"}

	// annalisaTraditionalChineseName represents annalisa traditional chinese name.
	annalisaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小雅"}
)

// annalisaName represents Annalisa's name in different languages.
var (
	// annalisaName represents annalisa name.
	annalisaName = nook.Languages{
		language.AmericanEnglish:      annalisaAmericanEnglishName,
		language.CanadianFrench:       annalisaCanadianFrenchName,
		language.Dutch:                annalisaDutchName,
		language.French:               annalisaFrenchName,
		language.German:               annalisaGermanName,
		language.Italian:              annalisaItalianName,
		language.Japanese:             annalisaJapaneseName,
		language.Korean:               annalisaKoreanName,
		language.LatinAmericanSpanish: annalisaLatinAmericanSpanishName,
		language.Russian:              annalisaRussianName,
		language.SimplifiedChinese:    annalisaSimplifiedChineseName,
		language.Spanish:              annalisaSpanishName,
		language.TraditionalChinese:   annalisaTraditionalChineseName}
)

// annalisaCharacter represents Annalisa's character information.
var (
	// annalisaCharacter represents annalisa character.
	annalisaCharacter = nook.Character{
		Animal:   animal.Anteater,
		Birthday: annalisaBirthday,
		Code:     annalisaCode,
		Key:      character.Annalisa,
		Gender:   gender.Female,
		Name:     annalisaName,
		Special:  false}
)

// Different phrases spoken by Annalisa in various languages.
var (
	// annalisaAmericanEnglishPhrase represents annalisa american english phrase.
	annalisaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "gumdrop"}

	// annalisaCanadianFrenchPhrase represents annalisa canadian french phrase.
	annalisaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tap tap"}

	// annalisaDutchPhrase represents annalisa dutch phrase.
	annalisaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "dropje"}

	// annalisaFrenchPhrase represents annalisa french phrase.
	annalisaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tap tap"}

	// annalisaGermanPhrase represents annalisa german phrase.
	annalisaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnüffel"}

	// annalisaItalianPhrase represents annalisa italian phrase.
	annalisaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "chic"}

	// annalisaJapanesePhrase represents annalisa japanese phrase.
	annalisaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "をかし"}

	// annalisaKoreanPhrase represents annalisa korean phrase.
	annalisaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "하양하양"}

	// annalisaLatinAmericanSpanishPhrase represents annalisa latin american spanish phrase.
	annalisaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "snifffi"}

	// annalisaRussianPhrase represents annalisa russian phrase.
	annalisaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "жуй-жуй"}

	// annalisaSimplifiedChinesePhrase represents annalisa simplified chinese phrase.
	annalisaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "优雅"}

	// annalisaSpanishPhrase represents annalisa spanish phrase.
	annalisaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "nubecillas"}

	// annalisaTraditionalChinesePhrase represents annalisa traditional chinese phrase.
	annalisaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "優雅"}
)

// annalisaPhrase represents Annalisa's phrases in different languages.
var (
	// annalisaPhrase represents annalisa phrase.
	annalisaPhrase = nook.Languages{
		language.AmericanEnglish:      annalisaAmericanEnglishPhrase,
		language.CanadianFrench:       annalisaCanadianFrenchPhrase,
		language.Dutch:                annalisaDutchPhrase,
		language.French:               annalisaFrenchPhrase,
		language.German:               annalisaGermanPhrase,
		language.Italian:              annalisaItalianPhrase,
		language.Japanese:             annalisaJapanesePhrase,
		language.Korean:               annalisaKoreanPhrase,
		language.LatinAmericanSpanish: annalisaLatinAmericanSpanishPhrase,
		language.Russian:              annalisaRussianPhrase,
		language.SimplifiedChinese:    annalisaSimplifiedChinesePhrase,
		language.Spanish:              annalisaSpanishPhrase,
		language.TraditionalChinese:   annalisaTraditionalChinesePhrase}
)

// Annalisa represents the character Annalisa with her complete information.
var (
	// Annalisa represents annalisa.
	Annalisa = nook.Villager{
		Character:   annalisaCharacter,
		Personality: personality.Normal,
		Phrase:      annalisaPhrase}
)

package bearcub

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
	// pudgeBirthday represents pudge birthday.
	pudgeBirthday = nook.Birthday{
		Day:   11,
		Month: time.June}
)

var (
	// pudgeCode represents pudge code.
	pudgeCode = nook.Code{
		Value: "cbr03"}
)

var (
	// pudgeAmericanEnglishName represents pudge american english name.
	pudgeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pudge"}

	// pudgeCanadianFrenchName represents pudge canadian french name.
	pudgeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gradub"}

	// pudgeDutchName represents pudge dutch name.
	pudgeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pudge"}

	// pudgeFrenchName represents pudge french name.
	pudgeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gradub"}

	// pudgeGermanName represents pudge german name.
	pudgeGermanName = nook.Name{
		Language: language.German,
		Value:    "Bertram"}

	// pudgeItalianName represents pudge italian name.
	pudgeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Tombolo"}

	// pudgeJapaneseName represents pudge japanese name.
	pudgeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "きんぞう"}

	// pudgeKoreanName represents pudge korean name.
	pudgeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "우띠"}

	// pudgeLatinAmericanSpanishName represents pudge latin american spanish name.
	pudgeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bollito"}

	// pudgeRussianName represents pudge russian name.
	pudgeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Падж"}

	// pudgeSimplifiedChineseName represents pudge simplified chinese name.
	pudgeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿钦"}

	// pudgeSpanishName represents pudge spanish name.
	pudgeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bollito"}

	// pudgeTraditionalChineseName represents pudge traditional chinese name.
	pudgeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿欽"}
)

var (
	// pudgeName represents pudge name.
	pudgeName = nook.Languages{
		language.AmericanEnglish:      pudgeAmericanEnglishName,
		language.CanadianFrench:       pudgeCanadianFrenchName,
		language.Dutch:                pudgeDutchName,
		language.French:               pudgeFrenchName,
		language.German:               pudgeGermanName,
		language.Italian:              pudgeItalianName,
		language.Japanese:             pudgeJapaneseName,
		language.Korean:               pudgeKoreanName,
		language.LatinAmericanSpanish: pudgeLatinAmericanSpanishName,
		language.Russian:              pudgeRussianName,
		language.SimplifiedChinese:    pudgeSimplifiedChineseName,
		language.Spanish:              pudgeSpanishName,
		language.TraditionalChinese:   pudgeTraditionalChineseName}
)

var (
	// pudgeCharacter represents pudge character.
	pudgeCharacter = nook.Character{
		Animal:   animal.Bearcub,
		Birthday: pudgeBirthday,
		Code:     pudgeCode,
		Key:      character.Pudge,
		Gender:   gender.Male,
		Name:     pudgeName,
		Special:  false}
)

var (
	// pudgeAmericanEnglishPhrase represents pudge american english phrase.
	pudgeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "pudgy"}

	// pudgeCanadianFrenchPhrase represents pudge canadian french phrase.
	pudgeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "louloute"}

	// pudgeDutchPhrase represents pudge dutch phrase.
	pudgeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "pudding"}

	// pudgeFrenchPhrase represents pudge french phrase.
	pudgeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "louloute"}

	// pudgeGermanPhrase represents pudge german phrase.
	pudgeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "griesgram"}

	// pudgeItalianPhrase represents pudge italian phrase.
	pudgeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "pulcino"}

	// pudgeJapanesePhrase represents pudge japanese phrase.
	pudgeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "んもう"}

	// pudgeKoreanPhrase represents pudge korean phrase.
	pudgeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아이참"}

	// pudgeLatinAmericanSpanishPhrase represents pudge latin american spanish phrase.
	pudgeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "bruah"}

	// pudgeRussianPhrase represents pudge russian phrase.
	pudgeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "карапуз"}

	// pudgeSimplifiedChinesePhrase represents pudge simplified chinese phrase.
	pudgeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "真是的"}

	// pudgeSpanishPhrase represents pudge spanish phrase.
	pudgeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "rebollos"}

	// pudgeTraditionalChinesePhrase represents pudge traditional chinese phrase.
	pudgeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "真是的"}
)

var (
	// pudgePhrase represents pudge phrase.
	pudgePhrase = nook.Languages{
		language.AmericanEnglish:      pudgeAmericanEnglishPhrase,
		language.CanadianFrench:       pudgeCanadianFrenchPhrase,
		language.Dutch:                pudgeDutchPhrase,
		language.French:               pudgeFrenchPhrase,
		language.German:               pudgeGermanPhrase,
		language.Italian:              pudgeItalianPhrase,
		language.Japanese:             pudgeJapanesePhrase,
		language.Korean:               pudgeKoreanPhrase,
		language.LatinAmericanSpanish: pudgeLatinAmericanSpanishPhrase,
		language.Russian:              pudgeRussianPhrase,
		language.SimplifiedChinese:    pudgeSimplifiedChinesePhrase,
		language.Spanish:              pudgeSpanishPhrase,
		language.TraditionalChinese:   pudgeTraditionalChinesePhrase}
)

var (
	// Pudge represents pudge.
	Pudge = nook.Villager{
		Character:   pudgeCharacter,
		Personality: personality.Lazy,
		Phrase:      pudgePhrase}
)

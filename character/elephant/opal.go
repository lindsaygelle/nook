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
	// opalBirthday represents opal birthday.
	opalBirthday = nook.Birthday{
		Day:   20,
		Month: time.January}
)

var (
	// opalCode represents opal code.
	opalCode = nook.Code{
		Value: "elp00"}
)

var (
	// opalAmericanEnglishName represents opal american english name.
	opalAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Opal"}

	// opalCanadianFrenchName represents opal canadian french name.
	opalCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Opaline"}

	// opalDutchName represents opal dutch name.
	opalDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Opal"}

	// opalFrenchName represents opal french name.
	opalFrenchName = nook.Name{
		Language: language.French,
		Value:    "Opaline"}

	// opalGermanName represents opal german name.
	opalGermanName = nook.Name{
		Language: language.German,
		Value:    "Olga"}

	// opalItalianName represents opal italian name.
	opalItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Opal"}

	// opalJapaneseName represents opal japanese name.
	opalJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "オパール"}

	// opalKoreanName represents opal korean name.
	opalKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "오팔"}

	// opalLatinAmericanSpanishName represents opal latin american spanish name.
	opalLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ópalo"}

	// opalRussianName represents opal russian name.
	opalRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Опал"}

	// opalSimplifiedChineseName represents opal simplified chinese name.
	opalSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "露露"}

	// opalSpanishName represents opal spanish name.
	opalSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ópalo"}

	// opalTraditionalChineseName represents opal traditional chinese name.
	opalTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "露露"}
)

var (
	// opalName represents opal name.
	opalName = nook.Languages{
		language.AmericanEnglish:      opalAmericanEnglishName,
		language.CanadianFrench:       opalCanadianFrenchName,
		language.Dutch:                opalDutchName,
		language.French:               opalFrenchName,
		language.German:               opalGermanName,
		language.Italian:              opalItalianName,
		language.Japanese:             opalJapaneseName,
		language.Korean:               opalKoreanName,
		language.LatinAmericanSpanish: opalLatinAmericanSpanishName,
		language.Russian:              opalRussianName,
		language.SimplifiedChinese:    opalSimplifiedChineseName,
		language.Spanish:              opalSpanishName,
		language.TraditionalChinese:   opalTraditionalChineseName}
)

var (
	// opalCharacter represents opal character.
	opalCharacter = nook.Character{
		Animal:   animal.Elephant,
		Birthday: opalBirthday,
		Code:     opalCode,
		Key:      character.Opal,
		Gender:   gender.Female,
		Name:     opalName,
		Special:  false}
)

var (
	// opalAmericanEnglishPhrase represents opal american english phrase.
	opalAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "snoot"}

	// opalCanadianFrenchPhrase represents opal canadian french phrase.
	opalCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "gropif"}

	// opalDutchPhrase represents opal dutch phrase.
	opalDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snoet"}

	// opalFrenchPhrase represents opal french phrase.
	opalFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "gropif"}

	// opalGermanPhrase represents opal german phrase.
	opalGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "tröööt"}

	// opalItalianPhrase represents opal italian phrase.
	opalItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "snoob"}

	// opalJapanesePhrase represents opal japanese phrase.
	opalJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ヨン"}

	// opalKoreanPhrase represents opal korean phrase.
	opalKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "땡"}

	// opalLatinAmericanSpanishPhrase represents opal latin american spanish phrase.
	opalLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "flip-flop"}

	// opalRussianPhrase represents opal russian phrase.
	opalRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хобот"}

	// opalSimplifiedChinesePhrase represents opal simplified chinese phrase.
	opalSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "勇"}

	// opalSpanishPhrase represents opal spanish phrase.
	opalSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "flip-flop"}

	// opalTraditionalChinesePhrase represents opal traditional chinese phrase.
	opalTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "勇"}
)

var (
	// opalPhrase represents opal phrase.
	opalPhrase = nook.Languages{
		language.AmericanEnglish:      opalAmericanEnglishPhrase,
		language.CanadianFrench:       opalCanadianFrenchPhrase,
		language.Dutch:                opalDutchPhrase,
		language.French:               opalFrenchPhrase,
		language.German:               opalGermanPhrase,
		language.Italian:              opalItalianPhrase,
		language.Japanese:             opalJapanesePhrase,
		language.Korean:               opalKoreanPhrase,
		language.LatinAmericanSpanish: opalLatinAmericanSpanishPhrase,
		language.Russian:              opalRussianPhrase,
		language.SimplifiedChinese:    opalSimplifiedChinesePhrase,
		language.Spanish:              opalSpanishPhrase,
		language.TraditionalChinese:   opalTraditionalChinesePhrase}
)

var (
	// Opal represents opal.
	Opal = nook.Villager{
		Character:   opalCharacter,
		Personality: personality.Snooty,
		Phrase:      opalPhrase}
)

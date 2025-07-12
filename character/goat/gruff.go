package goat

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
	// gruffBirthday represents gruff birthday.
	gruffBirthday = nook.Birthday{
		Day:   29,
		Month: time.August}
)

var (
	// gruffCode represents gruff code.
	gruffCode = nook.Code{
		Value: "goa04"}
)

var (
	// gruffAmericanEnglishName represents gruff american english name.
	gruffAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Gruff"}

	// gruffCanadianFrenchName represents gruff canadian french name.
	gruffCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Grognon"}

	// gruffDutchName represents gruff dutch name.
	gruffDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Gruff"}

	// gruffFrenchName represents gruff french name.
	gruffFrenchName = nook.Name{
		Language: language.French,
		Value:    "Grognon"}

	// gruffGermanName represents gruff german name.
	gruffGermanName = nook.Name{
		Language: language.German,
		Value:    "Gregor"}

	// gruffItalianName represents gruff italian name.
	gruffItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Capriolé"}

	// gruffJapaneseName represents gruff japanese name.
	gruffJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ビリー"}

	// gruffKoreanName represents gruff korean name.
	gruffKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "빌리"}

	// gruffLatinAmericanSpanishName represents gruff latin american spanish name.
	gruffLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Sátiro"}

	// gruffRussianName represents gruff russian name.
	gruffRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Графф"}

	// gruffSimplifiedChineseName represents gruff simplified chinese name.
	gruffSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "比利"}

	// gruffSpanishName represents gruff spanish name.
	gruffSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Sátiro"}

	// gruffTraditionalChineseName represents gruff traditional chinese name.
	gruffTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "比利"}
)

var (
	// gruffName represents gruff name.
	gruffName = nook.Languages{
		language.AmericanEnglish:      gruffAmericanEnglishName,
		language.CanadianFrench:       gruffCanadianFrenchName,
		language.Dutch:                gruffDutchName,
		language.French:               gruffFrenchName,
		language.German:               gruffGermanName,
		language.Italian:              gruffItalianName,
		language.Japanese:             gruffJapaneseName,
		language.Korean:               gruffKoreanName,
		language.LatinAmericanSpanish: gruffLatinAmericanSpanishName,
		language.Russian:              gruffRussianName,
		language.SimplifiedChinese:    gruffSimplifiedChineseName,
		language.Spanish:              gruffSpanishName,
		language.TraditionalChinese:   gruffTraditionalChineseName}
)

var (
	// gruffCharacter represents gruff character.
	gruffCharacter = nook.Character{
		Animal:   animal.Goat,
		Birthday: gruffBirthday,
		Code:     gruffCode,
		Key:      character.Gruff,
		Gender:   gender.Male,
		Name:     gruffName,
		Special:  false}
)

var (
	// gruffAmericanEnglishPhrase represents gruff american english phrase.
	gruffAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bleh eh eh"}

	// gruffCanadianFrenchPhrase represents gruff canadian french phrase.
	gruffCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "hé bêêê"}

	// gruffDutchPhrase represents gruff dutch phrase.
	gruffDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "mekker"}

	// gruffFrenchPhrase represents gruff french phrase.
	gruffFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "hé bêêê"}

	// gruffGermanPhrase represents gruff german phrase.
	gruffGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mähmääh"}

	// gruffItalianPhrase represents gruff italian phrase.
	gruffItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bleh eh"}

	// gruffJapanesePhrase represents gruff japanese phrase.
	gruffJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "イェーイ"}

	// gruffKoreanPhrase represents gruff korean phrase.
	gruffKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "예이예"}

	// gruffLatinAmericanSpanishPhrase represents gruff latin american spanish phrase.
	gruffLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "meeeh"}

	// gruffRussianPhrase represents gruff russian phrase.
	gruffRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ме-хе-хе"}

	// gruffSimplifiedChinesePhrase represents gruff simplified chinese phrase.
	gruffSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "耶"}

	// gruffSpanishPhrase represents gruff spanish phrase.
	gruffSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "zampoña"}

	// gruffTraditionalChinesePhrase represents gruff traditional chinese phrase.
	gruffTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "耶"}
)

var (
	// gruffPhrase represents gruff phrase.
	gruffPhrase = nook.Languages{
		language.AmericanEnglish:      gruffAmericanEnglishPhrase,
		language.CanadianFrench:       gruffCanadianFrenchPhrase,
		language.Dutch:                gruffDutchPhrase,
		language.French:               gruffFrenchPhrase,
		language.German:               gruffGermanPhrase,
		language.Italian:              gruffItalianPhrase,
		language.Japanese:             gruffJapanesePhrase,
		language.Korean:               gruffKoreanPhrase,
		language.LatinAmericanSpanish: gruffLatinAmericanSpanishPhrase,
		language.Russian:              gruffRussianPhrase,
		language.SimplifiedChinese:    gruffSimplifiedChinesePhrase,
		language.Spanish:              gruffSpanishPhrase,
		language.TraditionalChinese:   gruffTraditionalChinesePhrase}
)

var (
	// Gruff represents gruff.
	Gruff = nook.Villager{
		Character:   gruffCharacter,
		Personality: personality.Cranky,
		Phrase:      gruffPhrase}
)

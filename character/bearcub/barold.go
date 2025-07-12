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
	// baroldBirthday represents barold birthday.
	baroldBirthday = nook.Birthday{
		Day:   2,
		Month: time.March}
)

var (
	// baroldCode represents barold code.
	baroldCode = nook.Code{
		Value: "cbr16"}
)

var (
	// baroldAmericanEnglishName represents barold american english name.
	baroldAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Barold"}

	// baroldCanadianFrenchName represents barold canadian french name.
	baroldCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Manu"}

	// baroldDutchName represents barold dutch name.
	baroldDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Barold"}

	// baroldFrenchName represents barold french name.
	baroldFrenchName = nook.Name{
		Language: language.French,
		Value:    "Manu"}

	// baroldGermanName represents barold german name.
	baroldGermanName = nook.Name{
		Language: language.German,
		Value:    "Hubert"}

	// baroldItalianName represents barold italian name.
	baroldItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Placido"}

	// baroldJapaneseName represents barold japanese name.
	baroldJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ニッシー"}

	// baroldKoreanName represents barold korean name.
	baroldKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "곰시"}

	// baroldLatinAmericanSpanishName represents barold latin american spanish name.
	baroldLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Plácido"}

	// baroldRussianName represents barold russian name.
	baroldRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Барольд"}

	// baroldSimplifiedChineseName represents barold simplified chinese name.
	baroldSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿西"}

	// baroldSpanishName represents barold spanish name.
	baroldSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Plácido"}

	// baroldTraditionalChineseName represents barold traditional chinese name.
	baroldTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿西"}
)

var (
	// baroldName represents barold name.
	baroldName = nook.Languages{
		language.AmericanEnglish:      baroldAmericanEnglishName,
		language.CanadianFrench:       baroldCanadianFrenchName,
		language.Dutch:                baroldDutchName,
		language.French:               baroldFrenchName,
		language.German:               baroldGermanName,
		language.Italian:              baroldItalianName,
		language.Japanese:             baroldJapaneseName,
		language.Korean:               baroldKoreanName,
		language.LatinAmericanSpanish: baroldLatinAmericanSpanishName,
		language.Russian:              baroldRussianName,
		language.SimplifiedChinese:    baroldSimplifiedChineseName,
		language.Spanish:              baroldSpanishName,
		language.TraditionalChinese:   baroldTraditionalChineseName}
)

var (
	// baroldCharacter represents barold character.
	baroldCharacter = nook.Character{
		Animal:   animal.Bearcub,
		Birthday: baroldBirthday,
		Code:     baroldCode,
		Key:      character.Barold,
		Gender:   gender.Male,
		Name:     baroldName,
		Special:  false}
)

var (
	// baroldAmericanEnglishPhrase represents barold american english phrase.
	baroldAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cubby"}

	// baroldCanadianFrenchPhrase represents barold canadian french phrase.
	baroldCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "dodododu"}

	// baroldDutchPhrase represents barold dutch phrase.
	baroldDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "welpje"}

	// baroldFrenchPhrase represents barold french phrase.
	baroldFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "dodododu"}

	// baroldGermanPhrase represents barold german phrase.
	baroldGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mwuuu"}

	// baroldItalianPhrase represents barold italian phrase.
	baroldItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "babalù"}

	// baroldJapanesePhrase represents barold japanese phrase.
	baroldJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "いっそ"}

	// baroldKoreanPhrase represents barold korean phrase.
	baroldKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아니그냥"}

	// baroldLatinAmericanSpanishPhrase represents barold latin american spanish phrase.
	baroldLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "apurri"}

	// baroldRussianPhrase represents barold russian phrase.
	baroldRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "берлога"}

	// baroldSimplifiedChinesePhrase represents barold simplified chinese phrase.
	baroldSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "算了"}

	// baroldSpanishPhrase represents barold spanish phrase.
	baroldSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "apurri"}

	// baroldTraditionalChinesePhrase represents barold traditional chinese phrase.
	baroldTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "算了"}
)

var (
	// baroldPhrase represents barold phrase.
	baroldPhrase = nook.Languages{
		language.AmericanEnglish:      baroldAmericanEnglishPhrase,
		language.CanadianFrench:       baroldCanadianFrenchPhrase,
		language.Dutch:                baroldDutchPhrase,
		language.French:               baroldFrenchPhrase,
		language.German:               baroldGermanPhrase,
		language.Italian:              baroldItalianPhrase,
		language.Japanese:             baroldJapanesePhrase,
		language.Korean:               baroldKoreanPhrase,
		language.LatinAmericanSpanish: baroldLatinAmericanSpanishPhrase,
		language.Russian:              baroldRussianPhrase,
		language.SimplifiedChinese:    baroldSimplifiedChinesePhrase,
		language.Spanish:              baroldSpanishPhrase,
		language.TraditionalChinese:   baroldTraditionalChinesePhrase}
)

var (
	// Barold represents barold.
	Barold = nook.Villager{
		Character:   baroldCharacter,
		Personality: personality.Lazy,
		Phrase:      baroldPhrase}
)

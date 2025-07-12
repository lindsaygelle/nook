package lion

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
	// lionelBirthday represents lionel birthday.
	lionelBirthday = nook.Birthday{
		Day:   29,
		Month: time.July}
)

var (
	// lionelCode represents lionel code.
	lionelCode = nook.Code{
		Value: "lon08"}
)

var (
	// lionelAmericanEnglishName represents lionel american english name.
	lionelAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Lionel"}

	// lionelCanadianFrenchName represents lionel canadian french name.
	lionelCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Léopold"}

	// lionelDutchName represents lionel dutch name.
	lionelDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Lionel"}

	// lionelFrenchName represents lionel french name.
	lionelFrenchName = nook.Name{
		Language: language.French,
		Value:    "Léopold"}

	// lionelGermanName represents lionel german name.
	lionelGermanName = nook.Name{
		Language: language.German,
		Value:    "Lorenz"}

	// lionelItalianName represents lionel italian name.
	lionelItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Leonida"}

	// lionelJapaneseName represents lionel japanese name.
	lionelJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ライオネル"}

	// lionelKoreanName represents lionel korean name.
	lionelKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "라이오넬"}

	// lionelLatinAmericanSpanishName represents lionel latin american spanish name.
	lionelLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lionel"}

	// lionelRussianName represents lionel russian name.
	lionelRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лайонел"}

	// lionelSimplifiedChineseName represents lionel simplified chinese name.
	lionelSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "赖恩睿"}

	// lionelSpanishName represents lionel spanish name.
	lionelSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lionel"}

	// lionelTraditionalChineseName represents lionel traditional chinese name.
	lionelTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "賴恩睿"}
)

var (
	// lionelName represents lionel name.
	lionelName = nook.Languages{
		language.AmericanEnglish:      lionelAmericanEnglishName,
		language.CanadianFrench:       lionelCanadianFrenchName,
		language.Dutch:                lionelDutchName,
		language.French:               lionelFrenchName,
		language.German:               lionelGermanName,
		language.Italian:              lionelItalianName,
		language.Japanese:             lionelJapaneseName,
		language.Korean:               lionelKoreanName,
		language.LatinAmericanSpanish: lionelLatinAmericanSpanishName,
		language.Russian:              lionelRussianName,
		language.SimplifiedChinese:    lionelSimplifiedChineseName,
		language.Spanish:              lionelSpanishName,
		language.TraditionalChinese:   lionelTraditionalChineseName}
)

var (
	// lionelCharacter represents lionel character.
	lionelCharacter = nook.Character{
		Animal:   animal.Lion,
		Birthday: lionelBirthday,
		Code:     lionelCode,
		Key:      character.Lionel,
		Gender:   gender.Male,
		Name:     lionelName,
		Special:  false}
)

var (
	// lionelAmericanEnglishPhrase represents lionel american english phrase.
	lionelAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "precisely"}

	// lionelCanadianFrenchPhrase represents lionel canadian french phrase.
	lionelCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "papatte"}

	// lionelDutchPhrase represents lionel dutch phrase.
	lionelDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "exact"}

	// lionelFrenchPhrase represents lionel french phrase.
	lionelFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "papatte"}

	// lionelGermanPhrase represents lionel german phrase.
	lionelGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "grrroooh"}

	// lionelItalianPhrase represents lionel italian phrase.
	lionelItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ahora"}

	// lionelJapanesePhrase represents lionel japanese phrase.
	lionelJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "まさしく"}

	// lionelKoreanPhrase represents lionel korean phrase.
	lionelKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "옳거니"}

	// lionelLatinAmericanSpanishPhrase represents lionel latin american spanish phrase.
	lionelLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "picoplás"}

	// lionelRussianPhrase represents lionel russian phrase.
	lionelRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "честь имею"}

	// lionelSimplifiedChinesePhrase represents lionel simplified chinese phrase.
	lionelSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "正确"}

	// lionelSpanishPhrase represents lionel spanish phrase.
	lionelSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "picoplás"}

	// lionelTraditionalChinesePhrase represents lionel traditional chinese phrase.
	lionelTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "正確"}
)

var (
	// lionelPhrase represents lionel phrase.
	lionelPhrase = nook.Languages{
		language.AmericanEnglish:      lionelAmericanEnglishPhrase,
		language.CanadianFrench:       lionelCanadianFrenchPhrase,
		language.Dutch:                lionelDutchPhrase,
		language.French:               lionelFrenchPhrase,
		language.German:               lionelGermanPhrase,
		language.Italian:              lionelItalianPhrase,
		language.Japanese:             lionelJapanesePhrase,
		language.Korean:               lionelKoreanPhrase,
		language.LatinAmericanSpanish: lionelLatinAmericanSpanishPhrase,
		language.Russian:              lionelRussianPhrase,
		language.SimplifiedChinese:    lionelSimplifiedChinesePhrase,
		language.Spanish:              lionelSpanishPhrase,
		language.TraditionalChinese:   lionelTraditionalChinesePhrase}
)

var (
	// Lionel represents lionel.
	Lionel = nook.Villager{
		Character:   lionelCharacter,
		Personality: personality.Smug,
		Phrase:      lionelPhrase}
)

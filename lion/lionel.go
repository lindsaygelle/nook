package lion

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	lionelBirthday = nook.Birthday{
		Day:   29,
		Month: time.July}
)

var (
	lionelCode = nook.Code{
		Value: "lon08"}
)

var (
	lionelAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Lionel"}

	lionelCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Léopoldpapatte"}

	lionelDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Lionelexact"}

	lionelFrenchName = nook.Name{
		Language: language.French,
		Value:    "Léopoldpapatte"}

	lionelGermanName = nook.Name{
		Language: language.German,
		Value:    "Lorenzgrrroooh"}

	lionelItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Leonidaahora"}

	lionelJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ライオネルまさしく"}

	lionelKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "라이오넬옳거니"}

	lionelLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lionelpicoplás"}

	lionelRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лайонелчесть имею"}

	lionelSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "赖恩睿正确"}

	lionelSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lionelpicoplás"}

	lionelTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "賴恩睿正確"}
)

var (
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
	lionelCharacter = nook.Character{
		Animal:   Lion,
		Birthday: lionelBirthday,
		Code:     lionelCode,
		Gender:   nook.Male,
		Name:     lionelName}
)

var (
	lionelAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "precisely"}

	lionelCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "papatte"}

	lionelDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "exact"}

	lionelFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "papatte"}

	lionelGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "grrroooh"}

	lionelItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ahora"}

	lionelJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "まさしく"}

	lionelKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "옳거니"}

	lionelLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "picoplás"}

	lionelRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "честь имею"}

	lionelSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "正确"}

	lionelSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "picoplás"}

	lionelTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "正確"}
)

var (
	lionelPhrase = nook.Languages{
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
	Lionel = nook.Villager{
		Character:   lionelCharacter,
		Personality: nook.Smug,
		Phrase:      lionelPhrase}
)

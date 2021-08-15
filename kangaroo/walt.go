package kangaroo

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	waltBirthday = nook.Birthday{
		Day:   24,
		Month: time.April}
)

var (
	waltCode = nook.Code{
		Value: "kgr08"}
)

var (
	waltAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Walt"}

	waltCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Waltpopoche"}

	waltDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Waltzakkies"}

	waltFrenchName = nook.Name{
		Language: language.French,
		Value:    "Waltpopoche"}

	waltGermanName = nook.Name{
		Language: language.German,
		Value:    "Sinanboingboing"}

	waltItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Zompoboing"}

	waltJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "カンロクじゃのう"}

	waltKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "관록어험"}

	waltLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Brincoboing"}

	waltRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Уолткарманчик"}

	waltSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "袋钢别装了"}

	waltSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Brincoboing"}

	waltTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "袋鋼別裝了"}
)

var (
	waltName = nook.Languages{
		language.AmericanEnglish:      waltAmericanEnglishName,
		language.CanadianFrench:       waltCanadianFrenchName,
		language.Dutch:                waltDutchName,
		language.French:               waltFrenchName,
		language.German:               waltGermanName,
		language.Italian:              waltItalianName,
		language.Japanese:             waltJapaneseName,
		language.Korean:               waltKoreanName,
		language.LatinAmericanSpanish: waltLatinAmericanSpanishName,
		language.Russian:              waltRussianName,
		language.SimplifiedChinese:    waltSimplifiedChineseName,
		language.Spanish:              waltSpanishName,
		language.TraditionalChinese:   waltTraditionalChineseName}
)

var (
	waltCharacter = nook.Character{
		Animal:   Kangaroo,
		Birthday: waltBirthday,
		Code:     waltCode,
		Gender:   nook.Male,
		Name:     waltName}
)

var (
	waltAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "pockets"}

	waltCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "popoche"}

	waltDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "zakkies"}

	waltFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "popoche"}

	waltGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "boingboing"}

	waltItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "boing"}

	waltJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "じゃのう"}

	waltKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "어험"}

	waltLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "boing"}

	waltRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "карманчик"}

	waltSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "别装了"}

	waltSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "boing"}

	waltTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "別裝了"}
)

var (
	waltPhrase = nook.Languages{
		language.AmericanEnglish:      waltAmericanEnglishName,
		language.CanadianFrench:       waltCanadianFrenchName,
		language.Dutch:                waltDutchName,
		language.French:               waltFrenchName,
		language.German:               waltGermanName,
		language.Italian:              waltItalianName,
		language.Japanese:             waltJapaneseName,
		language.Korean:               waltKoreanName,
		language.LatinAmericanSpanish: waltLatinAmericanSpanishName,
		language.Russian:              waltRussianName,
		language.SimplifiedChinese:    waltSimplifiedChineseName,
		language.Spanish:              waltSpanishName,
		language.TraditionalChinese:   waltTraditionalChineseName}
)

var (
	Walt = nook.Villager{
		Character:   waltCharacter,
		Personality: nook.Cranky,
		Phrase:      waltPhrase}
)

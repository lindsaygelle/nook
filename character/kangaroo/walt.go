package kangaroo

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
		Value:    "Walt"}

	waltDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Walt"}

	waltFrenchName = nook.Name{
		Language: language.French,
		Value:    "Walt"}

	waltGermanName = nook.Name{
		Language: language.German,
		Value:    "Sinan"}

	waltItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Zompo"}

	waltJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "カンロク"}

	waltKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "관록"}

	waltLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Brinco"}

	waltRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Уолт"}

	waltSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "袋钢"}

	waltSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Brinco"}

	waltTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "袋鋼"}
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
		Animal:   animal.Kangaroo,
		Birthday: waltBirthday,
		Code:     waltCode,
		Key:      character.Walt,
		Gender:   gender.Male,
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
		language.AmericanEnglish:      waltAmericanEnglishPhrase,
		language.CanadianFrench:       waltCanadianFrenchPhrase,
		language.Dutch:                waltDutchPhrase,
		language.French:               waltFrenchPhrase,
		language.German:               waltGermanPhrase,
		language.Italian:              waltItalianPhrase,
		language.Japanese:             waltJapanesePhrase,
		language.Korean:               waltKoreanPhrase,
		language.LatinAmericanSpanish: waltLatinAmericanSpanishPhrase,
		language.Russian:              waltRussianPhrase,
		language.SimplifiedChinese:    waltSimplifiedChinesePhrase,
		language.Spanish:              waltSpanishPhrase,
		language.TraditionalChinese:   waltTraditionalChinesePhrase}
)

var (
	Walt = nook.Villager{
		Character:   waltCharacter,
		Personality: personality.Cranky,
		Phrase:      waltPhrase}
)

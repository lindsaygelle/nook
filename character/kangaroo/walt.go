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
	// waltBirthday represents walt birthday.
	waltBirthday = nook.Birthday{
		Day:   24,
		Month: time.April}
)

var (
	// waltCode represents walt code.
	waltCode = nook.Code{
		Value: "kgr08"}
)

var (
	// waltAmericanEnglishName represents walt american english name.
	waltAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Walt"}

	// waltCanadianFrenchName represents walt canadian french name.
	waltCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Walt"}

	// waltDutchName represents walt dutch name.
	waltDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Walt"}

	// waltFrenchName represents walt french name.
	waltFrenchName = nook.Name{
		Language: language.French,
		Value:    "Walt"}

	// waltGermanName represents walt german name.
	waltGermanName = nook.Name{
		Language: language.German,
		Value:    "Sinan"}

	// waltItalianName represents walt italian name.
	waltItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Zompo"}

	// waltJapaneseName represents walt japanese name.
	waltJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "カンロク"}

	// waltKoreanName represents walt korean name.
	waltKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "관록"}

	// waltLatinAmericanSpanishName represents walt latin american spanish name.
	waltLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Brinco"}

	// waltRussianName represents walt russian name.
	waltRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Уолт"}

	// waltSimplifiedChineseName represents walt simplified chinese name.
	waltSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "袋钢"}

	// waltSpanishName represents walt spanish name.
	waltSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Brinco"}

	// waltTraditionalChineseName represents walt traditional chinese name.
	waltTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "袋鋼"}
)

var (
	// waltName represents walt name.
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
	// waltCharacter represents walt character.
	waltCharacter = nook.Character{
		Animal:   animal.Kangaroo,
		Birthday: waltBirthday,
		Code:     waltCode,
		Key:      character.Walt,
		Gender:   gender.Male,
		Name:     waltName,
		Special:  false}
)

var (
	// waltAmericanEnglishPhrase represents walt american english phrase.
	waltAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "pockets"}

	// waltCanadianFrenchPhrase represents walt canadian french phrase.
	waltCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "popoche"}

	// waltDutchPhrase represents walt dutch phrase.
	waltDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "zakkies"}

	// waltFrenchPhrase represents walt french phrase.
	waltFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "popoche"}

	// waltGermanPhrase represents walt german phrase.
	waltGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "boingboing"}

	// waltItalianPhrase represents walt italian phrase.
	waltItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "boing"}

	// waltJapanesePhrase represents walt japanese phrase.
	waltJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "じゃのう"}

	// waltKoreanPhrase represents walt korean phrase.
	waltKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "어험"}

	// waltLatinAmericanSpanishPhrase represents walt latin american spanish phrase.
	waltLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "boing"}

	// waltRussianPhrase represents walt russian phrase.
	waltRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "карманчик"}

	// waltSimplifiedChinesePhrase represents walt simplified chinese phrase.
	waltSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "别装了"}

	// waltSpanishPhrase represents walt spanish phrase.
	waltSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "boing"}

	// waltTraditionalChinesePhrase represents walt traditional chinese phrase.
	waltTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "別裝了"}
)

var (
	// waltPhrase represents walt phrase.
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
	// Walt represents walt.
	Walt = nook.Villager{
		Character:   waltCharacter,
		Personality: personality.Cranky,
		Phrase:      waltPhrase}
)

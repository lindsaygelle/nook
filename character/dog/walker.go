package dog

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
	// walkerBirthday represents walker birthday.
	walkerBirthday = nook.Birthday{
		Day:   10,
		Month: time.June}
)

var (
	// walkerCode represents walker code.
	walkerCode = nook.Code{
		Value: "dog06"}
)

var (
	// walkerAmericanEnglishName represents walker american english name.
	walkerAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Walker"}

	// walkerCanadianFrenchName represents walker canadian french name.
	walkerCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "George"}

	// walkerDutchName represents walker dutch name.
	walkerDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Walker"}

	// walkerFrenchName represents walker french name.
	walkerFrenchName = nook.Name{
		Language: language.French,
		Value:    "George"}

	// walkerGermanName represents walker german name.
	walkerGermanName = nook.Name{
		Language: language.German,
		Value:    "Fido"}

	// walkerItalianName represents walker italian name.
	walkerItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Walter"}

	// walkerJapaneseName represents walker japanese name.
	walkerJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ベン"}

	// walkerKoreanName represents walker korean name.
	walkerKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "벤"}

	// walkerLatinAmericanSpanishName represents walker latin american spanish name.
	walkerLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Miki"}

	// walkerRussianName represents walker russian name.
	walkerRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Уокер"}

	// walkerSimplifiedChineseName represents walker simplified chinese name.
	walkerSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿笨"}

	// walkerSpanishName represents walker spanish name.
	walkerSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Miki"}

	// walkerTraditionalChineseName represents walker traditional chinese name.
	walkerTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿笨"}
)

var (
	// walkerName represents walker name.
	walkerName = nook.Languages{
		language.AmericanEnglish:      walkerAmericanEnglishName,
		language.CanadianFrench:       walkerCanadianFrenchName,
		language.Dutch:                walkerDutchName,
		language.French:               walkerFrenchName,
		language.German:               walkerGermanName,
		language.Italian:              walkerItalianName,
		language.Japanese:             walkerJapaneseName,
		language.Korean:               walkerKoreanName,
		language.LatinAmericanSpanish: walkerLatinAmericanSpanishName,
		language.Russian:              walkerRussianName,
		language.SimplifiedChinese:    walkerSimplifiedChineseName,
		language.Spanish:              walkerSpanishName,
		language.TraditionalChinese:   walkerTraditionalChineseName}
)

var (
	// walkerCharacter represents walker character.
	walkerCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: walkerBirthday,
		Code:     walkerCode,
		Key:      character.Walker,
		Gender:   gender.Male,
		Name:     walkerName,
		Special:  false}
)

var (
	// walkerAmericanEnglishPhrase represents walker american english phrase.
	walkerAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "wuh"}

	// walkerCanadianFrenchPhrase represents walker canadian french phrase.
	walkerCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ouh"}

	// walkerDutchPhrase represents walker dutch phrase.
	walkerDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "waf"}

	// walkerFrenchPhrase represents walker french phrase.
	walkerFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ouh"}

	// walkerGermanPhrase represents walker german phrase.
	walkerGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "wafff"}

	// walkerItalianPhrase represents walker italian phrase.
	walkerItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "wuff"}

	// walkerJapanesePhrase represents walker japanese phrase.
	walkerJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "バウ"}

	// walkerKoreanPhrase represents walker korean phrase.
	walkerKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "컹컹"}

	// walkerLatinAmericanSpanishPhrase represents walker latin american spanish phrase.
	walkerLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "guuu-ah"}

	// walkerRussianPhrase represents walker russian phrase.
	walkerRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ау-у"}

	// walkerSimplifiedChinesePhrase represents walker simplified chinese phrase.
	walkerSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咆"}

	// walkerSpanishPhrase represents walker spanish phrase.
	walkerSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "huesitos"}

	// walkerTraditionalChinesePhrase represents walker traditional chinese phrase.
	walkerTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咆"}
)

var (
	// walkerPhrase represents walker phrase.
	walkerPhrase = nook.Languages{
		language.AmericanEnglish:      walkerAmericanEnglishPhrase,
		language.CanadianFrench:       walkerCanadianFrenchPhrase,
		language.Dutch:                walkerDutchPhrase,
		language.French:               walkerFrenchPhrase,
		language.German:               walkerGermanPhrase,
		language.Italian:              walkerItalianPhrase,
		language.Japanese:             walkerJapanesePhrase,
		language.Korean:               walkerKoreanPhrase,
		language.LatinAmericanSpanish: walkerLatinAmericanSpanishPhrase,
		language.Russian:              walkerRussianPhrase,
		language.SimplifiedChinese:    walkerSimplifiedChinesePhrase,
		language.Spanish:              walkerSpanishPhrase,
		language.TraditionalChinese:   walkerTraditionalChinesePhrase}
)

var (
	// Walker represents walker.
	Walker = nook.Villager{
		Character:   walkerCharacter,
		Personality: personality.Lazy,
		Phrase:      walkerPhrase}
)

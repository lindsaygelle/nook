package cat

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
	// kikiBirthday represents kiki birthday.
	kikiBirthday = nook.Birthday{
		Day:   8,
		Month: time.October}
)

var (
	// kikiCode represents kiki code.
	kikiCode = nook.Code{
		Value: "cat04"}
)

var (
	// kikiAmericanEnglishName represents kiki american english name.
	kikiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Kiki"}

	// kikiCanadianFrenchName represents kiki canadian french name.
	kikiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Kiki"}

	// kikiDutchName represents kiki dutch name.
	kikiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Kiki"}

	// kikiFrenchName represents kiki french name.
	kikiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Kiki"}

	// kikiGermanName represents kiki german name.
	kikiGermanName = nook.Name{
		Language: language.German,
		Value:    "Kiki"}

	// kikiItalianName represents kiki italian name.
	kikiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Kiki"}

	// kikiJapaneseName represents kiki japanese name.
	kikiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "キャビア"}

	// kikiKoreanName represents kiki korean name.
	kikiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "캐비어"}

	// kikiLatinAmericanSpanishName represents kiki latin american spanish name.
	kikiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ágata"}

	// kikiRussianName represents kiki russian name.
	kikiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кики"}

	// kikiSimplifiedChineseName represents kiki simplified chinese name.
	kikiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "余子酱"}

	// kikiSpanishName represents kiki spanish name.
	kikiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ágata"}

	// kikiTraditionalChineseName represents kiki traditional chinese name.
	kikiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "余子醬"}
)

var (
	// kikiName represents kiki name.
	kikiName = nook.Languages{
		language.AmericanEnglish:      kikiAmericanEnglishName,
		language.CanadianFrench:       kikiCanadianFrenchName,
		language.Dutch:                kikiDutchName,
		language.French:               kikiFrenchName,
		language.German:               kikiGermanName,
		language.Italian:              kikiItalianName,
		language.Japanese:             kikiJapaneseName,
		language.Korean:               kikiKoreanName,
		language.LatinAmericanSpanish: kikiLatinAmericanSpanishName,
		language.Russian:              kikiRussianName,
		language.SimplifiedChinese:    kikiSimplifiedChineseName,
		language.Spanish:              kikiSpanishName,
		language.TraditionalChinese:   kikiTraditionalChineseName}
)

var (
	// kikiCharacter represents kiki character.
	kikiCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: kikiBirthday,
		Code:     kikiCode,
		Key:      character.Kiki,
		Gender:   gender.Female,
		Name:     kikiName,
		Special:  false}
)

var (
	// kikiAmericanEnglishPhrase represents kiki american english phrase.
	kikiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "kitty cat"}

	// kikiCanadianFrenchPhrase represents kiki canadian french phrase.
	kikiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "chti minou"}

	// kikiDutchPhrase represents kiki dutch phrase.
	kikiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "poesjemauw"}

	// kikiFrenchPhrase represents kiki french phrase.
	kikiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "chti minou"}

	// kikiGermanPhrase represents kiki german phrase.
	kikiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "miausi"}

	// kikiItalianPhrase represents kiki italian phrase.
	kikiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "mieaoo"}

	// kikiJapanesePhrase represents kiki japanese phrase.
	kikiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だニ"}

	// kikiKoreanPhrase represents kiki korean phrase.
	kikiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "뭐라지요"}

	// kikiLatinAmericanSpanishPhrase represents kiki latin american spanish phrase.
	kikiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "michifiú"}

	// kikiRussianPhrase represents kiki russian phrase.
	kikiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кисуля"}

	// kikiSimplifiedChinesePhrase represents kiki simplified chinese phrase.
	kikiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "酱酱"}

	// kikiSpanishPhrase represents kiki spanish phrase.
	kikiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "galletita"}

	// kikiTraditionalChinesePhrase represents kiki traditional chinese phrase.
	kikiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "醬醬"}
)

var (
	// kikiPhrase represents kiki phrase.
	kikiPhrase = nook.Languages{
		language.AmericanEnglish:      kikiAmericanEnglishPhrase,
		language.CanadianFrench:       kikiCanadianFrenchPhrase,
		language.Dutch:                kikiDutchPhrase,
		language.French:               kikiFrenchPhrase,
		language.German:               kikiGermanPhrase,
		language.Italian:              kikiItalianPhrase,
		language.Japanese:             kikiJapanesePhrase,
		language.Korean:               kikiKoreanPhrase,
		language.LatinAmericanSpanish: kikiLatinAmericanSpanishPhrase,
		language.Russian:              kikiRussianPhrase,
		language.SimplifiedChinese:    kikiSimplifiedChinesePhrase,
		language.Spanish:              kikiSpanishPhrase,
		language.TraditionalChinese:   kikiTraditionalChinesePhrase}
)

var (
	// Kiki represents kiki.
	Kiki = nook.Villager{
		Character:   kikiCharacter,
		Personality: personality.Normal,
		Phrase:      kikiPhrase}
)

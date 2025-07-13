package duck

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
	// ketchupBirthday represents ketchup birthday.
	ketchupBirthday = nook.Birthday{
		Day:   27,
		Month: time.July}
)

var (
	// ketchupCode represents ketchup code.
	ketchupCode = nook.Code{
		Value: "duk13"}
)

var (
	// ketchupAmericanEnglishName represents ketchup american english name.
	ketchupAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Ketchup"}

	// ketchupCanadianFrenchName represents ketchup canadian french name.
	ketchupCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Ketchup"}

	// ketchupDutchName represents ketchup dutch name.
	ketchupDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Ketchup"}

	// ketchupFrenchName represents ketchup french name.
	ketchupFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ketchup"}

	// ketchupGermanName represents ketchup german name.
	ketchupGermanName = nook.Name{
		Language: language.German,
		Value:    "Pullunda"}

	// ketchupItalianName represents ketchup italian name.
	ketchupItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ketchup"}

	// ketchupJapaneseName represents ketchup japanese name.
	ketchupJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ケチャップ"}

	// ketchupKoreanName represents ketchup korean name.
	ketchupKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "케첩"}

	// ketchupLatinAmericanSpanishName represents ketchup latin american spanish name.
	ketchupLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Kétchup"}

	// ketchupRussianName represents ketchup russian name.
	ketchupRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кетчуп"}

	// ketchupSimplifiedChineseName represents ketchup simplified chinese name.
	ketchupSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "番茄酱"}

	// ketchupSpanishName represents ketchup spanish name.
	ketchupSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Kétchup"}

	// ketchupTraditionalChineseName represents ketchup traditional chinese name.
	ketchupTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "番茄醬"}
)

var (
	// ketchupName represents ketchup name.
	ketchupName = nook.Languages{
		language.AmericanEnglish:      ketchupAmericanEnglishName,
		language.CanadianFrench:       ketchupCanadianFrenchName,
		language.Dutch:                ketchupDutchName,
		language.French:               ketchupFrenchName,
		language.German:               ketchupGermanName,
		language.Italian:              ketchupItalianName,
		language.Japanese:             ketchupJapaneseName,
		language.Korean:               ketchupKoreanName,
		language.LatinAmericanSpanish: ketchupLatinAmericanSpanishName,
		language.Russian:              ketchupRussianName,
		language.SimplifiedChinese:    ketchupSimplifiedChineseName,
		language.Spanish:              ketchupSpanishName,
		language.TraditionalChinese:   ketchupTraditionalChineseName}
)

var (
	// ketchupCharacter represents ketchup character.
	ketchupCharacter = nook.Character{
		Animal:   animal.Duck,
		Birthday: ketchupBirthday,
		Code:     ketchupCode,
		Key:      character.Ketchup,
		Gender:   gender.Female,
		Name:     ketchupName,
		Special:  false}
)

var (
	// ketchupAmericanEnglishPhrase represents ketchup american english phrase.
	ketchupAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bitty"}

	// ketchupCanadianFrenchPhrase represents ketchup canadian french phrase.
	ketchupCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "plouf"}

	// ketchupDutchPhrase represents ketchup dutch phrase.
	ketchupDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "tomaatje"}

	// ketchupFrenchPhrase represents ketchup french phrase.
	ketchupFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "plouf"}

	// ketchupGermanPhrase represents ketchup german phrase.
	ketchupGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "pullilol"}

	// ketchupItalianPhrase represents ketchup italian phrase.
	ketchupItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "quaraquack"}

	// ketchupJapanesePhrase represents ketchup japanese phrase.
	ketchupJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "プチ"}

	// ketchupKoreanPhrase represents ketchup korean phrase.
	ketchupKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "찌익"}

	// ketchupLatinAmericanSpanishPhrase represents ketchup latin american spanish phrase.
	ketchupLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cuic-cuic"}

	// ketchupRussianPhrase represents ketchup russian phrase.
	ketchupRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "покрякушка"}

	// ketchupSimplifiedChinesePhrase represents ketchup simplified chinese phrase.
	ketchupSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "噗吱"}

	// ketchupSpanishPhrase represents ketchup spanish phrase.
	ketchupSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cuic-cuic"}

	// ketchupTraditionalChinesePhrase represents ketchup traditional chinese phrase.
	ketchupTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "噗吱"}
)

var (
	// ketchupPhrase represents ketchup phrase.
	ketchupPhrase = nook.Languages{
		language.AmericanEnglish:      ketchupAmericanEnglishPhrase,
		language.CanadianFrench:       ketchupCanadianFrenchPhrase,
		language.Dutch:                ketchupDutchPhrase,
		language.French:               ketchupFrenchPhrase,
		language.German:               ketchupGermanPhrase,
		language.Italian:              ketchupItalianPhrase,
		language.Japanese:             ketchupJapanesePhrase,
		language.Korean:               ketchupKoreanPhrase,
		language.LatinAmericanSpanish: ketchupLatinAmericanSpanishPhrase,
		language.Russian:              ketchupRussianPhrase,
		language.SimplifiedChinese:    ketchupSimplifiedChinesePhrase,
		language.Spanish:              ketchupSpanishPhrase,
		language.TraditionalChinese:   ketchupTraditionalChinesePhrase}
)

var (
	// Ketchup represents ketchup.
	Ketchup = nook.Villager{
		Character:   ketchupCharacter,
		Personality: personality.Peppy,
		Phrase:      ketchupPhrase}
)

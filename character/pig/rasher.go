package pig

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
	// rasherBirthday represents rasher birthday.
	rasherBirthday = nook.Birthday{
		Day:   7,
		Month: time.April}
)

var (
	// rasherCode represents rasher code.
	rasherCode = nook.Code{
		Value: "pig02"}
)

var (
	// rasherAmericanEnglishName represents rasher american english name.
	rasherAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rasher"}

	// rasherCanadianFrenchName represents rasher canadian french name.
	rasherCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Salami"}

	// rasherDutchName represents rasher dutch name.
	rasherDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rasher"}

	// rasherFrenchName represents rasher french name.
	rasherFrenchName = nook.Name{
		Language: language.French,
		Value:    "Salami"}

	// rasherGermanName represents rasher german name.
	rasherGermanName = nook.Name{
		Language: language.German,
		Value:    "Ede"}

	// rasherItalianName represents rasher italian name.
	rasherItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Broncio"}

	// rasherJapaneseName represents rasher japanese name.
	rasherJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "グレオ"}

	// rasherKoreanName represents rasher korean name.
	rasherKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "글레이"}

	// rasherLatinAmericanSpanishName represents rasher latin american spanish name.
	rasherLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Curtis"}

	// rasherRussianName represents rasher russian name.
	rasherRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рашер"}

	// rasherSimplifiedChineseName represents rasher simplified chinese name.
	rasherSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "古烈"}

	// rasherSpanishName represents rasher spanish name.
	rasherSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Curtis"}

	// rasherTraditionalChineseName represents rasher traditional chinese name.
	rasherTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "古烈"}
)

var (
	// rasherName represents rasher name.
	rasherName = nook.Languages{
		language.AmericanEnglish:      rasherAmericanEnglishName,
		language.CanadianFrench:       rasherCanadianFrenchName,
		language.Dutch:                rasherDutchName,
		language.French:               rasherFrenchName,
		language.German:               rasherGermanName,
		language.Italian:              rasherItalianName,
		language.Japanese:             rasherJapaneseName,
		language.Korean:               rasherKoreanName,
		language.LatinAmericanSpanish: rasherLatinAmericanSpanishName,
		language.Russian:              rasherRussianName,
		language.SimplifiedChinese:    rasherSimplifiedChineseName,
		language.Spanish:              rasherSpanishName,
		language.TraditionalChinese:   rasherTraditionalChineseName}
)

var (
	// rasherCharacter represents rasher character.
	rasherCharacter = nook.Character{
		Animal:   animal.Pig,
		Birthday: rasherBirthday,
		Code:     rasherCode,
		Key:      character.Rasher,
		Gender:   gender.Male,
		Name:     rasherName,
		Special:  false}
)

var (
	// rasherAmericanEnglishPhrase represents rasher american english phrase.
	rasherAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "swine"}

	// rasherCanadianFrenchPhrase represents rasher canadian french phrase.
	rasherCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "porcelet"}

	// rasherDutchPhrase represents rasher dutch phrase.
	rasherDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "big"}

	// rasherFrenchPhrase represents rasher french phrase.
	rasherFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "porcelet"}

	// rasherGermanPhrase represents rasher german phrase.
	rasherGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "gronk"}

	// rasherItalianPhrase represents rasher italian phrase.
	rasherItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "snoooink"}

	// rasherJapanesePhrase represents rasher japanese phrase.
	rasherJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "まんねん"}

	// rasherKoreanPhrase represents rasher korean phrase.
	rasherKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "꾸엑"}

	// rasherLatinAmericanSpanishPhrase represents rasher latin american spanish phrase.
	rasherLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "esnoink"}

	// rasherRussianPhrase represents rasher russian phrase.
	rasherRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хряк"}

	// rasherSimplifiedChinesePhrase represents rasher simplified chinese phrase.
	rasherSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "万年"}

	// rasherSpanishPhrase represents rasher spanish phrase.
	rasherSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "porcino"}

	// rasherTraditionalChinesePhrase represents rasher traditional chinese phrase.
	rasherTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "萬年"}
)

var (
	// rasherPhrase represents rasher phrase.
	rasherPhrase = nook.Languages{
		language.AmericanEnglish:      rasherAmericanEnglishPhrase,
		language.CanadianFrench:       rasherCanadianFrenchPhrase,
		language.Dutch:                rasherDutchPhrase,
		language.French:               rasherFrenchPhrase,
		language.German:               rasherGermanPhrase,
		language.Italian:              rasherItalianPhrase,
		language.Japanese:             rasherJapanesePhrase,
		language.Korean:               rasherKoreanPhrase,
		language.LatinAmericanSpanish: rasherLatinAmericanSpanishPhrase,
		language.Russian:              rasherRussianPhrase,
		language.SimplifiedChinese:    rasherSimplifiedChinesePhrase,
		language.Spanish:              rasherSpanishPhrase,
		language.TraditionalChinese:   rasherTraditionalChinesePhrase}
)

var (
	// Rasher represents rasher.
	Rasher = nook.Villager{
		Character:   rasherCharacter,
		Personality: personality.Cranky,
		Phrase:      rasherPhrase}
)

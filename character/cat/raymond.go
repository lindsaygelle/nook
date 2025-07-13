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
	// raymondBirthday represents raymond birthday.
	raymondBirthday = nook.Birthday{
		Day:   1,
		Month: time.October}
)

var (
	// raymondCode represents raymond code.
	raymondCode = nook.Code{
		Value: "cat23"}
)

var (
	// raymondAmericanEnglishName represents raymond american english name.
	raymondAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Raymond"}

	// raymondCanadianFrenchName represents raymond canadian french name.
	raymondCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Raymond"}

	// raymondDutchName represents raymond dutch name.
	raymondDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Raymond"}

	// raymondFrenchName represents raymond french name.
	raymondFrenchName = nook.Name{
		Language: language.French,
		Value:    "Raymond"}

	// raymondGermanName represents raymond german name.
	raymondGermanName = nook.Name{
		Language: language.German,
		Value:    "Gunnar"}

	// raymondItalianName represents raymond italian name.
	raymondItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Raimondo"}

	// raymondJapaneseName represents raymond japanese name.
	raymondJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジャック"}

	// raymondKoreanName represents raymond korean name.
	raymondKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "잭슨"}

	// raymondLatinAmericanSpanishName represents raymond latin american spanish name.
	raymondLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Narciso"}

	// raymondRussianName represents raymond russian name.
	raymondRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Реймонд"}

	// raymondSimplifiedChineseName represents raymond simplified chinese name.
	raymondSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "杰克"}

	// raymondSpanishName represents raymond spanish name.
	raymondSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Narciso"}

	// raymondTraditionalChineseName represents raymond traditional chinese name.
	raymondTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "傑克"}
)

var (
	// raymondName represents raymond name.
	raymondName = nook.Languages{
		language.AmericanEnglish:      raymondAmericanEnglishName,
		language.CanadianFrench:       raymondCanadianFrenchName,
		language.Dutch:                raymondDutchName,
		language.French:               raymondFrenchName,
		language.German:               raymondGermanName,
		language.Italian:              raymondItalianName,
		language.Japanese:             raymondJapaneseName,
		language.Korean:               raymondKoreanName,
		language.LatinAmericanSpanish: raymondLatinAmericanSpanishName,
		language.Russian:              raymondRussianName,
		language.SimplifiedChinese:    raymondSimplifiedChineseName,
		language.Spanish:              raymondSpanishName,
		language.TraditionalChinese:   raymondTraditionalChineseName}
)

var (
	// raymondCharacter represents raymond character.
	raymondCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: raymondBirthday,
		Code:     raymondCode,
		Key:      character.Raymond,
		Gender:   gender.Male,
		Name:     raymondName,
		Special:  false}
)

var (
	// raymondAmericanEnglishPhrase represents raymond american english phrase.
	raymondAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "crisp"}

	// raymondCanadianFrenchPhrase represents raymond canadian french phrase.
	raymondCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "panache"}

	// raymondDutchPhrase represents raymond dutch phrase.
	raymondDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "netjes"}

	// raymondFrenchPhrase represents raymond french phrase.
	raymondFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "panache"}

	// raymondGermanPhrase represents raymond german phrase.
	raymondGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "weeßte"}

	// raymondItalianPhrase represents raymond italian phrase.
	raymondItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "rrricooo"}

	// raymondJapanesePhrase represents raymond japanese phrase.
	raymondJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "キリッ"}

	// raymondKoreanPhrase represents raymond korean phrase.
	raymondKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "우쭐"}

	// raymondLatinAmericanSpanishPhrase represents raymond latin american spanish phrase.
	raymondLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "atilda"}

	// raymondRussianPhrase represents raymond russian phrase.
	raymondRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "стильненько"}

	// raymondSimplifiedChinesePhrase represents raymond simplified chinese phrase.
	raymondSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "严肃"}

	// raymondSpanishPhrase represents raymond spanish phrase.
	raymondSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "atilda"}

	// raymondTraditionalChinesePhrase represents raymond traditional chinese phrase.
	raymondTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嚴肅"}
)

var (
	// raymondPhrase represents raymond phrase.
	raymondPhrase = nook.Languages{
		language.AmericanEnglish:      raymondAmericanEnglishPhrase,
		language.CanadianFrench:       raymondCanadianFrenchPhrase,
		language.Dutch:                raymondDutchPhrase,
		language.French:               raymondFrenchPhrase,
		language.German:               raymondGermanPhrase,
		language.Italian:              raymondItalianPhrase,
		language.Japanese:             raymondJapanesePhrase,
		language.Korean:               raymondKoreanPhrase,
		language.LatinAmericanSpanish: raymondLatinAmericanSpanishPhrase,
		language.Russian:              raymondRussianPhrase,
		language.SimplifiedChinese:    raymondSimplifiedChinesePhrase,
		language.Spanish:              raymondSpanishPhrase,
		language.TraditionalChinese:   raymondTraditionalChinesePhrase}
)

var (
	// Raymond represents raymond.
	Raymond = nook.Villager{
		Character:   raymondCharacter,
		Personality: personality.Smug,
		Phrase:      raymondPhrase}
)

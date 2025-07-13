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
	// lucyBirthday represents lucy birthday.
	lucyBirthday = nook.Birthday{
		Day:   2,
		Month: time.June}
)

var (
	// lucyCode represents lucy code.
	lucyCode = nook.Code{
		Value: "pig04"}
)

var (
	// lucyAmericanEnglishName represents lucy american english name.
	lucyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Lucy"}

	// lucyCanadianFrenchName represents lucy canadian french name.
	lucyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Lucie"}

	// lucyDutchName represents lucy dutch name.
	lucyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Lucy"}

	// lucyFrenchName represents lucy french name.
	lucyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Lucie"}

	// lucyGermanName represents lucy german name.
	lucyGermanName = nook.Name{
		Language: language.German,
		Value:    "Larissa"}

	// lucyItalianName represents lucy italian name.
	lucyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lucy"}

	// lucyJapaneseName represents lucy japanese name.
	lucyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ルーシー"}

	// lucyKoreanName represents lucy korean name.
	lucyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "루시"}

	// lucyLatinAmericanSpanishName represents lucy latin american spanish name.
	lucyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Aurelia"}

	// lucyRussianName represents lucy russian name.
	lucyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Люси"}

	// lucySimplifiedChineseName represents lucy simplified chinese name.
	lucySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "露西"}

	// lucySpanishName represents lucy spanish name.
	lucySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Aurelia"}

	// lucyTraditionalChineseName represents lucy traditional chinese name.
	lucyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "露西"}
)

var (
	// lucyName represents lucy name.
	lucyName = nook.Languages{
		language.AmericanEnglish:      lucyAmericanEnglishName,
		language.CanadianFrench:       lucyCanadianFrenchName,
		language.Dutch:                lucyDutchName,
		language.French:               lucyFrenchName,
		language.German:               lucyGermanName,
		language.Italian:              lucyItalianName,
		language.Japanese:             lucyJapaneseName,
		language.Korean:               lucyKoreanName,
		language.LatinAmericanSpanish: lucyLatinAmericanSpanishName,
		language.Russian:              lucyRussianName,
		language.SimplifiedChinese:    lucySimplifiedChineseName,
		language.Spanish:              lucySpanishName,
		language.TraditionalChinese:   lucyTraditionalChineseName}
)

var (
	// lucyCharacter represents lucy character.
	lucyCharacter = nook.Character{
		Animal:   animal.Pig,
		Birthday: lucyBirthday,
		Code:     lucyCode,
		Key:      character.Lucy,
		Gender:   gender.Female,
		Name:     lucyName,
		Special:  false}
)

var (
	// lucyAmericanEnglishPhrase represents lucy american english phrase.
	lucyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "snoooink"}

	// lucyCanadianFrenchPhrase represents lucy canadian french phrase.
	lucyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "groingroin"}

	// lucyDutchPhrase represents lucy dutch phrase.
	lucyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "knorrie"}

	// lucyFrenchPhrase represents lucy french phrase.
	lucyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "groingroin"}

	// lucyGermanPhrase represents lucy german phrase.
	lucyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnoink"}

	// lucyItalianPhrase represents lucy italian phrase.
	lucyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "snoink"}

	// lucyJapanesePhrase represents lucy japanese phrase.
	lucyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "よぅ"}

	// lucyKoreanPhrase represents lucy korean phrase.
	lucyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "예예"}

	// lucyLatinAmericanSpanishPhrase represents lucy latin american spanish phrase.
	lucyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cochín"}

	// lucyRussianPhrase represents lucy russian phrase.
	lucyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хрю-ю-ю"}

	// lucySimplifiedChinesePhrase represents lucy simplified chinese phrase.
	lucySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "唷"}

	// lucySpanishPhrase represents lucy spanish phrase.
	lucySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cochín"}

	// lucyTraditionalChinesePhrase represents lucy traditional chinese phrase.
	lucyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "唷"}
)

var (
	// lucyPhrase represents lucy phrase.
	lucyPhrase = nook.Languages{
		language.AmericanEnglish:      lucyAmericanEnglishPhrase,
		language.CanadianFrench:       lucyCanadianFrenchPhrase,
		language.Dutch:                lucyDutchPhrase,
		language.French:               lucyFrenchPhrase,
		language.German:               lucyGermanPhrase,
		language.Italian:              lucyItalianPhrase,
		language.Japanese:             lucyJapanesePhrase,
		language.Korean:               lucyKoreanPhrase,
		language.LatinAmericanSpanish: lucyLatinAmericanSpanishPhrase,
		language.Russian:              lucyRussianPhrase,
		language.SimplifiedChinese:    lucySimplifiedChinesePhrase,
		language.Spanish:              lucySpanishPhrase,
		language.TraditionalChinese:   lucyTraditionalChinesePhrase}
)

var (
	// Lucy represents lucy.
	Lucy = nook.Villager{
		Character:   lucyCharacter,
		Personality: personality.Normal,
		Phrase:      lucyPhrase}
)

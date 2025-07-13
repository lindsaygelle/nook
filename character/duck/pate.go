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
	// pateBirthday represents pate birthday.
	pateBirthday = nook.Birthday{
		Day:   23,
		Month: time.February}
)

var (
	// pateCode represents pate code.
	pateCode = nook.Code{
		Value: "duk02"}
)

var (
	// pateAmericanEnglishName represents pate american english name.
	pateAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pate"}

	// pateCanadianFrenchName represents pate canadian french name.
	pateCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Terrine"}

	// pateDutchName represents pate dutch name.
	pateDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pate"}

	// pateFrenchName represents pate french name.
	pateFrenchName = nook.Name{
		Language: language.French,
		Value:    "Terrine"}

	// pateGermanName represents pate german name.
	pateGermanName = nook.Name{
		Language: language.German,
		Value:    "Daune"}

	// pateItalianName represents pate italian name.
	pateItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Camilla"}

	// pateJapaneseName represents pate japanese name.
	pateJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ナッキー"}

	// pateKoreanName represents pate korean name.
	pateKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "나키"}

	// pateLatinAmericanSpanishName represents pate latin american spanish name.
	pateLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pati"}

	// pateRussianName represents pate russian name.
	pateRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пэйт"}

	// pateSimplifiedChineseName represents pate simplified chinese name.
	pateSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "爱哭鬼"}

	// pateSpanishName represents pate spanish name.
	pateSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pati"}

	// pateTraditionalChineseName represents pate traditional chinese name.
	pateTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "愛哭鬼"}
)

var (
	// pateName represents pate name.
	pateName = nook.Languages{
		language.AmericanEnglish:      pateAmericanEnglishName,
		language.CanadianFrench:       pateCanadianFrenchName,
		language.Dutch:                pateDutchName,
		language.French:               pateFrenchName,
		language.German:               pateGermanName,
		language.Italian:              pateItalianName,
		language.Japanese:             pateJapaneseName,
		language.Korean:               pateKoreanName,
		language.LatinAmericanSpanish: pateLatinAmericanSpanishName,
		language.Russian:              pateRussianName,
		language.SimplifiedChinese:    pateSimplifiedChineseName,
		language.Spanish:              pateSpanishName,
		language.TraditionalChinese:   pateTraditionalChineseName}
)

var (
	// pateCharacter represents pate character.
	pateCharacter = nook.Character{
		Animal:   animal.Duck,
		Birthday: pateBirthday,
		Code:     pateCode,
		Key:      character.Pate,
		Gender:   gender.Female,
		Name:     pateName,
		Special:  false}
)

var (
	// pateAmericanEnglishPhrase represents pate american english phrase.
	pateAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "quackle"}

	// pateCanadianFrenchPhrase represents pate canadian french phrase.
	pateCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "coin coin"}

	// pateDutchPhrase represents pate dutch phrase.
	pateDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kwakel"}

	// pateFrenchPhrase represents pate french phrase.
	pateFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "coin coin"}

	// pateGermanPhrase represents pate german phrase.
	pateGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "quaak"}

	// pateItalianPhrase represents pate italian phrase.
	pateItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "quackquack"}

	// pateJapanesePhrase represents pate japanese phrase.
	pateJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "メソメソ"}

	// pateKoreanPhrase represents pate korean phrase.
	pateKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아잉"}

	// pateLatinAmericanSpanishPhrase represents pate latin american spanish phrase.
	pateLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cuaqui"}

	// pateRussianPhrase represents pate russian phrase.
	pateRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кряки-кряк"}

	// pateSimplifiedChinesePhrase represents pate simplified chinese phrase.
	pateSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "呜呜"}

	// pateSpanishPhrase represents pate spanish phrase.
	pateSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cuaqui"}

	// pateTraditionalChinesePhrase represents pate traditional chinese phrase.
	pateTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嗚嗚"}
)

var (
	// patePhrase represents pate phrase.
	patePhrase = nook.Languages{
		language.AmericanEnglish:      pateAmericanEnglishPhrase,
		language.CanadianFrench:       pateCanadianFrenchPhrase,
		language.Dutch:                pateDutchPhrase,
		language.French:               pateFrenchPhrase,
		language.German:               pateGermanPhrase,
		language.Italian:              pateItalianPhrase,
		language.Japanese:             pateJapanesePhrase,
		language.Korean:               pateKoreanPhrase,
		language.LatinAmericanSpanish: pateLatinAmericanSpanishPhrase,
		language.Russian:              pateRussianPhrase,
		language.SimplifiedChinese:    pateSimplifiedChinesePhrase,
		language.Spanish:              pateSpanishPhrase,
		language.TraditionalChinese:   pateTraditionalChinesePhrase}
)

var (
	// Pate represents pate.
	Pate = nook.Villager{
		Character:   pateCharacter,
		Personality: personality.Peppy,
		Phrase:      patePhrase}
)

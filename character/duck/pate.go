package duck

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	pateBirthday = nook.Birthday{
		Day:   23,
		Month: time.February}
)

var (
	pateCode = nook.Code{
		Value: "duk02"}
)

var (
	pateAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pate"}

	pateCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Terrine"}

	pateDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pate"}

	pateFrenchName = nook.Name{
		Language: language.French,
		Value:    "Terrine"}

	pateGermanName = nook.Name{
		Language: language.German,
		Value:    "Daune"}

	pateItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Camilla"}

	pateJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ナッキー"}

	pateKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "나키"}

	pateLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pati"}

	pateRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пэйт"}

	pateSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "爱哭鬼"}

	pateSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pati"}

	pateTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "愛哭鬼"}
)

var (
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
	pateCharacter = nook.Character{
		Animal:   animal.Duck,
		Birthday: pateBirthday,
		Code:     pateCode,
		Gender:   gender.Female,
		Name:     pateName}
)

var (
	pateAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "quackle"}

	pateCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "coin coin"}

	pateDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kwakel"}

	pateFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "coin coin"}

	pateGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "quaak"}

	pateItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "quackquack"}

	pateJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "メソメソ"}

	pateKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아잉"}

	pateLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cuaqui"}

	pateRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кряки-кряк"}

	pateSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "呜呜"}

	pateSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cuaqui"}

	pateTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嗚嗚"}
)

var (
	patePhrase = nook.Languages{
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
	Pate = nook.Villager{
		Character:   pateCharacter,
		Personality: personality.Peppy,
		Phrase:      patePhrase}
)

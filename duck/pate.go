package duck

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Terrinecoin coin"}

	pateDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Patekwakel"}

	pateFrenchName = nook.Name{
		Language: language.French,
		Value:    "Terrinecoin coin"}

	pateGermanName = nook.Name{
		Language: language.German,
		Value:    "Daunequaak"}

	pateItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Camillaquackquack"}

	pateJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ナッキーメソメソ"}

	pateKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "나키아잉"}

	pateLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Paticuaqui"}

	pateRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пэйткряки-кряк"}

	pateSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "爱哭鬼呜呜"}

	pateSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Paticuaqui"}

	pateTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "愛哭鬼嗚嗚"}
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
		Animal:   Duck,
		Birthday: pateBirthday,
		Code:     pateCode,
		Gender:   nook.Female,
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
		Personality: nook.Peppy,
		Phrase:      patePhrase}
)

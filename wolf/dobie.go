package wolf

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	dobieBirthday = nook.Birthday{
		Day:   17,
		Month: time.February}
)

var (
	dobieCode = nook.Code{
		Value: "wol04"}
)

var (
	dobieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Dobie"}

	dobieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Loupiotloudonc"}

	dobieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Dobieeh"}

	dobieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Loupiotpeuchère"}

	dobieGermanName = nook.Name{
		Language: language.German,
		Value:    "Sigmundheuuul"}

	dobieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Sigmundgnams"}

	dobieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "けんゴホゴホ"}

	dobieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "켄음음"}

	dobieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Germánooomm"}

	dobieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Добиом-м-м..."}

	dobieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿研咳呵"}

	dobieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Germánooomm"}

	dobieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿研咳呵"}
)

var (
	dobieName = nook.Languages{
		language.AmericanEnglish:      dobieAmericanEnglishName,
		language.CanadianFrench:       dobieCanadianFrenchName,
		language.Dutch:                dobieDutchName,
		language.French:               dobieFrenchName,
		language.German:               dobieGermanName,
		language.Italian:              dobieItalianName,
		language.Japanese:             dobieJapaneseName,
		language.Korean:               dobieKoreanName,
		language.LatinAmericanSpanish: dobieLatinAmericanSpanishName,
		language.Russian:              dobieRussianName,
		language.SimplifiedChinese:    dobieSimplifiedChineseName,
		language.Spanish:              dobieSpanishName,
		language.TraditionalChinese:   dobieTraditionalChineseName}
)

var (
	dobieCharacter = nook.Character{
		Animal:   Wolf,
		Birthday: dobieBirthday,
		Code:     dobieCode,
		Gender:   nook.Male,
		Name:     dobieName}
)

var (
	dobieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ohmmm"}

	dobieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "loudonc"}

	dobieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "eh"}

	dobieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "peuchère"}

	dobieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "heuuul"}

	dobieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "gnams"}

	dobieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ゴホゴホ"}

	dobieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "음음"}

	dobieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ooomm"}

	dobieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ом-м-м..."}

	dobieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咳呵"}

	dobieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ooomm"}

	dobieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咳呵"}
)

var (
	dobiePhrase = nook.Languages{
		language.AmericanEnglish:      dobieAmericanEnglishName,
		language.CanadianFrench:       dobieCanadianFrenchName,
		language.Dutch:                dobieDutchName,
		language.French:               dobieFrenchName,
		language.German:               dobieGermanName,
		language.Italian:              dobieItalianName,
		language.Japanese:             dobieJapaneseName,
		language.Korean:               dobieKoreanName,
		language.LatinAmericanSpanish: dobieLatinAmericanSpanishName,
		language.Russian:              dobieRussianName,
		language.SimplifiedChinese:    dobieSimplifiedChineseName,
		language.Spanish:              dobieSpanishName,
		language.TraditionalChinese:   dobieTraditionalChineseName}
)

var (
	Dobie = nook.Villager{
		Character:   dobieCharacter,
		Personality: nook.Cranky,
		Phrase:      dobiePhrase}
)

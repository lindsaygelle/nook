package wolf

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
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
		Value:    "Loupiot"}

	dobieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Dobie"}

	dobieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Loupiot"}

	dobieGermanName = nook.Name{
		Language: language.German,
		Value:    "Sigmund"}

	dobieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Sigmund"}

	dobieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "けん"}

	dobieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "켄"}

	dobieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Germán"}

	dobieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Доби"}

	dobieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿研"}

	dobieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Germán"}

	dobieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿研"}
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
		Animal:   animal.Wolf,
		Birthday: dobieBirthday,
		Code:     dobieCode,
		Gender:   gender.Male,
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
		Personality: personality.Cranky,
		Phrase:      dobiePhrase}
)

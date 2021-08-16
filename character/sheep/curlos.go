package sheep

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
	curlosBirthday = nook.Birthday{
		Day:   8,
		Month: time.May}
)

var (
	curlosCode = nook.Code{
		Value: "shp08"}
)

var (
	curlosAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Curlos"}

	curlosCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Tonton"}

	curlosDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Curlos"}

	curlosFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tonton"}

	curlosGermanName = nook.Name{
		Language: language.German,
		Value:    "Locke"}

	curlosItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Boccolo"}

	curlosJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "カルロス"}

	curlosKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "카를로스"}

	curlosLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Carnerio"}

	curlosRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Керлос"}

	curlosSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "贾洛斯"}

	curlosSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Carnerio"}

	curlosTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "賈洛斯"}
)

var (
	curlosName = nook.Languages{
		language.AmericanEnglish:      curlosAmericanEnglishName,
		language.CanadianFrench:       curlosCanadianFrenchName,
		language.Dutch:                curlosDutchName,
		language.French:               curlosFrenchName,
		language.German:               curlosGermanName,
		language.Italian:              curlosItalianName,
		language.Japanese:             curlosJapaneseName,
		language.Korean:               curlosKoreanName,
		language.LatinAmericanSpanish: curlosLatinAmericanSpanishName,
		language.Russian:              curlosRussianName,
		language.SimplifiedChinese:    curlosSimplifiedChineseName,
		language.Spanish:              curlosSpanishName,
		language.TraditionalChinese:   curlosTraditionalChineseName}
)

var (
	curlosCharacter = nook.Character{
		Animal:   animal.Sheep,
		Birthday: curlosBirthday,
		Code:     curlosCode,
		Key:      character.Curlos,
		Gender:   gender.Male,
		Name:     curlosName}
)

var (
	curlosAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "shearly"}

	curlosCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "toison"}

	curlosDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "scheerlijk"}

	curlosFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "toison"}

	curlosGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "zottel"}

	curlosItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "hermoso"}

	curlosJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ベイビー"}

	curlosKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "자기야"}

	curlosLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "laaaná"}

	curlosRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "завитушка"}

	curlosSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "北鼻"}

	curlosSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "lanitas"}

	curlosTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "北鼻"}
)

var (
	curlosPhrase = nook.Languages{
		language.AmericanEnglish:      curlosAmericanEnglishName,
		language.CanadianFrench:       curlosCanadianFrenchName,
		language.Dutch:                curlosDutchName,
		language.French:               curlosFrenchName,
		language.German:               curlosGermanName,
		language.Italian:              curlosItalianName,
		language.Japanese:             curlosJapaneseName,
		language.Korean:               curlosKoreanName,
		language.LatinAmericanSpanish: curlosLatinAmericanSpanishName,
		language.Russian:              curlosRussianName,
		language.SimplifiedChinese:    curlosSimplifiedChineseName,
		language.Spanish:              curlosSpanishName,
		language.TraditionalChinese:   curlosTraditionalChineseName}
)

var (
	Curlos = nook.Villager{
		Character:   curlosCharacter,
		Personality: personality.Smug,
		Phrase:      curlosPhrase}
)

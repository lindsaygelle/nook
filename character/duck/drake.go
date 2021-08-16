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
	drakeBirthday = nook.Birthday{
		Day:   25,
		Month: time.June}
)

var (
	drakeCode = nook.Code{
		Value: "duk09"}
)

var (
	drakeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Drake"}

	drakeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Colvert"}

	drakeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Drake"}

	drakeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Colvert"}

	drakeGermanName = nook.Name{
		Language: language.German,
		Value:    "Gustav"}

	drakeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Drago"}

	drakeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フォアグラ"}

	drakeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "푸아그라"}

	drakeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Draco"}

	drakeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дрейк"}

	drakeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "肥肝"}

	drakeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Draco"}

	drakeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "肥肝"}
)

var (
	drakeName = nook.Languages{
		language.AmericanEnglish:      drakeAmericanEnglishName,
		language.CanadianFrench:       drakeCanadianFrenchName,
		language.Dutch:                drakeDutchName,
		language.French:               drakeFrenchName,
		language.German:               drakeGermanName,
		language.Italian:              drakeItalianName,
		language.Japanese:             drakeJapaneseName,
		language.Korean:               drakeKoreanName,
		language.LatinAmericanSpanish: drakeLatinAmericanSpanishName,
		language.Russian:              drakeRussianName,
		language.SimplifiedChinese:    drakeSimplifiedChineseName,
		language.Spanish:              drakeSpanishName,
		language.TraditionalChinese:   drakeTraditionalChineseName}
)

var (
	drakeCharacter = nook.Character{
		Animal:   animal.Duck,
		Birthday: drakeBirthday,
		Code:     drakeCode,
		Gender:   gender.Male,
		Name:     drakeName}
)

var (
	drakeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "quacko"}

	drakeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ben couac"}

	drakeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snabbel"}

	drakeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "canard"}

	drakeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnabbl"}

	drakeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "quaqua"}

	drakeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "かもね"}

	drakeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "어쩌면"}

	drakeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cui-cuac"}

	drakeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кра-а"}

	drakeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "或许鸭"}

	drakeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "polluelo"}

	drakeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "或許鴨"}
)

var (
	drakePhrase = nook.Languages{
		language.AmericanEnglish:      drakeAmericanEnglishName,
		language.CanadianFrench:       drakeCanadianFrenchName,
		language.Dutch:                drakeDutchName,
		language.French:               drakeFrenchName,
		language.German:               drakeGermanName,
		language.Italian:              drakeItalianName,
		language.Japanese:             drakeJapaneseName,
		language.Korean:               drakeKoreanName,
		language.LatinAmericanSpanish: drakeLatinAmericanSpanishName,
		language.Russian:              drakeRussianName,
		language.SimplifiedChinese:    drakeSimplifiedChineseName,
		language.Spanish:              drakeSpanishName,
		language.TraditionalChinese:   drakeTraditionalChineseName}
)

var (
	Drake = nook.Villager{
		Character:   drakeCharacter,
		Personality: personality.Lazy,
		Phrase:      drakePhrase}
)

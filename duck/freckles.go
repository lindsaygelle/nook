package duck

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	frecklesBirthday = nook.Birthday{
		Day:   19,
		Month: time.February}
)

var (
	frecklesCode = nook.Code{
		Value: "duk07"}
)

var (
	frecklesAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Freckles"}

	frecklesCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Caropoussin"}

	frecklesDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Freckleseendje"}

	frecklesFrenchName = nook.Name{
		Language: language.French,
		Value:    "Caropoussin"}

	frecklesGermanName = nook.Name{
		Language: language.German,
		Value:    "Quacksschnatter"}

	frecklesItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Semolaqua qua"}

	frecklesJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マグロぎょぎょ"}

	frecklesKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "다랑어물물"}

	frecklesLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rebecacui-cuá"}

	frecklesRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Фреклсутенок"}

	frecklesSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小鲔鱼鱼"}

	frecklesSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rebecacui-cuá"}

	frecklesTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小鮪魚魚"}
)

var (
	frecklesName = nook.Languages{
		language.AmericanEnglish:      frecklesAmericanEnglishName,
		language.CanadianFrench:       frecklesCanadianFrenchName,
		language.Dutch:                frecklesDutchName,
		language.French:               frecklesFrenchName,
		language.German:               frecklesGermanName,
		language.Italian:              frecklesItalianName,
		language.Japanese:             frecklesJapaneseName,
		language.Korean:               frecklesKoreanName,
		language.LatinAmericanSpanish: frecklesLatinAmericanSpanishName,
		language.Russian:              frecklesRussianName,
		language.SimplifiedChinese:    frecklesSimplifiedChineseName,
		language.Spanish:              frecklesSpanishName,
		language.TraditionalChinese:   frecklesTraditionalChineseName}
)

var (
	frecklesCharacter = nook.Character{
		Animal:   Duck,
		Birthday: frecklesBirthday,
		Code:     frecklesCode,
		Gender:   nook.Female,
		Name:     frecklesName}
)

var (
	frecklesAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ducky"}

	frecklesCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "poussin"}

	frecklesDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "eendje"}

	frecklesFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "poussin"}

	frecklesGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnatter"}

	frecklesItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "qua qua"}

	frecklesJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ぎょぎょ"}

	frecklesKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "물물"}

	frecklesLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cui-cuá"}

	frecklesRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "утенок"}

	frecklesSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "鱼鱼"}

	frecklesSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cui-cuá"}

	frecklesTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "魚魚"}
)

var (
	frecklesPhrase = nook.Languages{
		language.AmericanEnglish:      frecklesAmericanEnglishName,
		language.CanadianFrench:       frecklesCanadianFrenchName,
		language.Dutch:                frecklesDutchName,
		language.French:               frecklesFrenchName,
		language.German:               frecklesGermanName,
		language.Italian:              frecklesItalianName,
		language.Japanese:             frecklesJapaneseName,
		language.Korean:               frecklesKoreanName,
		language.LatinAmericanSpanish: frecklesLatinAmericanSpanishName,
		language.Russian:              frecklesRussianName,
		language.SimplifiedChinese:    frecklesSimplifiedChineseName,
		language.Spanish:              frecklesSpanishName,
		language.TraditionalChinese:   frecklesTraditionalChineseName}
)

var (
	Freckles = nook.Villager{
		Character:   frecklesCharacter,
		Personality: nook.Peppy,
		Phrase:      frecklesPhrase}
)

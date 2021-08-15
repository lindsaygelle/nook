package cat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	katieBirthday = nook.Birthday{
		Day:   22,
		Month: time.October}
)

var (
	katieCode = nook.Code{
		Value: "los/lom"}
)

var (
	katieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Katie"}

	katieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Cathie"}

	katieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Katie"}

	katieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Cathie"}

	katieGermanName = nook.Name{
		Language: language.German,
		Value:    "Katja"}

	katieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Micetta"}

	katieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "まいこちゃん"}

	katieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "미야"}

	katieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cati"}

	katieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кейти"}

	katieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咪露"}

	katieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cati"}

	katieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咪露"}
)

var (
	katieName = nook.Languages{
		language.AmericanEnglish:      katieAmericanEnglishName,
		language.CanadianFrench:       katieCanadianFrenchName,
		language.Dutch:                katieDutchName,
		language.French:               katieFrenchName,
		language.German:               katieGermanName,
		language.Italian:              katieItalianName,
		language.Japanese:             katieJapaneseName,
		language.Korean:               katieKoreanName,
		language.LatinAmericanSpanish: katieLatinAmericanSpanishName,
		language.Russian:              katieRussianName,
		language.SimplifiedChinese:    katieSimplifiedChineseName,
		language.Spanish:              katieSpanishName,
		language.TraditionalChinese:   katieTraditionalChineseName}
)

var (
	katieCharacter = nook.Character{
		Animal:   Cat,
		Birthday: katieBirthday,
		Code:     katieCode,
		Gender:   nook.Female,
		Name:     katieName}
)

var (
	Katie = nook.Resident{
		Character: katieCharacter}
)
